package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

const layout = "2006-01-02 15:04:05"

func main() {
	http.HandleFunc("/", handleWeb)
	http.HandleFunc("/api/convert", handleAPIConvert)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWeb(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

type ConvertRequest struct {
	Time     string `json:"time"`
	FromZone string `json:"from"` // jst or utc
}

type ConvertResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func handleAPIConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var req ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var res ConvertResponse
	jst, _ := time.LoadLocation("Asia/Tokyo")

	switch req.FromZone {
	case "jst":
		t, err := time.ParseInLocation(layout, req.Time, jst)
		if err != nil {
			res.Error = "Invalid JST time format"
		} else {
			res.Result = t.UTC().Format(layout)
		}
	case "utc":
		t, err := time.Parse(layout, req.Time)
		if err != nil {
			res.Error = "Invalued UTC time format"
		} else {
			res.Result = t.In(jst).Format(layout)
		}
	default:
		res.Error = "Invalid from zone"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
