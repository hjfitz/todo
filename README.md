# TODO

`todo` is a Go-based command-line tool that scans your project for all TODO comments and generates a neat TODO list. It automatically ignores directories like `node_modules` and `.git`.

## Features

- Scans your project for TODO comments.
- Outputs a clean TODO list, formatted for easy reading.
- Automatically ignores common directories like `node_modules` and `.git`.

## Installation

To install `todo`, you need to have Go installed on your machine. Then, run the following command:

```bash
go install github.com/hjfitz/todo@latest
```

## Usage

Simply run the following command in the root of your project directory:

```bash
todo
```

### Example Output

```plaintext
file1.go:
  12: Fix the error handling for invalid input
  45: Refactor this function to reduce complexity

file2.js:
  27: Add unit tests for this module
  89: Remove deprecated code in the next release
```

## Configuration

Currently, `todo` is not configurable. Future updates may include options for custom ignore patterns and output formats.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request if you have suggestions or improvements.

## License

This project is licensed under the [MIT License](LICENSE).
