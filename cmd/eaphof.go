package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rownrepo.duckdns.org/roberto/eaphof-back/data/jugador"
	"rownrepo.duckdns.org/roberto/eaphof-back/data/respuesta"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var jugadores jugador.Jugadores

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	hub := NewHub()

	go hub.run()

	e.GET("/ws", func(c echo.Context) error {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
		if !errors.Is(err, nil) {
			log.Println(err)
		}
		defer func() {
			delete(hub.clients, ws)
			ws.Close()
			log.Printf("Closed!")
		}()

		hub.clients[ws] = true

		log.Println("Connected!")
		ws.WriteJSON(jugadores)

		read(hub, ws)
		return nil
	})

	e.Logger.Fatal(e.Start("192.168.1.151:8000"))
}

func read(hub *Hub, client *websocket.Conn) {
	for {
		var respuesta respuesta.Respuesta
		err := client.ReadJSON(&respuesta)
		if !errors.Is(err, nil) {
			log.Printf("error courred: %v", err)
			delete(hub.clients, client)
			break
		}
		if respuesta.Corrige() {
			encontrado := false
			for _, j := range jugadores {
				log.Printf("%s == %s -> %t", j.Nick, respuesta.Nick, j.Nick == respuesta.Nick)
				if j.Nick == respuesta.Nick {
					j.AddPoint()
					encontrado = true
					log.Println(jugadores)
					break
				}
			}
			if !encontrado {
				jugadores = append(jugadores, &jugador.Jugador{Nick: respuesta.Nick, Puntos: 1})
			}
		}
		log.Println(respuesta)
		hub.broadcast <- jugadores
	}
}
