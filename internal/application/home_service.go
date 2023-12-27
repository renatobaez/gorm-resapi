package application

// Esta es la forma de escribir un servicio en una arquitectura limpia.
// Tiene un puerto al que el controlador le manda información (HomeService - Inferface) y la implementación no depende de nada del controlador.
// Recuerda que las capas deben depender de capas más internas, en este caso solo del dominio y la infraestructura

type HomeService interface {
	Home() error
}
type homeService struct {
}

func NewHomeService() HomeService {
	return &homeService{}
}

func (hs *homeService) Home() error {
	return nil
}
