package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

// Standard middleware
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("standard middleware")
		next.ServeHTTP(w, r)
	})
}

// Standard handler
func handler(w http.ResponseWriter, r *http.Request) {
	println("standard handler")
}

func main() {
	e := echo.New()
	e.Use(standard.WrapMiddleware(middleware))
	e.GET("/", standard.WrapHandler(http.HandlerFunc(handler)))
	e.Run(standard.New(":1323"))
}
