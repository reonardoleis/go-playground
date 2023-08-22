package strategy

type PaymentProvider struct {
	PaymentStrategy
}

func (p PaymentProvider) getStrategy(countryCode string) PaymentStrategy {
	switch countryCode {
	case "BR":
		return BrazilPaymentStrategy{[]float32{1.5, 2.2}}
	case "MC":
		return MonacoPaymentStrategy{1.3}
	default:
		return nil
	}
}

func (p PaymentProvider) Pay(countryCode string, amount float32) float32 {
	strategy := p.getStrategy(countryCode)
	return strategy.Pay(amount)
}
