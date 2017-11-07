package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println("vim-go")
	var prof Prof
	b, _ := json.Marshal(prof)
	sjson := string(b)
	fmt.Println(sjson)

	params := make(map[string]string)
	params["path"] = os.Getenv("HOME")

	t := reflect.TypeOf(prof)
	recursive(t)
}

func recursive(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name)
		fmt.Println(field.Tag.Get("gq"))
		if field.Type.Kind() == reflect.Struct && field.Type.NumField() > 0 {
			recursive(field.Type)
		}
	}

}
