package spread

import "github.com/onwsk8r/goption-pricing/formula"

// Strangle is an options spread that involves buying a call and a put.
// The call and put have the same expiration, but different strike prices.
// Generally a strangle involves buying options that are equidistant from
// the money, although sometimes one option will be further out. This
// strategy rarely creates a delta zero situation. An example would be
// buying a Feb 120 call and a Feb 160 put when the stock is at 140, however
// it would also be common to, for example, buy a 120 call and a 180 put.
type Strangle struct {
	formula.Calculator
	CallStrike float64
	PutStrike  float64
}

func GetStrangle(c formula.Calculator) Strangle {
	return Strangle{
		Calculator: c,
	}
}

// Long implements the Spread interface
func (s *Strangle) Long() float64 {
	s.SetStrikePrice(s.CallStrike)
	call := s.Call()
	s.SetStrikePrice(s.PutStrike)
	return call + s.Put()
}

// Short implements the Spread interface
func (s *Strangle) Short() float64 {
	s.SetStrikePrice(s.CallStrike)
	call := -s.Call()
	s.SetStrikePrice(s.PutStrike)
	return call - s.Put()
}
