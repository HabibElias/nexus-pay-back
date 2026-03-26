package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// SetupDatabase initializes the DB connection using cfg and stores it in the database package.
func SetupDatabase(cfg Config) *gorm.DB {
	user := firstNonEmpty([]string{cfg.DBUser, getEnv("DB_USER", "")})
	password := firstNonEmpty([]string{cfg.DBPassword, getEnv("DB_PASSWORD", "")})
	name := firstNonEmpty([]string{cfg.DBName, getEnv("DB_NAME", "")})
	host := firstNonEmpty([]string{cfg.DBHost, getEnv("DB_HOST", "")})
	port := firstNonEmpty([]string{cfg.DBPort, getEnv("DB_PORT", "")})

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, name, host, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	return db
}

func firstNonEmpty(values []string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func Migrate() {
	err := DB.AutoMigrate(
	// &entities.Todo{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
		return
	}

}

// Todo: Implement seeding
func Seed() {
	// seedFile, err := os.ReadFile("./data.json")
	// if err != nil {
	// 	log.Fatalf("can't read file, error: %s", err)
	// 	return
	// }

	// err = Seeder(seedFile)
	// if err != nil {
	// 	log.Fatalf("can't seed file, error: %s", err)
	// 	return
	// }
}
