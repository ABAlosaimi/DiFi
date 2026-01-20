# DiFi - Dynamic features reuse

DiFi is a CLI (command line interface) tool that enable you to copy your pervious written feature(s) to a brwon/new projects by removing the project-related details (like package statments) and writing the filtered content to a new location (your project). This project utilizes the provided operations in the loacl OS (Operating System) you are using to do its functions. 

## Features

- Open and read any text file
- Filter file content by skipping a specified number of initial lines don't needed 
- Write filtered content to a target directory with the same filename

## Project Structure

```
DiFi/
├── cmd/
│   └── main/
│       └── main.go           # Entry point of the application
├── internal/
│   └── combiner/
│       └── projectCombiner.go # Core filtering logic
├── go.mod                     # Go module definition
└── README.md                 
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd DiFi
```

2. Build the project:
```bash
go build -o difi ./cmd/main
```

## Usage

Run the compiled binary or use `go run`:

```bash
go run ./cmd/main/main.go <filePath> <filteringStartLineNum> <projectDir>
```

### Arguments

- `filePath`: Absolute path to the file you want to filter
- `filteringStartLineNum`: Number of lines to skip from the beginning (0-indexed)
- `projectDir`: Target directory where the filtered file will be written

### Example

```bash
go run ./cmd/main/main.go "/path/to/input.txt" 5 "/path/to/output/dir"
```

This will:
1. Open `/path/to/input.txt`
2. Skip the first 5 lines
3. Write the remaining content to `/path/to/output/dir/input.txt`

## Use Cases

### 1. Removing Headers from Data Files
Skip header rows in CSV or text files before processing:
```bash
difi data.csv 1 ./processed
```

### 2. Extracting Content from Templates
Remove boilerplate or template headers from code files:
```bash
difi template.go 10 ./src
```

### 3. Cleaning Log Files
Remove initial metadata or timestamps from log files:
```bash
difi app.log 3 ./clean-logs
```

## How It Works

DiFi operates in three main steps:

1. **Open**: Reads the source file from the specified path
2. **Filter**: Scans the file line-by-line and skips the specified number of initial lines
3. **Write**: Creates a new file in the target directory with the filtered content

## Requirements

- Go 1.16 or higher

## Author
Claude code, reviewed by eng.Abdulrahman F. Alosaimi 
