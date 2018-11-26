/*package spread provides a simple interface to working with spreads.
Each spread has a type that implements the Spread interface. Each spread
embeds a formula.Calculator() that it uses to calculate option prices. It may
be necessary to set several properties in addition to the four formula.Formula
properties in order to effectively use a Spread. This design decision keeps
method signatures shorter and provides a simpler way to extend or swap parts
at runtime.*/
package spread

import "github.com/onwsk8r/goption-pricing/formula"

// Spread is the base interface for a spread.
// Since spreads are not universal, they can only be abstracted so far-
// if you're not sure whether you have a Straddle or a Strangle, you will
// end up with unexpected results in the latter casr.
type Spread interface {
	formula.Calculator
	Long() float64
	Short() float64
}
