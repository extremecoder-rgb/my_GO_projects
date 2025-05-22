# Log Parser

A Go-based log parsing utility that helps analyze and filter log files with features like timestamp filtering, log level filtering, and summary statistics.

## Features

- Parse log files with timestamp and log level information
- Filter logs by level (INFO, ERROR, etc.)
- Filter logs by time range
- Generate summary statistics including:
  - Total number of entries
  - Count by log level
  - Earliest and latest log timestamps

## Installation

```bash
# Clone the repository
git clone <repository-url>

# Navigate to the project directory
cd log_parser

# Build the project
go build
```

## Usage

```bash
./log_parser <log_file_path>
```

### Expected Log Format

The parser expects logs in the following format:
```
[YYYY-MM-DD HH:MM:SS] LEVEL MESSAGE
```

Example:
```
[2025-05-22 20:21:34] INFO User logged in
[2025-05-22 20:21:35] ERROR Database connection failed
```

## API Documentation

### LogParser

The main struct that handles log parsing and analysis.

#### Methods

##### `NewLogParser() *LogParser`
Creates a new instance of LogParser.

##### `ParseLogFile(filePath string) error`
Parses a log file at the specified path.

##### `FilterByLevel(level string) []LogEntry`
Filters log entries by the specified log level (case-insensitive).

##### `FilterByTimeRange(start, end time.Time) []LogEntry`
Filters log entries within the specified time range.

##### `PrintSummary()`
Prints a summary of the parsed logs including:
- Total number of entries
- Count of entries by log level
- Timestamp of earliest and latest entries

### LogEntry

Represents a single log entry.

#### Fields

- `Timestamp time.Time`: The timestamp of the log entry
- `Level string`: The log level (INFO, ERROR, etc.)
- `Message string`: The log message

## Example

```go
parser := NewLogParser()
err := parser.ParseLogFile("sample.log")
if err != nil {
    log.Fatal(err)
}

// Print summary statistics
parser.PrintSummary()

// Filter INFO level entries
infoEntries := parser.FilterByLevel("INFO")

// Filter entries by time range
startTime := time.Parse("2006-01-02 15:04:05", "2025-05-22 00:00:00")
endTime := time.Parse("2006-01-02 15:04:05", "2025-05-23 00:00:00")
timeFiltered := parser.FilterByTimeRange(startTime, endTime)
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.