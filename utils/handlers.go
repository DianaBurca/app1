package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	output := map[string]string{
		"message": "I am the second app and I work fine",
		"status":  "200",
	}
	c.JSON(http.StatusOK, output)
}

func Health(c *gin.Context) {
	c.Status(http.StatusOK)
}

func StoreData(c *gin.Context) {
	payload := map[string]string{}
	c.BindJSON(&payload)

	if strings.Compare(payload["id"], "err") == 0 {
		c.JSON(401, gin.H{
			"message": "ID not allowed",
		})
	}

	f, err := os.OpenFile("./data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	defer f.Close()

	if _, err = f.WriteString(toString(payload)); err != nil {
		c.Status(http.StatusUnprocessableEntity)
	}

	c.Status(http.StatusCreated)
}

func toString(p map[string]string) string {
	return fmt.Sprintf("ID: %s || Name: %s\n", p["id"], p["name"])
}

func TestHandler(c *gin.Context) {
	output := map[string]string{
		"message": "test",
		"status":  "200",
	}
	c.JSON(http.StatusOK, output)
}
