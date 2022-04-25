package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecoundTimeout = 10 * time.Second

func Racer(website1, website2 string) (winner string, error error) {
	return ConfigurableRacer(website1, website2, tenSecoundTimeout)
}

func ConfigurableRacer(website1, website2 string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(website1):
		return website1, nil
	case <-ping(website2):
		return website2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", website1, website2)
	}
}

func ping(website string) chan struct{} {
	channel := make(chan struct{})
	go func() {
		http.Get(website)
		close(channel)
	}()
	return channel
}
