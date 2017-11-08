package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	URL     string  `toml:"url"`
	GQParam GQParam `toml:"gq_param"`
}

type GQParam struct {
	DiskPath string `toml:"disk_path"`
	FilePath string `toml:"file_path"`
	Num      int    `toml:"num"`
}

func run() error {
	var config Config
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		return err
	}

	query, err := createQuery(config.GQParam)
	if err != nil {
		return err
	}

	postBody := fmt.Sprintf(`{"query": %s}`, strconv.QuoteToASCII(query))

	log.Println(postBody)
	resp, err := http.Post(config.URL, "application/json", strings.NewReader(postBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var prof Prof
	err = decoder.Decode(&prof)
	if err != nil && err != io.EOF {
		return err
	}

	fmt.Println(prof)
	return nil
}

func createQuery(param GQParam) (string, error) {
	tmpl, err := template.ParseFiles("./template/template.json")
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	tmpl.Execute(&b, param)

	return b.String(), nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
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
