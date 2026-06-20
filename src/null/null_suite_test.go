package NullError_test

import (
	NullError "github.com/go-composites/error/src/null"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("NullError", func() {
	var error NullError.Interface
	var kindValue = `NullError.Interface`
	//var methodName = ``
	var message = ``
	ginkgo.BeforeEach(func() {
		error = NullError.New()
	})
	ginkgo.It("method Kind() can return the Interface kind", func() {
		gomega.Expect(
			error.Kind(),
		).To(gomega.BeIdenticalTo(kindValue))
	})
	ginkgo.It("method Message() can return the associated message", func() {
		gomega.Expect(
			error.Message(),
		).To(gomega.BeIdenticalTo(message))
	})
	ginkgo.It(`method IsNull() return true as this is a NullError.Interface`, func() {
		gomega.Expect(
			error.IsNull(),
		).To(gomega.BeTrueBecause(`this is a NullError.Interface`))
	})
})
