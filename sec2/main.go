package main

import (
	"bufio"
	"fmt" // formating
	"os"
	"os/signal" // work with signals
	"strings"   // remove blanc space
	"syscall"   // const system calls such like signals
)

func main() {

	sig := make(chan os.Signal, 1) // make one chan capacity 1

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM) // registring a signal for term

	go prompter()

	// loop wait a signal to termiante

	for {
		select {
		case res := <-sig:
			signal.Stop(sig)
			fmt.Printf("%s signal recieved\n", res)
			os.Exit(0)
		}
	}

}

func prompter() {
	fmt.Print(">> ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Printf("<- %s\n", line)
		fmt.Print(">> ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
}
