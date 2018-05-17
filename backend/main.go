package main

import (
	"log"
	"net/http"

	"github.com/Jim-Lin/bee-bee-alert/backend/db"
	"github.com/Jim-Lin/bee-bee-alert/backend/model"
)

func comparisonHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}()

	if r.Method != "POST" {
		http.Error(w, "Allowed POST method only", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Allowed application/json Content-Type only", http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	prod := model.DecodeJson(r.Body)
	log.Println(prod)

	w.Write(db.MostLike(prod))
	db.IncrCounter(prod)
}

func main() {
	http.HandleFunc("/comparison", comparisonHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
