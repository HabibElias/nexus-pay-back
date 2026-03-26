package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/HabibElias/nexus-pay-back/internal/config"
	"github.com/HabibElias/nexus-pay-back/internal/domain/entities"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it, continuing with environment variables")
	}

	cfg := config.LoadConfig()

	// Setup database using empty config, SetupDatabase falls back to env vars
	db := config.SetupDatabase(*cfg)

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please provide a command: migrate, seed, clean")
		os.Exit(1)
	}

	command := args[0]

	switch command {
	case "migrate":
		fmt.Println("Running migrations...")
		config.Migrate()
		fmt.Println("Migrations completed successfully.")
	case "seed":
		fmt.Println("Running seeder...")
		config.Seed()
		fmt.Println("Seeding completed successfully.")
	case "clean":
		fmt.Println("Cleaning database...")
		// Assuming we want to drop tables
		err := db.Migrator().DropTable(&entities.Payment{})
		if err != nil {
			log.Fatalf("Failed to drop tables: %v", err)
		}
		fmt.Println("Database cleaned successfully.")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: migrate, seed, clean")
		os.Exit(1)
	}
}
