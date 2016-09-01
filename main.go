package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	switches := map[string]statusSwitch{
		"🌇":  statusSwitch{os.Getenv("ONEON"), os.Getenv("ONEOFF")},
		"💡":  statusSwitch{os.Getenv("TWOON"), os.Getenv("TWOOFF")},
		"🔦":  statusSwitch{os.Getenv("THREEON"), os.Getenv("THREEOFF")},
		"🛰 ": statusSwitch{os.Getenv("FOURON"), os.Getenv("FOUROFF")},
		"👍🏻": statusSwitch{os.Getenv("ONEAON"), os.Getenv("ONEAOFF")},
		"🐕":  statusSwitch{os.Getenv("TWOAON"), os.Getenv("TWOAOFF")},
		"🌎":  statusSwitch{os.Getenv("THREEAON"), os.Getenv("THREEAOFF")},
		"📬":  statusSwitch{os.Getenv("FOURAON"), os.Getenv("FOURAOFF")},
		"🎤":  statusSwitch{os.Getenv("AON"), os.Getenv("AOFF")},
	}

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
		s, exists := switches[dat.Switch]
		if !exists {
			fmt.Println(w, dat.Switch)
			fmt.Fprintln(w, "🔕")
			return
		}
		if dat.Action == "on" {
			s.On()
			fmt.Fprintln(w, "👍💡")
			return
		}
		if dat.Action == "off" {
			s.Off()
			fmt.Fprintln(w, "👍🌆")
			return
		}
		fmt.Fprintln(w, "👾")
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
