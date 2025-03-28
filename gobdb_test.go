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

func TestDelete(t *testing.T) {
	t.Parallel()

	path := t.TempDir() + "/db.gobdb"
	db, err := gobdb.Open[string](path)
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}
	want := []string{"barbara", "victor", "walter", "walter"}
	err = db.Add(want...)
	if err != nil {
		t.Errorf("unable to add data %v: %s", want, err)
	}

	// extra validation, check everything, check the state, tests should not
	// depend on other tests.
	want = []string{"barbara", "victor", "walter", "walter"}
	got := db.List()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

	err = db.Delete("walter")
	if err != nil {
		t.Fatalf("unable to delete: %s", err)
	}

	want = []string{"barbara", "victor"}
	got = db.List()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDeleteAll(t *testing.T) {
	t.Parallel()

	path := t.TempDir() + "/db.gobdb"
	db, err := gobdb.Open[string](path)
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}
	want := []string{"barbara", "victor", "walter", "walter"}
	err = db.Add(want...)
	if err != nil {
		t.Errorf("unable to add data %v: %s", want, err)
	}

	// extra validation, check everything, check the state, tests should not
	// depend on other tests.
	want = []string{"barbara", "victor", "walter", "walter"}
	got := db.List()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

	err = db.DeleteAll()
	if err != nil {
		t.Fatalf("unable to delete all: %s", err)
	}

	want = []string{}
	got = db.List()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
