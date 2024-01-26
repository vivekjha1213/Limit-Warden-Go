package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
	"github.com/vivekjha1213/Limit-Warden-Go/routes"
)

func main() {
    r := gin.Default()
    routes.LoadRoutes(r)

    port := 8080

    // Run the Gin application on port 8080
    r.Run(fmt.Sprintf(":%d", port))
    fmt.Printf("Go-Gin App Running on port %d...\n", port)
}
