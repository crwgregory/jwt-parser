package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"encoding/base64"
	"os"
)

func main() {
	a := os.Args[1]
	jwt := strings.Split(a, ".")
	parts := []map[string]interface{}{}
	for i := 0; i < 2; i++ {
		b, err := base64.StdEncoding.DecodeString(jwt[i])
		if err != nil {
			panic(err)
		}
		d := make(map[string]interface{})
		if err := json.Unmarshal(b, &d); err != nil {
			panic(err)
		}
		parts = append(parts, d)
	}
	j, err := json.MarshalIndent(parts, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}