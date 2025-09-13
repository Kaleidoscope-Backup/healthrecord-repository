package resolver_unit_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResolverUnit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resolver Unit Suite")
}
