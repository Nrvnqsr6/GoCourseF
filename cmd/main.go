package main

import "final/internal/service"

func main() {
	router := service.SetupRouter()

	router.Run("localhost:8080")
}
