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

func testArticles(t *testing.T) {
	t.Parallel()

	query := Articles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testArticlesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
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

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticlesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Articles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticlesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArticleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticlesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ArticleExists(ctx, tx, o.UID)
	if err != nil {
		t.Errorf("Unable to check if Article exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ArticleExists to return true, but got false.")
	}
}

func testArticlesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	articleFound, err := FindArticle(ctx, tx, o.UID)
	if err != nil {
		t.Error(err)
	}

	if articleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testArticlesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Articles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testArticlesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Articles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testArticlesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	articleOne := &Article{}
	articleTwo := &Article{}
	if err = randomize.Struct(seed, articleOne, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}
	if err = randomize.Struct(seed, articleTwo, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = articleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = articleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Articles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testArticlesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	articleOne := &Article{}
	articleTwo := &Article{}
	if err = randomize.Struct(seed, articleOne, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}
	if err = randomize.Struct(seed, articleTwo, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = articleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = articleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func articleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func testArticlesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Article{}
	o := &Article{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, articleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Article object: %s", err)
	}

	AddArticleHook(boil.BeforeInsertHook, articleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	articleBeforeInsertHooks = []ArticleHook{}

	AddArticleHook(boil.AfterInsertHook, articleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	articleAfterInsertHooks = []ArticleHook{}

	AddArticleHook(boil.AfterSelectHook, articleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	articleAfterSelectHooks = []ArticleHook{}

	AddArticleHook(boil.BeforeUpdateHook, articleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	articleBeforeUpdateHooks = []ArticleHook{}

	AddArticleHook(boil.AfterUpdateHook, articleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	articleAfterUpdateHooks = []ArticleHook{}

	AddArticleHook(boil.BeforeDeleteHook, articleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	articleBeforeDeleteHooks = []ArticleHook{}

	AddArticleHook(boil.AfterDeleteHook, articleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	articleAfterDeleteHooks = []ArticleHook{}

	AddArticleHook(boil.BeforeUpsertHook, articleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	articleBeforeUpsertHooks = []ArticleHook{}

	AddArticleHook(boil.AfterUpsertHook, articleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	articleAfterUpsertHooks = []ArticleHook{}
}

func testArticlesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArticlesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(articleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArticleToManyFromArticleMaps(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Article
	var b, c ArticleMap

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.From = a.UID
	c.From = a.UID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.FromArticleMaps().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.From == b.From {
			bFound = true
		}
		if v.From == c.From {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ArticleSlice{&a}
	if err = a.L.LoadFromArticleMaps(ctx, tx, false, (*[]*Article)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FromArticleMaps); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FromArticleMaps = nil
	if err = a.L.LoadFromArticleMaps(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FromArticleMaps); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testArticleToManyToArticleMaps(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Article
	var b, c ArticleMap

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, articleMapDBTypes, false, articleMapColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.To = a.UID
	c.To = a.UID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ToArticleMaps().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.To == b.To {
			bFound = true
		}
		if v.To == c.To {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ArticleSlice{&a}
	if err = a.L.LoadToArticleMaps(ctx, tx, false, (*[]*Article)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ToArticleMaps); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ToArticleMaps = nil
	if err = a.L.LoadToArticleMaps(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ToArticleMaps); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testArticleToManyAddOpFromArticleMaps(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Article
	var b, c, d, e ArticleMap

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ArticleMap{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, articleMapDBTypes, false, strmangle.SetComplement(articleMapPrimaryKeyColumns, articleMapColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ArticleMap{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFromArticleMaps(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.UID != first.From {
			t.Error("foreign key was wrong value", a.UID, first.From)
		}
		if a.UID != second.From {
			t.Error("foreign key was wrong value", a.UID, second.From)
		}

		if first.R.FromArticle != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.FromArticle != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FromArticleMaps[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FromArticleMaps[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FromArticleMaps().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testArticleToManyAddOpToArticleMaps(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Article
	var b, c, d, e ArticleMap

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ArticleMap{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, articleMapDBTypes, false, strmangle.SetComplement(articleMapPrimaryKeyColumns, articleMapColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ArticleMap{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddToArticleMaps(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.UID != first.To {
			t.Error("foreign key was wrong value", a.UID, first.To)
		}
		if a.UID != second.To {
			t.Error("foreign key was wrong value", a.UID, second.To)
		}

		if first.R.ToArticle != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ToArticle != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ToArticleMaps[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ToArticleMaps[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ToArticleMaps().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testArticleToOneUserUsingOwnerUIDUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Article
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OwnerUID = foreign.UID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.OwnerUIDUser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.UID != foreign.UID {
		t.Errorf("want: %v, got %v", foreign.UID, check.UID)
	}

	slice := ArticleSlice{&local}
	if err = local.L.LoadOwnerUIDUser(ctx, tx, false, (*[]*Article)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OwnerUIDUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OwnerUIDUser = nil
	if err = local.L.LoadOwnerUIDUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OwnerUIDUser == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testArticleToOneSetOpUserUsingOwnerUIDUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Article
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetOwnerUIDUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OwnerUIDUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OwnerUIDArticles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OwnerUID != x.UID {
			t.Error("foreign key was wrong value", a.OwnerUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OwnerUID))
		reflect.Indirect(reflect.ValueOf(&a.OwnerUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OwnerUID != x.UID {
			t.Error("foreign key was wrong value", a.OwnerUID, x.UID)
		}
	}
}

func testArticlesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
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

func testArticlesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArticleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testArticlesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Articles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	articleDBTypes = map[string]string{`UID`: `int`, `Title`: `char`, `OwnerUID`: `int`}
	_              = bytes.MinRead
)

func testArticlesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(articleAllColumns) == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, articleDBTypes, true, articlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testArticlesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(articleAllColumns) == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, articleDBTypes, true, articlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(articleAllColumns, articlePrimaryKeyColumns) {
		fields = articleAllColumns
	} else {
		fields = strmangle.SetComplement(
			articleAllColumns,
			articlePrimaryKeyColumns,
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

	slice := ArticleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testArticlesUpsert(t *testing.T) {
	t.Parallel()

	if len(articleAllColumns) == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLArticleUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Article{}
	if err = randomize.Struct(seed, &o, articleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Article: %s", err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, articleDBTypes, false, articlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Article: %s", err)
	}

	count, err = Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}