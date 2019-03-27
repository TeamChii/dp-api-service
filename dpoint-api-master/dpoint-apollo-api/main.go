package main

import (
	"fmt"
	"net/http"

	"./src/th.co.droppoint/schema"
	"github.com/graphql-go/handler"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		//AllowCredentials: true,
	})
	h := handler.New(&handler.Config{
		Schema:   &Schema.GetSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	var handler = c.Handler(h)
	http.Handle("/graphql", handler)

	fmt.Println("Now GraphQL is running on port 14002")
	fmt.Println("http://localhost:14002/graphql")

	http.ListenAndServe(":14002", nil)

}
