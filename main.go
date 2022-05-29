package main

import (
	"assignment3/controllers"
	"fmt"
	"net/http"
)

func main() {
	PORT := ":3000"
	fmt.Println("assignment 3 koinworks")
	statusController := controllers.NewStatusController()

	http.HandleFunc("/status", statusController.GetStatus)

	fmt.Println("application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
