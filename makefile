# Define the Go source file
GOFILE = filter.go

# Define the executable name
EXECUTABLE = filter

# Default target to build the Go script
all: build

# Build the Go script
build:
	go build -o $(EXECUTABLE) $(GOFILE)

# Run the Go script with a benchmark flag
run: build
	./$(EXECUTABLE) -t=$(TIME)

# Clean up the executable
clean:
	rm -f $(EXECUTABLE)

# Phony targets to prevent conflicts with files
.PHONY: all build run clean
