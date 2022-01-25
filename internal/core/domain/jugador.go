package domain

type Jugador struct {
	Nick   string `json:"nick"`
	Puntos int    `json:"puntos"`
}

type Jugadores []*Jugador

func (j *Jugador) AddPoint() {
	j.Puntos += 1
}
