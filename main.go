package main

import "net/http"
import "fmt"
import "time"

func sayHello(w http.ResponseWriter, r *http.Request) {
	hour := time.Now().Hour()

	if hour < 12 {
		fmt.Fprint(w, "Hello, Good morning.")
	} else if hour < 18 {
		fmt.Fprint(w, "Hello, Good afternoon.")
	} else {
		fmt.Fprint(w, "Hello, Good evening.")
	}
}

func main() {
	fmt.Println("Started call API on http://localhost:8888/greetings")

	http.HandleFunc("/greetings", sayHello)
	http.ListenAndServe(":8888", nil)
}
