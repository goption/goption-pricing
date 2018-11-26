package spread

import "github.com/onwsk8r/goption-pricing/formula"

// IronCondor is almost a combination bear call and bull put spread.
// Like those two it is also a credit spread, although unlike those spreads
// it performs best when the stock does not move at all. In the case of the
// Iron Condor, you sell one OTM call and one OTM put, and then buy one of
// each a little farther from the money. As all options are out of the money,
// no movement is great. These spreads can be used with options that are quite
// a ways from the money. For example, you could sell an Aug 420 put on CMG and
// buy an Aug 400 CMG put for the first half, and sell an Aug 480 Jun CMG call
// and buy an Aug 500 CMG call. This gives the stock a bit of berth while
// keeping your risk to $2000/contract which you will pay because you did
// this spread with Chipotle!!?!?
type IronCondor struct {
	BullPut
	BearCall
	formula.Calculator
}

func GetIronCondor(c formula.Calculator) IronCondor {
	return IronCondor{
		Calculator: c,
		BullPut: BullPut{
			Calculator: c,
		},
		BearCall: BearCall{
			Calculator: c,
		},
	}
}

// Long fulfills the Spread interface
func (i *IronCondor) Long() float64 {
	// The longs
	i.SetStockPrice(i.LongCall)
	i.SetStockPrice(i.LongPut)
	call := i.Call()
	put := i.Put()

	// the shorts
	i.SetStockPrice(i.ShortCall)
	i.SetStockPrice(i.ShortPut)
	return call - i.Call() + put - i.Put()

}

// Short fulfills the Spread interface
func (i *IronCondor) Short() float64 {
	// The "longs"
	i.SetStockPrice(i.ShortCall)
	i.SetStockPrice(i.ShortPut)
	call := i.Call()
	put := i.Put()

	// the "shorts"
	i.SetStockPrice(i.LongCall)
	i.SetStockPrice(i.LongPut)
	return call - i.Call() + put - i.Put()
}
