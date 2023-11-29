package env

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/time/rate"
)

const (
	defaultRateLimit = 20
	defaultPort      = ":8080"
	defaultSecret    = "my-secret"
)

type env struct {
	dbUrl     string
	port      string
	rateLimit string
	jwtSecret string
}

func Data() env {
	return env{
		dbUrl:     os.Getenv("DB_URL"),
		port:      os.Getenv("PORT"),
		rateLimit: os.Getenv("RATE_LIMIT"),
		jwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func (e env) Port() string {
	if e.port == "" {
		return defaultPort
	}
	return ":" + e.port
}

func (e env) RateLimit() rate.Limit {
	parsedRateLimit, err := strconv.ParseFloat(e.rateLimit, 64)
	if err != nil {
		log.Fatal("Invalid rate limit value")
		parsedRateLimit = defaultRateLimit
	}
	return rate.Limit(parsedRateLimit)
}

func (e env) DbUrl() string {
	return e.dbUrl
}

func (e env) JwtSecret() string {
	if e.jwtSecret == "" {
		return defaultSecret
	}
	return e.jwtSecret
}
