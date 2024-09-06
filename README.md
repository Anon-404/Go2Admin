# Go2Admin 🔎
**Web Admin panel finder**  
*Written in Go*

<div align="center">
  <img src="https://raw.githubusercontent.com/Anon-404/My-assets/main/Go2Admin/Go2Admin.jpg" alt="GoFuzzer Logo" width="200"/>
</div>

## Features ✨

- 🔍 **Web admin-panl Brute-Forcing**  
  Performs targeted brute-forcing to discover hidden directories, files and admin-panel on a web server.

- ⚡ **Efficient Scanning:**  
  Utilizes concurrent requests to speed up the scanning process, allowing for faster discovery of paths and files.

- 🕵️‍♂️ **Detailed Status Reporting:**  
  Provides comprehensive feedback on server responses with color-coded status codes for easy identification.

- 🎯 **User-Friendly Interface:**  
  Includes a clear and concise command-line interface with helpful usage instructions and examples.

- 📖 **Comprehensive User Manual:**  
  Well-documented help and version information to assist users in navigating and utilizing the tool effectively.

## Installation 🛠️


#### Step 1: Install Go
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
  git clone https://github.com/Anon-404/GoFuzzer 
  cd GoFuzzer 
  go build -o Go2Admin 
  ```
- **For Termux:**
  ```bash
  git clone https://github.com/Anon-404/GoFuzzer 
  cd GoFuzzer 
  go build -o Go2Admin 
  ```

## Usage 🧑‍💻

```bash
GoFuzzer <option> <domain> <wordlist>
```

### Main Options:

- **`-d`, `--domain`**  
  🛠 **Debug Mode**  
  Enables debugging mode to show detailed responses for all status codes from 100 to 599.

### Additional Options:

- **`-h`, `--help`**  
  📝 **Help**  
  Displays this help page with descriptions of all available commands and options.

- **`-v`, `--version`**  
  🆚 **Version**  
  Prints the current version number of GoFuzzer.

### Examples:

- **Enable debugging mode:**
  ```bash
  GoFuzzer -d example.com wordlist.txt
  ```

For more information or assistance, use the `-h` option to display help page.


## Contributions 🤝
Contributions are welcome! Feel free to open issues or submit pull requests.
