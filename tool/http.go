package tool

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Get(url string) (io.ReadCloser, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("error recover")
		}
	}()
	c := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, err
}
