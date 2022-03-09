package main

import (
	"encoding/json"
	"fmt"
)

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
