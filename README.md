# Option pricing

A library of option pricing formulae, currently limited to Black-Scholes.

### Some Badges

[![Build Status](https://travis-ci.org/goption/option-pricing.svg?branch=master)](https://travis-ci.org/goption/option-pricing) [![Coverage Status](https://coveralls.io/repos/github/goption/option-pricing/badge.svg?branch=master)](https://coveralls.io/github/goption/option-pricing?branch=master) [![Documentation](https://godoc.org/github.com/goption/option-pricing?status.svg)](http://godoc.org/github.com/goption/option-pricing) ![Code Size](https://img.shields.io/github/languages/code-size/goption/option-pricing.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/goption/option-pricing)](https://goreportcard.com/report/github.com/goption/option-pricing) [![GitHub issues](https://img.shields.io/github/issues/goption/option-pricing.svg)](https://github.com/goption/option-pricing/issues) [![license](https://img.shields.io/github/license/goption/option-pricing.svg?maxAge=2592000)](https://github.com/goption/option-pricing/LICENSE)

## Documentation

Use the [Source](https://godoc.org/github.com/goption/option-pricing/blackscholes)!

## Usage

The Black-Scholes logic lives in the `blackscholes` package. There are three main function:

- Call returns the price of a call
- Put returns the price of a put
- Calculate returns both prices without doing any extra calculations.

There are also three additional functions that are exported in case you want them:

- D1 calculates the interim d1 value
- D2 calculates the interim d2 value
- N is an adapter to the normalized CDF.

N is an adapter because it does not contain an implementation of the function- it just calls the function in another library. My attempt to define it as a variable was unsuccessful because the function it calls has a defined type in its own library, so we have not-an-interface-type-that-needs-casting issue.

## Help Us Grow

If you have functionality that you'd like to add to this library or decouple from an app, feel free to open a feature request with some sample code. If you have a great idea, we'd love to hear about it!

## Issues/Contributions

- Fork the repo
- Create a feature branch or don't, it doesn't matter- it's your repo.
- Write your tests or ask for help doing so (I use [Ginkgo](https://onsi.github.io/ginkgo/))
- Write your functionality
- Open a PR

## TODO

I'd like to get the binomial model in here as well, and it would be cool to have a function that goes out and gets the current risk-free interest rate (that's the 3-month T-Bill rate).
