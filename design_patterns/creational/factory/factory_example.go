package factory

type SodaBottleGenerator struct {
	Bottles []SodaBottle
	Factory SodaBottleFactory
}

func (s *SodaBottleGenerator) AddBottle(bottleBrand string) {
	bottle := s.Factory.CreateBottle(bottleBrand)
	s.Bottles = append(s.Bottles, bottle)
}

func InteractWithBottles(bottles []SodaBottle) (cocaColaBottles int, pepsiBottles int) {
	for _, bottle := range bottles {
		bottleBrand := bottle.Interact()
		if bottleBrand == CocaCola {
			cocaColaBottles++
		}

		if bottleBrand == Pepsi {
			pepsiBottles++
		}
	}

	return
}
