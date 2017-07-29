package main

import (
	"log"
	"sync"
	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:2794/"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ws, err := websocket.Dial(url, "/", origin)
			defer ws.Close()
			if err != nil {
				log.Fatal(err)
			}

			message := []byte("")
			for i := 0; i < 1000; i++ {
				_, err = ws.Write(message)
				if err != nil {
					log.Fatal(err)
				}
			}
		}()
	}
	wg.Wait()
}
