package taxonomy

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestValidateSchema(t *testing.T) {
	g := NewGomegaWithT(t)

	err := ValidateSchema("./catalog.structs.schema.json")
	g.Expect(err).NotTo(HaveOccurred())
}
