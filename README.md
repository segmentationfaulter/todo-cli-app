# Tasks - A Terminal-Based Todo List Manager

A simple, fast, and reliable command-line todo list application built with Go and Cobra. Manage your tasks directly from the terminal with full CRUD operations and data integrity guarantees.

## Project Specification

This project is based on the todo list specification from [Dreams of Code's Go Projects](https://github.com/dreamsofcode-io/goprojects/tree/main/01-todo-list). The implementation follows the requirements outlined in that specification while adding additional technical features for production readiness.

## Features

- ✅ **Add, list, complete, and delete tasks**
- 📁 **Persistent CSV storage** in `~/.tasks/tasks.csv`
- 🔒 **File locking** prevents data corruption from concurrent access
- ⚛️ **Atomic operations** ensure data integrity during updates
- 🚀 **Fast and lightweight** - single binary with no dependencies
- 🎯 **Simple CLI interface** with intuitive commands
- 📊 **Tabular output** for easy reading
- 🏷️ **Task filtering** (show all vs pending only)

## Installation

### From Source

```bash
git clone https://github.com/segmentationfaulter/todo-cli-app
cd todo-cli-app
go build -o tasks
sudo mv tasks /usr/local/bin/  # Optional: make available system-wide
```

### Prerequisites

- Go 1.24+ installed on your system

## Usage

### Adding Tasks

```bash
# Add a new task
tasks add "Buy groceries"
tasks add "Finish the Go project"
tasks add "Call mom"
```

### Listing Tasks

```bash
# Show all pending tasks
tasks list

# Show all tasks (including completed)
tasks list --all
tasks list -a
```

Example output:
```
ID   Task                    Created at           Done
1    Buy groceries          02 Jan 25 14:30 MST  false
2    Finish the Go project  02 Jan 25 14:31 MST  false
3    Call mom               02 Jan 25 14:32 MST  false
```

### Completing Tasks

```bash
# Mark a task as completed
tasks complete 1
```

### Deleting Tasks

```bash
# Delete a task by ID
tasks delete 2
```

### Other Commands

```bash
# Show version
tasks --version

# Show help
tasks --help
tasks [command] --help
```

## Technical Features

### Data Integrity

- **File Locking**: Uses `syscall.Flock` to prevent concurrent read/write operations
- **Atomic Updates**: Writes to temporary files and atomically renames them to prevent data corruption
- **Error Handling**: Comprehensive error handling with meaningful messages

### Storage

- **Format**: CSV (Comma-Separated Values)
- **Location**: `~/.tasks/tasks.csv`
- **Schema**: `ID,Description,CreatedAt,IsComplete`
- **Automatic Directory Creation**: Creates `~/.tasks/` if it doesn't exist

### Architecture

```
todo-cli-app/
├── main.go              # Application entry point
├── cmd/                 # Command definitions (Cobra pattern)
│   ├── root.go         # Root command setup
│   ├── add.go          # Add task command
│   ├── list.go         # List tasks command
│   ├── complete.go     # Complete task command
│   ├── delete.go       # Delete task command
│   └── utils.go        # Shared utilities and atomic operations
├── storage/             # Data persistence layer
│   └── storage.go      # File operations and CSV handling
├── go.mod              # Go module definition
└── README.md           # This file
```

## Example Data File

The CSV file structure:
```csv
ID,Description,CreatedAt,IsComplete
1,Buy groceries,1704207019,false
2,Finish the Go project,1704207079,true
3,Call mom,1704207139,false
```

## Development

### Building

```bash
# Build for current platform
go build -o tasks

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o tasks-linux
GOOS=windows GOARCH=amd64 go build -o tasks-windows.exe
GOOS=darwin GOARCH=amd64 go build -o tasks-macos
```

### Project Structure

- **cmd/**: Contains all CLI command implementations using the Cobra library pattern
- **storage/**: Handles all file I/O operations with proper locking and error handling
- **main.go**: Simple entry point that delegates to the cmd package

### Dependencies

- `github.com/spf13/cobra` - CLI framework for building command-line applications
- `github.com/spf13/pflag` - POSIX-compliant command-line flag parsing (Cobra dependency)

## Design Decisions

### Why CSV?

- **Simplicity**: Human-readable format that's easy to debug and inspect
- **Portability**: Can be opened in spreadsheet applications if needed
- **Lightweight**: No database dependencies or complex setup required
- **Atomic Operations**: Easy to implement safe concurrent access patterns

### Why File Locking?

Prevents data corruption when multiple instances of the application run simultaneously:

```go
// Exclusive lock obtained on the file descriptor
if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
    return nil, err
}
```

### Why Atomic Operations?

Updates are performed by writing to a temporary file and then atomically renaming it:

```go
// Write to temporary file first, then rename (atomic operation)
tmpFile := filepath + ".tmp"
// ... write all data to tmpFile ...
os.Rename(tmpFile, filepath) // Atomic on most filesystems
```

This ensures that the data file is never in a corrupted state, even if the application crashes during an update.

## Error Handling

The application provides clear error messages for common scenarios:

- Invalid task IDs
- File permission issues
- Corrupted data files
- Missing tasks

## Future Enhancements

- [ ] Task due dates and priorities
- [ ] Task categories/tags
- [ ] Search and filter functionality
- [ ] Export to different formats (JSON, YAML)
- [ ] Task completion timestamps
- [ ] Recurring tasks
- [ ] Task notes/descriptions

## Contributing

This is a learning project, but feedback and suggestions are welcome! Please feel free to:

1. Open issues for bugs or feature requests
2. Submit pull requests for improvements
3. Provide feedback on code structure and Go best practices

## License

This project is open source and available under the [MIT License](LICENSE).

## Author

Built as a learning project to explore Go programming, CLI development with Cobra, and concurrent file operations.
