package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	component := flag.String("component", "server", "defines type of instance to be launched (server/client)")
	flag.Parse()

	switch *component {
	case "server":
		server := NewServer()

		go func() {
			fmt.Println("Write message to be published to open connections")
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("-> ")
				text, _ := reader.ReadString('\n')

				// convert CRLF to LF
				text = strings.Replace(text, "\n", "", -1)
				server.Send(text)
			}
		}()

		http.ListenAndServe(":8000", server)
	case "client":
		client := &Client{}
		fmt.Printf("Client error: %s", client.Connect("http://127.0.0.1:8000").Error())
	}

}
