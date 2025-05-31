package tmpl

import (
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/lucasepe/cirql/internal/vcards"
)

type Encoder interface {
	Render(w io.Writer, c vcards.Card) error
}

func New(txt string) (Encoder, error) {
	tpl, err := template.New("cirql").Parse(txt)
	if err != nil {
		return nil, err
	}

	return &encoderImpl{t: tpl}, nil
}

var _ Encoder = (*encoderImpl)(nil)

type encoderImpl struct {
	t *template.Template
}

func (enc *encoderImpl) Render(w io.Writer, c vcards.Card) error {
	if c.Name() == nil {
		return fmt.Errorf("missing name field in vCard")
	}

	eml, tel := vcards.EMAIL(c), vcards.TEL(c)
	if eml == "" && tel == "" {
		return fmt.Errorf("at least email or phone number is required")
	}

	ds := map[string]any{}
	ds["GivenName"] = c.Name().GivenName
	ds["FamilyName"] = c.Name().FamilyName

	buf := bytes.Buffer{}
	if err := enc.t.Execute(&buf, ds); err != nil {
		return err
	}

	to := []Recipient{}
	if eml != "" {
		to = append(to, Recipient{
			Type: Email, Value: eml,
		})
	}
	if tel != "" {
		to = append(to, Recipient{
			Type: Phone, Value: tel,
		})
	}

	_, err := fmt.Fprintf(w, "=== %s ===\n", Join(to, ","))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%s\n\n", bytes.TrimSpace(buf.Bytes()))
	return err
}
