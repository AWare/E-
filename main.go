package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/switch/", switchHandler())
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	fmt.Println("HANDLING HTTP")
	http.ListenAndServe(":8080", nil)
}

func switchHandler() func(w http.ResponseWriter, r *http.Request) {
	switches := map[string]statusSwitch{"one": statusSwitch{7289615, 7289607}}
	fmt.Println("Making handler.")
	return func(w http.ResponseWriter, r *http.Request) {
		k := r.PostFormValue("switch")
		action := r.PostFormValue("action")
		fmt.Println(k, action)
		s, exists := switches[k]
		if !exists {
			fmt.Fprintln(w, "¯\\_(ツ)_/¯")
			return
		}
		if action == "on" {
			s.On()
			fmt.Fprintln(w, "👍💡")
			return
		}
		if action == "off" {
			s.Off()
			fmt.Fprintln(w, "👍𝍈")
			return
		}
		fmt.Fprintln(w, "¯\\_(ツ)_/¯")
		return
	}
}
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HELLO")
}
