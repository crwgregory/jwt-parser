package main

import (
	"log"
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

		// make sure the string is divisable by 4
		part := jwt[i]
		m := len(part) % 4
		if m != 0 {
			for i := 0; i < 4 - m; i++ {
				part += `=`
			}
		}

		b, err := base64.StdEncoding.DecodeString(part)
		if err != nil {
			log.Fatalf("base(32|64) decode error: part[%d] %s \n64:%s", i, part, err)				
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