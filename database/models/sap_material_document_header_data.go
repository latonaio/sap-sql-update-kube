// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SapMaterialDocumentHeaderDatum is an object representing the database table.
type SapMaterialDocumentHeaderDatum struct {
	MaterialDocumentYear       string      `boil:"MaterialDocumentYear" json:"MaterialDocumentYear" toml:"MaterialDocumentYear" yaml:"MaterialDocumentYear"`
	MaterialDocument           string      `boil:"MaterialDocument" json:"MaterialDocument" toml:"MaterialDocument" yaml:"MaterialDocument"`
	InventoryTransactionType   null.String `boil:"InventoryTransactionType" json:"InventoryTransactionType,omitempty" toml:"InventoryTransactionType" yaml:"InventoryTransactionType,omitempty"`
	DocumentDate               null.String `boil:"DocumentDate" json:"DocumentDate,omitempty" toml:"DocumentDate" yaml:"DocumentDate,omitempty"`
	PostingDate                null.String `boil:"PostingDate" json:"PostingDate,omitempty" toml:"PostingDate" yaml:"PostingDate,omitempty"`
	CreationDate               null.String `boil:"CreationDate" json:"CreationDate,omitempty" toml:"CreationDate" yaml:"CreationDate,omitempty"`
	CreationTime               null.String `boil:"CreationTime" json:"CreationTime,omitempty" toml:"CreationTime" yaml:"CreationTime,omitempty"`
	MaterialDocumentHeaderText null.String `boil:"MaterialDocumentHeaderText" json:"MaterialDocumentHeaderText,omitempty" toml:"MaterialDocumentHeaderText" yaml:"MaterialDocumentHeaderText,omitempty"`
	ReferenceDocument          null.String `boil:"ReferenceDocument" json:"ReferenceDocument,omitempty" toml:"ReferenceDocument" yaml:"ReferenceDocument,omitempty"`
	GoodsMovementCode          null.String `boil:"GoodsMovementCode" json:"GoodsMovementCode,omitempty" toml:"GoodsMovementCode" yaml:"GoodsMovementCode,omitempty"`

	R *sapMaterialDocumentHeaderDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapMaterialDocumentHeaderDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapMaterialDocumentHeaderDatumColumns = struct {
	MaterialDocumentYear       string
	MaterialDocument           string
	InventoryTransactionType   string
	DocumentDate               string
	PostingDate                string
	CreationDate               string
	CreationTime               string
	MaterialDocumentHeaderText string
	ReferenceDocument          string
	GoodsMovementCode          string
}{
	MaterialDocumentYear:       "MaterialDocumentYear",
	MaterialDocument:           "MaterialDocument",
	InventoryTransactionType:   "InventoryTransactionType",
	DocumentDate:               "DocumentDate",
	PostingDate:                "PostingDate",
	CreationDate:               "CreationDate",
	CreationTime:               "CreationTime",
	MaterialDocumentHeaderText: "MaterialDocumentHeaderText",
	ReferenceDocument:          "ReferenceDocument",
	GoodsMovementCode:          "GoodsMovementCode",
}

var SapMaterialDocumentHeaderDatumTableColumns = struct {
	MaterialDocumentYear       string
	MaterialDocument           string
	InventoryTransactionType   string
	DocumentDate               string
	PostingDate                string
	CreationDate               string
	CreationTime               string
	MaterialDocumentHeaderText string
	ReferenceDocument          string
	GoodsMovementCode          string
}{
	MaterialDocumentYear:       "sap_material_document_header_data.MaterialDocumentYear",
	MaterialDocument:           "sap_material_document_header_data.MaterialDocument",
	InventoryTransactionType:   "sap_material_document_header_data.InventoryTransactionType",
	DocumentDate:               "sap_material_document_header_data.DocumentDate",
	PostingDate:                "sap_material_document_header_data.PostingDate",
	CreationDate:               "sap_material_document_header_data.CreationDate",
	CreationTime:               "sap_material_document_header_data.CreationTime",
	MaterialDocumentHeaderText: "sap_material_document_header_data.MaterialDocumentHeaderText",
	ReferenceDocument:          "sap_material_document_header_data.ReferenceDocument",
	GoodsMovementCode:          "sap_material_document_header_data.GoodsMovementCode",
}

// Generated where

var SapMaterialDocumentHeaderDatumWhere = struct {
	MaterialDocumentYear       whereHelperstring
	MaterialDocument           whereHelperstring
	InventoryTransactionType   whereHelpernull_String
	DocumentDate               whereHelpernull_String
	PostingDate                whereHelpernull_String
	CreationDate               whereHelpernull_String
	CreationTime               whereHelpernull_String
	MaterialDocumentHeaderText whereHelpernull_String
	ReferenceDocument          whereHelpernull_String
	GoodsMovementCode          whereHelpernull_String
}{
	MaterialDocumentYear:       whereHelperstring{field: "`sap_material_document_header_data`.`MaterialDocumentYear`"},
	MaterialDocument:           whereHelperstring{field: "`sap_material_document_header_data`.`MaterialDocument`"},
	InventoryTransactionType:   whereHelpernull_String{field: "`sap_material_document_header_data`.`InventoryTransactionType`"},
	DocumentDate:               whereHelpernull_String{field: "`sap_material_document_header_data`.`DocumentDate`"},
	PostingDate:                whereHelpernull_String{field: "`sap_material_document_header_data`.`PostingDate`"},
	CreationDate:               whereHelpernull_String{field: "`sap_material_document_header_data`.`CreationDate`"},
	CreationTime:               whereHelpernull_String{field: "`sap_material_document_header_data`.`CreationTime`"},
	MaterialDocumentHeaderText: whereHelpernull_String{field: "`sap_material_document_header_data`.`MaterialDocumentHeaderText`"},
	ReferenceDocument:          whereHelpernull_String{field: "`sap_material_document_header_data`.`ReferenceDocument`"},
	GoodsMovementCode:          whereHelpernull_String{field: "`sap_material_document_header_data`.`GoodsMovementCode`"},
}

// SapMaterialDocumentHeaderDatumRels is where relationship names are stored.
var SapMaterialDocumentHeaderDatumRels = struct {
}{}

// sapMaterialDocumentHeaderDatumR is where relationships are stored.
type sapMaterialDocumentHeaderDatumR struct {
}

// NewStruct creates a new relationship struct
func (*sapMaterialDocumentHeaderDatumR) NewStruct() *sapMaterialDocumentHeaderDatumR {
	return &sapMaterialDocumentHeaderDatumR{}
}

// sapMaterialDocumentHeaderDatumL is where Load methods for each relationship are stored.
type sapMaterialDocumentHeaderDatumL struct{}

var (
	sapMaterialDocumentHeaderDatumAllColumns            = []string{"MaterialDocumentYear", "MaterialDocument", "InventoryTransactionType", "DocumentDate", "PostingDate", "CreationDate", "CreationTime", "MaterialDocumentHeaderText", "ReferenceDocument", "GoodsMovementCode"}
	sapMaterialDocumentHeaderDatumColumnsWithoutDefault = []string{"MaterialDocumentYear", "MaterialDocument", "InventoryTransactionType", "DocumentDate", "PostingDate", "CreationDate", "CreationTime", "MaterialDocumentHeaderText", "ReferenceDocument", "GoodsMovementCode"}
	sapMaterialDocumentHeaderDatumColumnsWithDefault    = []string{}
	sapMaterialDocumentHeaderDatumPrimaryKeyColumns     = []string{"MaterialDocumentYear", "MaterialDocument"}
	sapMaterialDocumentHeaderDatumGeneratedColumns      = []string{}
)

type (
	// SapMaterialDocumentHeaderDatumSlice is an alias for a slice of pointers to SapMaterialDocumentHeaderDatum.
	// This should almost always be used instead of []SapMaterialDocumentHeaderDatum.
	SapMaterialDocumentHeaderDatumSlice []*SapMaterialDocumentHeaderDatum
	// SapMaterialDocumentHeaderDatumHook is the signature for custom SapMaterialDocumentHeaderDatum hook methods
	SapMaterialDocumentHeaderDatumHook func(context.Context, boil.ContextExecutor, *SapMaterialDocumentHeaderDatum) error

	sapMaterialDocumentHeaderDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapMaterialDocumentHeaderDatumType                 = reflect.TypeOf(&SapMaterialDocumentHeaderDatum{})
	sapMaterialDocumentHeaderDatumMapping              = queries.MakeStructMapping(sapMaterialDocumentHeaderDatumType)
	sapMaterialDocumentHeaderDatumPrimaryKeyMapping, _ = queries.BindMapping(sapMaterialDocumentHeaderDatumType, sapMaterialDocumentHeaderDatumMapping, sapMaterialDocumentHeaderDatumPrimaryKeyColumns)
	sapMaterialDocumentHeaderDatumInsertCacheMut       sync.RWMutex
	sapMaterialDocumentHeaderDatumInsertCache          = make(map[string]insertCache)
	sapMaterialDocumentHeaderDatumUpdateCacheMut       sync.RWMutex
	sapMaterialDocumentHeaderDatumUpdateCache          = make(map[string]updateCache)
	sapMaterialDocumentHeaderDatumUpsertCacheMut       sync.RWMutex
	sapMaterialDocumentHeaderDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapMaterialDocumentHeaderDatumAfterSelectHooks []SapMaterialDocumentHeaderDatumHook

var sapMaterialDocumentHeaderDatumBeforeInsertHooks []SapMaterialDocumentHeaderDatumHook
var sapMaterialDocumentHeaderDatumAfterInsertHooks []SapMaterialDocumentHeaderDatumHook

var sapMaterialDocumentHeaderDatumBeforeUpdateHooks []SapMaterialDocumentHeaderDatumHook
var sapMaterialDocumentHeaderDatumAfterUpdateHooks []SapMaterialDocumentHeaderDatumHook

var sapMaterialDocumentHeaderDatumBeforeDeleteHooks []SapMaterialDocumentHeaderDatumHook
var sapMaterialDocumentHeaderDatumAfterDeleteHooks []SapMaterialDocumentHeaderDatumHook

var sapMaterialDocumentHeaderDatumBeforeUpsertHooks []SapMaterialDocumentHeaderDatumHook
var sapMaterialDocumentHeaderDatumAfterUpsertHooks []SapMaterialDocumentHeaderDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapMaterialDocumentHeaderDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapMaterialDocumentHeaderDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapMaterialDocumentHeaderDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapMaterialDocumentHeaderDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapMaterialDocumentHeaderDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapMaterialDocumentHeaderDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapMaterialDocumentHeaderDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapMaterialDocumentHeaderDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapMaterialDocumentHeaderDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialDocumentHeaderDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapMaterialDocumentHeaderDatumHook registers your hook function for all future operations.
func AddSapMaterialDocumentHeaderDatumHook(hookPoint boil.HookPoint, sapMaterialDocumentHeaderDatumHook SapMaterialDocumentHeaderDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sapMaterialDocumentHeaderDatumAfterSelectHooks = append(sapMaterialDocumentHeaderDatumAfterSelectHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.BeforeInsertHook:
		sapMaterialDocumentHeaderDatumBeforeInsertHooks = append(sapMaterialDocumentHeaderDatumBeforeInsertHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.AfterInsertHook:
		sapMaterialDocumentHeaderDatumAfterInsertHooks = append(sapMaterialDocumentHeaderDatumAfterInsertHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.BeforeUpdateHook:
		sapMaterialDocumentHeaderDatumBeforeUpdateHooks = append(sapMaterialDocumentHeaderDatumBeforeUpdateHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.AfterUpdateHook:
		sapMaterialDocumentHeaderDatumAfterUpdateHooks = append(sapMaterialDocumentHeaderDatumAfterUpdateHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.BeforeDeleteHook:
		sapMaterialDocumentHeaderDatumBeforeDeleteHooks = append(sapMaterialDocumentHeaderDatumBeforeDeleteHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.AfterDeleteHook:
		sapMaterialDocumentHeaderDatumAfterDeleteHooks = append(sapMaterialDocumentHeaderDatumAfterDeleteHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.BeforeUpsertHook:
		sapMaterialDocumentHeaderDatumBeforeUpsertHooks = append(sapMaterialDocumentHeaderDatumBeforeUpsertHooks, sapMaterialDocumentHeaderDatumHook)
	case boil.AfterUpsertHook:
		sapMaterialDocumentHeaderDatumAfterUpsertHooks = append(sapMaterialDocumentHeaderDatumAfterUpsertHooks, sapMaterialDocumentHeaderDatumHook)
	}
}

// One returns a single sapMaterialDocumentHeaderDatum record from the query.
func (q sapMaterialDocumentHeaderDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapMaterialDocumentHeaderDatum, error) {
	o := &SapMaterialDocumentHeaderDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_material_document_header_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapMaterialDocumentHeaderDatum records from the query.
func (q sapMaterialDocumentHeaderDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapMaterialDocumentHeaderDatumSlice, error) {
	var o []*SapMaterialDocumentHeaderDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapMaterialDocumentHeaderDatum slice")
	}

	if len(sapMaterialDocumentHeaderDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapMaterialDocumentHeaderDatum records in the query.
func (q sapMaterialDocumentHeaderDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_material_document_header_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapMaterialDocumentHeaderDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_material_document_header_data exists")
	}

	return count > 0, nil
}

// SapMaterialDocumentHeaderData retrieves all the records using an executor.
func SapMaterialDocumentHeaderData(mods ...qm.QueryMod) sapMaterialDocumentHeaderDatumQuery {
	mods = append(mods, qm.From("`sap_material_document_header_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`sap_material_document_header_data`.*"})
	}

	return sapMaterialDocumentHeaderDatumQuery{q}
}

// FindSapMaterialDocumentHeaderDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapMaterialDocumentHeaderDatum(ctx context.Context, exec boil.ContextExecutor, materialDocumentYear string, materialDocument string, selectCols ...string) (*SapMaterialDocumentHeaderDatum, error) {
	sapMaterialDocumentHeaderDatumObj := &SapMaterialDocumentHeaderDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_material_document_header_data` where `MaterialDocumentYear`=? AND `MaterialDocument`=?", sel,
	)

	q := queries.Raw(query, materialDocumentYear, materialDocument)

	err := q.Bind(ctx, exec, sapMaterialDocumentHeaderDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_material_document_header_data")
	}

	if err = sapMaterialDocumentHeaderDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapMaterialDocumentHeaderDatumObj, err
	}

	return sapMaterialDocumentHeaderDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapMaterialDocumentHeaderDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_material_document_header_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapMaterialDocumentHeaderDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapMaterialDocumentHeaderDatumInsertCacheMut.RLock()
	cache, cached := sapMaterialDocumentHeaderDatumInsertCache[key]
	sapMaterialDocumentHeaderDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapMaterialDocumentHeaderDatumAllColumns,
			sapMaterialDocumentHeaderDatumColumnsWithDefault,
			sapMaterialDocumentHeaderDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapMaterialDocumentHeaderDatumType, sapMaterialDocumentHeaderDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapMaterialDocumentHeaderDatumType, sapMaterialDocumentHeaderDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_material_document_header_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_material_document_header_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_material_document_header_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapMaterialDocumentHeaderDatumPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into sap_material_document_header_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.MaterialDocumentYear,
		o.MaterialDocument,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_material_document_header_data")
	}

CacheNoHooks:
	if !cached {
		sapMaterialDocumentHeaderDatumInsertCacheMut.Lock()
		sapMaterialDocumentHeaderDatumInsertCache[key] = cache
		sapMaterialDocumentHeaderDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapMaterialDocumentHeaderDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapMaterialDocumentHeaderDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapMaterialDocumentHeaderDatumUpdateCacheMut.RLock()
	cache, cached := sapMaterialDocumentHeaderDatumUpdateCache[key]
	sapMaterialDocumentHeaderDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapMaterialDocumentHeaderDatumAllColumns,
			sapMaterialDocumentHeaderDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_material_document_header_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_material_document_header_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapMaterialDocumentHeaderDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapMaterialDocumentHeaderDatumType, sapMaterialDocumentHeaderDatumMapping, append(wl, sapMaterialDocumentHeaderDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_material_document_header_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_material_document_header_data")
	}

	if !cached {
		sapMaterialDocumentHeaderDatumUpdateCacheMut.Lock()
		sapMaterialDocumentHeaderDatumUpdateCache[key] = cache
		sapMaterialDocumentHeaderDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapMaterialDocumentHeaderDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_material_document_header_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_material_document_header_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapMaterialDocumentHeaderDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaterialDocumentHeaderDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_material_document_header_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaterialDocumentHeaderDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapMaterialDocumentHeaderDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapMaterialDocumentHeaderDatum")
	}
	return rowsAff, nil
}

var mySQLSapMaterialDocumentHeaderDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapMaterialDocumentHeaderDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_material_document_header_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapMaterialDocumentHeaderDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapMaterialDocumentHeaderDatumUniqueColumns, o)

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

	sapMaterialDocumentHeaderDatumUpsertCacheMut.RLock()
	cache, cached := sapMaterialDocumentHeaderDatumUpsertCache[key]
	sapMaterialDocumentHeaderDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapMaterialDocumentHeaderDatumAllColumns,
			sapMaterialDocumentHeaderDatumColumnsWithDefault,
			sapMaterialDocumentHeaderDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sapMaterialDocumentHeaderDatumAllColumns,
			sapMaterialDocumentHeaderDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_material_document_header_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_material_document_header_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_material_document_header_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapMaterialDocumentHeaderDatumType, sapMaterialDocumentHeaderDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapMaterialDocumentHeaderDatumType, sapMaterialDocumentHeaderDatumMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for sap_material_document_header_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapMaterialDocumentHeaderDatumType, sapMaterialDocumentHeaderDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_material_document_header_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_material_document_header_data")
	}

CacheNoHooks:
	if !cached {
		sapMaterialDocumentHeaderDatumUpsertCacheMut.Lock()
		sapMaterialDocumentHeaderDatumUpsertCache[key] = cache
		sapMaterialDocumentHeaderDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapMaterialDocumentHeaderDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapMaterialDocumentHeaderDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapMaterialDocumentHeaderDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapMaterialDocumentHeaderDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_material_document_header_data` WHERE `MaterialDocumentYear`=? AND `MaterialDocument`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_material_document_header_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_material_document_header_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapMaterialDocumentHeaderDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapMaterialDocumentHeaderDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_material_document_header_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_material_document_header_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapMaterialDocumentHeaderDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapMaterialDocumentHeaderDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaterialDocumentHeaderDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_material_document_header_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaterialDocumentHeaderDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapMaterialDocumentHeaderDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_material_document_header_data")
	}

	if len(sapMaterialDocumentHeaderDatumAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SapMaterialDocumentHeaderDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapMaterialDocumentHeaderDatum(ctx, exec, o.MaterialDocumentYear, o.MaterialDocument)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapMaterialDocumentHeaderDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapMaterialDocumentHeaderDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaterialDocumentHeaderDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_material_document_header_data`.* FROM `sap_material_document_header_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaterialDocumentHeaderDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapMaterialDocumentHeaderDatumSlice")
	}

	*o = slice

	return nil
}

// SapMaterialDocumentHeaderDatumExists checks if the SapMaterialDocumentHeaderDatum row exists.
func SapMaterialDocumentHeaderDatumExists(ctx context.Context, exec boil.ContextExecutor, materialDocumentYear string, materialDocument string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_material_document_header_data` where `MaterialDocumentYear`=? AND `MaterialDocument`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, materialDocumentYear, materialDocument)
	}
	row := exec.QueryRowContext(ctx, sql, materialDocumentYear, materialDocument)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_material_document_header_data exists")
	}

	return exists, nil
}
