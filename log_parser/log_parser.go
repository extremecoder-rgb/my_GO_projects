package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

type LogParser struct {
	Entries       []LogEntry
	LevelCounts   map[string]int
	EarliestTime  time.Time
	LatestTime    time.Time
}

func NewLogParser() *LogParser {
	return &LogParser{
		Entries:     []LogEntry{},
		LevelCounts: make(map[string]int),
	}
}

func (lp *LogParser) ParseLogFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Regular expression for common log format: [TIMESTAMP] LEVEL MESSAGE
	// Example: [2025-05-22 20:21:34] INFO User logged in
	logPattern := regexp.MustCompile(`^\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})\] (\w+) (.*)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := logPattern.FindStringSubmatch(line)
		if len(matches) != 4 {
			continue 
		}

		// Parse timestamp
		timestamp, err := time.Parse("2006-01-02 15:04:05", matches[1])
		if err != nil {
			continue 
		}

		
		entry := LogEntry{
			Timestamp: timestamp,
			Level:     matches[2],
			Message:   matches[3],
		}

		lp.Entries = append(lp.Entries, entry)
		lp.LevelCounts[entry.Level]++
		
	
		if lp.EarliestTime.IsZero() || timestamp.Before(lp.EarliestTime) {
			lp.EarliestTime = timestamp
		}
		if lp.LatestTime.IsZero() || timestamp.After(lp.LatestTime) {
			lp.LatestTime = timestamp
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	return nil
}


func (lp *LogParser) FilterByLevel(level string) []LogEntry {
	var filtered []LogEntry
	for _, entry := range lp.Entries {
		if strings.ToUpper(entry.Level) == strings.ToUpper(level) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}


func (lp *LogParser) FilterByTimeRange(start, end time.Time) []LogEntry {
	var filtered []LogEntry
	for _, entry := range lp.Entries {
		if (entry.Timestamp.Equal(start) || entry.Timestamp.After(start)) && 
		   (entry.Timestamp.Equal(end) || entry.Timestamp.Before(end)) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}


func (lp *LogParser) PrintSummary() {
	fmt.Println("Log Analysis Summary:")
	fmt.Printf("Total Entries: %d\n", len(lp.Entries))
	fmt.Println("Entries by Level:")
	for level, count := range lp.LevelCounts {
		fmt.Printf("  %s: %d\n", level, count)
	}
	if !lp.EarliestTime.IsZero() {
		fmt.Printf("Earliest Entry: %s\n", lp.EarliestTime.Format("2006-01-02 15:04:05"))
	}
	if !lp.LatestTime.IsZero() {
		fmt.Printf("Latest Entry: %s\n", lp.LatestTime.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: log_parser <log_file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	parser := NewLogParser()

	
	err := parser.ParseLogFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	parser.PrintSummary()

	// Example: Filter by log level
	fmt.Println("\nExample: INFO level entries:")
	infoEntries := parser.FilterByLevel("INFO")
	for _, entry := range infoEntries {
		fmt.Printf("[%s] %s %s\n", 
			entry.Timestamp.Format("2006-01-02 15:04:05"), 
			entry.Level, 
			entry.Message)
	}

	// Example: Filter by time range
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2025-05-22 00:00:00")
	endTime, _ := time.Parse("2006-01-02 15:04:05", "2025-05-23 00:00:00")
	fmt.Println("\nExample: Entries in time range:")
	timeFiltered := parser.FilterByTimeRange(startTime, endTime)
	for _, entry := range timeFiltered {
		fmt.Printf("[%s] %s %s\n", 
			entry.Timestamp.Format("2006-01-02 15:04:05"), 
			entry.Level, 
			entry.Message)
	}
}