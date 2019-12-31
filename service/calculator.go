package service

// Calculator util, just for fun
type Calculator struct {
}

// Sum two integers
func (c *Calculator) Sum(a, b int) int {
	return a + b
}

// Pow a^b
func (c *Calculator) Pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
