package main

import (
	"log"
	"os"

	"github.com/comfysweet/go-in-action/chapter2/sample/search"
)

// init is called prior to main
func init() {
	// Change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
	// Perform the search for a specified term
	search.Run("president")
}
