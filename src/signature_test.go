package rdiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignatureHash(t *testing.T) {
	alpha := Compute("../testdata/sample")

	if len(alpha) != 6 {
		t.Errorf("Alpha is wrong! Source test data is broken!")
	}

	t.Run(
		"Signature - make a signature on delta and verify it",
		func(t *testing.T) {

			delta := Compare("../testdata/changed-sample.txt", alpha)
			if len(delta) != 2 {
				t.Errorf("delta is wrong!")
			}
			deltaBody := delta.ToBytes()
			signature := Signature(deltaBody)
			verified := SignatureVerify(deltaBody, signature)
			assert.True(t, verified)
		})
}
