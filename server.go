package main

import (
	"net/http"
	"io"
	"encoding/json"

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

	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(string(data)), &jsonMap)
	if err != nil {
		return err
	}
	
	fileUrl := jsonMap["message"].(string)
	
	imageRes, err := http.Get(string(fileUrl))
	if err != nil {
		return err
	}
	defer imageRes.Body.Close()

	c.Response().Header().Set(echo.HeaderContentType, "image/png")

	_, err = io.Copy(c.Response(), imageRes.Body)
	if err != nil {
		return err
	}
	
	return nil
}