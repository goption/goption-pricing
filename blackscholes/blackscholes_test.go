package blackscholes_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/onwsk8r/goption-pricing/blackscholes"
)

var _ = Describe("Black-Scholes pricing", func() {
	BeforeEach(func() {
		RiskFree = 2.34
	})

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

	Context("Call", func() { //nolint: dupl
		It("should work with a handful of different values", func() {
			By("using 60 65 0.25y 30%")
			Expect(Call(60, 65, 0.25, 0.3)).To(BeNumerically("~", 1.8603, 0.0001))
			By("using 150 146  14d 30%")
			Expect(Call(150, 146, 14.0/365, 0.3)).To(BeNumerically("~", 5.9168, 0.0001))
			By("using 1560 1500 19mo 15%")
			Expect(Call(1560, 1500, 19.0/12, 0.15)).To(BeNumerically("~", 179.3330, 0.0001))
			By("using 145 125 55d 65% NVDAJAN19125") // This differs from Excel by +0.0017
			Expect(Call(145, 125, 55.0/365, 0.65)).To(BeNumerically("~", 26.1147, 0.0001))
			By("using 471.42 475 209d 33.8% CMGJUN19475") // Excel shows 49.1868 for this
			Expect(Call(471.42, 475, 209.0/365, 0.338)).To(BeNumerically("~", 49.2149, 0.0001))
			By("using 1023.88 1150 574d 34% GOOGJUN201150") // Excel shows 140.9677
			Expect(Call(1023.88, 1150, 574.0/365, 0.34)).To(BeNumerically("~", 140.7629, 0.0001))
			By("using 3.15 3 83d 76.6% CHKFEB193") // Excel shows 0.5320
			Expect(Call(3.15, 3, 83.0/365, 0.766)).To(BeNumerically("~", 0.5322, 0.0001))
		})
	})

	Context("Put", func() { // nolint: dupl
		It("should work with a handful of different values", func() {
			By("using 60 65 0.25y 30%")
			Expect(Put(60, 65, 0.25, 0.3)).To(BeNumerically("~", 6.4812, 0.0001))
			By("using 150 146  14d 30%")
			Expect(Put(150, 146, 14.0/365, 0.3)).To(BeNumerically("~", 1.7858, 0.0001))
			By("using 1560 1500 19mo 15%")
			Expect(Put(1560, 1500, 19.0/12, 0.15)).To(BeNumerically("~", 64.7749, 0.0001))
			By("using 145 125 55d 65% NVDAJAN19125") // This differs from Excel by +0.0016
			Expect(Put(145, 125, 55.0/365, 0.65)).To(BeNumerically("~", 5.6747, 0.0001))
			By("using 471.42 475 209d 33.8% CMGJUN19475") // Excel shows 46.4447 for this
			Expect(Put(471.42, 475, 209.0/365, 0.338)).To(BeNumerically("~", 46.4728, 0.0001))
			By("using 1023.88 1150 574d 34% GOOGJUN201150") // Excel shows 225.5382
			Expect(Put(1023.88, 1150, 574.0/365, 0.34)).To(BeNumerically("~", 225.3334, 0.0001))
			By("using 3.15 3 83d 76.6% CHKFEB193") // Excel shows 0.3660
			Expect(Put(3.15, 3, 83.0/365, 0.766)).To(BeNumerically("~", 0.3663, 0.0001))
		})
	})

	Context("Calculate", func() {
		It("should work with a handful of different values", func() {
			By("using 60 65 0.25y 30%")
			call, put := Calculate(60, 65, 0.25, 0.3)
			Expect(call).To(BeNumerically("~", 1.8603, 0.0001))
			Expect(put).To(BeNumerically("~", 6.4812, 0.0001))

			By("using 150 146  14d 30%")
			call, put = Calculate(150, 146, 14.0/365, 0.3)
			Expect(call).To(BeNumerically("~", 5.9168, 0.0001))
			Expect(put).To(BeNumerically("~", 1.7858, 0.0001))

			By("using 1560 1500 19mo 15%")
			call, put = Calculate(1560, 1500, 19.0/12, 0.15)
			Expect(call).To(BeNumerically("~", 179.3330, 0.0001))
			Expect(put).To(BeNumerically("~", 64.7749, 0.0001))

			By("using 145 125 55d 65% NVDAJAN19125")
			call, put = Calculate(145, 125, 55.0/365, 0.65)
			Expect(call).To(BeNumerically("~", 26.1147, 0.0001))
			Expect(put).To(BeNumerically("~", 5.6747, 0.0001))

			By("using 471.42 475 209d 33.8% CMGJUN19475")
			call, put = Calculate(471.42, 475, 209.0/365, 0.338)
			Expect(call).To(BeNumerically("~", 49.2149, 0.0001))
			Expect(put).To(BeNumerically("~", 46.4728, 0.0001))

			By("using 1023.88 1150 574d 34% GOOGJUN201150") // Excel shows 140.9677
			call, put = Calculate(1023.88, 1150, 574.0/365, 0.34)
			Expect(call).To(BeNumerically("~", 140.7629, 0.0001))
			Expect(put).To(BeNumerically("~", 225.3334, 0.0001))

			By("using 3.15 3 83d 76.6% CHKFEB193") // Excel shows 0.5320
			call, put = Calculate(3.15, 3, 83.0/365, 0.766)
			Expect(call).To(BeNumerically("~", 0.5322, 0.0001))
			Expect(put).To(BeNumerically("~", 0.3663, 0.0001))
		})
	})

	Context("d1", func() {
		It("should work with a handful of different values", func() {
			// TODO: Find an online calculator to double-check these. Some of them
			// are less than the price in ToS, GOOG differs from the Excel spreadsheet
			// (maybe AMZN would as well), and the higher option prices are significantly
			// different than what ToS reports.
			By("using 60 65 0.25y 30%")
			Expect(D1(60, 65, 0.25, 0.3)).To(BeNumerically("~", -0.4196, 0.0001))
			By("using 150 146  14d 30%")
			Expect(D1(150, 146, 14.0/365, 0.3)).To(BeNumerically("~", 0.5047, 0.0001))
			By("using 1560 1500 19mo 15%")
			Expect(D1(1560, 1500, 19.0/12, 0.15)).To(BeNumerically("~", 0.4985, 0.0001))
			By("using 145 125 55d 65% NVDAJAN19125")
			Expect(D1(145, 125, 55.0/365, 0.65)).To(BeNumerically("~", 0.7284, 0.0001))
			By("using 471.42 475 209d 33.8% CMGJUN19475")
			Expect(D1(471.42, 475, 209.0/365, 0.338)).To(BeNumerically("~", 0.1506, 0.0001))
			By("using 1023.88 1150 574d 34% GOOGJUN201150") // This differs from Excel by 0.005
			Expect(D1(1023.88, 1150, 574.0/365, 0.34)).To(BeNumerically("~", 0.0270, 0.0001))
			By("using 3.15 3 83d 76.6% CHKFEB193")
			Expect(D1(3.15, 3, 83.0/365, 0.766)).To(BeNumerically("~", 0.3308, 0.0001))
		})
	})
	Context("D2", func() {
		It("should work with a handful of different values", func() {
			// TODO: These differ like D1 does...
			By("using 60 65 0.25y 30%")
			d1 := D1(60, 65, 0.25, 0.30)
			Expect(D2(d1, 0.25, 0.30)).To(BeNumerically("~", -0.5696, 0.0001))

			By("using 150 146  14d 30%")
			d1 = D1(150, 146, 14.0/365, 0.3)
			Expect(D2(d1, 14.0/365, 0.3)).To(BeNumerically("~", 0.4459, 0.0001))

			By("using 1560 1500 19mo 15%")
			d1 = D1(1560, 1500, 19.0/12, 0.15)
			Expect(D2(d1, 19.0/12, 0.15)).To(BeNumerically("~", 0.3097, 0.0001))

			By("using 145 125 55d 65% NVDAJAN19125")
			d1 = D1(145, 125, 55.0/365, 0.65) // This is .4762 in Excel
			Expect(D2(d1, 55.0/365, 0.65)).To(BeNumerically("~", 0.4760, 0.0001))

			By("using 471.42 475 209d 33.8% CMGJUN19475")
			d1 = D1(471.42, 475, 209.0/365, 0.338)
			Expect(D2(d1, 209.0/365, 0.338)).To(BeNumerically("~", -0.1050, 0.0001))

			By("using 1023.88 1150 574d 34% GOOGJUN201150") // This oddly enough is identical
			d1 = D1(1023.88, 1150, 574.0/365, 0.34)
			Expect(D2(d1, 574.0/365, 0.34)).To(BeNumerically("~", -0.3994, 0.0001))

			By("using 3.15 3 83d 76.6% CHKFEB193")
			d1 = D1(3.15, 3, 83.0/365, 0.766) // This is -0.0343 in Excel
			Expect(D2(d1, 83.0/365, 0.766)).To(BeNumerically("~", -0.0345, 0.0001))
		})
	})
})
