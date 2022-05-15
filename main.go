package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	// "github.com/gin-gonic/gin"
	"fmt"
	"reflect"
	"github.com/neurocollective/wuxia/paths"
	"github.com/neurocollective/wuxia/components"
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

		fmt.Println("doc.Paths:", doc.Paths)
		fmt.Println("doc.Components:", doc.Components)
		slashPets, present := doc.Paths["/pets"]
		fmt.Println("slashPets:", slashPets)
		fmt.Println("present:", present)
		fmt.Println("reflect.TypeOf(slashPets)", reflect.TypeOf(slashPets))

		components := components.BuildComponents(doc.Components)

		paths.BuildPaths(components, doc.Paths)

		// consumes, consumesPresent := doc["consumes"]
		// definitions, definitionsPresent := doc["definitions"]

		// fmt.Println(consumes)
		// fmt.Println(definitions)
	}
}
