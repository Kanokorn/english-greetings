package main

import "net/http"
import "fmt"
import "time"

func sayHello(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	midnigth := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	midday := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.Local)
	sunset := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.Local)

	if (now.After(midnigth) || now.Equal(midnigth)) && now.Before(midday) {
		fmt.Fprint(w, "Hello, Good morning.")
	}

	if now.After(midday) && now.Before(sunset) {
		fmt.Fprint(w, "Hello, Good afternoon.")
	}

	if now.After(sunset) && now.Before(midnigth) {
		fmt.Fprint(w, "Hello, Good evening.")
	}
}

func main() {
	http.HandleFunc("/greetings", sayHello)
	http.ListenAndServe(":8888", nil)
}
