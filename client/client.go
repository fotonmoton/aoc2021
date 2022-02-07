package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type client struct {
	session string
}

func NewClient(session string) client {
	return client{session}
}

func (c *client) Day(day int) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2021/day/%v/input", day), nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", c.session))
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return body
}
