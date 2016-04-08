package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

func fetchBusStopData(c echo.Context) error {
	id := c.Param("busttopid")
	busStopURL := makeBusStopURL(id)
	bs := &busStop{}
	err := getJson(busStopURL, bs)
	if err != nil {
		log.Println(err)
		return err
	}
	c.JSON(http.StatusOK, bs)
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Get("/bus-stop/:busttopid", fetchBusStopData)
	log.Printf("Started Server at port 3000")
	e.Run(standard.New(":3000"))
}
