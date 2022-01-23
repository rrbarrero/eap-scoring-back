package respuesta

type Respuesta struct {
	Nick      string
	Reto      string
	Respuesta string
}

func (r Respuesta) Corrige() bool {
	return true
}
