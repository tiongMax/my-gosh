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
- [ ] **Day 2: Navigation**
  - [ ] Commands: `pwd`, `cd`
  - [ ] Concepts: Process state, `os.Getwd`, `os.Chdir`
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
    goshell> echo hello world
    hello world
    goshell> exit
    Goodbye!
    ```

---

## 🧠 Technical Learnings

*Documenting the specific Go concepts and System Programming challenges mastered.*

### **1. Input Handling (`bufio` vs `fmt`)**
I learned that standard `fmt.Scanln` is insufficient for a shell because it stops reading at the first whitespace. To support commands with arguments (like `echo hello world`), I used `bufio.NewScanner(os.Stdin)`, which captures the entire input stream up to the newline character.

### **2. System Interaction**
*Upcoming: Notes on how `os.Getwd` interacts with the kernel and why `cd` must be a shell builtin.*

### **3. File Streams**
*Upcoming: Notes on `io.Copy`, file descriptors, and using `defer` for resource cleanup.*

---

## 📝 Dev Log

<details>
<summary><strong>Click to expand Daily Logs</strong></summary>

### Day 1: The Skeleton
* **Progress:** Built the main infinite `for` loop and the command parser.
* **Key Learning:** `strings.Fields()` is much better than `strings.Split()` for CLI parsing because it automatically ignores multiple spaces between arguments.
* **Challenge:** Had to ensure the program exits gracefully without panic when the user inputs an empty string.

### Day 2: Navigation
* *Pending...*

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
