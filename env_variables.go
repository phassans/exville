package main

import "os"

const (
	rocketURL           = "http://65.255.36.179:3000"
	rocketAdminUser     = "pramod"
	rocketAdminPassword = "123456"

	phantomURL = "https://phantombuster.com"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_DATABASE")
)
