package blackscholes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBlackscholes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Black-Scholes Suite")
}
