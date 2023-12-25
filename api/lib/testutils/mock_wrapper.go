package testutils

import (
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

type (
	MockWrapper[T any] interface {
		AsMockX() *mock.Mock
		Start()
		Stop()
	}

	defaultMockWrapper[T any] struct {
		MockWrapper[T]

		mock     T
		original T
		replacer func(T)
	}
)

func StartMockWrapper[T any](mock T, original T, replacer func(T)) MockWrapper[T] {
	result := defaultMockWrapper[T]{
		mock:     mock,
		original: original,
		replacer: replacer,
	}
	result.Start()
	return &result
}

func (wrapper defaultMockWrapper[T]) Start() {
	wrapper.replacer(wrapper.mock)
}

func (wrapper defaultMockWrapper[T]) Stop() {
	wrapper.replacer(wrapper.original)
}

func (wrapper defaultMockWrapper[T]) AsMockX() *mock.Mock {
	var intermediate interface{} = wrapper.mock
	result, ok := (intermediate).(*mock.Mock)
	Expect(ok).To(BeTrue())
	return result
}
