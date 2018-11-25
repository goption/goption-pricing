package formula_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/goption-pricing/formula"
)

var _ = Describe("Formula", func() {
	Context("Years()", func() {
		It("should work with a value of one year", func() {
			res := Years(time.Now(), time.Now().Add(365*24*time.Hour))
			Expect(res).To(BeNumerically("~", 1, 0.01))
		})
		It("should work with a value of six months", func() {
			res := Years(time.Now(), time.Now().Add(182.5*24*time.Hour))
			Expect(res).To(BeNumerically("~", 0.5, 0.01))
		})
		It("should work with a value of one month", func() {
			res := Years(time.Now(), time.Now().Add(30*24*time.Hour))
			Expect(res).To(BeNumerically("~", 0.083, 0.001))
		})
		It("should work with a value of one day", func() {
			res := Years(time.Now(), time.Now().Add(24*time.Hour))
			Expect(res).To(BeNumerically("~", 0.0027, 0.0001))
		})
		It("should work with a value of ten days", func() {
			res := Years(time.Now(), time.Now().Add(10*24*time.Hour))
			Expect(res).To(BeNumerically("~", 0.027, 0.001))
		})
		It("should work with a value of negative one year", func() {
			res := Years(time.Now(), time.Now().Add(-365*24*time.Hour))
			Expect(res).To(BeNumerically("~", -1, 0.01))
		})
	})

	Context("N", func() {
		It("should work with a value of 1", func() {
			Expect(N(1)).To(BeNumerically("~", 0.84134, 0.00001))
		})
		It("should work with a value of -1", func() {
			Expect(N(-1)).To(BeNumerically("~", 0.15865, 0.00001))
		})
		It("should work with a value of 0.25", func() {
			Expect(N(0.25)).To(BeNumerically("~", 0.5987, 0.00001))
		})
		It("should work with a value of 0.5", func() {
			Expect(N(0.5)).To(BeNumerically("~", 0.69146, 0.00001))
		})
		It("should work with a value of 0.75", func() {
			Expect(N(0.75)).To(BeNumerically("~", 0.77337, 0.00001))
		})
		It("should work with a value of 0.015", func() {
			Expect(N(0.015)).To(BeNumerically("~", 0.50598, 0.00001))
		})
		It("should work with a value of 0.009", func() {
			Expect(N(0.009)).To(BeNumerically("~", 0.50359, 0.00001))
		})
		It("should work with a value of 0", func() {
			Expect(N(0)).To(BeNumerically("==", 0.5))
		})
	})

	Context("Formula", func() {
		var f Formula
		BeforeEach(func() {
			f = Formula{}
		})

		It("should calculate D1 properly", func() {
			// TODO: Find an online calculator to double-check these. Some of them
			// are less than the price in ToS, GOOG differs from the Excel spreadsheet
			// (maybe AMZN would as well), and the higher option prices are significantly
			// different than what ToS reports.
			By("using 60 65 0.25y 30%")
			f = Formula{
				StockPrice:  60,
				StrikePrice: 65,
				TTL:         0.25,
				Volatility:  0.3,
			}
			Expect(f.D1()).To(BeNumerically("~", -0.4196, 0.0001))
			By("using 150 146  14d 30%")
			f = Formula{
				StockPrice:  150,
				StrikePrice: 146,
				TTL:         14.0 / 365,
				Volatility:  0.3,
			}
			Expect(f.D1()).To(BeNumerically("~", 0.5047, 0.0001))
			By("using 1560 1500 19mo 15%")
			f = Formula{
				StockPrice:  1560,
				StrikePrice: 1500,
				TTL:         19.0 / 12,
				Volatility:  0.15,
			}
			Expect(f.D1()).To(BeNumerically("~", 0.4985, 0.0001))
			By("using 145 125 55d 65% NVDAJAN19125")
			f = Formula{
				StockPrice:  145,
				StrikePrice: 125,
				TTL:         55.0 / 365,
				Volatility:  0.65,
			}
			Expect(f.D1()).To(BeNumerically("~", 0.7284, 0.0001))
			By("using 471.42 475 209d 33.8% CMGJUN19475")
			f = Formula{
				StockPrice:  471.42,
				StrikePrice: 475,
				TTL:         209.0 / 365,
				Volatility:  0.338,
			}
			Expect(f.D1()).To(BeNumerically("~", 0.1506, 0.0001))
			By("using 1023.88 1150 574d 34% GOOGJUN201150") // This differs from Excel by 0.005
			f = Formula{
				StockPrice:  1023.88,
				StrikePrice: 1150,
				TTL:         574.0 / 365,
				Volatility:  0.34,
			}
			Expect(f.D1()).To(BeNumerically("~", 0.0270, 0.0001))
			By("using 3.15 3 83d 76.6% CHKFEB193")
			f = Formula{
				StockPrice:  3.15,
				StrikePrice: 3,
				TTL:         83.0 / 365,
				Volatility:  0.766,
			}
			Expect(f.D1()).To(BeNumerically("~", 0.3308, 0.0001))
		})
		It("should calculate D2 properly", func() {
			// TODO: These don't quite differ like D1 does...
			By("using 60 65 0.25y 30%")
			f = Formula{
				StockPrice:  60,
				StrikePrice: 65,
				TTL:         0.25,
				Volatility:  0.30}
			Expect(f.D2()).To(BeNumerically("~", -0.5696, 0.0001))

			By("using 150 146  14d 30%")
			f = Formula{
				StockPrice:  150,
				StrikePrice: 146,
				TTL:         14.0 / 365,
				Volatility:  0.3}
			Expect(f.D2()).To(BeNumerically("~", 0.4459, 0.0001))

			By("using 1560 1500 19mo 15%")
			f = Formula{
				StockPrice:  1560,
				StrikePrice: 1500,
				TTL:         19.0 / 12,
				Volatility:  0.15}
			Expect(f.D2()).To(BeNumerically("~", 0.3097, 0.0001))

			By("using 145 125 55d 65% NVDAJAN19125")
			f = Formula{
				StockPrice:  145,
				StrikePrice: 125,
				TTL:         55.0 / 365,
				Volatility:  0.65} // This is .4762 in Excel
			Expect(f.D2()).To(BeNumerically("~", 0.4760, 0.0001))

			By("using 471.42 475 209d 33.8% CMGJUN19475")
			f = Formula{
				StockPrice:  471.42,
				StrikePrice: 475,
				TTL:         209.0 / 365,
				Volatility:  0.338}
			Expect(f.D2()).To(BeNumerically("~", -0.1050, 0.0001))

			By("using 1023.88 1150 574d 34% GOOGJUN201150") // This oddly enough is identical
			f = Formula{
				StockPrice:  1023.88,
				StrikePrice: 1150,
				TTL:         574.0 / 365,
				Volatility:  0.34}
			Expect(f.D2()).To(BeNumerically("~", -0.3994, 0.0001))

			By("using 3.15 3 83d 76.6% CHKFEB193")
			f = Formula{
				StockPrice:  3.15,
				StrikePrice: 3,
				TTL:         83.0 / 365,
				Volatility:  0.766} // This is -0.0343 in Excel
			Expect(f.D2()).To(BeNumerically("~", -0.0345, 0.0001))
		})
	})
})
