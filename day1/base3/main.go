package main

import (
	"fmt"
	"gee"
	"net/http"
)
func handleHello(w http.ResponseWriter, req * http.Request){
	fmt.Fprintf(w, "hello")
}
func main(){
	r := gee.New()
	// r.GET("/hello", handleHello)
	// r.Run("localhost:9999")
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.Run(":9999")
}
