package gobdb_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/gobdb"
)

func TestList(t *testing.T){
	t.Parallel()
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
	t.Parallel()
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
		t.Errorf("unable to add data %v: %s", want, err)
	}
	got := db.List()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddStruct(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir() + "/addStruct.gobdb"
	db, err := gobdb.Open(tmp)
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}

	type Prio struct {
		Name string
	}

	type Todo struct {
		Name string
		Prio Prio
	}

	todos := []Todo{
		{Name: "buy milk", Prio: Prio{"High"}},
		{Name: "tax calculation", Prio: Prio{"High"}},
	}

	want := gobdb.Data{
		"todos": todos,
	}

	//db.AddSchema(todos)
	err = db.Add(want)
	if err != nil {
		t.Errorf("unable to add data %v: %s", want, err)
	}
	got := db.List()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
