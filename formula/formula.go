/*package formula contains functions for calculating option prices.


based on the pricing model developed by Fischer Black and Myron Scholes
in 1973. For many options, the Black-Scholes model is reasonably close
to the actual market price.

This package uses S to represent the stock price, X to represent the strike
price, T to represent time in number of years, and v to represent volatility
in all of the functions. The risk free interest rate is set as a global
variable for both convenience and to keep it out of your code unless you
want to set it.
*/
package formula

import (
	"math"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

// TODO: Write a function that updates the RiskFree rate.

// RiskFree represents r in the Black-Scholes formula.
// It is generally defined as the rate of a 3-month T-Bill.
var RiskFree = 2.34 // nolint: gochecknoglobals

// N is the Normalized CDF (NORMSDIST in Excel/Sheets) function.
// The current implementation comes from gonum, although the use
// of this variable allows the implementation to vary without
// modifying the interface (aka The Adapter Pattern)
var N func(float64) float64 = distuv.UnitNormal.CDF // nolint: gochecknoglobals

// Years returns the number of years from one time to another.
// You can also use time.Since(<someTime>).Hours() / 8760 to do this.
func Years(from, to time.Time) float64 {
	return to.Sub(from).Hours() / 8760 // 24 * 365
}

// Calculator types are capable of calculating prices for options.
// BlackScholes and Binomial would be examples of Calculators.
type Calculator interface {
	Call() float64
	Put() float64
	Calculate() (float64, float64)
	SetStockPrice(float64)
	SetStrikePrice(float64)
	SetVolatility(float64)
	SetTTL(time.Duration)
}

// Formula is a base type that the formulae embed.
// It provides methods that are common to all pricing types such as Greeks
// and some interim calculations (for example d1, which is used by
// Black-Scholes). This type can't calculate option prices by itself,
// but it provides a lot of handy information.
type Formula struct {
	StockPrice  float64
	StrikePrice float64
	Volatility  float64
	TTL         float64
}

// TODO: make these functions StockPrice, make the var stockPrice, and add GetStockPrice()

func (f *Formula) SetStockPrice(p float64) {
	f.StockPrice = p
}

func (f *Formula) SetStrikePrice(p float64) {
	f.StrikePrice = p
}

// SetVolatility sets the volatility in %
func (f *Formula) SetVolatility(v float64) {
	f.Volatility = v / 100
}

// SetTTL sets the TTL in days for now
func (f *Formula) SetTTL(t time.Duration) {
	f.TTL = float64(t) / 365
}

// D1 calculates the interim D1 value
func (f *Formula) D1() float64 {
	r := RiskFree / 100
	return (math.Log(f.StockPrice/f.StrikePrice) + (r+math.Pow(f.Volatility, 2)/2)*f.TTL) /
		(f.Volatility * math.Sqrt(f.TTL))
}

// D2 calculates the interim D2 value
func (f *Formula) D2() float64 {
	return f.D1() - f.Volatility*math.Sqrt(f.TTL)
}

// TODO: Add the greeks!
// TODO: Add SetVolatility()
