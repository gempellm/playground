package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type myStruct struct {
	privateField string
	PublicField  string
}

type otherStruct struct {
	privateField string
	PublicField  string
}

func (s *myStruct) MarshalJSON() ([]byte, error) {
	const prefix, suffix = `{"privatefield": "`, `"}`

	buf := new(bytes.Buffer)
	buf.Grow(len(prefix) + len(s.privateField) + len(suffix))

	buf.WriteString(prefix)
	buf.WriteString(s.privateField)
	buf.WriteString(suffix)

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*myStruct)(nil)

func main() {
	fmt.Println()

	var data = &myStruct{
		privateField: "privateField",
		PublicField:  "PublicField",
	}

	var publicData = otherStruct{
		privateField: "privateField",
		PublicField:  "PublicField",
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))

	jsonBytes, err = json.Marshal(publicData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}
