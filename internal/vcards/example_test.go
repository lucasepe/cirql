package vcards_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/lucasepe/cirql/internal/vcards"
)

func ExampleNewDecoder() {
	f, err := os.Open("../../testdata/sample.vcf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := vcards.NewDecoder(f)
	for {
		card, err := dec.Decode()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if card.Name() == nil {
			continue
		}

		adr := vcards.ADR(card)
		adr = strings.ReplaceAll(adr, ";", " ")
		adr = strings.TrimSpace(adr)

		fn, gn := card.Name().FamilyName, card.Name().GivenName

		fmt.Println(fn)
		fmt.Println(gn)
		fmt.Println(vcards.EMAIL(card))
		fmt.Println(vcards.TEL(card))
		fmt.Println(vcards.BDAY(card))
		fmt.Println(adr)
		fmt.Println(vcards.GEO(card))
	}

	// Output:
	// Pallo
	// Pinco
	// pinco.pallo@gmail.com
	// +39344988755
	// 19710126
	// Via Torricelli 4 81030 Teverola (Caserta) IT
	// 40.999524 14.209458
}

// encoding a vcard can be done as follows

func ExampleNewEncoder() {
	destFile, err := os.Create("cards.vcf")
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	// data in order: first name, middle name, last name, telephone number
	contacts := [][4]string{
		{"John", "Webber", "Maxwell", "(+1) 199 8714"},
		{"Donald", "", "Ron", "(+44) 421 8913"},
		{"Eric", "E.", "Peter", "(+37) 221 9903"},
		{"Nelson", "D.", "Patrick", "(+1) 122 8810"},
	}

	var (
		// card is a map of strings to []*vcard.Field objects
		card = make(vcards.Card)

		// destination where the vcard will be encoded to
		enc = vcards.NewEncoder(destFile)
	)

	for _, entry := range contacts {
		// set only the value of a field by using card.SetValue.
		// This does not set parameters
		card.SetValue(vcards.FieldFormattedName, strings.Join(entry[:3], " "))
		card.SetValue(vcards.FieldTelephone, entry[3])

		// set the value of a field and other parameters by using card.Set
		card.Set(vcards.FieldName, &vcards.Field{
			Value: strings.Join(entry[:3], ";"),
			Params: map[string][]string{
				vcards.ParamSortAs: {
					entry[0] + " " + entry[2],
				},
			},
		})

		// make the vCard version 4 compliant
		vcards.ToV4(card)
		err := enc.Encode(card)
		if err != nil {
			log.Fatal(err)
		}
	}
}
