package main

import (
	"errors"
	"log"

	"github.com/gorilla/websocket"
	"rownrepo.duckdns.org/roberto/eaphof-back/internal/core/domain"
)

type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan domain.Jugadores
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan domain.Jugadores),
	}
}

func (h *Hub) run() {
	for {
		select {
		case jugadores := <-h.broadcast:
			for client := range h.clients {
				if err := client.WriteJSON(jugadores); !errors.Is(err, nil) {
					log.Printf("error ocurred: %v", err)
				}
			}
		}
	}
}
