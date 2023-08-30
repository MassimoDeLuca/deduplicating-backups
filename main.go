package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Directories []string `json:"directories"`
	OutputDir   string   `json:"outputDir"`
}

var config Config
var db *sql.DB

func handleFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(path, err) // can't walk here,
		return nil             // but continue walking elsewhere
	}

	if info.IsDir() {
		fmt.Println("Directory:", path)
	} else {
		fmt.Println("File:", path)

		if err := ProcessFile(path, config.OutputDir); err != nil {
			fmt.Println("Process file error:", err)
			return err
		}

	}
	return nil
}

func main() {
	fmt.Println("Hello World")
	startTime := time.Now()

	// ############################
	// Initialize Database
	// ############################
	db, err := InitDatabase("database.sqlite")
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	// insertDataSQL := `INSERT INTO users (username, email) VALUES (?, ?);`
	// _, err = db.Exec(insertDataSQL, "john_doe2", "john2@example.com")
	// if err != nil {
	// 	fmt.Println("Error inserting data:", err)
	// 	return
	// }

	// ############################
	// Parse config
	// ############################
	config, err := ParseConfig()
	if err != nil {
		fmt.Println("Parsing config error:", err)
		return
	}

	// ############################
	// Init output folder
	// ############################
	if err := os.MkdirAll(config.OutputDir, os.ModePerm); err != nil {
		fmt.Println("Error creating destination folder:", err)
		return
	}

	// ############################
	// Itterate directories
	// ############################
	fmt.Println("Directories:")
	for _, dir := range config.Directories {

		err := filepath.Walk(dir, handleFile)
		if err != nil {
			fmt.Printf("error walking the path %v: %v\n", dir, err)
		}
	}

	fmt.Println("Files copied successfully!")

	// // Record the end time
	fmt.Println("Elapsed time:", time.Since(startTime))
}
