package env

import "os"

type env struct {
	dbUrl     string
	Port      string
	jwtSecret string
}

func Data() env {
	return env{
		dbUrl:     os.Getenv("DB_URL"),
		Port:      os.Getenv("PORT"),
		jwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func (e env) DbUrl() string {
	return e.dbUrl
}

func (e env) JwtSecret() string {
	return e.jwtSecret
}
