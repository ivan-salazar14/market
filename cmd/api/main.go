package main

import (
	"os"

	"backend_crudgo/infrastructure"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting API CMD")
	infrastructure.InitLogger()

	port := os.Getenv("API_PORT")
	infrastructure.Start(port)
}
