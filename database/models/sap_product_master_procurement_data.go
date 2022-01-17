// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// SapProductMasterProcurementDatum is an object representing the database table.
type SapProductMasterProcurementDatum struct {
	Product                     string      `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
	Plant                       string      `boil:"Plant" json:"Plant" toml:"Plant" yaml:"Plant"`
	PurchaseOrderQuantityUnit   null.String `boil:"PurchaseOrderQuantityUnit" json:"PurchaseOrderQuantityUnit,omitempty" toml:"PurchaseOrderQuantityUnit" yaml:"PurchaseOrderQuantityUnit,omitempty"`
	IsAutoPurOrdCreationAllowed null.Bool   `boil:"IsAutoPurOrdCreationAllowed" json:"IsAutoPurOrdCreationAllowed,omitempty" toml:"IsAutoPurOrdCreationAllowed" yaml:"IsAutoPurOrdCreationAllowed,omitempty"`
	IsSourceListRequired        null.Bool   `boil:"IsSourceListRequired" json:"IsSourceListRequired,omitempty" toml:"IsSourceListRequired" yaml:"IsSourceListRequired,omitempty"`

	R *sapProductMasterProcurementDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapProductMasterProcurementDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapProductMasterProcurementDatumColumns = struct {
	Product                     string
	Plant                       string
	PurchaseOrderQuantityUnit   string
	IsAutoPurOrdCreationAllowed string
	IsSourceListRequired        string
}{
	Product:                     "Product",
	Plant:                       "Plant",
	PurchaseOrderQuantityUnit:   "PurchaseOrderQuantityUnit",
	IsAutoPurOrdCreationAllowed: "IsAutoPurOrdCreationAllowed",
	IsSourceListRequired:        "IsSourceListRequired",
}

var SapProductMasterProcurementDatumTableColumns = struct {
	Product                     string
	Plant                       string
	PurchaseOrderQuantityUnit   string
	IsAutoPurOrdCreationAllowed string
	IsSourceListRequired        string
}{
	Product:                     "sap_product_master_procurement_data.Product",
	Plant:                       "sap_product_master_procurement_data.Plant",
	PurchaseOrderQuantityUnit:   "sap_product_master_procurement_data.PurchaseOrderQuantityUnit",
	IsAutoPurOrdCreationAllowed: "sap_product_master_procurement_data.IsAutoPurOrdCreationAllowed",
	IsSourceListRequired:        "sap_product_master_procurement_data.IsSourceListRequired",
}

// Generated where

var SapProductMasterProcurementDatumWhere = struct {
	Product                     whereHelperstring
	Plant                       whereHelperstring
	PurchaseOrderQuantityUnit   whereHelpernull_String
	IsAutoPurOrdCreationAllowed whereHelpernull_Bool
	IsSourceListRequired        whereHelpernull_Bool
}{
	Product:                     whereHelperstring{field: "`sap_product_master_procurement_data`.`Product`"},
	Plant:                       whereHelperstring{field: "`sap_product_master_procurement_data`.`Plant`"},
	PurchaseOrderQuantityUnit:   whereHelpernull_String{field: "`sap_product_master_procurement_data`.`PurchaseOrderQuantityUnit`"},
	IsAutoPurOrdCreationAllowed: whereHelpernull_Bool{field: "`sap_product_master_procurement_data`.`IsAutoPurOrdCreationAllowed`"},
	IsSourceListRequired:        whereHelpernull_Bool{field: "`sap_product_master_procurement_data`.`IsSourceListRequired`"},
}

// SapProductMasterProcurementDatumRels is where relationship names are stored.
var SapProductMasterProcurementDatumRels = struct {
	ProductSapProductMasterGeneralDatum string
}{
	ProductSapProductMasterGeneralDatum: "ProductSapProductMasterGeneralDatum",
}

// sapProductMasterProcurementDatumR is where relationships are stored.
type sapProductMasterProcurementDatumR struct {
	ProductSapProductMasterGeneralDatum *SapProductMasterGeneralDatum `boil:"ProductSapProductMasterGeneralDatum" json:"ProductSapProductMasterGeneralDatum" toml:"ProductSapProductMasterGeneralDatum" yaml:"ProductSapProductMasterGeneralDatum"`
}

// NewStruct creates a new relationship struct
func (*sapProductMasterProcurementDatumR) NewStruct() *sapProductMasterProcurementDatumR {
	return &sapProductMasterProcurementDatumR{}
}

// sapProductMasterProcurementDatumL is where Load methods for each relationship are stored.
type sapProductMasterProcurementDatumL struct{}

var (
	sapProductMasterProcurementDatumAllColumns            = []string{"Product", "Plant", "PurchaseOrderQuantityUnit", "IsAutoPurOrdCreationAllowed", "IsSourceListRequired"}
	sapProductMasterProcurementDatumColumnsWithoutDefault = []string{"Product", "Plant", "PurchaseOrderQuantityUnit", "IsAutoPurOrdCreationAllowed", "IsSourceListRequired"}
	sapProductMasterProcurementDatumColumnsWithDefault    = []string{}
	sapProductMasterProcurementDatumPrimaryKeyColumns     = []string{"Product", "Plant"}
)

type (
	// SapProductMasterProcurementDatumSlice is an alias for a slice of pointers to SapProductMasterProcurementDatum.
	// This should almost always be used instead of []SapProductMasterProcurementDatum.
	SapProductMasterProcurementDatumSlice []*SapProductMasterProcurementDatum
	// SapProductMasterProcurementDatumHook is the signature for custom SapProductMasterProcurementDatum hook methods
	SapProductMasterProcurementDatumHook func(context.Context, boil.ContextExecutor, *SapProductMasterProcurementDatum) error

	sapProductMasterProcurementDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapProductMasterProcurementDatumType                 = reflect.TypeOf(&SapProductMasterProcurementDatum{})
	sapProductMasterProcurementDatumMapping              = queries.MakeStructMapping(sapProductMasterProcurementDatumType)
	sapProductMasterProcurementDatumPrimaryKeyMapping, _ = queries.BindMapping(sapProductMasterProcurementDatumType, sapProductMasterProcurementDatumMapping, sapProductMasterProcurementDatumPrimaryKeyColumns)
	sapProductMasterProcurementDatumInsertCacheMut       sync.RWMutex
	sapProductMasterProcurementDatumInsertCache          = make(map[string]insertCache)
	sapProductMasterProcurementDatumUpdateCacheMut       sync.RWMutex
	sapProductMasterProcurementDatumUpdateCache          = make(map[string]updateCache)
	sapProductMasterProcurementDatumUpsertCacheMut       sync.RWMutex
	sapProductMasterProcurementDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapProductMasterProcurementDatumBeforeInsertHooks []SapProductMasterProcurementDatumHook
var sapProductMasterProcurementDatumBeforeUpdateHooks []SapProductMasterProcurementDatumHook
var sapProductMasterProcurementDatumBeforeDeleteHooks []SapProductMasterProcurementDatumHook
var sapProductMasterProcurementDatumBeforeUpsertHooks []SapProductMasterProcurementDatumHook

var sapProductMasterProcurementDatumAfterInsertHooks []SapProductMasterProcurementDatumHook
var sapProductMasterProcurementDatumAfterSelectHooks []SapProductMasterProcurementDatumHook
var sapProductMasterProcurementDatumAfterUpdateHooks []SapProductMasterProcurementDatumHook
var sapProductMasterProcurementDatumAfterDeleteHooks []SapProductMasterProcurementDatumHook
var sapProductMasterProcurementDatumAfterUpsertHooks []SapProductMasterProcurementDatumHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapProductMasterProcurementDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapProductMasterProcurementDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapProductMasterProcurementDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapProductMasterProcurementDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapProductMasterProcurementDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapProductMasterProcurementDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapProductMasterProcurementDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapProductMasterProcurementDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapProductMasterProcurementDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapProductMasterProcurementDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapProductMasterProcurementDatumHook registers your hook function for all future operations.
func AddSapProductMasterProcurementDatumHook(hookPoint boil.HookPoint, sapProductMasterProcurementDatumHook SapProductMasterProcurementDatumHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		sapProductMasterProcurementDatumBeforeInsertHooks = append(sapProductMasterProcurementDatumBeforeInsertHooks, sapProductMasterProcurementDatumHook)
	case boil.BeforeUpdateHook:
		sapProductMasterProcurementDatumBeforeUpdateHooks = append(sapProductMasterProcurementDatumBeforeUpdateHooks, sapProductMasterProcurementDatumHook)
	case boil.BeforeDeleteHook:
		sapProductMasterProcurementDatumBeforeDeleteHooks = append(sapProductMasterProcurementDatumBeforeDeleteHooks, sapProductMasterProcurementDatumHook)
	case boil.BeforeUpsertHook:
		sapProductMasterProcurementDatumBeforeUpsertHooks = append(sapProductMasterProcurementDatumBeforeUpsertHooks, sapProductMasterProcurementDatumHook)
	case boil.AfterInsertHook:
		sapProductMasterProcurementDatumAfterInsertHooks = append(sapProductMasterProcurementDatumAfterInsertHooks, sapProductMasterProcurementDatumHook)
	case boil.AfterSelectHook:
		sapProductMasterProcurementDatumAfterSelectHooks = append(sapProductMasterProcurementDatumAfterSelectHooks, sapProductMasterProcurementDatumHook)
	case boil.AfterUpdateHook:
		sapProductMasterProcurementDatumAfterUpdateHooks = append(sapProductMasterProcurementDatumAfterUpdateHooks, sapProductMasterProcurementDatumHook)
	case boil.AfterDeleteHook:
		sapProductMasterProcurementDatumAfterDeleteHooks = append(sapProductMasterProcurementDatumAfterDeleteHooks, sapProductMasterProcurementDatumHook)
	case boil.AfterUpsertHook:
		sapProductMasterProcurementDatumAfterUpsertHooks = append(sapProductMasterProcurementDatumAfterUpsertHooks, sapProductMasterProcurementDatumHook)
	}
}

// One returns a single sapProductMasterProcurementDatum record from the query.
func (q sapProductMasterProcurementDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapProductMasterProcurementDatum, error) {
	o := &SapProductMasterProcurementDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_product_master_procurement_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapProductMasterProcurementDatum records from the query.
func (q sapProductMasterProcurementDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapProductMasterProcurementDatumSlice, error) {
	var o []*SapProductMasterProcurementDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapProductMasterProcurementDatum slice")
	}

	if len(sapProductMasterProcurementDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapProductMasterProcurementDatum records in the query.
func (q sapProductMasterProcurementDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_product_master_procurement_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapProductMasterProcurementDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_product_master_procurement_data exists")
	}

	return count > 0, nil
}

// ProductSapProductMasterGeneralDatum pointed to by the foreign key.
func (o *SapProductMasterProcurementDatum) ProductSapProductMasterGeneralDatum(mods ...qm.QueryMod) sapProductMasterGeneralDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`Product` = ?", o.Product),
	}

	queryMods = append(queryMods, mods...)

	query := SapProductMasterGeneralData(queryMods...)
	queries.SetFrom(query.Query, "`sap_product_master_general_data`")

	return query
}

// LoadProductSapProductMasterGeneralDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (sapProductMasterProcurementDatumL) LoadProductSapProductMasterGeneralDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSapProductMasterProcurementDatum interface{}, mods queries.Applicator) error {
	var slice []*SapProductMasterProcurementDatum
	var object *SapProductMasterProcurementDatum

	if singular {
		object = maybeSapProductMasterProcurementDatum.(*SapProductMasterProcurementDatum)
	} else {
		slice = *maybeSapProductMasterProcurementDatum.(*[]*SapProductMasterProcurementDatum)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &sapProductMasterProcurementDatumR{}
		}
		args = append(args, object.Product)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &sapProductMasterProcurementDatumR{}
			}

			for _, a := range args {
				if a == obj.Product {
					continue Outer
				}
			}

			args = append(args, obj.Product)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`sap_product_master_general_data`),
		qm.WhereIn(`sap_product_master_general_data.Product in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load SapProductMasterGeneralDatum")
	}

	var resultSlice []*SapProductMasterGeneralDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice SapProductMasterGeneralDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for sap_product_master_general_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sap_product_master_general_data")
	}

	if len(sapProductMasterProcurementDatumAfterSelectHooks) != 0 {
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
		object.R.ProductSapProductMasterGeneralDatum = foreign
		if foreign.R == nil {
			foreign.R = &sapProductMasterGeneralDatumR{}
		}
		foreign.R.ProductSapProductMasterProcurementData = append(foreign.R.ProductSapProductMasterProcurementData, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Product == foreign.Product {
				local.R.ProductSapProductMasterGeneralDatum = foreign
				if foreign.R == nil {
					foreign.R = &sapProductMasterGeneralDatumR{}
				}
				foreign.R.ProductSapProductMasterProcurementData = append(foreign.R.ProductSapProductMasterProcurementData, local)
				break
			}
		}
	}

	return nil
}

// SetProductSapProductMasterGeneralDatum of the sapProductMasterProcurementDatum to the related item.
// Sets o.R.ProductSapProductMasterGeneralDatum to related.
// Adds o to related.R.ProductSapProductMasterProcurementData.
func (o *SapProductMasterProcurementDatum) SetProductSapProductMasterGeneralDatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SapProductMasterGeneralDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `sap_product_master_procurement_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"Product"}),
		strmangle.WhereClause("`", "`", 0, sapProductMasterProcurementDatumPrimaryKeyColumns),
	)
	values := []interface{}{related.Product, o.Product, o.Plant}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Product = related.Product
	if o.R == nil {
		o.R = &sapProductMasterProcurementDatumR{
			ProductSapProductMasterGeneralDatum: related,
		}
	} else {
		o.R.ProductSapProductMasterGeneralDatum = related
	}

	if related.R == nil {
		related.R = &sapProductMasterGeneralDatumR{
			ProductSapProductMasterProcurementData: SapProductMasterProcurementDatumSlice{o},
		}
	} else {
		related.R.ProductSapProductMasterProcurementData = append(related.R.ProductSapProductMasterProcurementData, o)
	}

	return nil
}

// SapProductMasterProcurementData retrieves all the records using an executor.
func SapProductMasterProcurementData(mods ...qm.QueryMod) sapProductMasterProcurementDatumQuery {
	mods = append(mods, qm.From("`sap_product_master_procurement_data`"))
	return sapProductMasterProcurementDatumQuery{NewQuery(mods...)}
}

// FindSapProductMasterProcurementDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapProductMasterProcurementDatum(ctx context.Context, exec boil.ContextExecutor, product string, plant string, selectCols ...string) (*SapProductMasterProcurementDatum, error) {
	sapProductMasterProcurementDatumObj := &SapProductMasterProcurementDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_product_master_procurement_data` where `Product`=? AND `Plant`=?", sel,
	)

	q := queries.Raw(query, product, plant)

	err := q.Bind(ctx, exec, sapProductMasterProcurementDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_product_master_procurement_data")
	}

	if err = sapProductMasterProcurementDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapProductMasterProcurementDatumObj, err
	}

	return sapProductMasterProcurementDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapProductMasterProcurementDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_product_master_procurement_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapProductMasterProcurementDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapProductMasterProcurementDatumInsertCacheMut.RLock()
	cache, cached := sapProductMasterProcurementDatumInsertCache[key]
	sapProductMasterProcurementDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapProductMasterProcurementDatumAllColumns,
			sapProductMasterProcurementDatumColumnsWithDefault,
			sapProductMasterProcurementDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapProductMasterProcurementDatumType, sapProductMasterProcurementDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapProductMasterProcurementDatumType, sapProductMasterProcurementDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_product_master_procurement_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_product_master_procurement_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_product_master_procurement_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapProductMasterProcurementDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_product_master_procurement_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Product,
		o.Plant,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_product_master_procurement_data")
	}

CacheNoHooks:
	if !cached {
		sapProductMasterProcurementDatumInsertCacheMut.Lock()
		sapProductMasterProcurementDatumInsertCache[key] = cache
		sapProductMasterProcurementDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapProductMasterProcurementDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapProductMasterProcurementDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapProductMasterProcurementDatumUpdateCacheMut.RLock()
	cache, cached := sapProductMasterProcurementDatumUpdateCache[key]
	sapProductMasterProcurementDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapProductMasterProcurementDatumAllColumns,
			sapProductMasterProcurementDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_product_master_procurement_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_product_master_procurement_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapProductMasterProcurementDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapProductMasterProcurementDatumType, sapProductMasterProcurementDatumMapping, append(wl, sapProductMasterProcurementDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_product_master_procurement_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_product_master_procurement_data")
	}

	if !cached {
		sapProductMasterProcurementDatumUpdateCacheMut.Lock()
		sapProductMasterProcurementDatumUpdateCache[key] = cache
		sapProductMasterProcurementDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapProductMasterProcurementDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_product_master_procurement_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_product_master_procurement_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapProductMasterProcurementDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapProductMasterProcurementDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_product_master_procurement_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapProductMasterProcurementDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapProductMasterProcurementDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapProductMasterProcurementDatum")
	}
	return rowsAff, nil
}

var mySQLSapProductMasterProcurementDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapProductMasterProcurementDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_product_master_procurement_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapProductMasterProcurementDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapProductMasterProcurementDatumUniqueColumns, o)

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

	sapProductMasterProcurementDatumUpsertCacheMut.RLock()
	cache, cached := sapProductMasterProcurementDatumUpsertCache[key]
	sapProductMasterProcurementDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapProductMasterProcurementDatumAllColumns,
			sapProductMasterProcurementDatumColumnsWithDefault,
			sapProductMasterProcurementDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			sapProductMasterProcurementDatumAllColumns,
			sapProductMasterProcurementDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_product_master_procurement_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_product_master_procurement_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_product_master_procurement_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapProductMasterProcurementDatumType, sapProductMasterProcurementDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapProductMasterProcurementDatumType, sapProductMasterProcurementDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_product_master_procurement_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapProductMasterProcurementDatumType, sapProductMasterProcurementDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_product_master_procurement_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_product_master_procurement_data")
	}

CacheNoHooks:
	if !cached {
		sapProductMasterProcurementDatumUpsertCacheMut.Lock()
		sapProductMasterProcurementDatumUpsertCache[key] = cache
		sapProductMasterProcurementDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapProductMasterProcurementDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapProductMasterProcurementDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapProductMasterProcurementDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapProductMasterProcurementDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_product_master_procurement_data` WHERE `Product`=? AND `Plant`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_product_master_procurement_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_product_master_procurement_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapProductMasterProcurementDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapProductMasterProcurementDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_product_master_procurement_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_product_master_procurement_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapProductMasterProcurementDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapProductMasterProcurementDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapProductMasterProcurementDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_product_master_procurement_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapProductMasterProcurementDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapProductMasterProcurementDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_product_master_procurement_data")
	}

	if len(sapProductMasterProcurementDatumAfterDeleteHooks) != 0 {
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
func (o *SapProductMasterProcurementDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapProductMasterProcurementDatum(ctx, exec, o.Product, o.Plant)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapProductMasterProcurementDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapProductMasterProcurementDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapProductMasterProcurementDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_product_master_procurement_data`.* FROM `sap_product_master_procurement_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapProductMasterProcurementDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapProductMasterProcurementDatumSlice")
	}

	*o = slice

	return nil
}

// SapProductMasterProcurementDatumExists checks if the SapProductMasterProcurementDatum row exists.
func SapProductMasterProcurementDatumExists(ctx context.Context, exec boil.ContextExecutor, product string, plant string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_product_master_procurement_data` where `Product`=? AND `Plant`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, product, plant)
	}
	row := exec.QueryRowContext(ctx, sql, product, plant)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_product_master_procurement_data exists")
	}

	return exists, nil
}