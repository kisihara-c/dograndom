package main

import (
	"net/http"
	"io"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", dog)
  e.GET("/hello", hello)

  e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func dog(c echo.Context) error {
	res, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	
  fmt.Printf(string(data))
	return c.JSON(http.StatusOK, string(data))

	// 画像を表示
	// return c.File(http.StatusOK, file)
}