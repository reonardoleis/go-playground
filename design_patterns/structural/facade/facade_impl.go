package facade

func ReallyComplexFunctionFacade(x, a, b, c int) int {
	fakeComplexInput := FakeComplexInput{
		a: &a,
		b: &b,
		c: &c,
	}

	return ReallyComplexFunction(x, fakeComplexInput)
}
