package main

import (
	"github.com/joho/godotenv"
	"github.com/ryuji-cre8ive/super-suica/internal/database"
	"github.com/ryuji-cre8ive/super-suica/internal/stores"
	"github.com/ryuji-cre8ive/super-suica/internal/ui"
	"github.com/ryuji-cre8ive/super-suica/internal/usecase"
	"golang.org/x/xerrors"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	db, err := database.New()
	if err != nil {
		log.Fatal(xerrors.Errorf("failed to connect to database: %w", err))
	}
	postgres, err := db.DB()
	defer postgres.Close()

	e := ui.Echo()

	s := stores.New(db)
	ss := usecase.New(s)
	h := ui.New(ss)

	ui.SetApi(e, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err != nil {
		log.Fatal(xerrors.Errorf("failed to create new bot: %w", err))
	}

	e.Logger.Fatal(e.Start(":" + port))
}
