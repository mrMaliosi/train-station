package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mrMaliosi/train-station/train-station/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Train Station backend is up and running!")
	})

	router := gin.Default()
	router.GET("/employees", handler.GetFilteredEmployees(employeeRepo))
	router.Run(":8080")

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
