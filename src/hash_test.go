package rdiff

import (
	"testing"
)

func TestRDiffHash(t *testing.T) {
	alpha := Compute("../testdata/sample")

	if len(alpha) != 6 {
		t.Errorf("Alpha is wrong! Source test data is broken!")
	}

	t.Run(
		"rdiff - text changed in liness",
		func(t *testing.T) {

			delta := Compare("../testdata/changed-sample.txt", alpha)
			if len(delta) != 2 {
				t.Errorf("delta is wrong!")
			}
			if _, ok := delta[0]; !ok {
				// key 0 is not founded
				t.Errorf("The file name has been changed.")
			}
			if _, ok := delta[3]; !ok {
				// key 3 is not founded
				t.Errorf("The line number 3 has been changed.")
			}
		})

	t.Run(
		"rdiff - line removed from file",
		func(t *testing.T) {

			delta := Compare("../testdata/deleted-sample.txt", alpha)
			if len(delta) != 2 {
				t.Errorf("delta is wrong!")
			}
			if _, ok := delta[0]; !ok {
				// key 0 is not founded
				t.Errorf("The file name has been changed.")
			}
			if _, ok := delta[3]; !ok {
				// key 3 is not founded
				t.Errorf("The line number 3 has been changed/removed.")
			}
		})

	t.Run(
		"rdiff - line added to file",
		func(t *testing.T) {

			delta := Compare("../testdata/added-sample.txt", alpha)
			if len(delta) != 2 {
				t.Errorf("delta is wrong!")
			}
			if _, ok := delta[0]; !ok {
				// key 0 is not founded
				t.Errorf("The file name has been changed.")
			}
			if _, ok := delta[4]; !ok {
				// key 4 is not founded
				t.Errorf("The line number 4 has been added.")
			}
		})

	t.Run(
		"rdiff - mix of changes",
		func(t *testing.T) {
			delta := Compare("../testdata/mix-sample.txt", alpha)
			if len(delta) != 6 {
				t.Errorf("delta is wrong!")
			}
			if _, ok := delta[0]; !ok {
				// key 0 is not founded
				t.Errorf("The file name has been changed.")
			}
			if _, ok := delta[1]; !ok {
				// key 1 is not founded
				t.Errorf("The line number 1 has been changed.")
			}
			if _, ok := delta[2]; !ok {
				// key 2 is not founded
				t.Errorf("The line number 2 has been removed.")
			}
			if _, ok := delta[4]; !ok {
				// key 4 is not founded
				t.Errorf("The line number 4 has been added.")
			}
			if _, ok := delta[6]; !ok {
				// key 6 is not founded
				t.Errorf("The line number 6 has been added.")
			}
			if _, ok := delta[7]; !ok {
				// key 7 is not founded
				t.Errorf("The line number 7 has been added.")
			}
		})

	t.Run(
		"rdiff - mix 2 of changes (more removes at the end)",
		func(t *testing.T) {

			delta := Compare("../testdata/mix-2-sample.txt", alpha)
			if len(delta) != 4 {
				t.Errorf("delta is wrong!")
			}
			if _, ok := delta[0]; !ok {
				// key 0 is not founded
				t.Errorf("The file name has been changed.")
			}
			if _, ok := delta[1]; !ok {
				// key 1 is not founded
				t.Errorf("The line number 1 has been added.")
			}
			if _, ok := delta[3]; !ok {
				// key 3 is not founded
				t.Errorf("The line number 3 has been removed.")
			}
			if _, ok := delta[5]; !ok {
				// key 5 is not founded
				t.Errorf("The line number 5 has been removed.")
			}

		})
}
