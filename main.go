package main

import (
	"assignment-three/client"
	"assignment-three/controller"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	go func() {
		for {
			time.Sleep(15 * time.Second)
			client.DoPostAPI()
		}
	}()

	e := echo.New()
	e.POST("/update", controller.UpdateData)

	e.Logger.Fatal(e.Start(":1323"))
}
