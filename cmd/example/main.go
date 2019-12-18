package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Fields map[string]interface{}
type MyStruct struct {
	Name string
	Age  int
}

func makeString(f Fields) string {
	var data []string
	for k, v := range f {
		data = append(data, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(data, " ")
}

func makeJson(f Fields) string {
	b, e := json.Marshal(f)
	if e == nil {
		return string(b)
	} else {
		return ""
	}
}

func Args(v ...interface{}) {
	len := len(v)
	fmt.Println("arguments length:", len)
	fmt.Println(v...)
	i := 0
	for _, item := range v {
		switch item.(type) {
		case Fields:
			v[i] = makeString(item.(Fields))
			//fmt.Println(makeString(item.(Fields)))
			//fmt.Println(makeJson(item.(Fields)))
		}
		i++
	}
	fmt.Println(v...)
}

func main() {
	f := Fields{"a": 1, "b": 2, "people": []MyStruct{MyStruct{
		Name: "William",
		Age:  39,
	}}}
	Args(1, 2, f)
	fmt.Println("json:", makeJson(f), errors.New("hello errors"))
}
