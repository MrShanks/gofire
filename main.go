package main

import (
	"encoding/json"
	"fmt"
	"gofire/templates"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func main() {

	historicData := []float64{1503.25, 1555.01, 1564.00, 1594.09, 1650.73, 16509.44, 165092.79}

	historicJSON, _ := json.Marshal(historicData)

	nw := templates.NetWorth{
		Balance:          165092.79,
		HistoricDataJSON: string(historicJSON),
	}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/dashboard", templ.Handler(templates.Dashboard(nw)))
	http.Handle("/", templ.Handler(templates.Home("Simone")))

	fmt.Println("Listening on :8080")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		os.Exit(1)
		fmt.Println("Aborting...")
	}
}
