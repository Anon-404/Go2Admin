
# Go2Admin 🔎
**Web Admin Panel Finder**  
*Written in Go*

<div align="center">
  <img src="https://raw.githubusercontent.com/Anon-404/My-assets/main/Go2Admin/Go2Admin.jpg" alt="Go2Admin Logo" width="200"/>
</div>

## Features ✨

- 🔍 **Web Admin Panel Brute-Forcing:**  
  Discovers hidden directories, files, and admin panels on a web server using brute-forcing techniques.

- ⚡ **Efficient Scanning:**  
  Utilizes concurrent requests to speed up the scanning process, ensuring faster discovery of paths and files.

- 🕵️‍♂️ **Detailed Status Reporting:**  
  Provides feedback on server responses with color-coded status codes for easy identification.

- 🎯 **User-Friendly Interface:**  
  Offers a clear and concise command-line interface with helpful usage instructions and examples.

- 📖 **Comprehensive User Manual:**  
  Includes well-documented help and version information to assist users in navigating and utilizing the tool effectively.

## Installation 🛠️

#### Step 1: Install Golang
- **Arch-based Linux:**
  ```bash
  sudo pacman -S go
  ```
- **Debian-based Linux:**
  ```bash
  sudo apt install golang -y
  ```
- **Fedora:**
  ```bash
  sudo dnf install golang
  ```
- **Termux:**
  ```bash
  pkg install golang -y
  ```
- **OpenSUSE:**
  ```bash
  sudo zypper install go
  ```
- **Void Linux:**
  ```bash
  sudo xbps-install go
  ```

#### Step 2: Clone and Build
- **For Linux:**
  ```bash
  git clone https://github.com/Anon-404/Go2Admin 
  cd Go2Admin 
  go build -o Go2Admin 
  ```
- **For Termux:**
  ```bash
  git clone https://github.com/Anon-404/Go2Admin
  cd Go2Admin
  go build -o Go2Admin
  ```

## Usage 🧑‍💻

```bash
./Go2Admin -d <domain>
```

- **Using a custom wordlist:**
  ```bash
  ./Go2Admin -d <domain> wordlist.txt
  ```

### Main Options:

- **`-d`, `--domain`**  
  🛠 **Domain Input:**  
  Specify the target domain for scanning.

### Additional Options:

- **`-h`, `--help`**  
  📝 **Help:**  
  Displays this help page with descriptions of all available commands and options.

- **`-v`, `--version`**  
  🆚 **Version:**  
  Prints the current version number of Go2Admin.

### Examples:

- **Using the built-in wordlist:**
  ```bash
  ./Go2Admin -d <domain>
  ```
- **Using a custom wordlist:**
  ```bash
  ./Go2Admin -d <domain> wordlist.txt
  ```

For more information or assistance, use the `-h` option to display the help page.

## Contributions 🤝
Contributions are welcome! Feel free to open issues or submit pull requests.
