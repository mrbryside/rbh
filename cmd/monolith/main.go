package main

import (
	"fmt"

	"github.com/brpaz/echozap"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment/gormappointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment/gormcomment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator/libcreator"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history/gormhistory"
	"github.com/mrbryside/rbh/domain/interview/appointment/rest/echo/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/rest/echo/comment"
	appService "github.com/mrbryside/rbh/domain/interview/appointment/service"
	myJwt "github.com/mrbryside/rbh/domain/user/domain/authorization/jwt"
	"github.com/mrbryside/rbh/domain/user/domain/user/gormuser"
	"github.com/mrbryside/rbh/domain/user/rest/echo/auth"
	"github.com/mrbryside/rbh/domain/user/service"
	"github.com/mrbryside/rbh/pkg/claim"
	"github.com/mrbryside/rbh/pkg/db/mygorm"
	"github.com/mrbryside/rbh/pkg/env"
	"github.com/mrbryside/rbh/pkg/logger"
	"github.com/mrbryside/rbh/pkg/mymiddleware"
	"gorm.io/gorm"
)

func main() {
	// load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// init gorm db
	db, close, err := mygorm.BasicConnection()
	if err != nil {
		panic(err.Error())
	}
	defer close()

	e, g := initEcho()

	authService := initUserService(db, e)
	// lib-relation between user-service and interview-appointment-service
	initInterviewAppointmentService(db, g, authService)

	e.Start(env.Data().Port())
}

func initEcho() (*echo.Echo, *echo.Group) {
	// init echo and register log, middleware, create group
	e := echo.New()
	e.Use(echozap.ZapLogger(logger.Log))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(env.Data().RateLimit())))

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(claim.CustomClaims)
		},
		SigningKey: []byte(env.Data().JwtSecret()),
	}
	g := e.Group("")
	g.Use(echojwt.WithConfig(config))
	g.Use(mymiddleware.ExtractFromClaims)
	return e, g
}

func initUserService(db *gorm.DB, e *echo.Echo) service.AuthServicer {
	// init all domain
	userDomain := gormuser.NewRepository(db)
	jwtDomain := myJwt.NewJwt(env.Data().JwtSecret())

	// init all service
	jwtService := service.NewJwtService(jwtDomain)
	authService := service.NewAuthService(userDomain, jwtService)

	// init handler with service
	authHandler := auth.NewAuthHandler(authService)

	// register all routes
	authHandler.RegisterRoutes(e)

	return authService
}

func initInterviewAppointmentService(db *gorm.DB, g *echo.Group, authService service.AuthServicer) {
	// init all domain
	appointmentDomain := gormappointment.NewRepository(db)
	commentDomain := gormcomment.NewRepository(db)
	// right here is relationship between user-service and interview-appointment-service (now is a lib integration)
	// we can change this to be client integration and split to micro-service later
	creatorDomain := libcreator.NewRepository(authService)
	historyDomain := gormhistory.NewRepository(db)

	// init all service
	creatorService := appService.NewCreatorService(creatorDomain)
	commentService := appService.NewCommentService(commentDomain, creatorService)
	historyService := appService.NewHistoryService(historyDomain)
	appointmentService := appService.NewAppointmentService(appointmentDomain, historyService, creatorService, commentService)

	// init handler with service
	appointmentHandler := appointment.NewHandler(appointmentService)
	commentHandler := comment.NewHandler(commentService)

	// register all routes
	appointmentHandler.RegisterRoutes(g)
	commentHandler.RegisterRoutes(g)
}
