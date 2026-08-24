package config

import "os"

type Config struct {
	Port        string
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	DB_Name     string
}

func Load() Config {
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_Host")
	dbPort := os.Getenv("DB_Port")
	dbUser := os.Getenv("DB_User")
	dbPassword := os.Getenv("DB_Password")
	dbName := os.Getenv("DB_Name")

	if port == "" {
		port = "8080"
	}

	if dbHost == "" {
		dbHost = "localhost"
	}

	if dbPort == "" {
		dbPort = "5432"
	}

	if dbUser == "" {
		dbUser = "postgres"
	}

	if dbPassword == "" {
		dbPassword = "postgres"
	}

	if dbName == "" {
		dbName = "user-service"
	}

	return Config{
		Port:        port,
		DB_Host:     dbHost,
		DB_Port:     dbPort,
		DB_User:     dbUser,
		DB_Password: dbPassword,
		DB_Name:     dbName,
	}
}
