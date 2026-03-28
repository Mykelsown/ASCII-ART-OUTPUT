# ASCII-Art Output

A command-line tool written in Go that renders text as ASCII art using banner styles, with support for writing the output directly to a file via the `--output` flag.

---

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
- [Installation](#installation)
- [Usage](#usage)
  - [Print to Terminal](#print-to-terminal)
  - [Write to File](#write-to-file)
- [Banner Styles](#banner-styles)
- [Flag Format Rules](#flag-format-rules)
- [Error Handling](#error-handling)
- [Examples](#examples)
- [Testing](#testing)
- [Allowed Packages](#allowed-packages)

---

## Overview

ASCII-Art Output extends the base ASCII-art project by adding file output support. It reads a banner `.txt` file, maps each printable ASCII character (32–126) to its corresponding 8-line art block, and either prints the result to `stdout` or writes it to a named file using the `--output` flag.

---

## Project Structure

```
.
├── main.go                  # Entry point — flag parsing, routing, file writing
├── MethodsAndTesting/
│   ├── file-reader.go       # Reads the correct banner .txt file from disk
│   └── printer.go           # Renders input string into ASCII art
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── README.md
```

---

## How It Works

### `FileReader(styleName string) ([]byte, bool)`
- Accepts a banner style name (`standard`, `shadow`, or `thinkertoy`).
- Defaults to `standard` if the provided style is unrecognised.
- Reads the corresponding file from the `banners/` directory.
- Returns the raw file bytes and a boolean indicating success.

### `FormatPrinter(input, contentRead string, readingStatus bool) string`
- Receives the user's input string and the banner file content as a string.
- Splits the input on the literal `\n` sequence to handle multi-line art.
- For each character, calculates its position in the banner using:
  ```
  charIndex = char - 32
  lineIndex = charIndex * 9 + 1 + row
  ```
  Each character occupies 9 lines in the banner file (8 art lines + 1 blank separator).
- Builds and returns the full rendered output string.

### `main()`
- Parses the `--output=<fileName>` flag using Go's `flag` package.
- Separates the input string from the banner style (last argument).
- If `--output` is set, writes the rendered art to the named file.
- If no flag is set, prints the rendered art to `stdout`.

---

## Installation

**Prerequisites:** Go 1.18 or later.

```bash
git clone https://github.com/your-username/ascii-art-output.git
cd ascii-art-output
```

No external dependencies — only the Go standard library is used.

---

## Usage

### Print to Terminal

```bash
go run . "string"
go run . "string" [BANNER]
```

**Examples:**
```bash
go run . "hello"
go run . "hello" shadow
go run . "Hello There!" thinkertoy
```

### Write to File

```bash
go run . --output=<fileName.txt> "string" [BANNER]
```

**Examples:**
```bash
go run . --output=banner.txt "hello" standard
go run . --output=result.txt "Hello There!" shadow
```

The output file will contain the rendered ASCII art followed by a trailing newline, matching the format visible via `cat -e`.

---

## Banner Styles

| Style        | Description                          |
|--------------|--------------------------------------|
| `standard`   | Classic block letters (default)      |
| `shadow`     | Shadow-style letters with depth      |
| `thinkertoy` | Lightweight, symbolic characters     |

If no banner style is specified, `standard` is used automatically.

---

## Flag Format Rules

The `--output` flag must follow this **exact** format:

```
--output=<fileName.txt>
```

Any other format will print the usage message and exit:

```
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard
```

**Invalid formats that will be rejected:**

| Invalid Input           | Reason                        |
|-------------------------|-------------------------------|
| `-output=file.txt`      | Single dash instead of double |
| `output=file.txt`       | Missing dashes entirely       |
| `.output=file.txt`      | Dot prefix instead of dashes  |
| `--output file.txt`     | Space instead of `=`          |

---

## Error Handling

| Scenario                          | Behaviour                                              |
|-----------------------------------|--------------------------------------------------------|
| No arguments provided             | Prints usage message and exits                         |
| Only flag given, no input string  | Prints usage message and exits                         |
| Invalid flag format               | Prints usage message and exits                         |
| Banner file not found             | Prints `Error` and returns empty result                |
| File write failure                | Logs fatal error with `log.Fatalf`                     |
| Non-printable ASCII characters    | Characters outside range 32–126 are silently skipped   |

---

## Examples

### Terminal output — standard banner
```bash
go run . "hello" standard
```
```
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \ 
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/ 
                               
                               
```

### File output — shadow banner
```bash
go run . --output=banner.txt "Hello There!" shadow
cat -e banner.txt
```
```
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _|$
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _|$
...
```

### Multi-line input
```bash
go run . "Hello\nWorld"
```
Renders `Hello` and `World` on separate ASCII-art lines.

### Single string, no banner specified
```bash
go run . "Hi"
```
Defaults to `standard` banner automatically.

---

## Testing

Unit tests are recommended for both `FileReader` and `FormatPrinter`. To run any test files in the package:

```bash
cd MethodsAndTesting
go test ./...
```

Suggested test cases:
- Single character input
- Multi-line input using `\n`
- All three banner styles
- Non-printable characters (should be skipped)
- Invalid/missing banner name (should default to `standard`)
- Empty string input

---

## Allowed Packages

Only Go standard library packages are used:

- `os` — file reading and writing
- `flag` — CLI flag parsing
- `fmt` — formatted output
- `log` — fatal error logging
- `strings` — string splitting and building

---

## Learning Outcomes

This project covers:

- Go's file system (`fs`) API — reading and writing files with `os.ReadFile` and `os.WriteFile`
- Data manipulation — mapping characters to banner positions and building output strings
- CLI flag parsing with Go's `flag` package
- Clean separation of concerns across packages