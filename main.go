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

	params := make([string]string)
	params["path"] = os.Getenv("HOME")

	reflect.TypeOf(prof)
}
