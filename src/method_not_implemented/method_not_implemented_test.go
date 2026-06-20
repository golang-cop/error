package MethodNotImplementedError_test

import (
	"fmt"

	MethodNotImplementedError "github.com/go-composites/error/src/method_not_implemented"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("MethodNotImplementedError", func() {
	var error MethodNotImplementedError.Interface
	var kindValue = `MethodNotImplementedError.Interface`
	var methodName = `someMethod`
	var message = fmt.Sprintf(`The method %s is not implemented`, methodName)
	ginkgo.BeforeEach(func() {
		error = MethodNotImplementedError.New(methodName)
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
	ginkgo.It(`method IsNull() return false as this is not a NullError.Interface`, func() {
		gomega.Expect(
			error.IsNull(),
		).To(gomega.BeFalseBecause(`this is not a NullError.interface`))
	})
})
