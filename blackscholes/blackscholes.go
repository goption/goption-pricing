/*package blackscholes contains functions for calculating option prices
based on the pricing model developed by Fischer Black and Myron Scholes
in 1973. For many options, the Black-Scholes model is reasonably close
to the actual market price.

This package uses S to represent the stock price, X to represent the strike
price, T to represent time in number of years, and v to represent volatility
in all of the functions. The risk free interest rate is set as a global
variable for both convenience and to keep it out of your code unless you
want to set it.
*/
package blackscholes

import (
	"math"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

// TODO: Make N a variable of type func instead of a func
// The compiler complains because this function has a type of distuv.Normal
// in its library. It would make a lot more sense to use it like this as a
// variable so the implementation could be altered without requiring changes
// to the source code here....
// var N func(float64) float64 = distuv.UnitNormal

// TODO: Write a function that updates the RiskFree rate.

// RiskFree represents r in the Black-Scholes formula.
// It is generally defined as the rate of a 3-month T-Bill.
var RiskFree = 2.34 // nolint: gochecknoglobals

// Years returns the number of years from one time to another.
// You can also use time.Since(<someTime>).Hours() / 8760 to do this.
func Years(from, to time.Time) float64 {
	return to.Sub(from).Hours() / 8760 // 24 * 365
}

// N is the Cumulative Normal Distributionn function.
// It is currently an adapter to Gonum's CDF function.
func N(n float64) float64 {
	return distuv.UnitNormal.CDF(n)
}

// Call returns the Black-Scholes call price
func Call(S, X, T, v float64) float64 {
	done := D1(S, X, T, v)
	r := RiskFree / 100
	return S*N(done) - X*math.Exp(-r*T)*N(D2(done, T, v))
}

// Put returns the Black-Scholes put price
func Put(S, X, T, v float64) float64 {
	done := D1(S, X, T, v)
	r := RiskFree / 100
	return X*math.Exp(-r*T)*N(-D2(done, T, v)) - S*N(-done)
}

// Calculate returns the Black-Scholes call and put price, respectively
func Calculate(S, X, T, v float64) (call, put float64) {
	r := RiskFree / 100
	done := D1(S, X, T, v)
	dtwo := D2(done, T, v)

	call = S*N(done) - X*math.Exp(-r*T)*N(dtwo)
	put = X*math.Exp(-r*T)*N(-dtwo) - S*N(-done)
	return
}

// D1 calculates the interim D1 value
func D1(S, X, T, v float64) float64 {
	r := RiskFree / 100
	return (math.Log(S/X) + (r+math.Pow(v, 2)/2)*T) / (v * math.Sqrt(T))
}

// D2 calculates the interim D2 value
func D2(D1, T, v float64) float64 {
	return D1 - v*math.Sqrt(T)
}
