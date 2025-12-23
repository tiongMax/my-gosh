# GoShell 🐚

![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat&logo=go)
![Status](https://img.shields.io/badge/Status-In%20Progress-yellow)
![License](https://img.shields.io/badge/License-MIT-green)

Welcome to **GoShell**, a lightweight, custom implementation of a Linux shell written in Go. Designed to explore system programming concepts, this shell re-implements core Linux utilities from scratch, offering a hands-on way to understand how shells interact with the operating system.

## Key Features

- **Interactive REPL**: A fully functional Read-Eval-Print Loop for command execution.
- **Navigation Tools**: Built-in commands like `pwd` and `cd` for seamless file system navigation.
- **File Inspection**: Utilities like `ls` and `cat` to view directory contents and file data.
- **File Management**: Create and manage files and directories with `mkdir`, `touch`, `cp`, and `mv`.
- **Safe Destruction**: Securely remove files and directories with `rm` and `rmdir` (including recursive delete).
- **Search & History**: Built-in `grep` for text searching and session `history` tracking.
- **Cross-Platform**: Works efficiently on Windows, Linux, and macOS.

## Core Technologies

GoShell is built using the Go Standard Library, focusing on low-level system interactions without third-party dependencies:

- **Language**: [Go (Golang)](https://go.dev/) - Chosen for its strong system programming capabilities.
- **Standard Library Packages**:
  - `os`: For direct system calls, file manipulation, and process management.
  - `bufio`: For efficient buffered I/O and input scanning.
  - `strings`: For powerful text processing and command parsing.
  - `io`: For streaming data between files.

## Setup Instructions

### 1. Install Go

Download and install Go (version 1.20 or higher) from [here](https://go.dev/dl/).

### 2. Clone the Repository

Clone the project to your local machine:

```bash
git clone https://github.com/tiongMax/my-gosh.git
cd my-gosh
```

### 3. Run the Shell

You can run the shell directly using the Go toolchain:

```bash
go run main.go
```

### 4. Example Usage

Once the shell is running, you can execute commands just like in a standard terminal:

```bash
gosh> echo hello world
hello world

gosh> pwd
C:\Users\USER\Desktop\my-gosh

gosh> ls
main.go
README.md

gosh> mkdir testdir
gosh> touch newfile.txt
gosh> ls
main.go
README.md
testdir/
newfile.txt

gosh> exit
Goodbye!
```

---

For detailed development progress, technical learnings, and daily logs, please refer to [DEVLOG.md](./DEVLOG.md).
