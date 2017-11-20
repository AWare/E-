package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	switches := [...]statusSwitch{
		statusSwitch{os.Getenv("ONEON"), os.Getenv("ONEOFF"), "🛏"},
		statusSwitch{os.Getenv("TWOON"), os.Getenv("TWOOFF"), "📚"},
		statusSwitch{os.Getenv("FOURAON"), os.Getenv("FOURAOFF"), "🛋"},
	}
	switchMap := make(map[string]switcher)

	for _, s := range switches {
		spew.Dump(s)
		switchMap[s.name] = s
	}
	switchMap["all"] = multiswitch{switches[:]}

	key := os.Getenv("RFKEY")

	switchHandler := func(w http.ResponseWriter, r *http.Request) {
		var dat struct {
			Switch string
			Action string
		}
		secret := r.Header.Get("badlykeptsecret")

		if secret != key {
			fmt.Fprintln(w, "⛔")
			return
		}
		d := json.NewDecoder(r.Body)
		d.Decode(&dat)
		fmt.Println(dat)
		s, exists := switchMap[dat.Switch]
		if !exists {
			fmt.Println(w, dat.Switch)
			fmt.Fprintln(w, "🔕")
			return
		}
		if dat.Action == "on" {
			go s.On()
			fmt.Fprintln(w, "👍🌄")
			return
		}
		if dat.Action == "off" {
			go s.Off()
			fmt.Fprintln(w, "👍🌌")
			return
		}
		fmt.Fprintln(w, "👾")
		return
	}
	listSwitchesHandler := func(w http.ResponseWriter, r *http.Request) {
		s, _ := json.Marshal(switchMap)
		fmt.Fprint(w, string(s))
	}

	http.HandleFunc("/switch/", switchHandler)
	http.HandleFunc("/list/", listSwitchesHandler)
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	fmt.Println("HANDLING HTTP")
	http.ListenAndServe(":8080", nil)
}
