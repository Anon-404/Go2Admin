package main

import (
    "bufio"
    "fmt"
    "net/http"
    "net"
    "os"
    "sync"
    "time"
    //"strings"

    "github.com/fatih/color"
)

var i = 0

// color blur
var blue = color.New(color.FgBlue)
var boldBlue = blue.Add(color.Bold)

// color red
var red = color.New(color.FgRed)
var boldRed = red.Add(color.Bold)

// color green
var Green = color.New(color.FgGreen)
var boldGreen = Green.Add(color.Bold)

// color yellow
var Yellow = color.New(color.FgYellow)
var boldYellow = Yellow.Add(color.Bold)

//color cyan
var cyan = color.New(color.FgCyan)
var boldCyan = cyan.Add(color.Bold)

//color magenta
var magenta = color.New(color.FgMagenta)
var boldMagenta = magenta.Add(color.Bold)

// color white
var white = color.New(color.FgWhite)
var boldWhite = white.Add(color.Bold)

var wg sync.WaitGroup
var mutex sync.Mutex

func banner() {
    banner := `
   ______     ___   ___       __          _     
  / ____/___ |__ \ /   | ____/ /___ ___  (_)___ 
 / / __/ __ \__/ // /| |/ __  / __ '__ \/ / __ \
/ /_/ / /_/ / __// ___ / /_/ / / / / / / / / / /
\____/\____/____/_/  |_\__,_/_/ /_/ /_/_/_/ /_/ 

 `
    boldCyan.Println(banner)
    boldGreen.Println("William Steven(Anon404)\n")
}

func version() {
	boldGreen.Println("version 1.0")
	return
}

func help() {
    boldBlue.Println("╔════════════════════════════════════════════════════╗")
	boldBlue.Println("║                      Go2Admin                      ║")
	boldBlue.Println("║       Web Admin panel Brute-Forcing  Utility       ║")
	boldBlue.Println("╚════════════════════════════════════════════════════╝")
    boldBlue.Println()

    boldBlue.Println("Usage:")                                                                                                                                                       
    boldWhite.Println("  Go2Admin -s <domain> <wordlist>")
    boldBlue.Println()

    boldBlue.Println("Options:")
    boldGreen.Print("  -h ")
    boldBlue.Print(" --help")
    boldWhite.Println(": Show this help manual")

    boldGreen.Print("  -v ")
    boldBlue.Print(" --version      ")
    boldWhite.Println(": Show version info")
    fmt.Println()

    boldBlue.Println("Examples:")
    boldWhite.Println("  Go2Admin example.com")
    boldWhite.Println("  Go2Admin example.com [wordlist]")
    boldBlue.Println()

    boldBlue.Println("╔════════════════════════════════════════════════════╗")
    boldBlue.Println("║            Thank you for using Go2Admin!           ║")
    boldBlue.Println("╚════════════════════════════════════════════════════╝")
}


func printProgressBar(current, total int) {
	// Calculate percentage
	progress := float64(current) / float64(total)
	percentage := int(progress * 100)

	// Save cursor position, print the progress, and restore cursor position
	fmt.Print("\0337") // Save cursor position
	fmt.Printf("\r[%d/%d] %d%%", current, total, percentage)
	fmt.Print("\0338") // Restore cursor position
}


func secret(wg *sync.WaitGroup, domain string, sword string){
	defer wg.Done()
    	dom := "https://" + domain + "/" + sword
    	conn, err := http.Get(dom)
    	time.Sleep(500 * time.Millisecond)

    	if err != nil {
        	//fmt.Println(err)
        	return
    	}

    	defer conn.Body.Close()

    	if conn.StatusCode >= 200 && conn.StatusCode <= 399 {
        	boldCyan.Println("[!] Found :", dom)
        }	
}

func dbgbrute(domain string, word string) {
    // defer wg.Done()
    dom := "https://" + domain + "/" + word
    conn, err := http.Get(dom)
    time.Sleep(500 * time.Millisecond)

    if err != nil {
        //fmt.Println(err)
        return
    }

    defer conn.Body.Close()

    if conn.StatusCode >= 100 && conn.StatusCode <= 199 {
        boldYellow.Printf("%-40s status : %d\n", dom, conn.StatusCode)
		i++
        // 200-299
    } else if conn.StatusCode >= 200 && conn.StatusCode <= 299 {
    	boldGreen.Printf("%-40s status : %d\n", dom, conn.StatusCode)
		i++
        // 300-399
    } else if conn.StatusCode >= 300 && conn.StatusCode <= 399 {
    	boldBlue.Printf("%-40s status : %d\n", dom, conn.StatusCode)

		i++
        // 400-499
    } else if conn.StatusCode >= 400 && conn.StatusCode <= 499 {
        boldRed.Printf("%-40s status : %d\n", dom, conn.StatusCode)

        // 500-599
    } else if conn.StatusCode >= 500 && conn.StatusCode <= 599 {
        boldMagenta.Printf("%-40s status : %d\n", dom, conn.StatusCode)
        i++
    }
}

func main() {
    banner()

    if len(os.Args) < 2 {
        boldBlue.Println("Usage: Go2Admin -s [domain] ")
        boldBlue.Println("Manual: Go2Admin -h")
        boldBlue.Println("Version: Go2Admin -v")
        return
    }

    option := os.Args[1]
    if option == "-h" || option == "--help" {
        help()
        return
    } else if option == "-v" || option == "--version" {
        version()
        return
    }

    if len(os.Args) < 3 {
        boldRed.Println("Error: Domain missing")
        boldBlue.Println("Usage: Go2Admin -s <domain>")
        boldBlue.Println("Manual: Go2Admin -h")
        boldBlue.Println("Version: Go2Admin -v")
        return
    }

    domain := os.Args[2]

    _, err := net.LookupIP(domain)
    if err != nil {
        boldRed.Println("[-] Error:", err)
        return
    }
    

    if len(os.Args) < 4 {
        boldRed.Println("Error: Wordlist missing")
        boldBlue.Println("Usage: GoFuzzer <option> <domain/ip> <wordlist>")
        boldBlue.Println("Manual: GoFuzzer -h")
        boldBlue.Println("Version: GoFuzzer -v")
        return
    }
    
    filename := os.Args[3]

    file, err := os.Open(filename)
    if err != nil {
        boldRed.Println("[-] Error opening file:", err)
        return
    }
    defer file.Close()
    

    var wd []string
    
    scret := []string{
		"robots.txt",      // Often lists directories or files meant to be hidden from search engines
		".htaccess",       // Apache configuration file
		".env",            // Environment variables file
		"admin",           // Admin interface directory
		"config.php",      // PHP configuration file
		"db-config.php",   // Database configuration file
		"backup",          // Backup files directory
		"includes",        // Includes directory for code
		"uploads",         // File uploads directory
		"temp",            // Temporary files directory
		".git",            // Git directory, exposing repository files
		".svn",            // SVN (Subversion) directory
		"install",         // Installation directory
		"logs",            // Log files directory
		"scripts",         // Scripts directory
		"cache",           // Cache files directory
		"private",         // Private or restricted directory
		"cgi-bin",         // CGI scripts directory
		"phpinfo.php",     // PHP information page
		"status",          // Status page or file
		"readme.txt",      // README file
		"error.log",       // Error log file
		"debug.log",       // Debugging log file
		"webdav",          // WebDAV directory
		"old",             // Old or legacy files directory
		"dump.sql",        // SQL database dump file
		"backup.sql",      // Backup of SQL database
		"user/.ssh",       // SSH configuration directory
		"settings.php",    // Settings file for applications
		"metadata",        // Metadata directory
		"sensitive-data",  // Sensitive data directory
		"config.json",     // JSON configuration file
		"database.yml",    // YAML database configuration file
		"appsettings.json",// Application settings file
		"web.config",      // IIS Web configuration file
		".bashrc",         // Bash configuration file
		".bash_profile",   // Bash profile file
		"site_backup.zip", // Compressed site backup
		".htpasswd",       // Password file for basic HTTP authentication
		"index.php.bak",   // Backup of index.php
		"wp-config.php",   // WordPress configuration file
		"local.xml",       // Magento configuration file
		"main.js.bak",     // Backup of main JavaScript file
		".DS_Store",       // macOS directory attribute file
		"composer.json",   // Composer dependency manager file
		"package-lock.json", // NPM package lock file
		"yarn.lock",       // Yarn lock file
		"swagger.json",    // Swagger API definition file
		"backup.tar.gz",   // Tarball backup file
		"license.txt",     // License file
		"CONTRIBUTING.md", // Contribution guidelines file
		"keys.json",       // API keys or credentials file
		"ftpconfig.json",  // FTP configuration file
		"error_log",
		"sitemap.xml",
	}
	
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wd = append(wd, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        boldRed.Println("[-] Error reading file:", err)
        return
    }

    boldBlue.Println("[+] Target:", domain)
    boldBlue.Println("[+] Wordlist:", filename)
    boldBlue.Println("[+] Words:", len(wd))
    boldGreen.Println("[+] Status: Demo version\n")
    
    boldYellow.Println("[+] Looking for secret path/file\n")
    
    for _, sword := range scret {
        wg.Add(1)
        go secret(&wg, domain, sword)
    }
    wg.Wait()

	boldYellow.Println("\n[+] Starting main process\n")

    
    
    var completed int
    total := len(wd)
    
    
    wg := sync.WaitGroup{}
    semaphore := make(chan struct{}, 150)

    for num, word := range wd {
        wg.Add(1)
        go func(num int, word string) {
            defer wg.Done()
            semaphore <- struct{}{}
            defer func() { <-semaphore }()
            
            
            switch option {
            case "-d", "--debug":
                dbgbrute(domain, word)
            default:
                boldRed.Println("[-] Unknown option", option)
                return
            }
            
            mutex.Lock()
            completed++
            printProgressBar(completed, total)
            mutex.Unlock()
            
        }(num, word)
        
    }
    
    wg.Wait()
    fmt.Println()
    boldBlue.Println("[+] total found :",i)
    
}
