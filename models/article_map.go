// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ArticleMap is an object representing the database table.
type ArticleMap struct {
	UID  int `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	From int `boil:"from" json:"from" toml:"from" yaml:"from"`
	To   int `boil:"to" json:"to" toml:"to" yaml:"to"`

	R *articleMapR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L articleMapL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArticleMapColumns = struct {
	UID  string
	From string
	To   string
}{
	UID:  "uid",
	From: "from",
	To:   "to",
}

// Generated where

var ArticleMapWhere = struct {
	UID  whereHelperint
	From whereHelperint
	To   whereHelperint
}{
	UID:  whereHelperint{field: "`article_map`.`uid`"},
	From: whereHelperint{field: "`article_map`.`from`"},
	To:   whereHelperint{field: "`article_map`.`to`"},
}

// ArticleMapRels is where relationship names are stored.
var ArticleMapRels = struct {
	FromArticle string
	ToArticle   string
}{
	FromArticle: "FromArticle",
	ToArticle:   "ToArticle",
}

// articleMapR is where relationships are stored.
type articleMapR struct {
	FromArticle *Article `boil:"FromArticle" json:"FromArticle" toml:"FromArticle" yaml:"FromArticle"`
	ToArticle   *Article `boil:"ToArticle" json:"ToArticle" toml:"ToArticle" yaml:"ToArticle"`
}

// NewStruct creates a new relationship struct
func (*articleMapR) NewStruct() *articleMapR {
	return &articleMapR{}
}

// articleMapL is where Load methods for each relationship are stored.
type articleMapL struct{}

var (
	articleMapAllColumns            = []string{"uid", "from", "to"}
	articleMapColumnsWithoutDefault = []string{"from", "to"}
	articleMapColumnsWithDefault    = []string{"uid"}
	articleMapPrimaryKeyColumns     = []string{"uid"}
)

type (
	// ArticleMapSlice is an alias for a slice of pointers to ArticleMap.
	// This should generally be used opposed to []ArticleMap.
	ArticleMapSlice []*ArticleMap
	// ArticleMapHook is the signature for custom ArticleMap hook methods
	ArticleMapHook func(context.Context, boil.ContextExecutor, *ArticleMap) error

	articleMapQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	articleMapType                 = reflect.TypeOf(&ArticleMap{})
	articleMapMapping              = queries.MakeStructMapping(articleMapType)
	articleMapPrimaryKeyMapping, _ = queries.BindMapping(articleMapType, articleMapMapping, articleMapPrimaryKeyColumns)
	articleMapInsertCacheMut       sync.RWMutex
	articleMapInsertCache          = make(map[string]insertCache)
	articleMapUpdateCacheMut       sync.RWMutex
	articleMapUpdateCache          = make(map[string]updateCache)
	articleMapUpsertCacheMut       sync.RWMutex
	articleMapUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var articleMapBeforeInsertHooks []ArticleMapHook
var articleMapBeforeUpdateHooks []ArticleMapHook
var articleMapBeforeDeleteHooks []ArticleMapHook
var articleMapBeforeUpsertHooks []ArticleMapHook

var articleMapAfterInsertHooks []ArticleMapHook
var articleMapAfterSelectHooks []ArticleMapHook
var articleMapAfterUpdateHooks []ArticleMapHook
var articleMapAfterDeleteHooks []ArticleMapHook
var articleMapAfterUpsertHooks []ArticleMapHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ArticleMap) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ArticleMap) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ArticleMap) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ArticleMap) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ArticleMap) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ArticleMap) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ArticleMap) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ArticleMap) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ArticleMap) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range articleMapAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddArticleMapHook registers your hook function for all future operations.
func AddArticleMapHook(hookPoint boil.HookPoint, articleMapHook ArticleMapHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		articleMapBeforeInsertHooks = append(articleMapBeforeInsertHooks, articleMapHook)
	case boil.BeforeUpdateHook:
		articleMapBeforeUpdateHooks = append(articleMapBeforeUpdateHooks, articleMapHook)
	case boil.BeforeDeleteHook:
		articleMapBeforeDeleteHooks = append(articleMapBeforeDeleteHooks, articleMapHook)
	case boil.BeforeUpsertHook:
		articleMapBeforeUpsertHooks = append(articleMapBeforeUpsertHooks, articleMapHook)
	case boil.AfterInsertHook:
		articleMapAfterInsertHooks = append(articleMapAfterInsertHooks, articleMapHook)
	case boil.AfterSelectHook:
		articleMapAfterSelectHooks = append(articleMapAfterSelectHooks, articleMapHook)
	case boil.AfterUpdateHook:
		articleMapAfterUpdateHooks = append(articleMapAfterUpdateHooks, articleMapHook)
	case boil.AfterDeleteHook:
		articleMapAfterDeleteHooks = append(articleMapAfterDeleteHooks, articleMapHook)
	case boil.AfterUpsertHook:
		articleMapAfterUpsertHooks = append(articleMapAfterUpsertHooks, articleMapHook)
	}
}

// OneG returns a single articleMap record from the query using the global executor.
func (q articleMapQuery) OneG(ctx context.Context) (*ArticleMap, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single articleMap record from the query.
func (q articleMapQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ArticleMap, error) {
	o := &ArticleMap{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for article_map")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ArticleMap records from the query using the global executor.
func (q articleMapQuery) AllG(ctx context.Context) (ArticleMapSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ArticleMap records from the query.
func (q articleMapQuery) All(ctx context.Context, exec boil.ContextExecutor) (ArticleMapSlice, error) {
	var o []*ArticleMap

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ArticleMap slice")
	}

	if len(articleMapAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ArticleMap records in the query, and panics on error.
func (q articleMapQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ArticleMap records in the query.
func (q articleMapQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count article_map rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q articleMapQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q articleMapQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if article_map exists")
	}

	return count > 0, nil
}

// FromArticle pointed to by the foreign key.
func (o *ArticleMap) FromArticle(mods ...qm.QueryMod) articleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`uid` = ?", o.From),
	}

	queryMods = append(queryMods, mods...)

	query := Articles(queryMods...)
	queries.SetFrom(query.Query, "`article`")

	return query
}

// ToArticle pointed to by the foreign key.
func (o *ArticleMap) ToArticle(mods ...qm.QueryMod) articleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`uid` = ?", o.To),
	}

	queryMods = append(queryMods, mods...)

	query := Articles(queryMods...)
	queries.SetFrom(query.Query, "`article`")

	return query
}

// LoadFromArticle allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (articleMapL) LoadFromArticle(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArticleMap interface{}, mods queries.Applicator) error {
	var slice []*ArticleMap
	var object *ArticleMap

	if singular {
		object = maybeArticleMap.(*ArticleMap)
	} else {
		slice = *maybeArticleMap.(*[]*ArticleMap)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &articleMapR{}
		}
		args = append(args, object.From)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &articleMapR{}
			}

			for _, a := range args {
				if a == obj.From {
					continue Outer
				}
			}

			args = append(args, obj.From)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`article`),
		qm.WhereIn(`article.uid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Article")
	}

	var resultSlice []*Article
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Article")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for article")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for article")
	}

	if len(articleMapAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.FromArticle = foreign
		if foreign.R == nil {
			foreign.R = &articleR{}
		}
		foreign.R.FromArticleMaps = append(foreign.R.FromArticleMaps, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.From == foreign.UID {
				local.R.FromArticle = foreign
				if foreign.R == nil {
					foreign.R = &articleR{}
				}
				foreign.R.FromArticleMaps = append(foreign.R.FromArticleMaps, local)
				break
			}
		}
	}

	return nil
}

// LoadToArticle allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (articleMapL) LoadToArticle(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArticleMap interface{}, mods queries.Applicator) error {
	var slice []*ArticleMap
	var object *ArticleMap

	if singular {
		object = maybeArticleMap.(*ArticleMap)
	} else {
		slice = *maybeArticleMap.(*[]*ArticleMap)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &articleMapR{}
		}
		args = append(args, object.To)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &articleMapR{}
			}

			for _, a := range args {
				if a == obj.To {
					continue Outer
				}
			}

			args = append(args, obj.To)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`article`),
		qm.WhereIn(`article.uid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Article")
	}

	var resultSlice []*Article
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Article")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for article")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for article")
	}

	if len(articleMapAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ToArticle = foreign
		if foreign.R == nil {
			foreign.R = &articleR{}
		}
		foreign.R.ToArticleMaps = append(foreign.R.ToArticleMaps, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.To == foreign.UID {
				local.R.ToArticle = foreign
				if foreign.R == nil {
					foreign.R = &articleR{}
				}
				foreign.R.ToArticleMaps = append(foreign.R.ToArticleMaps, local)
				break
			}
		}
	}

	return nil
}

// SetFromArticleG of the articleMap to the related item.
// Sets o.R.FromArticle to related.
// Adds o to related.R.FromArticleMaps.
// Uses the global database handle.
func (o *ArticleMap) SetFromArticleG(ctx context.Context, insert bool, related *Article) error {
	return o.SetFromArticle(ctx, boil.GetContextDB(), insert, related)
}

// SetFromArticle of the articleMap to the related item.
// Sets o.R.FromArticle to related.
// Adds o to related.R.FromArticleMaps.
func (o *ArticleMap) SetFromArticle(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Article) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `article_map` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"from"}),
		strmangle.WhereClause("`", "`", 0, articleMapPrimaryKeyColumns),
	)
	values := []interface{}{related.UID, o.UID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.From = related.UID
	if o.R == nil {
		o.R = &articleMapR{
			FromArticle: related,
		}
	} else {
		o.R.FromArticle = related
	}

	if related.R == nil {
		related.R = &articleR{
			FromArticleMaps: ArticleMapSlice{o},
		}
	} else {
		related.R.FromArticleMaps = append(related.R.FromArticleMaps, o)
	}

	return nil
}

// SetToArticleG of the articleMap to the related item.
// Sets o.R.ToArticle to related.
// Adds o to related.R.ToArticleMaps.
// Uses the global database handle.
func (o *ArticleMap) SetToArticleG(ctx context.Context, insert bool, related *Article) error {
	return o.SetToArticle(ctx, boil.GetContextDB(), insert, related)
}

// SetToArticle of the articleMap to the related item.
// Sets o.R.ToArticle to related.
// Adds o to related.R.ToArticleMaps.
func (o *ArticleMap) SetToArticle(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Article) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `article_map` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"to"}),
		strmangle.WhereClause("`", "`", 0, articleMapPrimaryKeyColumns),
	)
	values := []interface{}{related.UID, o.UID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.To = related.UID
	if o.R == nil {
		o.R = &articleMapR{
			ToArticle: related,
		}
	} else {
		o.R.ToArticle = related
	}

	if related.R == nil {
		related.R = &articleR{
			ToArticleMaps: ArticleMapSlice{o},
		}
	} else {
		related.R.ToArticleMaps = append(related.R.ToArticleMaps, o)
	}

	return nil
}

// ArticleMaps retrieves all the records using an executor.
func ArticleMaps(mods ...qm.QueryMod) articleMapQuery {
	mods = append(mods, qm.From("`article_map`"))
	return articleMapQuery{NewQuery(mods...)}
}

// FindArticleMapG retrieves a single record by ID.
func FindArticleMapG(ctx context.Context, uID int, selectCols ...string) (*ArticleMap, error) {
	return FindArticleMap(ctx, boil.GetContextDB(), uID, selectCols...)
}

// FindArticleMap retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArticleMap(ctx context.Context, exec boil.ContextExecutor, uID int, selectCols ...string) (*ArticleMap, error) {
	articleMapObj := &ArticleMap{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `article_map` where `uid`=?", sel,
	)

	q := queries.Raw(query, uID)

	err := q.Bind(ctx, exec, articleMapObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from article_map")
	}

	return articleMapObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ArticleMap) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ArticleMap) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no article_map provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(articleMapColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	articleMapInsertCacheMut.RLock()
	cache, cached := articleMapInsertCache[key]
	articleMapInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			articleMapAllColumns,
			articleMapColumnsWithDefault,
			articleMapColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(articleMapType, articleMapMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(articleMapType, articleMapMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `article_map` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `article_map` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `article_map` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, articleMapPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into article_map")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.UID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == articleMapMapping["uid"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.UID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for article_map")
	}

CacheNoHooks:
	if !cached {
		articleMapInsertCacheMut.Lock()
		articleMapInsertCache[key] = cache
		articleMapInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ArticleMap record using the global executor.
// See Update for more documentation.
func (o *ArticleMap) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ArticleMap.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ArticleMap) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	articleMapUpdateCacheMut.RLock()
	cache, cached := articleMapUpdateCache[key]
	articleMapUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			articleMapAllColumns,
			articleMapPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update article_map, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `article_map` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, articleMapPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(articleMapType, articleMapMapping, append(wl, articleMapPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update article_map row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for article_map")
	}

	if !cached {
		articleMapUpdateCacheMut.Lock()
		articleMapUpdateCache[key] = cache
		articleMapUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q articleMapQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q articleMapQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for article_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for article_map")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ArticleMapSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArticleMapSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), articleMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `article_map` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, articleMapPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in articleMap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all articleMap")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ArticleMap) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLArticleMapUniqueColumns = []string{
	"uid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ArticleMap) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no article_map provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(articleMapColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLArticleMapUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	articleMapUpsertCacheMut.RLock()
	cache, cached := articleMapUpsertCache[key]
	articleMapUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			articleMapAllColumns,
			articleMapColumnsWithDefault,
			articleMapColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			articleMapAllColumns,
			articleMapPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert article_map, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "article_map", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `article_map` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(articleMapType, articleMapMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(articleMapType, articleMapMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for article_map")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.UID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == articleMapMapping["uid"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(articleMapType, articleMapMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for article_map")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for article_map")
	}

CacheNoHooks:
	if !cached {
		articleMapUpsertCacheMut.Lock()
		articleMapUpsertCache[key] = cache
		articleMapUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ArticleMap record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ArticleMap) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ArticleMap record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ArticleMap) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ArticleMap provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), articleMapPrimaryKeyMapping)
	sql := "DELETE FROM `article_map` WHERE `uid`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from article_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for article_map")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q articleMapQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q articleMapQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no articleMapQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from article_map")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for article_map")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ArticleMapSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArticleMapSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(articleMapBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), articleMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `article_map` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, articleMapPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from articleMap slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for article_map")
	}

	if len(articleMapAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ArticleMap) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ArticleMap provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ArticleMap) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindArticleMap(ctx, exec, o.UID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArticleMapSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ArticleMapSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArticleMapSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ArticleMapSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), articleMapPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `article_map`.* FROM `article_map` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, articleMapPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ArticleMapSlice")
	}

	*o = slice

	return nil
}

// ArticleMapExistsG checks if the ArticleMap row exists.
func ArticleMapExistsG(ctx context.Context, uID int) (bool, error) {
	return ArticleMapExists(ctx, boil.GetContextDB(), uID)
}

// ArticleMapExists checks if the ArticleMap row exists.
func ArticleMapExists(ctx context.Context, exec boil.ContextExecutor, uID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `article_map` where `uid`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, uID)
	}
	row := exec.QueryRowContext(ctx, sql, uID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if article_map exists")
	}

	return exists, nil
}