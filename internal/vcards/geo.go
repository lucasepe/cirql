package vcards

import (
	"strconv"
	"strings"
)

func (c Card) Geo() *Geo {
	f := c.Preferred(FieldGeolocation)
	if f == nil {
		return nil
	}
	return newGeo(f)
}

func (c Card) AddGeo(g *Geo) {
	c.Add(FieldGeolocation, g.field())
}

func (c Card) SetGeo(g *Geo) {
	c.Set(FieldGeolocation, g.field())
}

// Geo rappresenta una posizione geografica in un Field.
type Geo struct {
	*Field
	Lat float64
	Lon float64
}

// newGeo costruisce un Geo a partire da un Field.
func newGeo(f *Field) *Geo {
	parts := strings.SplitN(f.Value, ";", 2)
	lat, _ := strconv.ParseFloat(maybeGet(parts, 0), 64)
	lon, _ := strconv.ParseFloat(maybeGet(parts, 1), 64)
	return &Geo{
		Field: f,
		Lat:   lat,
		Lon:   lon,
	}
}

// field aggiorna il Field con i valori correnti di Lat e Lon.
func (g *Geo) field() *Field {
	if g.Field == nil {
		g.Field = new(Field)
	}
	g.Field.Value = strconv.FormatFloat(g.Lat, 'f', 6, 64) + ";" + strconv.FormatFloat(g.Lon, 'f', 6, 64)
	return g.Field
}
