// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// SapSupplierInvoiceTaxDatum is an object representing the database table.
type SapSupplierInvoiceTaxDatum struct {
	SupplierInvoice          string      `boil:"SupplierInvoice" json:"SupplierInvoice" toml:"SupplierInvoice" yaml:"SupplierInvoice"`
	FiscalYear               string      `boil:"FiscalYear" json:"FiscalYear" toml:"FiscalYear" yaml:"FiscalYear"`
	TaxCode                  null.String `boil:"TaxCode" json:"TaxCode,omitempty" toml:"TaxCode" yaml:"TaxCode,omitempty"`
	DocumentCurrency         null.String `boil:"DocumentCurrency" json:"DocumentCurrency,omitempty" toml:"DocumentCurrency" yaml:"DocumentCurrency,omitempty"`
	TaxAmount                null.String `boil:"TaxAmount" json:"TaxAmount,omitempty" toml:"TaxAmount" yaml:"TaxAmount,omitempty"`
	TaxBaseAmountInTransCrcy null.String `boil:"TaxBaseAmountInTransCrcy" json:"TaxBaseAmountInTransCrcy,omitempty" toml:"TaxBaseAmountInTransCrcy" yaml:"TaxBaseAmountInTransCrcy,omitempty"`

	R *sapSupplierInvoiceTaxDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapSupplierInvoiceTaxDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapSupplierInvoiceTaxDatumColumns = struct {
	SupplierInvoice          string
	FiscalYear               string
	TaxCode                  string
	DocumentCurrency         string
	TaxAmount                string
	TaxBaseAmountInTransCrcy string
}{
	SupplierInvoice:          "SupplierInvoice",
	FiscalYear:               "FiscalYear",
	TaxCode:                  "TaxCode",
	DocumentCurrency:         "DocumentCurrency",
	TaxAmount:                "TaxAmount",
	TaxBaseAmountInTransCrcy: "TaxBaseAmountInTransCrcy",
}

var SapSupplierInvoiceTaxDatumTableColumns = struct {
	SupplierInvoice          string
	FiscalYear               string
	TaxCode                  string
	DocumentCurrency         string
	TaxAmount                string
	TaxBaseAmountInTransCrcy string
}{
	SupplierInvoice:          "sap_supplier_invoice_tax_data.SupplierInvoice",
	FiscalYear:               "sap_supplier_invoice_tax_data.FiscalYear",
	TaxCode:                  "sap_supplier_invoice_tax_data.TaxCode",
	DocumentCurrency:         "sap_supplier_invoice_tax_data.DocumentCurrency",
	TaxAmount:                "sap_supplier_invoice_tax_data.TaxAmount",
	TaxBaseAmountInTransCrcy: "sap_supplier_invoice_tax_data.TaxBaseAmountInTransCrcy",
}

// Generated where

var SapSupplierInvoiceTaxDatumWhere = struct {
	SupplierInvoice          whereHelperstring
	FiscalYear               whereHelperstring
	TaxCode                  whereHelpernull_String
	DocumentCurrency         whereHelpernull_String
	TaxAmount                whereHelpernull_String
	TaxBaseAmountInTransCrcy whereHelpernull_String
}{
	SupplierInvoice:          whereHelperstring{field: "`sap_supplier_invoice_tax_data`.`SupplierInvoice`"},
	FiscalYear:               whereHelperstring{field: "`sap_supplier_invoice_tax_data`.`FiscalYear`"},
	TaxCode:                  whereHelpernull_String{field: "`sap_supplier_invoice_tax_data`.`TaxCode`"},
	DocumentCurrency:         whereHelpernull_String{field: "`sap_supplier_invoice_tax_data`.`DocumentCurrency`"},
	TaxAmount:                whereHelpernull_String{field: "`sap_supplier_invoice_tax_data`.`TaxAmount`"},
	TaxBaseAmountInTransCrcy: whereHelpernull_String{field: "`sap_supplier_invoice_tax_data`.`TaxBaseAmountInTransCrcy`"},
}

// SapSupplierInvoiceTaxDatumRels is where relationship names are stored.
var SapSupplierInvoiceTaxDatumRels = struct {
}{}

// sapSupplierInvoiceTaxDatumR is where relationships are stored.
type sapSupplierInvoiceTaxDatumR struct {
}

// NewStruct creates a new relationship struct
func (*sapSupplierInvoiceTaxDatumR) NewStruct() *sapSupplierInvoiceTaxDatumR {
	return &sapSupplierInvoiceTaxDatumR{}
}

// sapSupplierInvoiceTaxDatumL is where Load methods for each relationship are stored.
type sapSupplierInvoiceTaxDatumL struct{}

var (
	sapSupplierInvoiceTaxDatumAllColumns            = []string{"SupplierInvoice", "FiscalYear", "TaxCode", "DocumentCurrency", "TaxAmount", "TaxBaseAmountInTransCrcy"}
	sapSupplierInvoiceTaxDatumColumnsWithoutDefault = []string{"SupplierInvoice", "FiscalYear", "TaxCode", "DocumentCurrency", "TaxAmount", "TaxBaseAmountInTransCrcy"}
	sapSupplierInvoiceTaxDatumColumnsWithDefault    = []string{}
	sapSupplierInvoiceTaxDatumPrimaryKeyColumns     = []string{"SupplierInvoice", "FiscalYear"}
	sapSupplierInvoiceTaxDatumGeneratedColumns      = []string{}
)

type (
	// SapSupplierInvoiceTaxDatumSlice is an alias for a slice of pointers to SapSupplierInvoiceTaxDatum.
	// This should almost always be used instead of []SapSupplierInvoiceTaxDatum.
	SapSupplierInvoiceTaxDatumSlice []*SapSupplierInvoiceTaxDatum
	// SapSupplierInvoiceTaxDatumHook is the signature for custom SapSupplierInvoiceTaxDatum hook methods
	SapSupplierInvoiceTaxDatumHook func(context.Context, boil.ContextExecutor, *SapSupplierInvoiceTaxDatum) error

	sapSupplierInvoiceTaxDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapSupplierInvoiceTaxDatumType                 = reflect.TypeOf(&SapSupplierInvoiceTaxDatum{})
	sapSupplierInvoiceTaxDatumMapping              = queries.MakeStructMapping(sapSupplierInvoiceTaxDatumType)
	sapSupplierInvoiceTaxDatumPrimaryKeyMapping, _ = queries.BindMapping(sapSupplierInvoiceTaxDatumType, sapSupplierInvoiceTaxDatumMapping, sapSupplierInvoiceTaxDatumPrimaryKeyColumns)
	sapSupplierInvoiceTaxDatumInsertCacheMut       sync.RWMutex
	sapSupplierInvoiceTaxDatumInsertCache          = make(map[string]insertCache)
	sapSupplierInvoiceTaxDatumUpdateCacheMut       sync.RWMutex
	sapSupplierInvoiceTaxDatumUpdateCache          = make(map[string]updateCache)
	sapSupplierInvoiceTaxDatumUpsertCacheMut       sync.RWMutex
	sapSupplierInvoiceTaxDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapSupplierInvoiceTaxDatumAfterSelectHooks []SapSupplierInvoiceTaxDatumHook

var sapSupplierInvoiceTaxDatumBeforeInsertHooks []SapSupplierInvoiceTaxDatumHook
var sapSupplierInvoiceTaxDatumAfterInsertHooks []SapSupplierInvoiceTaxDatumHook

var sapSupplierInvoiceTaxDatumBeforeUpdateHooks []SapSupplierInvoiceTaxDatumHook
var sapSupplierInvoiceTaxDatumAfterUpdateHooks []SapSupplierInvoiceTaxDatumHook

var sapSupplierInvoiceTaxDatumBeforeDeleteHooks []SapSupplierInvoiceTaxDatumHook
var sapSupplierInvoiceTaxDatumAfterDeleteHooks []SapSupplierInvoiceTaxDatumHook

var sapSupplierInvoiceTaxDatumBeforeUpsertHooks []SapSupplierInvoiceTaxDatumHook
var sapSupplierInvoiceTaxDatumAfterUpsertHooks []SapSupplierInvoiceTaxDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapSupplierInvoiceTaxDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapSupplierInvoiceTaxDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapSupplierInvoiceTaxDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapSupplierInvoiceTaxDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapSupplierInvoiceTaxDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapSupplierInvoiceTaxDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapSupplierInvoiceTaxDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapSupplierInvoiceTaxDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapSupplierInvoiceTaxDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapSupplierInvoiceTaxDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapSupplierInvoiceTaxDatumHook registers your hook function for all future operations.
func AddSapSupplierInvoiceTaxDatumHook(hookPoint boil.HookPoint, sapSupplierInvoiceTaxDatumHook SapSupplierInvoiceTaxDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sapSupplierInvoiceTaxDatumAfterSelectHooks = append(sapSupplierInvoiceTaxDatumAfterSelectHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.BeforeInsertHook:
		sapSupplierInvoiceTaxDatumBeforeInsertHooks = append(sapSupplierInvoiceTaxDatumBeforeInsertHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.AfterInsertHook:
		sapSupplierInvoiceTaxDatumAfterInsertHooks = append(sapSupplierInvoiceTaxDatumAfterInsertHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.BeforeUpdateHook:
		sapSupplierInvoiceTaxDatumBeforeUpdateHooks = append(sapSupplierInvoiceTaxDatumBeforeUpdateHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.AfterUpdateHook:
		sapSupplierInvoiceTaxDatumAfterUpdateHooks = append(sapSupplierInvoiceTaxDatumAfterUpdateHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.BeforeDeleteHook:
		sapSupplierInvoiceTaxDatumBeforeDeleteHooks = append(sapSupplierInvoiceTaxDatumBeforeDeleteHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.AfterDeleteHook:
		sapSupplierInvoiceTaxDatumAfterDeleteHooks = append(sapSupplierInvoiceTaxDatumAfterDeleteHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.BeforeUpsertHook:
		sapSupplierInvoiceTaxDatumBeforeUpsertHooks = append(sapSupplierInvoiceTaxDatumBeforeUpsertHooks, sapSupplierInvoiceTaxDatumHook)
	case boil.AfterUpsertHook:
		sapSupplierInvoiceTaxDatumAfterUpsertHooks = append(sapSupplierInvoiceTaxDatumAfterUpsertHooks, sapSupplierInvoiceTaxDatumHook)
	}
}

// One returns a single sapSupplierInvoiceTaxDatum record from the query.
func (q sapSupplierInvoiceTaxDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapSupplierInvoiceTaxDatum, error) {
	o := &SapSupplierInvoiceTaxDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_supplier_invoice_tax_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapSupplierInvoiceTaxDatum records from the query.
func (q sapSupplierInvoiceTaxDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapSupplierInvoiceTaxDatumSlice, error) {
	var o []*SapSupplierInvoiceTaxDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapSupplierInvoiceTaxDatum slice")
	}

	if len(sapSupplierInvoiceTaxDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapSupplierInvoiceTaxDatum records in the query.
func (q sapSupplierInvoiceTaxDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_supplier_invoice_tax_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapSupplierInvoiceTaxDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_supplier_invoice_tax_data exists")
	}

	return count > 0, nil
}

// SapSupplierInvoiceTaxData retrieves all the records using an executor.
func SapSupplierInvoiceTaxData(mods ...qm.QueryMod) sapSupplierInvoiceTaxDatumQuery {
	mods = append(mods, qm.From("`sap_supplier_invoice_tax_data`"))
	return sapSupplierInvoiceTaxDatumQuery{NewQuery(mods...)}
}

// FindSapSupplierInvoiceTaxDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapSupplierInvoiceTaxDatum(ctx context.Context, exec boil.ContextExecutor, supplierInvoice string, fiscalYear string, selectCols ...string) (*SapSupplierInvoiceTaxDatum, error) {
	sapSupplierInvoiceTaxDatumObj := &SapSupplierInvoiceTaxDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_supplier_invoice_tax_data` where `SupplierInvoice`=? AND `FiscalYear`=?", sel,
	)

	q := queries.Raw(query, supplierInvoice, fiscalYear)

	err := q.Bind(ctx, exec, sapSupplierInvoiceTaxDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_supplier_invoice_tax_data")
	}

	if err = sapSupplierInvoiceTaxDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapSupplierInvoiceTaxDatumObj, err
	}

	return sapSupplierInvoiceTaxDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapSupplierInvoiceTaxDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_supplier_invoice_tax_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapSupplierInvoiceTaxDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapSupplierInvoiceTaxDatumInsertCacheMut.RLock()
	cache, cached := sapSupplierInvoiceTaxDatumInsertCache[key]
	sapSupplierInvoiceTaxDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapSupplierInvoiceTaxDatumAllColumns,
			sapSupplierInvoiceTaxDatumColumnsWithDefault,
			sapSupplierInvoiceTaxDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapSupplierInvoiceTaxDatumType, sapSupplierInvoiceTaxDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapSupplierInvoiceTaxDatumType, sapSupplierInvoiceTaxDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_supplier_invoice_tax_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_supplier_invoice_tax_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_supplier_invoice_tax_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapSupplierInvoiceTaxDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_supplier_invoice_tax_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.SupplierInvoice,
		o.FiscalYear,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_supplier_invoice_tax_data")
	}

CacheNoHooks:
	if !cached {
		sapSupplierInvoiceTaxDatumInsertCacheMut.Lock()
		sapSupplierInvoiceTaxDatumInsertCache[key] = cache
		sapSupplierInvoiceTaxDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapSupplierInvoiceTaxDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapSupplierInvoiceTaxDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapSupplierInvoiceTaxDatumUpdateCacheMut.RLock()
	cache, cached := sapSupplierInvoiceTaxDatumUpdateCache[key]
	sapSupplierInvoiceTaxDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapSupplierInvoiceTaxDatumAllColumns,
			sapSupplierInvoiceTaxDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_supplier_invoice_tax_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_supplier_invoice_tax_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapSupplierInvoiceTaxDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapSupplierInvoiceTaxDatumType, sapSupplierInvoiceTaxDatumMapping, append(wl, sapSupplierInvoiceTaxDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_supplier_invoice_tax_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_supplier_invoice_tax_data")
	}

	if !cached {
		sapSupplierInvoiceTaxDatumUpdateCacheMut.Lock()
		sapSupplierInvoiceTaxDatumUpdateCache[key] = cache
		sapSupplierInvoiceTaxDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapSupplierInvoiceTaxDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_supplier_invoice_tax_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_supplier_invoice_tax_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapSupplierInvoiceTaxDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapSupplierInvoiceTaxDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_supplier_invoice_tax_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapSupplierInvoiceTaxDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapSupplierInvoiceTaxDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapSupplierInvoiceTaxDatum")
	}
	return rowsAff, nil
}

var mySQLSapSupplierInvoiceTaxDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapSupplierInvoiceTaxDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_supplier_invoice_tax_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapSupplierInvoiceTaxDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapSupplierInvoiceTaxDatumUniqueColumns, o)

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

	sapSupplierInvoiceTaxDatumUpsertCacheMut.RLock()
	cache, cached := sapSupplierInvoiceTaxDatumUpsertCache[key]
	sapSupplierInvoiceTaxDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapSupplierInvoiceTaxDatumAllColumns,
			sapSupplierInvoiceTaxDatumColumnsWithDefault,
			sapSupplierInvoiceTaxDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sapSupplierInvoiceTaxDatumAllColumns,
			sapSupplierInvoiceTaxDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_supplier_invoice_tax_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_supplier_invoice_tax_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_supplier_invoice_tax_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapSupplierInvoiceTaxDatumType, sapSupplierInvoiceTaxDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapSupplierInvoiceTaxDatumType, sapSupplierInvoiceTaxDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_supplier_invoice_tax_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapSupplierInvoiceTaxDatumType, sapSupplierInvoiceTaxDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_supplier_invoice_tax_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_supplier_invoice_tax_data")
	}

CacheNoHooks:
	if !cached {
		sapSupplierInvoiceTaxDatumUpsertCacheMut.Lock()
		sapSupplierInvoiceTaxDatumUpsertCache[key] = cache
		sapSupplierInvoiceTaxDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapSupplierInvoiceTaxDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapSupplierInvoiceTaxDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapSupplierInvoiceTaxDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapSupplierInvoiceTaxDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_supplier_invoice_tax_data` WHERE `SupplierInvoice`=? AND `FiscalYear`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_supplier_invoice_tax_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_supplier_invoice_tax_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapSupplierInvoiceTaxDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapSupplierInvoiceTaxDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_supplier_invoice_tax_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_supplier_invoice_tax_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapSupplierInvoiceTaxDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapSupplierInvoiceTaxDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapSupplierInvoiceTaxDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_supplier_invoice_tax_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapSupplierInvoiceTaxDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapSupplierInvoiceTaxDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_supplier_invoice_tax_data")
	}

	if len(sapSupplierInvoiceTaxDatumAfterDeleteHooks) != 0 {
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
func (o *SapSupplierInvoiceTaxDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapSupplierInvoiceTaxDatum(ctx, exec, o.SupplierInvoice, o.FiscalYear)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapSupplierInvoiceTaxDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapSupplierInvoiceTaxDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapSupplierInvoiceTaxDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_supplier_invoice_tax_data`.* FROM `sap_supplier_invoice_tax_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapSupplierInvoiceTaxDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapSupplierInvoiceTaxDatumSlice")
	}

	*o = slice

	return nil
}

// SapSupplierInvoiceTaxDatumExists checks if the SapSupplierInvoiceTaxDatum row exists.
func SapSupplierInvoiceTaxDatumExists(ctx context.Context, exec boil.ContextExecutor, supplierInvoice string, fiscalYear string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_supplier_invoice_tax_data` where `SupplierInvoice`=? AND `FiscalYear`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, supplierInvoice, fiscalYear)
	}
	row := exec.QueryRowContext(ctx, sql, supplierInvoice, fiscalYear)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_supplier_invoice_tax_data exists")
	}

	return exists, nil
}
