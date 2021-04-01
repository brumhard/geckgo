package pkg

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MomentType", func() {
	Describe("UnmarshalJSON", func() {
		It("should work", func() {
			var t map[string]MomentType
			Expect(json.Unmarshal([]byte(`{"test":"start"}`), &t)).To(Succeed())
			Expect(t["test"]).To(Equal(MomentTypeStart))
		})
		It("should fail for invalid types", func() {
			var t MomentType
			Expect(json.Unmarshal([]byte("invalid"), &t)).ToNot(Succeed())
		})
	})
	Describe("MarshalJSON", func() {
		It("should work", func() {
			t, err := json.Marshal(MomentTypeStart)
			Expect(err).ToNot(HaveOccurred())
			Expect(t).To(Equal([]byte(`"start"`)))
		})
	})
})
