/*
A miniature RESTful service to serve randomly chosen quotes
*/
package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	Text   string `json:"quote"`
	Source string `json:"source"`
}

func loadQuotes() []Quote {
	file, err := os.Open("quotes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	quotes := make([]Quote, 0, 48391)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &quotes)
	if err != nil {
		log.Fatal(err)
	}
	return quotes
}

func randomQuote(q []Quote) Quote {
	return q[rand.Intn(len(q))]
}

func main() {
	rand.Seed(time.Now().Unix())
	quotes := loadQuotes()
	router := gin.Default()
	router.GET("/quote", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, randomQuote(quotes))
	})
	router.Run()
}
