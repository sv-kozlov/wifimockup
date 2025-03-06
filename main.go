package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"wifimokup/nets"
)

const (
	RedirectURL = "https://www.google.com" // redirect after connection
)

func main() {
	list := nets.Generator()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// networks list - scan imitation
	e.GET("/nets", func(c echo.Context) error {
		// time.Sleep(5 * time.Second) // scan timeout imitation

		ret := make([]nets.Network, 0, len(list))
		for _, n := range list {
			n.Password = ""
			ret = append(ret, n)
		}

		return c.JSON(http.StatusOK, ret)
	})

	// save network config
	e.POST("/nets", func(c echo.Context) error {
		ssid := c.FormValue("ssid")
		password := c.FormValue("password")

		net, ok := list[ssid]
		if !ok {
			net = nets.Network{
				Ssid:     ssid,
				Password: password,
				Stored:   true,
				RSSI:     -55,
			}
			list[ssid] = net
			slog.Info(fmt.Sprintf("Add new network ssid [%s], pwd:[%s]", ssid, password))
		} else {
			net.Password = password
			net.Stored = true
			slog.Info(fmt.Sprintf("Update password for ssid [%s], pwd:[%s]", net.Ssid, password))
		}

		// successful network connection
		resp := nets.NetworkConnectResponse{
			Ssid:        net.Ssid,
			RedirectURL: RedirectURL,
		}
		return c.JSON(http.StatusOK, resp)
	})

	// connection network
	e.PUT("/nets/:ssid", func(c echo.Context) error {
		ssid := c.Param("ssid")
		if ssid == "" {
			return c.JSON(http.StatusBadRequest, "ssid is empty")
		}

		net, ok := list[ssid]
		if !ok {
			return c.JSON(http.StatusBadRequest, "ssid not exists")
		}

		slog.Info(fmt.Sprintf("connect to net [%s]", net.Ssid))

		// successful network connection
		resp := nets.NetworkConnectResponse{
			Ssid:        net.Ssid,
			RedirectURL: RedirectURL,
		}
		return c.JSON(http.StatusOK, resp)
	})

	// network removal
	e.DELETE("/nets/:ssid", func(c echo.Context) error {
		ssid := c.Param("ssid")
		if ssid == "" {
			return c.JSON(http.StatusBadRequest, "ssid is empty")
		}

		net, ok := list[ssid]
		if !ok {
			return c.JSON(http.StatusBadRequest, "ssid not exists")
		}
		slog.Info(fmt.Sprintf("Delete ssid [%s]", net.Ssid))

		// simulate deletion from memory, but leave it in the scan list
		net.Stored = false
		net.Password = ""
		list[ssid] = net

		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start("localhost:8080"))
}
