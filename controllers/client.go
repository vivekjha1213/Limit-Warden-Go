package controllers

import (
	"errors"
	"fmt"
	token "github.com/vivekjha1213/Limit-Warden-Go/pkg/utils"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)



type Client struct {
	Key       string `json:"Key"`
	MaxTokens string `json:"MaxToken"`
	FillRate  int64  `json:"FillRate"`
}


var clientBuckets = make(map[string]*token.TokenBucket)

type Rule struct{
	MaxTokens int64
	Rate      int64
}

func GenerateClientKey(c *gin.Context){
	Key := uuid.New().String()
	clientBuckets[Key] = token.NewTokenBucket(1,10)
	c.JSON(200,gin.H{"Key":Key})

}

func GetBucket(key string)(*token.TokenBucket,error){
	fmt.Println(clientBuckets)
	clientBucket := clientBuckets[key]

	if clientBucket  == nil{
		return nil,errors.New("Client is not Found")
	}
	return clientBucket,nil
}