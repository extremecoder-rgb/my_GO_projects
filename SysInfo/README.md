# SysInfo

A Go-based system information utility that provides detailed information about your computer's hardware and operating system.

## Features

SysInfo displays comprehensive system information including:

- Basic System Details
  - Hostname
  - Operating System
  - System Architecture
  - OS Version

- CPU Information
  - Processor Model
  - Number of Cores
  - Current CPU Usage

- Memory Statistics
  - Total RAM
  - Used Memory
  - Free Memory
  - Memory Usage Percentage

- Disk Information
  - Drive Details (C: drive)
  - Total Space
  - Used Space
  - Free Space
  - Usage Percentage

## Prerequisites

- Go 1.24.3 or higher
- Required Go packages (automatically installed via go.mod):
  - github.com/shirou/gopsutil
  - github.com/go-ole/go-ole
  - github.com/yusufpapurcu/wmi
  - golang.org/x/sys

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/extremecoder-rgb/SysInfo.git
   cd SysInfo
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

## Building

To build the executable:

```sh
go build sysinfo.go
```

This will create an executable named `sysinfo.exe` (on Windows) or `sysinfo` (on Unix-based systems).

## Usage

Simply run the executable:

```sh
./sysinfo
```

The program will display detailed system information in a formatted output, including:

```
=== System Information ===
Hostname: YourHostname
Operating System: windows
Architecture: amd64
OS Version: 10.0.0

CPU Information:
  Model: Intel(R) Core(TM) i7-9750H
  Cores: 6
  Usage: 25.50%

Memory Information:
  Total: 16.00 GB
  Used: 8.50 GB
  Free: 7.50 GB
  Usage: 53.12%

Disk Information:
  Drive: C:\
  Total: 500.00 GB
  Used: 300.00 GB
  Free: 200.00 GB
  Usage: 60.00%
```

## Dependencies

This project uses the following external packages:

- [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc)
- [go-ole](https://github.com/go-ole/go-ole) - Go bindings for Windows COM
- [wmi](https://github.com/yusufpapurcu/wmi) - WMI interface for Go to retrieve system information

## License

This project is open source and available under the MIT License.

## Author

- extremecoder-rgb