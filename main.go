package main

import (
	"os"
	"unisun/api/auth-listener/src"
	"unisun/api/auth-listener/src/config"
	"unisun/api/auth-listener/src/constants"
)

func main() {
	if os.Getenv(constants.NODE) != constants.PRODUCTION {
		config.SetENV()
	}
	r := src.App()
	port := os.Getenv(constants.PORT)
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
