# GoShell 🐚

![Go Version](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=flat&logo=go)
![Status](https://img.shields.io/badge/Status-In%20Progress-yellow)
![License](https://img.shields.io/badge/License-MIT-green)

**GoShell** is a lightweight, custom implementation of a Linux shell written in Go.

This project was built to explore **System Programming** concepts and master the **Go Standard Library** (`os`, `bufio`, `strings`). Instead of acting as a wrapper for existing system calls, GoShell re-implements the logic of core Linux utilities (like `ls`, `cp`, `grep`) from scratch.

---

## 📅 Project Roadmap (7-Day Sprint)

I am building this project over one week, adding new commands and complexity daily.

- [x] **Day 1: The Engine**
  - [x] REPL (Read-Eval-Print Loop)
  - [x] Input parsing using `bufio`
  - [x] Commands: `exit`, `echo`
- [x] **Day 2: Navigation**
  - [x] Commands: `pwd`, `cd`
  - [x] Concepts: Process state, `os.Getwd`, `os.Chdir`
- [x] **Day 3: Inspection**
  - [x] Commands: `ls`, `cat`
  - [x] Concepts: File descriptors, `os.ReadDir`, `os.ReadFile`
- [x] **Day 4: Creation**
  - [x] Commands: `mkdir`, `touch`
  - [x] Concepts: File permissions (0755), `os.Create`, Resource management
- [x] **Day 5: Manipulation**
  - [x] Commands: `mv`, `cp`
  - [x] Concepts: IO Streaming (`io.Copy`), Buffer management
- [ ] **Day 6: Destruction**
  - [ ] Commands: `rm`, `rmdir`
  - [ ] Concepts: Safety checks, recursive deletion
- [ ] **Day 7: Search & History**
  - [ ] Commands: `grep`, `history`
  - [ ] Concepts: String processing, slice storage

---

## 🛠️ Installation & Usage

### Prerequisites
* [Go 1.20+](https://go.dev/dl/) installed on your machine.

### Quick Start
1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/yourusername/goshell.git](https://github.com/yourusername/goshell.git)
    cd goshell
    ```

2.  **Run the shell:**
    ```bash
    go run main.go
    ```

3.  **Example Usage:**
    ```bash
    gosh> echo hello world
    hello world
    gosh> pwd
    C:\Users\USER\Desktop\my-gosh
    gosh> ls
    main.go
    README.md
    .gitignore
    gosh> mkdir testdir
    gosh> touch newfile.txt
    gosh> ls
    main.go
    README.md
    testdir/
    newfile.txt
    gosh> cat newfile.txt
    (empty)
    gosh> cd ~
    gosh> pwd
    C:\Users\USER
    gosh> exit
    Goodbye!
    ```

---

## 🧠 Technical Learnings

*Documenting the specific Go concepts and System Programming challenges mastered.*

### **1. Input Handling (`bufio` vs `fmt`)**
I learned that standard `fmt.Scanln` is insufficient for a shell because it stops reading at the first whitespace. To support commands with arguments (like `echo hello world`), I used `bufio.NewScanner(os.Stdin)`, which captures the entire input stream up to the newline character.

**Key Implementation Details:**
* Created scanner once before the loop: `scanner := bufio.NewScanner(os.Stdin)`
* Used `scanner.Scan()` return value to detect EOF and exit gracefully
* Combined with `strings.Fields()` to intelligently parse input while handling edge cases (empty input, multiple spaces, tabs)

### **2. System Interaction & Process State**
I discovered that `cd` **must** be implemented as a shell builtin (not an external command) because it modifies the process's working directory. If `cd` were external, it would change the child process's directory, not the shell's.

**Key Implementation Details:**
* `os.Getwd()` retrieves the current working directory from the kernel
* `os.Chdir()` changes the process's working directory via system call
* Added home directory support: running `cd` with no arguments takes you home
* Implemented tilde expansion: `cd ~` and `cd ~/path` work correctly
* Used `os.UserHomeDir()` to portably get the user's home across Windows/Linux/macOS

### **3. File System Operations & Directory Inspection**
I learned how Go provides high-level abstractions for file system operations that hide the complexity of file descriptors and system calls.

**Key Implementation Details:**
* `os.ReadDir()` returns a slice of `os.DirEntry` (lightweight directory entries)
* Each `DirEntry` provides `.Name()` and `.IsDir()` methods without needing to stat the file
* `os.ReadFile()` handles the entire file lifecycle: open → read → close in one call
* Implemented visual distinction: directories get `/` suffix using `.IsDir()` check
* Both commands support multiple arguments (e.g., `ls dir1 dir2`, `cat file1.txt file2.txt`)

### **4. File Creation & Resource Management**
I learned that Go requires explicit resource management for file handles, unlike memory which is garbage collected automatically. This was my first encounter with the critical concept of **resource cleanup**.

**Key Implementation Details:**
* `os.Mkdir(path, 0755)` creates directories with Unix permissions (rwxr-xr-x)
* Permission `0755` means: owner can read/write/execute, group and others can read/execute
* `os.Create()` returns a file handle (`*os.File`) that must be closed to prevent resource leaks
* **Critical:** Always call `file.Close()` after creating files, or the OS will run out of file descriptors
* `os.Stat()` checks if a file/directory exists without opening it
* `os.IsNotExist(err)` distinguishes "file not found" from other errors (permission denied, etc.)
* `os.Chtimes()` updates file access and modification timestamps
* Advanced `touch` implementation: creates new files OR updates timestamps on existing files (matches Unix behavior)

### **5. File Streams & Manipulation**
I implemented `cp` and `mv`, learning how to efficiently stream data between files using buffers instead of loading everything into memory.

**Key Implementation Details:**
*   `io.Copy(dst, src)` handles the heavy lifting of transferring bytes
*   **Critical Bug Fix:** Learned NOT to `defer file.Close()` inside a loop (it only runs at function exit). Instead, I had to manually close files to prevent resource leaks.
*   **Atomic Operations:** `mv` uses `os.Rename()` which is an atomic system call (instant), whereas `cp` must physically copy data byte-by-byte.

---

## 📝 Dev Log

<details>
<summary><strong>Click to expand Daily Logs</strong></summary>

### Day 1: The Engine ✅
* **Progress:** Built the complete REPL (Read-Eval-Print Loop) with robust input parsing and command execution.
* **Commands Implemented:** `exit`, `echo`
* **Key Learning:** `strings.Fields()` is much better than `strings.Split()` for CLI parsing because it automatically ignores multiple spaces between arguments and handles tabs/newlines.
* **Challenges Solved:**
  * Gracefully handling EOF (Ctrl+D/Ctrl+Z) using `scanner.Scan()` return value
  * Preventing panic on empty input by checking `len(parts) == 0` after `strings.Fields()`
  * Handling whitespace-only input (spaces/tabs) that would create empty slices
  * Understanding `break` (exits loop/switch) vs `return` (exits function) for the `exit` command
* **Technical Insight:** The scanner must be created **once** before the loop (not inside it) for efficiency, as it maintains an internal buffer that would be lost if recreated each iteration.

### Day 2: Navigation ✅
* **Progress:** Implemented directory navigation commands with robust path handling.
* **Commands Implemented:** `pwd`, `cd`
* **Key Learning:** The `cd` command **must** be a shell builtin because it changes the shell process's own working directory. If it were an external program, it would only change the child process's directory and have no effect on the parent shell.
* **Challenges Solved:**
  * Handling `cd` with no arguments (should go to home directory)
  * Implementing tilde expansion (`~` and `~/path`) using `strings.HasPrefix()` and `strings.Replace()`
  * Cross-platform home directory detection using `os.UserHomeDir()`
  * Proper error handling when directory doesn't exist or permission denied
* **Technical Insight:** Windows uses backslashes (`\`) while Unix uses forward slashes (`/`) for paths, but Go's `os` package handles both transparently. The shell correctly displays Windows-style paths on Windows and Unix-style paths on Unix systems.

### Day 3: Inspection ✅
* **Progress:** Implemented file system inspection commands with support for multiple targets.
* **Commands Implemented:** `ls`, `cat`
* **Key Learning:** `os.ReadDir()` returns `DirEntry` objects (not `FileInfo`), which are lightweight and only contain basic metadata. 
* **Challenges Solved:**
  * Handling multiple directories/files in one command (e.g., `ls . .. ~/Desktop`)
* **Technical Insight:** `os.ReadFile()` is a convenience function that opens, reads entirely into memory, and closes the file automatically. 

### Day 4: Creation ✅
* **Progress:** Implemented file and directory creation commands with proper resource management.
* **Commands Implemented:** `mkdir`, `touch`
* **Key Learning:** Go does **not** automatically garbage collect file handles! Unlike memory management, you must explicitly close files or you'll leak resources and eventually crash when the OS runs out of file descriptors.
* **Challenges Solved:**
  * Understanding Unix file permissions in octal notation (`0755` = rwxr-xr-x)
  * Distinguishing between "file doesn't exist" vs "permission denied" errors using `os.IsNotExist()`
  * Implementing advanced `touch` behavior: updating timestamps on existing files vs creating new ones
  * Properly closing file handles immediately after creation (can't use `defer` in a loop)
* **Technical Insight:** Windows largely ignores Unix permission bits (0755), but Go accepts them for cross-platform compatibility. The real insight was learning that `os.Create()` returns a file handle that holds system resources - forgetting to close it is like opening a connection and never releasing it. This was my first real encounter with **manual resource management** in Go.

### Day 5: Manipulation ✅
* **Progress:** Added `cp` (copy) and `mv` (move) commands, introducing input/output streaming.
* **Commands Implemented:** `cp`, `mv`
* **Key Learning:** `io.Copy` is a powerful abstraction that streams data between any `Reader` and `Writer`. It uses a small buffer (usually 32KB) so you can copy a 10GB file without using 10GB of RAM.
* **Challenges Solved:**
  * **Panic Fix:** I initially tried to close files even when `os.Open` returned an error (which meant the file pointer was `nil`). This caused the shell to crash. I learned to only close resources that were successfully acquired.
  * **Defer Trap:** I reinforced the lesson that `defer` is dangerous inside a long-running REPL loop because cleanup happens too late.
  * **Atomic Moves:** Implemented `mv` using `os.Rename`, which is much faster than copying, but learned it has limitations (cannot move across different disk partitions).
* **Technical Insight:** The difference between "moving" data (changing a pointer in the filesystem table) and "copying" data (duplicating actual bytes) is fundamental to understanding OS performance.

</details>

---

## 📂 Project Structure

```text
.
├── main.go       # Entry point: REPL loop and command dispatch
├── README.md     # Documentation
└── ...           # (Future) Modularized command packages
```
