package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type data_t struct {
	Arr      []string       `json:"arr"`
	Data     string         `json:"data"`
	Map_data map[string]int `json:"map"`
}

func main() {
	var data data_t

	json_text, _ := os.ReadFile("./data.json")
	json.Unmarshal(json_text, &data)

	fmt.Printf("%v\n", data)

	output_json_text, _ := json.Marshal(data)

	f, _ := os.Create("out.json")

	f.Write(output_json_text)

	f.Close()
}
