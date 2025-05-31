package tmpl

import "strings"

func Join(elems []Recipient, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		dat, _ := elems[0].MarshalText()
		return string(dat)
	}

	dat, _ := elems[0].MarshalText()

	sb := strings.Builder{}
	sb.Write(dat)

	for _, r := range elems[1:] {
		dat, err := r.MarshalText()
		if err != nil {
			continue
		}
		sb.WriteString(sep)
		sb.Write(dat)
	}

	return sb.String()
}

func Split(s string, sep string) ([]Recipient, error) {
	if s == "" {
		return nil, nil
	}

	parts := strings.Split(s, sep)
	out := make([]Recipient, 0, len(parts))
	for _, p := range parts {
		item := Recipient{}
		err := item.UnmarshalText([]byte(p))
		if err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, nil
}
