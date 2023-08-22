package factory

const (
	CocaCola = "Coca-Cola"
	Pepsi    = "Pepsi"
)

type SodaBottle interface {
	Interact() (bottleBrand string)
}

type SodaBottleFactory struct {
}

func (s SodaBottleFactory) CreateBottle(bottleBrand string) SodaBottle {
	switch bottleBrand {
	case CocaCola:
		return CocaColaBottle{}
	case Pepsi:
		return PepsiBottle{}
	}

	return nil
}

type CocaColaBottle struct{}

func (c CocaColaBottle) Interact() string {
	return CocaCola
}

type PepsiBottle struct{}

func (p PepsiBottle) Interact() string {
	return Pepsi
}
