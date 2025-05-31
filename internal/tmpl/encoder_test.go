package tmpl_test

import (
	"os"
	"testing"

	"github.com/lucasepe/cirql/internal/tmpl"
	"github.com/lucasepe/cirql/internal/vcards"
)

func TestEncoder(t *testing.T) {
	const draft = `
Hello {{ .GivenName }}!

Here is your code for a 20% discount!

We look forward to seeing you at our shop.

All the best,
The Shop Team`

	cards := []vcards.Card{
		{
			"VERSION": []*vcards.Field{{Value: "3.0"}},
			"UID":     []*vcards.Field{{Value: "urn:uuid:000000001"}},
			"FN":      []*vcards.Field{{Value: "J. Doe"}},
			"N":       []*vcards.Field{{Value: "Doe;J.;;;"}},
			"EMAIL":   []*vcards.Field{{Value: "jdoe@example.com"}},
		},

		{
			"VERSION": []*vcards.Field{{Value: "3.0"}},
			"UID":     []*vcards.Field{{Value: "urn:uuid:000000002"}},
			"FN":      []*vcards.Field{{Value: "Scarlett Johansson"}},
			"N":       []*vcards.Field{{Value: "Johansson;Scarlett;;;"}},
			"EMAIL":   []*vcards.Field{{Value: "scarlett.johansson@example.com"}},
			"TEL":     []*vcards.Field{{Value: "+12135550198"}},
		},
	}

	enc, err := tmpl.New(draft)
	if err != nil {
		panic(err)
	}

	for _, c := range cards {
		err := enc.Render(os.Stdout, c)
		if err != nil {
			panic(err)
		}
	}
}
