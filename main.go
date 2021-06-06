package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Degine gin router
	router := gin.Default()

	// Create Sub Router for customized API version and basic auth
	subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{"basic_auth_user": "userpass"}))

	subRouterAuthenticated.GET("/:IdValue", GetMethod)

	listenPort := "8080"
	// Listen and Serve on the LocalHost:Port
	router.Run(":" + listenPort)
}

func GetMethod(c *gin.Context) {
	fmt.Println("\n'GetMethod' called")
	IdValue := c.Params.ByName("IdValue")
	message := "GetMethod Called With Param: " + IdValue
	c.JSON(http.StatusOK, message)

	// Print the Request Payload in console
	ReqPayload := make([]byte, 1024)
	ReqPayload, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Request Payload Data: ", string(ReqPayload))
}
