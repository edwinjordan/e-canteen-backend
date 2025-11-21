package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/edwinjordan/e-canteen-backend/pkg/mysql"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db := mysql.DBConnectGorm()

	// Run permission seeder
	log.Println("=== Starting Permission Seeder ===")
	permissionSeeder := NewPermissionSeeder(db)
	err = permissionSeeder.SeedPermissions()
	if err != nil {
		log.Fatalf("Permission seeder failed: %v", err)
	}

	log.Println("=== Seeding completed successfully ===")
	os.Exit(0)
}
