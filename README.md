# cchead – Head Clone

This is a simple Go implementation of the Unix `head` command, created as part of a coding challenge on [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-head).

## Features

-  Print a specified number of lines or bytes from the start of files.
-  Supports multiple files, displaying headers for each.

## Usage

### Options

-  `-n <number>`: Number of lines to print (default: 10).
-  `-c <number>`: Number of bytes to print.

### Examples

-  Print the first 10 lines of `file1.txt`:
  ```bash
  go run main.go -n 10 file1.txt
  ```

-  Print the first 100 bytes of `file1.txt`:
  ```bash
  go run main.go -c 100 file1.txt
  ```

-  Print the first 5 lines of `file1.txt` and `file2.txt`:
  ```bash
  go run main.go -n 5 file1.txt file2.txt
  ```

## Requirements

-  Go 1.16 or later

## License

MIT License.

