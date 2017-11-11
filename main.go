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
	"time"

	"github.com/BurntSushi/toml"
	ui "github.com/gizak/termui"
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

	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/timer/1s", func(e ui.Event) {
		ui.Body.Align()
		prof, err := req(config)
		if err != nil {
			ui.StopLoop()
		}
		if err = display(prof); err != nil {
			ui.StopLoop()
		}
		ui.Render(ui.Body)
		time.Sleep(1 * time.Second)
	})

	ui.Loop()
	return nil
}

func req(config Config) (Prof, error) {
	query, err := createQuery(config.GQParam)
	if err != nil {
		return Prof{}, err
	}

	postBody := fmt.Sprintf(`{"query": %s}`, strconv.QuoteToASCII(query))
	resp, err := http.Post(config.URL, "application/json", strings.NewReader(postBody))
	if err != nil {
		return Prof{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var prof Prof
	err = decoder.Decode(&prof)
	if err != nil && err != io.EOF {
		return Prof{}, err
	}
	return prof, nil
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
