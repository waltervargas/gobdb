package gobdb_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/gobdb"
)

func TestList(t *testing.T) {
	t.Parallel()
	db, err := gobdb.Open[string]("tests/list.gobdb")
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}
	want := []string{"barbara", "victor", "walter"}
	got := db.List()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/add.gobdb"
	db, err := gobdb.Open[string](path)
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}
	want := []string{"barbara", "victor", "walter"}
	err = db.Add(want...)
	if err != nil {
		t.Fatalf("unable to add data %v: %s", want, err)
	}
	got := db.List()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
