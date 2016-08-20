package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	switches := map[string]statusSwitch{"one": statusSwitch{7289615, 7289607}}
	switchHandler := func(w http.ResponseWriter, r *http.Request) {
		var dat struct {
			Switch string
			Action string
		}
		d := json.NewDecoder(r.Body)
		d.Decode(&dat)
		fmt.Println(dat)
		s, exists := switches[dat.Switch]
		if !exists {
			fmt.Println(w, dat.Switch)
			fmt.Fprintln(w, "¯\\_(ツ)_/¯ (no switch)")
			return
		}
		if dat.Action == "on" {
			s.On()
			fmt.Fprintln(w, "👍💡")
			return
		}
		if dat.Action == "off" {
			s.Off()
			fmt.Fprintln(w, "👍𝍈")
			return
		}
		fmt.Fprintln(w, "¯\\_(ツ)_/¯ (no action)")
		return
	}
	listSwitchesHandler := func(w http.ResponseWriter, r *http.Request) {
		s, _ := json.Marshal(switches)
		fmt.Fprint(w, string(s))
	}

	http.HandleFunc("/switch/", switchHandler)
	http.HandleFunc("/list/", listSwitchesHandler)
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	fmt.Println("HANDLING HTTP")
	http.ListenAndServe(":8080", nil)
}
