package sccommute

import (
	"fmt"
	"net/http"
	"strconv"
)

func handleLocPing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		latStr := r.FormValue("lat")
		longStr := r.FormValue("long")
		busId := r.FormValue("id")
		lat, err := strconv.ParseFloat(latStr, 32)
		long, err := strconv.ParseFloat(longStr, 32)
		if err != nil {
			fmt.Println("Error parsing float")
		}
		Post(LocPing{busId, float32(lat), float32(long)})
	}
}

func RunServer() {
	http.HandleFunc("/location/ping", handleLocPing)
	http.ListenAndServe(":80", nil)
}