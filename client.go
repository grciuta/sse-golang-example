package main

import (
	"fmt"
	"net/http"
)

type Client struct{}

func (c *Client) Connect(address string) error {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	for {
		data := make([]byte, 1024)
		_, err := resp.Body.Read(data)
		if err != nil {
			return err
		}

		fmt.Printf("Received message: \n%s\n", string(data))
	}
}
