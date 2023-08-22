package facade

type FakeComplexInput struct {
	a *int
	b *int
	c *int
}

func ReallyComplexFunction(x int, fakeComplexInput FakeComplexInput) int {
	return x + *fakeComplexInput.a + *fakeComplexInput.b + *fakeComplexInput.c
}
