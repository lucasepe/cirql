package ioutil

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
)

type Rotator interface {
	io.Writer
	io.Closer
	Rotate() error
}

type Option func(*fileRotator)

func RootDir(dir string) Option {
	return func(fr *fileRotator) {
		if dir == "" {
			cwd, err := os.Getwd()
			if err == nil {
				dir = cwd
			}
		}

		fr.root = dir
	}
}

func Prefix(fn func() string) Option {
	return func(fr *fileRotator) {
		if fn == nil {
			fn = func() string { return "data" }
		}

		fr.prefix = fn()
	}
}

// NewFileRotator creates a new Writer.  The files will be created in the
// root directory.  root will be created if necessary.  The
// filenames will start with prefix.
func NewFileRotator(opts ...Option) (Rotator, error) {
	l := &fileRotator{}

	for _, opt := range opts {
		opt(l)
	}

	err := l.setup()

	return l, err
}

// fileRotator implements the Rotator interface.
type fileRotator struct {
	root    string
	prefix  string
	counter int
	current *os.File
	sync.Mutex
}

// Write writes p to the current file, then checks to see if
// rotation is necessary.
func (r *fileRotator) Write(p []byte) (n int, err error) {
	r.Lock()
	defer r.Unlock()

	return r.current.Write(p)
}

// Close closes the current file.  Writer is unusable after this
// is called.
func (r *fileRotator) Close() error {
	r.Lock()
	defer r.Unlock()
	if err := r.current.Close(); err != nil {
		return err
	}
	r.current = nil

	return os.Rename(path.Join(r.root, "current"),
		path.Join(r.root, fmt.Sprintf("%s.txt", r.prefix)))
}

func (r *fileRotator) Rotate() error {
	r.Lock()
	defer r.Unlock()
	err := r.current.Close()
	if err != nil {
		return err
	}

	r.counter += 1

	filename := fmt.Sprintf("%s_%04d.txt", r.prefix, r.counter)

	err = os.Rename(path.Join(r.root, "current"), path.Join(r.root, filename))
	if err != nil {
		return err
	}

	return r.openCurrent()
}

// setup creates the root directory if necessary, then opens the
// current file.
func (r *fileRotator) setup() error {
	fi, err := os.Stat(r.root)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(r.root, os.FileMode(0755))
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else if !fi.IsDir() {
		return errors.New("root must be a directory")
	}

	// root exists, and it is a directory

	return r.openCurrent()
}

func (r *fileRotator) openCurrent() error {
	cp := path.Join(r.root, "current")
	var err error
	r.current, err = os.OpenFile(cp, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0666))
	if err != nil {
		return err
	}

	return nil
}
