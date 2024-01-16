package main

import (
	"Learning/concurrency/cocurrentAPIcalls/controllers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api-data-1", controllers.HandleApi1)
	http.HandleFunc("/api-data-2", controllers.HandleApi2)
	http.HandleFunc("/api-data-3", controllers.HandleApi3)
	http.HandleFunc("/concurrent-api", controllers.ConcurrentAPI)

	port := 8080
	serveAdr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(serveAdr, nil)
	if err == nil {
		fmt.Println("server started at port:", port)
	} else {
		fmt.Println("error starting http on localhost:", port,err)
	}
}
