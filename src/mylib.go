package rohitmodule

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func ReadCSV(file_path string) int {
	// Open the CSV file
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close() // Ensure the file is closed after we are done

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatalf("failed to read CSV: %s", err)
	}

	// Iterate over the records and print them
	for _, record := range records {
		fmt.Println(record) // Each record is a slice of strings
	}
	return 1
}

func ReadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_url := os.Getenv("DATABASE_URL")
	secret_key := os.Getenv("SECRET_KEY")
	fmt.Println(db_url)
	fmt.Println(secret_key)
}
