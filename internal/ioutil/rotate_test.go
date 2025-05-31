package ioutil

import (
	"os"
	"sort"
	"testing"
)

func TestCreate(t *testing.T) {
	root, err := os.MkdirTemp("", "multitest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(root)

	x, err := NewFileRotator(
		RootDir(root),
		Prefix(func() string { return "mt" }),
	)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := x.Write([]byte("hello\n")); err != nil {
		t.Fatal(err)
	}

	d, err := os.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	names, err := d.Readdirnames(1024)
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 1 {
		t.Errorf("number files in root: %d, expected 1", len(names))
	}
}

func TestRotate(t *testing.T) {
	const text = "Hello!\n"

	root, err := os.MkdirTemp("", "multitest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(root)

	x, err := NewFileRotator(
		RootDir(root),
		Prefix(func() string { return "mt" }),
	)
	if err != nil {
		t.Fatal(err)
	}

	splitAt := 2
	for i := 1; i < 20; i++ {
		_, err := x.Write([]byte(text))
		if err != nil {
			t.Fatal(err)
		}

		if i%splitAt != 0 {
			continue
		}

		if err := x.Rotate(); err != nil {
			t.Fatal(err)
		}
	}
	x.Close()

	d, err := os.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	names, err := d.Readdirnames(100)
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 10 {
		t.Errorf("number files in root: %d, expected 10", len(names))
		sort.Strings(names)
		for i, n := range names {
			t.Logf("%d: %q", i, n)
		}
	}
}
