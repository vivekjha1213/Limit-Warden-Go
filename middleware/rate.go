package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	clientController "github.com/vivekjha1213/Limit-Warden-Go/controllers"
)
func RateLimit(c *gin.Context){
	key := c.Request.Header.Get("X-Client-Key")
	if key == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request please provide X-Client-Key"})
        c.Abort()
		return
	}
	client, error := clientController.GetBucket(key)
	if error != nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"Unauthorized request please provide Valid X-Client-Key"})
	    c.Abort()
		return
	}

	if !client.IsRequestAllowed(1){
		c.JSON(http.StatusBadRequest,gin.H{"error":"Too Many Requests...."})
		c.Abort()
		return
	}
	c.Next()
}