package config

import (
	"log"

	"doc-editor/prisma/db"

	"github.com/joho/godotenv"
)

var Prisma *db.PrismaClient

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Prisma = db.NewClient()
	err = Prisma.Connect()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
}
