// Build a Go program that:

// Has a slice of log lines

// Spawns goroutines to analyze them

// Sends result back via channel

// Counts total errors

// Prints final summary

//	logs := []string{
//		"INFO: Server started",
//		"ERROR: Database connection failed",
//		"WARNING: Disk space low",
//		"ERROR: Timeout occurred",
//		"INFO: Request completed",
//	}
package main

import (
	"fmt"
	"strings"
	"time"
)

type logResult struct {
	Line    string
	IsError bool
}

func analyzeLog(line string, result chan logResult) {
	time.Sleep(2 * time.Second)
	isError := strings.Contains(line, "ERROR")
	result <- logResult{Line: line, IsError: isError}
}

func main() {
	fmt.Println("main func started")
	result := make(chan logResult)
	logs := []string{
		"INFO: Server started",
		"ERROR: Database connection failed",
		"WARNING: Disk space low",
		"ERROR: Timeout occurred",
		"INFO: Request completed",
	}
	for _, log := range logs {
		go analyzeLog(log, result)
	}
	errorCount := 0
	for i := 0; i < len(logs); i++ {
		r := <-result
		fmt.Println("Processed", r.Line)
		if r.IsError {
			errorCount++
		}
	}
	fmt.Println("Total Error found:", errorCount)
}
