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

	todo := Todo{Name: "buy milk", Prio: Prio{"High"}}
	want := gobdb.Data{
		"todo": todo,
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

func TestAddStructList(t *testing.T) {
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


func TestAddNestedStruct(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir() + "/addNestedStruct.gobdb"
	db, err := gobdb.Open(tmp)
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}


	type Reminder struct {
		What string
	}

	type Reminders []Reminder

	reminders := Reminders{
		{What: "buy milk"},
		{What: "tax calculation"},
	}

	want := gobdb.Data{
		"reminders": reminders,
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

// HELP wanted here
// tried with muttex but stil get dupicated panic error
// panic: gob: registering duplicate types for "[]gobdb_test.Todo": []gobdb_test.Todo != []gobdb_test.Todo [recovered]
// panic: gob: registering duplicate types for "[]gobdb_test.Todo": []gobdb_test.Todo != []gobdb_test.Todo
func TestDuplicatedTypes(t *testing.T) {
	t.Parallel()
	t.Skip()

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

func TestWithType(t *testing.T) {
	t.Parallel()

	type Reminder struct {
		What string
	}

	type Reminders2 []Reminder

	tmp := t.TempDir() + "/addNestedStruct.gobdb"
	db, err := gobdb.Open(tmp, gobdb.WithType(Reminders2{}))
	if err != nil {
		t.Fatalf("unable to open db: %s", err)
	}

	reminders := Reminders2{
		{What: "buy milk"},
		{What: "tax calculation"},
	}

	want := gobdb.Data{
		"reminders": reminders,
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
