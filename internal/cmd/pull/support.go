package pull

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasepe/cirql/internal/ioutil"
	"github.com/lucasepe/cirql/internal/tmpl"
	"github.com/lucasepe/cirql/internal/vcards"
)

func newFileRotator(filename, dir string) (ioutil.Rotator, error) {
	base := filepath.Base(filename)
	prefix := func() string {
		return strings.TrimSuffix(base, filepath.Ext(base))
	}

	wri, err := ioutil.NewFileRotator(
		ioutil.RootDir(dir),
		ioutil.Prefix(prefix),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create file writer: %w", err)
	}

	return wri, err
}

func newTemplateHandler(rot ioutil.Rotator, enc tmpl.Encoder, splitAt int) vcards.CardHandler {
	handler := &templateHandler{
		rot: rot, enc: enc, splitAt: splitAt,
	}

	return handler
}

func newDefaultHandler(wri io.Writer) vcards.CardHandler {
	return &defaultHandler{
		wri: wri,
		enc: vcards.NewEncoder(wri),
	}
}

func newTemplateEncoder(filename string) (tmpl.Encoder, error) {
	bin, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to load template file: %w", err)
	}

	enc, err := tmpl.New(string(bin))
	if err != nil {
		return nil, fmt.Errorf("unable to create template encoder: %w", err)
	}

	return enc, err
}

/***************** defaultHandler ***********************/

var _ vcards.CardHandler = (*defaultHandler)(nil)

type defaultHandler struct {
	wri io.Writer
	enc *vcards.Encoder
}

func (h defaultHandler) Handle(c vcards.Card) error {
	fmt.Fprintln(h.wri)
	return h.enc.Encode(c)
}

/***************** templateHandler ***********************/

var _ vcards.CardHandler = (*templateHandler)(nil)

type templateHandler struct {
	rot     ioutil.Rotator
	enc     tmpl.Encoder
	splitAt int
	total   int
}

func (h *templateHandler) Handle(c vcards.Card) error {
	err := h.enc.Render(h.rot, c)
	if err != nil {
		return err
	}

	if h.splitAt <= 0 {
		return nil
	}

	h.total += 1
	if h.total%h.splitAt == 0 {
		return h.rot.Rotate()
	}

	return nil
}
