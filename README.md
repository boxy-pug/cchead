# cchead

cchead is a simple Go implementation of the Unix `head` command, created as part of a coding challenge on [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-head).

## Features

-  Print a specified number of lines or bytes from the start of files.
-  Supports multiple files, displaying headers for each.

## Usage

### Options

-  `-n <number>`: Number of lines to print (default: 10).
-  `-c <number>`: Number of bytes to print.

### Examples

-  Print the first 10 lines of `./testdata/test.txt`:
  ```bash
  ./cchead -n 10 ./testdata/test.txt
  ```

-  Print the first 100 bytes of `./testdata/test.txt`:
  ```bash
  ./cchead -c 100 ./testdata/test.txt
  ```

-  Print the first 5 lines of `./testdata/test.txt` and `./testdata/test2.txt`:
  ```bash
  ./cchead -n 5 ./testdata/test.txt ./testdata/test2.txt
  ```

## Requirements

-  Go 1.16 or later

## License

MIT License.

