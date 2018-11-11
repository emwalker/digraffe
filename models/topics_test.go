// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testTopics(t *testing.T) {
	t.Parallel()

	query := Topics()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTopicsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
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

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTopicsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Topics().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTopicsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TopicSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTopicsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TopicExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Topic exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TopicExists to return true, but got false.")
	}
}

func testTopicsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	topicFound, err := FindTopic(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if topicFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTopicsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Topics().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTopicsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Topics().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTopicsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	topicOne := &Topic{}
	topicTwo := &Topic{}
	if err = randomize.Struct(seed, topicOne, topicDBTypes, false, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}
	if err = randomize.Struct(seed, topicTwo, topicDBTypes, false, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = topicOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = topicTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Topics().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTopicsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	topicOne := &Topic{}
	topicTwo := &Topic{}
	if err = randomize.Struct(seed, topicOne, topicDBTypes, false, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}
	if err = randomize.Struct(seed, topicTwo, topicDBTypes, false, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = topicOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = topicTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func topicBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func topicAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Topic) error {
	*o = Topic{}
	return nil
}

func testTopicsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Topic{}
	o := &Topic{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, topicDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Topic object: %s", err)
	}

	AddTopicHook(boil.BeforeInsertHook, topicBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	topicBeforeInsertHooks = []TopicHook{}

	AddTopicHook(boil.AfterInsertHook, topicAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	topicAfterInsertHooks = []TopicHook{}

	AddTopicHook(boil.AfterSelectHook, topicAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	topicAfterSelectHooks = []TopicHook{}

	AddTopicHook(boil.BeforeUpdateHook, topicBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	topicBeforeUpdateHooks = []TopicHook{}

	AddTopicHook(boil.AfterUpdateHook, topicAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	topicAfterUpdateHooks = []TopicHook{}

	AddTopicHook(boil.BeforeDeleteHook, topicBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	topicBeforeDeleteHooks = []TopicHook{}

	AddTopicHook(boil.AfterDeleteHook, topicAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	topicAfterDeleteHooks = []TopicHook{}

	AddTopicHook(boil.BeforeUpsertHook, topicBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	topicBeforeUpsertHooks = []TopicHook{}

	AddTopicHook(boil.AfterUpsertHook, topicAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	topicAfterUpsertHooks = []TopicHook{}
}

func testTopicsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTopicsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(topicColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTopicToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Topic
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, topicDBTypes, false, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OrganizationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organization().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TopicSlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*Topic)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organization = nil
	if err = local.L.LoadOrganization(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTopicToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Topic
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, topicDBTypes, false, strmangle.SetComplement(topicPrimaryKeyColumns, topicColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organization{&b, &c} {
		err = a.SetOrganization(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organization != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Topics[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrganizationID))
		reflect.Indirect(reflect.ValueOf(&a.OrganizationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID, x.ID)
		}
	}
}

func testTopicsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
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

func testTopicsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TopicSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTopicsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Topics().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	topicDBTypes = map[string]string{`Description`: `character varying`, `ID`: `uuid`, `OrganizationID`: `uuid`}
	_            = bytes.MinRead
)

func testTopicsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(topicPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(topicColumns) == len(topicPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, topicDBTypes, true, topicPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTopicsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(topicColumns) == len(topicPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Topic{}
	if err = randomize.Struct(seed, o, topicDBTypes, true, topicColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, topicDBTypes, true, topicPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(topicColumns, topicPrimaryKeyColumns) {
		fields = topicColumns
	} else {
		fields = strmangle.SetComplement(
			topicColumns,
			topicPrimaryKeyColumns,
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

	slice := TopicSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTopicsUpsert(t *testing.T) {
	t.Parallel()

	if len(topicColumns) == len(topicPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Topic{}
	if err = randomize.Struct(seed, &o, topicDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Topic: %s", err)
	}

	count, err := Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, topicDBTypes, false, topicPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Topic struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Topic: %s", err)
	}

	count, err = Topics().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
