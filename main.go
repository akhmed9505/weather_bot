package main

import (
	"context"
	"github.com/akhmed9505/weatherbot/clients/openweather"
	"github.com/akhmed9505/weatherbot/handler"
	"github.com/akhmed9505/weatherbot/repo"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Error ping db")
	}

	conn.QueryRow(context.Background(), "select id from users")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	owClient := openweather.New(os.Getenv("OPENWEATHERAPI_KEY"))

	userRepo := repo.New(conn)

	botHandler := handler.New(bot, owClient, userRepo)

	botHandler.Start()
}
