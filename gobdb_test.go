package gobdb_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/gobdb"
)

func TestList(t *testing.T){
	db, err := gobdb.Open("tests/list.gobdb")
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}
	want := gobdb.Data{
		1: "walter",
		2: "barbara",
		3: "victor",
	}
	got := db.List()
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
	want := gobdb.Data{
		"answer": 42,
	}
	err = db.Add(want)
	if err != nil {
		t.Errorf("unable to add Object %v: %s", want, err)
	}
	got := db.List()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
