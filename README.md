# Courier CLI

## Overview

`courier` is a simple command-line tool that takes input from stdin and sends it as a JSON payload to a specified API endpoint via a POST request. It's designed to be lightweight and easily integrated into shell scripts or command pipelines.

## Features

- Read input from stdin
- Send input as JSON payload
- Configurable API endpoint via `--url` flag
- Error handling and reporting
- Supports any API endpoint that accepts JSON

## Installation

### Prerequisites

- Go (golang) 1.16 or later

### Build from Source

1. Clone the repository
2. Navigate to the project directory
3. Build the binary:

```bash
go build -o courier main.go
```

4. (Optional) Move to a directory in your PATH:

```bash
sudo mv courier /usr/local/bin/
```

## Usage

### Basic Usage

```bash
echo "buy milk" | courier --url http://todo.app/todo
```

### Examples

1. Create a todo item:
```bash
echo "Clean the garage" | courier --url https://mytodoapp.com/todos
```

2. Send a message to a chat API:
```bash
echo "Hello, world!" | courier --url https://chat.example.com/messages
```

3. Log an event:
```bash
echo "User logged in" | courier --url https://logging.myservice.com/events
```

## JSON Payload Structure

The input is transformed into a JSON payload with a single `data` key:

```json
{
  "data": "Your input goes here"
}
```

## Error Handling

The tool provides detailed error messages:
- Missing `--url` flag
- Input reading errors
- JSON marshaling errors
- HTTP request errors
- Non-successful HTTP response codes

Errors are printed to stderr, making the tool friendly for scripting and logging.

## Environment and Compatibility

- Tested on Linux, macOS
- Requires Go 1.16+
- No external dependencies beyond standard library

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
