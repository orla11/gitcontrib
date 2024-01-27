package main

import (
	"flag"
	"gitcontrib/internal/scan"
	"gitcontrib/internal/stats"
)

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan.Scan(folder)
		return
	}

	stats.Stats(email)
}