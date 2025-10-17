package main

import (
	"fmt"
	"github.com/MrShanks/gofire/templates"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func main() {

	layout := templates.Layout()

	http.Handle("/", templ.Handler(layout))

	fmt.Println("Listening on :8080")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		os.Exit(1)
		fmt.Println("Aborting...")
	}
}
