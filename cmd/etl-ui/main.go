package main

import (
	"etl-ui/internal/ui"
	"etl-ui/pkg/config"
	"log"
)

func main() {
	// Load configuration
	err := config.LoadConfig("config.json")
	if err != nil {
		log.Printf("Failed to load config: %v. Using default configuration.", err)
	}

	// Initialize and run the UI
	termUI := ui.NewTerminalUI()
	if err := termUI.Run(); err != nil {
		log.Fatalf("Error running UI: %v", err)
	}
}
