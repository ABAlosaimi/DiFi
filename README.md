# DiFi
#### reuse your code without OOP stuff

DiFi is a CLI (command line interface) tool that enable you to reuse your pervious written feature in a brwon/new project(s) by cleaning up project-related details (like package statments) and write the cleaned content to your project directory.

## Features

- Open and read any file 
- Filter file content by skipping a specified number of initial lines don't needed 
- Write filtered content to a target directory with the same filename

Note: you should check about permissions on read & write to the file you are trying to reuse. 

## Project Structure

```
DiFi/
├── cmd/
│   └── main/
│       └── main.go           # Entry point of the application
├── internal/
│   └── combiner/
│       └── projectCombiner.go # Core logic
├── go.mod                     # Go module definition
└── README.md                 
```

## Local installation 

1. Clone the repo:
```bash
git clone https://github.com/ABAlosaimi/DiFi
```

2. Build the project:
```bash
go build -o difi ./cmd/main
```

## Usage

Run the compiled binary or use `go run`:

```bash
go run main.go <filePath> <filteringStartLineNum> <projectDir>
```
Note: You must be in the main directory

#### Arguments 

- `filePath`: Absolute path to the file you want to filter
- `filteringStartLineNum`: Number of lines to skip from the beginning of the file (0-indexed)
- `projectDir`: The target directory where the filtered file will be written

#### Example

```bash
go run main.go "/path/to/input.txt" 5 "/path/to/output/dir"
```

The workflow will be:
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

## Requirements to run locally

- Go 1.16 or higher

## Author
Claude code, reviewed by eng.Abdulrahman F. Alosaimi 
