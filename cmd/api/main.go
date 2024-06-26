package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"goapi/internal/handlers"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)
	fmt.Println("\nService is running on http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
