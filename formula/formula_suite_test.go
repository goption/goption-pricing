package formula_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFormula(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Formula Suite")
}
