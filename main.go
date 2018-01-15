package main

import (
	"fmt"
	"net"
	"log"
	"encoding/json"
)

const (
	port string = ":5000"
)

type Trace struct {
	ID      string `json:"id"`
	Data    string `json:"data"`
	Project string `json:"project"`
	Action  string `json:"action"`
}

func main() {
	ServerAddr,err := net.ResolveUDPAddr("udp", port)
	checkError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	checkError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		checkError(err)

		//todo: declare func
		go func() {
			data := string(buf[0:n]);
			fmt.Println("Json: ",data)
			traces := make(map[string][]Trace)

			err = json.Unmarshal([]byte(data), &traces)
			checkError(err)

			fmt.Println("Traces: ",traces)
			/*
				save to mongo
			 */
		}()

	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}