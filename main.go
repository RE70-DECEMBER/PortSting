package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
	"os/user"
	"strings"
)

var (
	ip = flag.String("t", "127.0.0.1", "Target ip address to scan")
	portCounter = 0
)

func scanner(ip string, port int, wg *sync.WaitGroup) {

	defer wg.Done()

	target := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.DialTimeout("tcp", target, 500*time.Millisecond)

	if err != nil {
		return
	}

	fmt.Println("Found Port --> ", port)
	portCounter++
	conn.Close()

}

func main() {
	flag.Parse()

	var wg sync.WaitGroup

	currentUser, err := user.Current()	

	username := currentUser.Username

	if err != nil {
		
		username = "!!!MEOW MEOW Could Not Find Your Username! MEOW MEOW!!!"

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
	fmt.Printf("\033[35m Author: RE70-DECEMBER\nSpecial Thanks To Yassinox\nWelcome %v Glad Your Here :)\033[0m\n\n", strings.Title(username))
	fmt.Printf("\033[34mScanning Target: %s\tThrough TCP Ports 1/65535 \033[0m\n\n", *ip)

	for currentPort := 1; currentPort <= 65535; currentPort++ {
		wg.Add(1)
		go scanner(*ip, currentPort, &wg)
	}

	wg.Wait()
	fmt.Println()

	

	if portCounter <= 0 {
		fmt.Println("\033[32mScanning Done! \033[0m -->\t\033[31m!!!MEOW MEOW\033[0m \033[32m--> No Found Ports Sorry :( <-- \033[31mMEOW MEOW!!!\033[0m\n\033")
	}else {
		fmt.Printf("\033[34mScanning Done! \033[0m \tFound Ports --> %v\n", portCounter)
	}
}
