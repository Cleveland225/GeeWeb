package main

import(
	// "fmt"
	// "log"
	"net/http"
	"time"
	"base/gee"
	"log"
)

func indexHandler(c *gee.Context){
	c.HTML(http.StatusOK, "<h1>Hello!!!</h1>")
}

func helloHandler(c *gee.Context) {
	c.String(http.StatusOK, "hello %s, you are at %s\n",c.Query("name"), c.Path)
}

func helloParam(c *gee.Context) {
	c.String(http.StatusOK, "hello %s, you are at %s\n",c.Param("name"), c.Path)
}

func Form(c *gee.Context){
	c.JSON(http.StatusOK, gee.H{
		"username":c.PostForm("username"),
		"password":c.PostForm("password"),
	})
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// c.Fail(200, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main(){
	r:=gee.New()
	// r.GET("/",indexHandler)
	// r.GET("/hello",helloHandler)
	// // r.POST("/login",Form)
	// r.GET("/hello/:name",helloParam)
	// r.GET("/assets/*filepath",func(c *gee.Context){
	// 	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	// })

	// r.GET("/index", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	// })
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/", func(c *gee.Context) {
	// 		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// 	})

	// 	v1.GET("/hello", func(c *gee.Context) {
	// 		// expect /hello?name=geektutu
	// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// 	})
	// }
	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/hello/:name", func(c *gee.Context) {
	// 		// expect /hello/geektutu
	// 		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	// 	})
	// 	v2.POST("/login", func(c *gee.Context) {
	// 		c.JSON(http.StatusOK, gee.H{
	// 			"username": c.PostForm("username"),
	// 			"password": c.PostForm("password"),
	// 		})
	// 	})

	// }

	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":8989")
	// http.HandleFunc("/",indexHandler)
	// http.HandleFunc("/hello",helloHandler)

	// log.Fatal(http.ListenAndServe(":8989",nil))
}
