package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/SahanaGowdaBV/go-microservices/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)
	response := map[string]string{
		"hostname": hostname,
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("server has started!!!")

	log.Fatal(http.ListenAndServe(":80", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("web server has started")
// 	http.ListenAndServe(":80", nil)
// }

// package main

// import "fmt"
// import "rsc.io/quote"
//  geo "github.com/SahanaGowdaBV/go-microservices/geometry"

// func rectprops(length , width float64) (float64, float64) {
// 	area := length * width
// 	perimeter := 2 * (length + width)
// 	return area, perimeter
// }

// func main() {
// 	name :="devops"
// 	var y, z = 2, 3
//     fmt.Println("Hello, World!")
// 	fmt.Println(quote.Go())
// 	fmt.Println(y, z, name)

// 	a, p:= rectprops(1, 2)
// 	fmt.Println("Area is %f and perimeter is %f, a, p")

// 	area := geo.Area(1, 2)
// 	diag := geo.Diagonal(1, 2)
// 	fmt.Println(area, diag)
// }
