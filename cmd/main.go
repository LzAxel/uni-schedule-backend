package main

import "uni-schedule-backend/internal/app"

// @title Uni Schedule API
// @version 0.1
// @description This is an API for Uni Schedule App.
// @host localhost:3000
// @BasePath /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	appInstance := app.New()
	appInstance.Run()
}
