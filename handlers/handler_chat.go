package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (lac *LocalApiConfig) HandlersSendMessage(c *gin.Context) {
	type NewMessage struct {
		Message  string `json:"message"`
		UserName string `json:"user_name"`
	}

	message := NewMessage{}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := lac.PusherClient.Trigger("my-channel", "my-event", message)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
