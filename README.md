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
- [ ] **Day 3: Inspection**
  - [ ] Commands: `ls`, `cat`
  - [ ] Concepts: File descriptors, `os.ReadDir`, `os.ReadFile`
- [ ] **Day 4: Creation**
  - [ ] Commands: `mkdir`, `touch`
  - [ ] Concepts: File permissions (0755), `os.Create`
- [ ] **Day 5: Manipulation**
  - [ ] Commands: `mv`, `cp`
  - [ ] Concepts: IO Streaming (`io.Copy`), Buffer management
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
    gosh> cd ~
    gosh> pwd
    C:\Users\USER
    gosh> cd Desktop/my-gosh
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

### **3. File Streams**
*Upcoming: Notes on `io.Copy`, file descriptors, and using `defer` for resource cleanup.*

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

### Day 3: Inspection
* *Pending...*

</details>

---

## 📂 Project Structure

```text
.
├── main.go       # Entry point: REPL loop and command dispatch
├── README.md     # Documentation
└── ...           # (Future) Modularized command packages
```
