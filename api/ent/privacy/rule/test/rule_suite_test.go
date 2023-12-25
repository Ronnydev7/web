package test

import (
	"api/ent/entgenerated"
	"api/ent/internal"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var TheT *testing.T

func createTestEntClient() *entgenerated.Client {
	return internal.CreateEntClientForTest(TheT)
}

func TestSchema(t *testing.T) {
	TheT = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rule Suite")
}
