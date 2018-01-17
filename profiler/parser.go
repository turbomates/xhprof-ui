package profiler

import (
	"fmt"
	"encoding/json"
	"../error"
)

type Trace struct {
	ID      string `json:"id"`
	Data    string `json:"data"`
	Project string `json:"project"`
	Action  string `json:"action"`
}

func ParseTrace(data string) (Trace) {
	fmt.Println("Json: ",data)
	var trace Trace
	err := json.Unmarshal([]byte(data), &trace)
	error.CheckError(err)

	return trace
}