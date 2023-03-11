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
	t.Log(got)
	if !cmp.Equal(want, got){
		t.Error(cmp.Diff(want, got))
	}
}

// func Add(c gobdb.Collection) {
// 	f, err := os.Create("tests/list.gobdb")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	enc := gob.NewEncoder(f)
// 	err = enc.Encode(c)
// 	if err != nil {
// 		panic(err)
// 	}
// }

