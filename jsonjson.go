package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

func convert(i interface{}) (result interface{}, err error) {
	result = i
	switch x := i.(type) {
	case string:
		var decoded []byte
		if decoded, err = base64.StdEncoding.DecodeString(x); err == nil {
			var i interface{}
			if err = json.Unmarshal(decoded, &i); err == nil {
				result = i
			} else {
				onlyASCIIchars := true
				for _, b := range decoded {
					if b > 126 || b < 32 {
						onlyASCIIchars = false
						break
					}
				}
				if onlyASCIIchars {
					result = string(decoded)
				}
			}
		}
	case []interface{}:
		for index, el := range x {
			var converted interface{}
			if converted, err = convert(el); err != nil {
				return
			}
			x[index] = converted
		}
	case map[string]interface{}:
		for k, v := range x {
			var converted interface{}
			if converted, err = convert(v); err != nil {
				return
			}
			x[k] = converted
		}
	}
	err = nil
	return
}

func main() {
	var input interface{}
	if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
		panic(err)
	}
	converted, err := convert(input)
	if err != nil {
		panic(err)
	}
	output, err := json.MarshalIndent(converted, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))
}
