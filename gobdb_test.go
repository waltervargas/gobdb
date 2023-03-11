package gobdb_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/gobdb"
)

func TestList(t *testing.T){
	// opendb
	db, err := gobdb.Open("tests/list.gobdb")
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}

	want := gobdb.Collection{
		gobdb.Object{
			Val: []string{"walter","barbara","victor"},
		},
	}
	got, err := db.List()
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(want, got){
		t.Error(cmp.Diff(want, got))
	}
}

func TestAdd(t *testing.T) {
	tmp := t.TempDir() + "/add.gobdb"
	db, err := gobdb.Open(tmp)
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}

	want := gobdb.Collection{
		gobdb.Object{
			Val: 42,
		},
	}
	err = db.Add(want[0])
	if err != nil {
		t.Errorf("unable to add Object %v: %s", want, err)
	}
	got, err := db.List()
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
