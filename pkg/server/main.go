package server

import (
	"log"
	"net/http"

	"github.com/bm-krishna/tenant-client/internal/handlers"
)

// define all api endpints here
func BootStrapServer() {
	log.Println("server boot straping")
	handlers := handlers.Service{}
	servMux := http.NewServeMux()
	servMux.Handle("/api", &handlers)
	log.Println("sever running on port :3030")
	http.ListenAndServe(":3030", servMux)
}
