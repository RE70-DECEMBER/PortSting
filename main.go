package main

import (
	"flag"
	"fmt"
	"net"
	"os/user"
	"strings"
	"sync"
	"time"
)

var (
	ip            = flag.String("t", "127.0.0.1", "Target ip address to scan")
	portCounter   = 0
	bannerCounter = 0
	mu            sync.Mutex // protects portCounter and bannerCounter
)

func scanner(ip string, port int) {
	target := fmt.Sprintf("%s:%d", ip, port)

	// Attempt TCP connection
	conn, err := net.DialTimeout("tcp", target, 2*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	// Send The HTTP Request
	if port == 80 || port == 8080 {
		conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	}

	// Grab Banner
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	banner := "\033[31m Unknown \033[0m"
	if err == nil && n > 0 {
		response := string(buf[:n])
		if port == 80 || port == 8080 {
			lines := strings.Split(response, "\r\n")
			for _, line := range lines {
				if strings.HasPrefix(strings.ToLower(line), "server:") {
					banner = strings.TrimSpace(line[len("server:"):])
					break
				}
			}
		} else {
			banner = strings.TrimSpace(response)
		}
		mu.Lock()
		bannerCounter++
		mu.Unlock()
	}

	mu.Lock()
	fmt.Printf("Found Port --> %v \tService Info: %v \n", port, banner)
	portCounter++
	mu.Unlock()
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	concurrencyLimit := 100 // limit concurrent scans
	sem := make(chan struct{}, concurrencyLimit)

	currentUser, err := user.Current() // detect username
	username := "!!!MEOW MEOW Could Not Find Your Username! MEOW MEOW!!!"
	if err == nil {
		username = currentUser.Username
	}

	fmt.Print("\033[34m")
	fmt.Print(`
  _____           _    _____ _   _             
 |  __ \         | |  / ____| | (_)            
 | |__) |__  _ __| |_| (___ | |_ _ _ __   __ _ 
 |  ___/ _ \| '__| __|\___ \| __| | '_ \ / _  |
 | |  | (_) | |  | |_ ____) | |_| | | | | (_| |
 |_|   \___/|_|   \__|_____/ \__|_|_| |_|\__, |
                                          __/ |
                                         |___/ 
`)
	fmt.Print("\033[0m")
	fmt.Printf("\033[35m Author: RE70-DECEMBER\nSpecial Thanks To Yassinox AND NoFace\nWelcome %v Glad Your Here :)\033[0m\n\n", strings.Title(username))
	fmt.Printf("\033[34mScanning Target: %s\tThrough TCP Ports 1/65535 \033[0m\n\n", *ip)

	for currentPort := 1; currentPort <= 65535; currentPort++ {
		wg.Add(1)
		sem <- struct{}{} // acquire slot
		go func(port int) {
			defer wg.Done()
			scanner(*ip, port)
			<-sem // release slot
		}(currentPort)
	}

	wg.Wait()
	fmt.Println()

	if portCounter <= 0 {
		fmt.Println("\033[32mScanning Done! \033[0m -->\t\033[31m!!!MEOW MEOW\033[0m \033[32m--> No Found Ports Sorry :( <-- \033[31mMEOW MEOW!!!\033[0m\n\033")
	} else {
		fmt.Printf("\033[34mScanning Done! \033[0m \tFound Ports --> %v \tFound Banners --> %v\n", portCounter, bannerCounter)
	}
}
