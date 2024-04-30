package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (lac *LocalApiConfig) HandlerWs(c *gin.Context) {

	data := map[string]string{"message": "hello world"}
	err := lac.PusherClient.Trigger("my-channel", "my-event", data)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
