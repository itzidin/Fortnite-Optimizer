# Building from Source

This document provides detailed instructions for building the Fortnite Optimizer from source code.

## Prerequisites

1. Go 1.18 or later
2. Windows 10/11 (for running the compiled executable)
3. Git (optional, for cloning the repository)

## Build Instructions

### Windows

1. Install Go from [golang.org](https://golang.org)
2. Open Command Prompt or PowerShell
3. Navigate to the `src` directory
4. Run:
   ```
   go build -o fortnite_optimizer.exe
   ```

### Cross-compilation from Linux/Mac

1. Install Go from [golang.org](https://golang.org)
2. Open Terminal
3. Navigate to the `src` directory
4. Run:
   ```
   GOOS=windows GOARCH=amd64 go build -o fortnite_optimizer.exe
   ```

## Verification

After building, verify the executable:
1. The file should be named `fortnite_optimizer.exe`
2. Right-click and select "Run as Administrator"
3. Verify that the menu displays correctly
4. Test both optimization modes

## Troubleshooting

If you encounter build errors:
1. Ensure Go is properly installed: `go version`
2. Verify all dependencies are present: `go mod tidy`
3. Check Windows SDK is available (for Windows builds)
