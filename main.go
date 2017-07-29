package main

import (
	"flag"
	"log"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:3012", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Fatal("dial:", err)
			}
			defer c.Close()

			go func() {
				for {
					_, _, err := c.ReadMessage()
					if err != nil {
						return
					}
				}
			}()

			for i := 0; i < 1000; i++ {
				err := c.WriteMessage(websocket.TextMessage, []byte(""))
				if err != nil {
					log.Println("write:", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}
