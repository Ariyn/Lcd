// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testArticleMaps(t *testing.T) {
	t.Parallel()

	query := ArticleMaps()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testArticleMapsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
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

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticleMapsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ArticleMaps().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticleMapsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArticleMapSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticleMapsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ArticleMapExists(ctx, tx, o.UID)
	if err != nil {
		t.Errorf("Unable to check if ArticleMap exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ArticleMapExists to return true, but got false.")
	}
}

func testArticleMapsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	articleMapFound, err := FindArticleMap(ctx, tx, o.UID)
	if err != nil {
		t.Error(err)
	}

	if articleMapFound == nil {
		t.Error("want a record, got nil")
	}
}

func testArticleMapsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ArticleMaps().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testArticleMapsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ArticleMaps().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testArticleMapsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	articleMapOne := &ArticleMap{}
	articleMapTwo := &ArticleMap{}
	if err = randomize.Struct(seed, articleMapOne, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}
	if err = randomize.Struct(seed, articleMapTwo, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = articleMapOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = articleMapTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ArticleMaps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testArticleMapsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	articleMapOne := &ArticleMap{}
	articleMapTwo := &ArticleMap{}
	if err = randomize.Struct(seed, articleMapOne, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}
	if err = randomize.Struct(seed, articleMapTwo, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = articleMapOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = articleMapTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func articleMapBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func articleMapAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ArticleMap) error {
	*o = ArticleMap{}
	return nil
}

func testArticleMapsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ArticleMap{}
	o := &ArticleMap{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, articleMapDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ArticleMap object: %s", err)
	}

	AddArticleMapHook(boil.BeforeInsertHook, articleMapBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	articleMapBeforeInsertHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.AfterInsertHook, articleMapAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	articleMapAfterInsertHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.AfterSelectHook, articleMapAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	articleMapAfterSelectHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.BeforeUpdateHook, articleMapBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	articleMapBeforeUpdateHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.AfterUpdateHook, articleMapAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	articleMapAfterUpdateHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.BeforeDeleteHook, articleMapBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	articleMapBeforeDeleteHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.AfterDeleteHook, articleMapAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	articleMapAfterDeleteHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.BeforeUpsertHook, articleMapBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	articleMapBeforeUpsertHooks = []ArticleMapHook{}

	AddArticleMapHook(boil.AfterUpsertHook, articleMapAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	articleMapAfterUpsertHooks = []ArticleMapHook{}
}

func testArticleMapsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArticleMapsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(articleMapColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArticleMapToOneArticleUsingFromArticle(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ArticleMap
	var foreign Article

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.From = foreign.UID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FromArticle().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UID != foreign.UID {
		t.Errorf("want: %v, got %v", foreign.UID, check.UID)
	}

	slice := ArticleMapSlice{&local}
	if err = local.L.LoadFromArticle(ctx, tx, false, (*[]*ArticleMap)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromArticle == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FromArticle = nil
	if err = local.L.LoadFromArticle(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromArticle == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testArticleMapToOneArticleUsingToArticle(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ArticleMap
	var foreign Article

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.To = foreign.UID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ToArticle().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UID != foreign.UID {
		t.Errorf("want: %v, got %v", foreign.UID, check.UID)
	}

	slice := ArticleMapSlice{&local}
	if err = local.L.LoadToArticle(ctx, tx, false, (*[]*ArticleMap)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ToArticle == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ToArticle = nil
	if err = local.L.LoadToArticle(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ToArticle == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testArticleMapToOneSetOpArticleUsingFromArticle(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ArticleMap
	var b, c Article

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, articleMapDBTypes, false, strmangle.SetComplement(articleMapPrimaryKeyColumns, articleMapColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Article{&b, &c} {
		err = a.SetFromArticle(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FromArticle != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FromArticleMaps[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.From != x.UID {
			t.Error("foreign key was wrong value", a.From)
		}

		zero := reflect.Zero(reflect.TypeOf(a.From))
		reflect.Indirect(reflect.ValueOf(&a.From)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.From != x.UID {
			t.Error("foreign key was wrong value", a.From, x.UID)
		}
	}
}
func testArticleMapToOneSetOpArticleUsingToArticle(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ArticleMap
	var b, c Article

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, articleMapDBTypes, false, strmangle.SetComplement(articleMapPrimaryKeyColumns, articleMapColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Article{&b, &c} {
		err = a.SetToArticle(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ToArticle != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ToArticleMaps[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.To != x.UID {
			t.Error("foreign key was wrong value", a.To)
		}

		zero := reflect.Zero(reflect.TypeOf(a.To))
		reflect.Indirect(reflect.ValueOf(&a.To)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.To != x.UID {
			t.Error("foreign key was wrong value", a.To, x.UID)
		}
	}
}

func testArticleMapsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
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

func testArticleMapsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArticleMapSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testArticleMapsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ArticleMaps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	articleMapDBTypes = map[string]string{`UID`: `int`, `From`: `int`, `To`: `int`}
	_                 = bytes.MinRead
)

func testArticleMapsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(articleMapPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(articleMapAllColumns) == len(articleMapPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testArticleMapsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(articleMapAllColumns) == len(articleMapPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ArticleMap{}
	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, articleMapDBTypes, true, articleMapPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(articleMapAllColumns, articleMapPrimaryKeyColumns) {
		fields = articleMapAllColumns
	} else {
		fields = strmangle.SetComplement(
			articleMapAllColumns,
			articleMapPrimaryKeyColumns,
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

	slice := ArticleMapSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testArticleMapsUpsert(t *testing.T) {
	t.Parallel()

	if len(articleMapAllColumns) == len(articleMapPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLArticleMapUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ArticleMap{}
	if err = randomize.Struct(seed, &o, articleMapDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ArticleMap: %s", err)
	}

	count, err := ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, articleMapDBTypes, false, articleMapPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ArticleMap struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ArticleMap: %s", err)
	}

	count, err = ArticleMaps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}