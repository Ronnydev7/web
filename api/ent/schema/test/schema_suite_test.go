package test

import (
	"api/ent/entgenerated"
	"api/ent/internal"
	"api/privacy/viewer"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testViewer struct {
	viewer.Viewer
}

func (testViewer) IsAdmin() bool {
	return true
}

var TheT *testing.T

func createTestEntClient() *entgenerated.Client {
	return internal.CreateEntClientForTest(TheT)
}

func TestSchema(t *testing.T) {
	TheT = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Schema Suite")
}
