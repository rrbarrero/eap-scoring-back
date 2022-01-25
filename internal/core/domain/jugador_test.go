package domain

import (
	"testing"
)

func TestAddPoint(t *testing.T) {
	des := 2
	j := Jugador{Nick: "jaime", Puntos: 0}
	j.AddPoint()
	j.AddPoint()
	if des != j.Puntos {
		t.Errorf("Expected %d, got %d\n", des, j.Puntos)
	}
}
