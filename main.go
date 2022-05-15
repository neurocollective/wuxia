package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	// "github.com/gin-gonic/gin"
	"fmt"
	"reflect"
)

func main() {
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello from golang on kubernetes!",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")


	doc, err := openapi3.NewLoader().LoadFromFile("openapi.yaml")

	if err != nil {
		fmt.Println(err.Error)
	} else {

		fmt.Println(doc.Paths)
		root, present := doc.Paths["/"]
		fmt.Println(root)
		fmt.Println(present)
		fmt.Println(reflect.TypeOf(root))

		// consumes, consumesPresent := doc["consumes"]
		// definitions, definitionsPresent := doc["definitions"]

		// fmt.Println(consumes)
		// fmt.Println(definitions)
	}
}
