package Error_test

import (
	Error "github.com/golang-cop/error/src"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Error", func() {
	/*
		ginkgo.It("constructor New(string) can create a new Error from Go native string", func() {
			gomega.Expect(
				Error.New(`something gone wrong`),
			).To()
		})
	*/
	ginkgo.It("method Kind() can return the Interface kind", func() {
		gomega.Expect(
			Error.New(`something gone wrong`).Kind(),
		).To(gomega.BeIdenticalTo(`Error.Interface`))
	})
	ginkgo.It("method Message() can return the associated message", func() {
		gomega.Expect(
			Error.New(`something gone wrong`).Message(),
		).To(gomega.BeIdenticalTo(`something gone wrong`))
	})
	ginkgo.It(`method IsNull() return false as this not a NullError.Interface`, func() {
		gomega.Expect(
			Error.New(`something gone wrong`).IsNull(),
		).To(gomega.BeFalseBecause(`this is not a NullError.interface`))
	})
})
