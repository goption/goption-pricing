package spread

import "github.com/onwsk8r/goption-pricing/formula"

// Straddle is an options spread that involves buying a call and a put.
// The call and put have the same strike price and expiration. For example,
// you might buy a Jun 140 call and put.
type Straddle struct {
	formula.Calculator
}

func GetStraddle(c formula.Calculator) Straddle {
	return Straddle{
		Calculator: c,
	}
}

// Long implements the Spread interface
func (s *Straddle) Long() float64 {
	return s.Call() + s.Put()
}

// Short implements the Spread interface
func (s *Straddle) Short() float64 {
	return -s.Call() - s.Put()
}
