package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
    Product []Product `json:"products"`
}

type Product struct {
    ID int `json:"id"`
    Name string `json:"name"`
    ShortDesc string `json:"short_desc"`
    Desc string `json:"desc"`
    Price float64 `json:"price"`
}

var sample_products Products
func getProducts(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, sample_products)
}

func main () {
    json_data_file, err := os.Open("MOCK_DATA.json")
    if err != nil {
        fmt.Println("Can't open Datafile")
    }
    defer json_data_file.Close()
    json_data, err := io.ReadAll(json_data_file)
    if err != nil {
        fmt.Println("Can't read data in json file.")
    }
    err = json.Unmarshal(json_data, &sample_products)
    if err != nil {
        fmt.Println("Unable to unmarshal the JSON")
    }


    router := gin.Default()
    router.GET("/products", getProducts)
    router.Run("localhost:8080")
}
