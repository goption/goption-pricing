package formula

import (
	"math"
)

type BlackScholes struct {
	Formula
}

// Call returns the Black-Scholes call price
func (b *BlackScholes) Call() float64 {
	r := RiskFree / 100
	return b.StockPrice*N(b.D1()) - b.StrikePrice*math.Exp(-r*b.TTL)*N(b.D2())
}

// Put returns the Black-Scholes put price
func (b *BlackScholes) Put() float64 {
	r := RiskFree / 100
	return b.StrikePrice*math.Exp(-r*b.TTL)*N(-b.D2()) - b.StockPrice*N(-b.D1())
}

// Calculate returns the Black-Scholes call and put price, respectively
func (b *BlackScholes) Calculate() (call, put float64) {
	r := RiskFree / 100
	done := b.D1()
	dtwo := b.D2()

	call = b.StockPrice*N(done) - b.StrikePrice*math.Exp(-r*b.TTL)*N(dtwo)
	put = b.StrikePrice*math.Exp(-r*b.TTL)*N(-dtwo) - b.StockPrice*N(-done)
	return
}
