package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"./server"
	"./error"
	"./profiler"
)

const (
	port string = ":5000"
)

func main() {
	err := godotenv.Load()
	error.CheckError(err)

	conn := server.Listen(os.Getenv("PORT"))

	defer conn.Close()
	fmt.Println(123)
	buf := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		error.CheckError(err)
		data := string(buf[0:n]);

		trace := profiler.ParseTrace(data)
		fmt.Println(trace)
	}
}

