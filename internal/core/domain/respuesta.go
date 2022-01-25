package domain

type Respuesta struct {
	Nick      string
	Reto      string
	Respuesta string
}

func (r Respuesta) Corrige() bool {
	return true
}

func (r Respuesta) Guarda() error {
	return nil
}
