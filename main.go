package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	ip = flag.String("t", "127.0.0.1", "Target ip address to scan")
)

func scanner(ip string, port int, wg *sync.WaitGroup) {

	defer wg.Done()

	target := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.DialTimeout("tcp", target, 500*time.Millisecond)

	if err != nil {
		return
	}

	fmt.Println("Found Port --> ", port)
	conn.Close()

}

func main() {
	flag.Parse()

	var wg sync.WaitGroup

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
	fmt.Print("\033[35m Author: RE70-DECEMBER\nSpecial Thanks To Yassinox \033[0m\n\n")
	fmt.Printf("\033[34mScanning Target: %s\tThrough TCP Ports 1/65535 \033[0m\n\n", *ip)

	for currentPort := 1; currentPort <= 65535; currentPort++ {
		wg.Add(1)
		go scanner(*ip, currentPort, &wg)
	}

	wg.Wait()
	fmt.Println()
	fmt.Println("\033[34mScanning Done! \033[0m")

}
