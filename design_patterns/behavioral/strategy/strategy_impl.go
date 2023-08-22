package strategy

type PaymentStrategy interface {
	Pay(amount float32) (finalAmount float32)
}

type BrazilPaymentStrategy struct {
	taxes []float32
}

func (b BrazilPaymentStrategy) Pay(amount float32) float32 {
	for _, tax := range b.taxes {
		amount *= tax
	}

	return amount
}

type MonacoPaymentStrategy struct {
	singleTax float32
}

func (m MonacoPaymentStrategy) Pay(amount float32) float32 {
	return amount * m.singleTax
}
