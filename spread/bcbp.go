package spread

import "github.com/onwsk8r/goption-pricing/formula"

// BearCall involves selling an ITM call and buying an OTM call.
// The is a credit spread, and is most effective when the stock deceeds the
// strike price of the ITM call and both options expire worthless. An example
// would be buying a Mar 130 Call and selling a Mar 100 Call with an underlying
// price of 115. Hopefully those options are pretty far out because the stock
// needs to go up 15% for you to maximize your profit.
type BearCall struct {
	formula.Calculator
	LongCallPrice  float64
	ShortCallPrice float64
}

func GetBearCall(c formula.Calculator) BearCall {
	return BearCall{
		Calculator: c,
	}
}

// Long implements the Spread interface
func (s *BearCall) Long() float64 {
	s.SetStrikePrice(s.LongCallPrice)
	longCall := s.Call()
	s.SetStrikePrice(s.ShortCallPrice)
	return longCall - s.Call()
}

// Short implements the Spread interface
func (s *BearCall) Short() float64 {
	// This is not confusing at all...
	s.SetStrikePrice(s.ShortCallPrice)
	longCall := s.Call()
	s.SetStrikePrice(s.LongCallPrice)
	return longCall - s.Call()
}

// BullPut involves selling an ITM put and buying an OTM put.
// The is a credit spread, and is most effective when the stock exceeds the
// strike price of the ITM put and both options expire worthless. An example
// would be buying a Mar 130 Put and selling a Mar 100 Put with an underlying
// price of 115. Hopefully those options are pretty far out because the stock
// needs to go up 15% for you to maximize your profit.
type BullPut struct {
	formula.Calculator
	LongPutPrice  float64
	ShortPutPrice float64
}

func GetBullPut(c formula.Calculator) BullPut {
	return BullPut{
		Calculator: c,
	}
}

// Long implements the Spread interface
func (s *BullPut) Long() float64 {
	s.SetStrikePrice(s.LongPutPrice)
	longPut := s.Put()
	s.SetStrikePrice(s.ShortPutPrice)
	return longPut - s.Put()
}

// Short implements the Spread interface
func (s *BullPut) Short() float64 {
	// This is not confusing at all...
	s.SetStrikePrice(s.ShortPutPrice)
	longPut := s.Put()
	s.SetStrikePrice(s.LongPutPrice)
	return longPut - s.Put()
}
