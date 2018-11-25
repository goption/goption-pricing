package formula_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/goption-pricing/formula"
)

var _ = Describe("Black-Scholes pricing", func() {
	var b BlackScholes
	BeforeEach(func() {
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  0,
				StrikePrice: 0,
				TTL:         0,
				Volatility:  0,
			},
		}
		RiskFree = 2.34
	})

	It("should calculate Call() prices properly", func() {
		By("using 60 65 0.25y 30%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  60,
				StrikePrice: 65,
				TTL:         0.25,
				Volatility:  0.3},
		}
		Expect(b.Call()).To(BeNumerically("~", 1.8603, 0.0001))

		By("using 150 146  14d 30%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  150,
				StrikePrice: 146,
				TTL:         14.0 / 365,
				Volatility:  0.3},
		}
		Expect(b.Call()).To(BeNumerically("~", 5.9168, 0.0001))

		By("using 1560 1500 19mo 15%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  1560,
				StrikePrice: 1500,
				TTL:         19.0 / 12,
				Volatility:  0.15},
		}
		Expect(b.Call()).To(BeNumerically("~", 179.3330, 0.0001))

		By("using 145 125 55d 65% NVDAJAN19125") // This differs from Excel by +0.0017
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  145,
				StrikePrice: 125,
				TTL:         55.0 / 365,
				Volatility:  0.65},
		}
		Expect(b.Call()).To(BeNumerically("~", 26.1147, 0.0001))

		By("using 471.42 475 209d 33.8% CMGJUN19475") // Excel shows 49.1868 for this
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  471.42,
				StrikePrice: 475,
				TTL:         209.0 / 365,
				Volatility:  0.338},
		}
		Expect(b.Call()).To(BeNumerically("~", 49.2149, 0.0001))

		By("using 1023.88 1150 574d 34% GOOGJUN201150") // Excel shows 140.9677
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  1023.88,
				StrikePrice: 1150,
				TTL:         574.0 / 365,
				Volatility:  0.34},
		}
		Expect(b.Call()).To(BeNumerically("~", 140.7629, 0.0001))

		By("using 3.15 3 83d 76.6% CHKFEB193") // Excel shows 0.5320
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  3.15,
				StrikePrice: 3,
				TTL:         83.0 / 365,
				Volatility:  0.766},
		}
		Expect(b.Call()).To(BeNumerically("~", 0.5322, 0.0001))
	})

	It("should calculate Put() prices properly", func() {
		By("using 60 65 0.25y 30%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  60,
				StrikePrice: 65,
				TTL:         0.25,
				Volatility:  0.3},
		}
		Expect(b.Put()).To(BeNumerically("~", 6.4812, 0.0001))

		By("using 150 146  14d 30%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  150,
				StrikePrice: 146,
				TTL:         14.0 / 365,
				Volatility:  0.3},
		}
		Expect(b.Put()).To(BeNumerically("~", 1.7858, 0.0001))

		By("using 1560 1500 19mo 15%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  1560,
				StrikePrice: 1500,
				TTL:         19.0 / 12,
				Volatility:  0.15},
		}
		Expect(b.Put()).To(BeNumerically("~", 64.7749, 0.0001))

		By("using 145 125 55d 65% NVDAJAN19125") // This differs from Excel by +0.0016
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  145,
				StrikePrice: 125,
				TTL:         55.0 / 365,
				Volatility:  0.65},
		}
		Expect(b.Put()).To(BeNumerically("~", 5.6747, 0.0001))

		By("using 471.42 475 209d 33.8% CMGJUN19475") // Excel shows 46.4447 for this
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  471.42,
				StrikePrice: 475,
				TTL:         209.0 / 365,
				Volatility:  0.338},
		}
		Expect(b.Put()).To(BeNumerically("~", 46.4728, 0.0001))

		By("using 1023.88 1150 574d 34% GOOGJUN201150") // Excel shows 225.5382
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  1023.88,
				StrikePrice: 1150,
				TTL:         574.0 / 365,
				Volatility:  0.34},
		}
		Expect(b.Put()).To(BeNumerically("~", 225.3334, 0.0001))

		By("using 3.15 3 83d 76.6% CHKFEB193") // Excel shows 0.3660
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  3.15,
				StrikePrice: 3,
				TTL:         83.0 / 365,
				Volatility:  0.766},
		}
		Expect(b.Put()).To(BeNumerically("~", 0.3663, 0.0001))
	})

	It("should Calculate() prices properly", func() {
		By("using 60 65 0.25y 30%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  60,
				StrikePrice: 65,
				TTL:         0.25,
				Volatility:  0.3},
		}
		call, put := b.Calculate()
		Expect(call).To(BeNumerically("~", 1.8603, 0.0001))
		Expect(put).To(BeNumerically("~", 6.4812, 0.0001))

		By("using 150 146  14d 30%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  150,
				StrikePrice: 146,
				TTL:         14.0 / 365,
				Volatility:  0.3},
		}
		call, put = b.Calculate()
		Expect(call).To(BeNumerically("~", 5.9168, 0.0001))
		Expect(put).To(BeNumerically("~", 1.7858, 0.0001))

		By("using 1560 1500 19mo 15%")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  1560,
				StrikePrice: 1500,
				TTL:         19.0 / 12,
				Volatility:  0.15},
		}
		call, put = b.Calculate()
		Expect(call).To(BeNumerically("~", 179.3330, 0.0001))
		Expect(put).To(BeNumerically("~", 64.7749, 0.0001))

		By("using 145 125 55d 65% NVDAJAN19125")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  145,
				StrikePrice: 125,
				TTL:         55.0 / 365,
				Volatility:  0.65},
		}
		call, put = b.Calculate()
		Expect(call).To(BeNumerically("~", 26.1147, 0.0001))
		Expect(put).To(BeNumerically("~", 5.6747, 0.0001))

		By("using 471.42 475 209d 33.8% CMGJUN19475")
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  471.42,
				StrikePrice: 475,
				TTL:         209.0 / 365,
				Volatility:  0.338},
		}
		call, put = b.Calculate()
		Expect(call).To(BeNumerically("~", 49.2149, 0.0001))
		Expect(put).To(BeNumerically("~", 46.4728, 0.0001))

		By("using 1023.88 1150 574d 34% GOOGJUN201150") // Excel shows 140.9677
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  1023.88,
				StrikePrice: 1150,
				TTL:         574.0 / 365,
				Volatility:  0.34},
		}
		call, put = b.Calculate()
		Expect(call).To(BeNumerically("~", 140.7629, 0.0001))
		Expect(put).To(BeNumerically("~", 225.3334, 0.0001))

		By("using 3.15 3 83d 76.6% CHKFEB193") // Excel shows 0.5320
		b = BlackScholes{
			Formula: Formula{
				StockPrice:  3.15,
				StrikePrice: 3,
				TTL:         83.0 / 365,
				Volatility:  0.766},
		}
		call, put = b.Calculate()
		Expect(call).To(BeNumerically("~", 0.5322, 0.0001))
		Expect(put).To(BeNumerically("~", 0.3663, 0.0001))
	})
})
