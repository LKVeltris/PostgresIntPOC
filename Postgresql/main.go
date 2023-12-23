package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func main() {
    PsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	initDB(PsqlInfo) // init DB

	// Router
    r := gin.Default()

    r.GET("/user", GetUser)
    r.POST("/service/router", AddRouter)
    r.PUT("/service/router/:id", UpdateRouter)
    r.DELETE("/service/router/:id", DeleteRouter)
    r.GET("/billing/info", GetBillingInfo)
    r.POST("/billing/update", UpdateBillingRecord)

    r.Run() // listen and serve on 0.0.0.0:8080
}
