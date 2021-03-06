// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testLinkTransitiveClosures(t *testing.T) {
	t.Parallel()

	query := LinkTransitiveClosures()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLinkTransitiveClosuresDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinkTransitiveClosuresQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := LinkTransitiveClosures().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinkTransitiveClosuresSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LinkTransitiveClosureSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLinkTransitiveClosuresExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LinkTransitiveClosureExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if LinkTransitiveClosure exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LinkTransitiveClosureExists to return true, but got false.")
	}
}

func testLinkTransitiveClosuresFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	linkTransitiveClosureFound, err := FindLinkTransitiveClosure(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if linkTransitiveClosureFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLinkTransitiveClosuresBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = LinkTransitiveClosures().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLinkTransitiveClosuresOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := LinkTransitiveClosures().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLinkTransitiveClosuresAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	linkTransitiveClosureOne := &LinkTransitiveClosure{}
	linkTransitiveClosureTwo := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, linkTransitiveClosureOne, linkTransitiveClosureDBTypes, false, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}
	if err = randomize.Struct(seed, linkTransitiveClosureTwo, linkTransitiveClosureDBTypes, false, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = linkTransitiveClosureOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = linkTransitiveClosureTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LinkTransitiveClosures().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLinkTransitiveClosuresCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	linkTransitiveClosureOne := &LinkTransitiveClosure{}
	linkTransitiveClosureTwo := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, linkTransitiveClosureOne, linkTransitiveClosureDBTypes, false, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}
	if err = randomize.Struct(seed, linkTransitiveClosureTwo, linkTransitiveClosureDBTypes, false, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = linkTransitiveClosureOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = linkTransitiveClosureTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testLinkTransitiveClosuresInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLinkTransitiveClosuresInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(linkTransitiveClosureColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLinkTransitiveClosureToOneLinkUsingChild(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LinkTransitiveClosure
	var foreign Link

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, linkTransitiveClosureDBTypes, false, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, linkDBTypes, false, linkColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Link struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ChildID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Child().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LinkTransitiveClosureSlice{&local}
	if err = local.L.LoadChild(ctx, tx, false, (*[]*LinkTransitiveClosure)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Child == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Child = nil
	if err = local.L.LoadChild(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Child == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLinkTransitiveClosureToOneTopicUsingParent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local LinkTransitiveClosure
	var foreign Topic

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, linkTransitiveClosureDBTypes, false, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, topicDBTypes, false, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ParentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Parent().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := LinkTransitiveClosureSlice{&local}
	if err = local.L.LoadParent(ctx, tx, false, (*[]*LinkTransitiveClosure)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Parent = nil
	if err = local.L.LoadParent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testLinkTransitiveClosureToOneSetOpLinkUsingChild(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LinkTransitiveClosure
	var b, c Link

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, linkTransitiveClosureDBTypes, false, strmangle.SetComplement(linkTransitiveClosurePrimaryKeyColumns, linkTransitiveClosureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, linkDBTypes, false, strmangle.SetComplement(linkPrimaryKeyColumns, linkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, linkDBTypes, false, strmangle.SetComplement(linkPrimaryKeyColumns, linkColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Link{&b, &c} {
		err = a.SetChild(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Child != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChildLinkTransitiveClosures[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChildID != x.ID {
			t.Error("foreign key was wrong value", a.ChildID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChildID))
		reflect.Indirect(reflect.ValueOf(&a.ChildID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChildID != x.ID {
			t.Error("foreign key was wrong value", a.ChildID, x.ID)
		}
	}
}
func testLinkTransitiveClosureToOneSetOpTopicUsingParent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a LinkTransitiveClosure
	var b, c Topic

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, linkTransitiveClosureDBTypes, false, strmangle.SetComplement(linkTransitiveClosurePrimaryKeyColumns, linkTransitiveClosureColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, topicDBTypes, false, strmangle.SetComplement(topicPrimaryKeyColumns, topicColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, topicDBTypes, false, strmangle.SetComplement(topicPrimaryKeyColumns, topicColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Topic{&b, &c} {
		err = a.SetParent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Parent != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ParentLinkTransitiveClosures[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ParentID != x.ID {
			t.Error("foreign key was wrong value", a.ParentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ParentID))
		reflect.Indirect(reflect.ValueOf(&a.ParentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ParentID != x.ID {
			t.Error("foreign key was wrong value", a.ParentID, x.ID)
		}
	}
}

func testLinkTransitiveClosuresReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLinkTransitiveClosuresReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LinkTransitiveClosureSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLinkTransitiveClosuresSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := LinkTransitiveClosures().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	linkTransitiveClosureDBTypes = map[string]string{`ID`: `integer`, `ParentID`: `uuid`, `ChildID`: `uuid`}
	_                            = bytes.MinRead
)

func testLinkTransitiveClosuresUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(linkTransitiveClosurePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(linkTransitiveClosureAllColumns) == len(linkTransitiveClosurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosurePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLinkTransitiveClosuresSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(linkTransitiveClosureAllColumns) == len(linkTransitiveClosurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &LinkTransitiveClosure{}
	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosureColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, linkTransitiveClosureDBTypes, true, linkTransitiveClosurePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(linkTransitiveClosureAllColumns, linkTransitiveClosurePrimaryKeyColumns) {
		fields = linkTransitiveClosureAllColumns
	} else {
		fields = strmangle.SetComplement(
			linkTransitiveClosureAllColumns,
			linkTransitiveClosurePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := LinkTransitiveClosureSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLinkTransitiveClosuresUpsert(t *testing.T) {
	t.Parallel()

	if len(linkTransitiveClosureAllColumns) == len(linkTransitiveClosurePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := LinkTransitiveClosure{}
	if err = randomize.Struct(seed, &o, linkTransitiveClosureDBTypes, true); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LinkTransitiveClosure: %s", err)
	}

	count, err := LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, linkTransitiveClosureDBTypes, false, linkTransitiveClosurePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize LinkTransitiveClosure struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert LinkTransitiveClosure: %s", err)
	}

	count, err = LinkTransitiveClosures().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
