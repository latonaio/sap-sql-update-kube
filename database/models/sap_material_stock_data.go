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

// SapMaterialStockDatum is an object representing the database table.
type SapMaterialStockDatum struct {
	Material                     string      `boil:"Material" json:"Material" toml:"Material" yaml:"Material"`
	Plant                        string      `boil:"Plant" json:"Plant" toml:"Plant" yaml:"Plant"`
	StorageLocation              string      `boil:"StorageLocation" json:"StorageLocation" toml:"StorageLocation" yaml:"StorageLocation"`
	Batch                        string      `boil:"Batch" json:"Batch" toml:"Batch" yaml:"Batch"`
	Supplier                     string      `boil:"Supplier" json:"Supplier" toml:"Supplier" yaml:"Supplier"`
	Customer                     string      `boil:"Customer" json:"Customer" toml:"Customer" yaml:"Customer"`
	WBSElementInternalID         string      `boil:"WBSElementInternalID" json:"WBSElementInternalID" toml:"WBSElementInternalID" yaml:"WBSElementInternalID"`
	SDDocument                   string      `boil:"SDDocument" json:"SDDocument" toml:"SDDocument" yaml:"SDDocument"`
	SDDocumentItem               string      `boil:"SDDocumentItem" json:"SDDocumentItem" toml:"SDDocumentItem" yaml:"SDDocumentItem"`
	InventorySpecialStockType    string      `boil:"InventorySpecialStockType" json:"InventorySpecialStockType" toml:"InventorySpecialStockType" yaml:"InventorySpecialStockType"`
	InventoryStockType           string      `boil:"InventoryStockType" json:"InventoryStockType" toml:"InventoryStockType" yaml:"InventoryStockType"`
	MaterialBaseUnit             null.String `boil:"MaterialBaseUnit" json:"MaterialBaseUnit,omitempty" toml:"MaterialBaseUnit" yaml:"MaterialBaseUnit,omitempty"`
	MatlWrhsStkQtyInMatlBaseUnit null.String `boil:"MatlWrhsStkQtyInMatlBaseUnit" json:"MatlWrhsStkQtyInMatlBaseUnit,omitempty" toml:"MatlWrhsStkQtyInMatlBaseUnit" yaml:"MatlWrhsStkQtyInMatlBaseUnit,omitempty"`

	R *sapMaterialStockDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapMaterialStockDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapMaterialStockDatumColumns = struct {
	Material                     string
	Plant                        string
	StorageLocation              string
	Batch                        string
	Supplier                     string
	Customer                     string
	WBSElementInternalID         string
	SDDocument                   string
	SDDocumentItem               string
	InventorySpecialStockType    string
	InventoryStockType           string
	MaterialBaseUnit             string
	MatlWrhsStkQtyInMatlBaseUnit string
}{
	Material:                     "Material",
	Plant:                        "Plant",
	StorageLocation:              "StorageLocation",
	Batch:                        "Batch",
	Supplier:                     "Supplier",
	Customer:                     "Customer",
	WBSElementInternalID:         "WBSElementInternalID",
	SDDocument:                   "SDDocument",
	SDDocumentItem:               "SDDocumentItem",
	InventorySpecialStockType:    "InventorySpecialStockType",
	InventoryStockType:           "InventoryStockType",
	MaterialBaseUnit:             "MaterialBaseUnit",
	MatlWrhsStkQtyInMatlBaseUnit: "MatlWrhsStkQtyInMatlBaseUnit",
}

var SapMaterialStockDatumTableColumns = struct {
	Material                     string
	Plant                        string
	StorageLocation              string
	Batch                        string
	Supplier                     string
	Customer                     string
	WBSElementInternalID         string
	SDDocument                   string
	SDDocumentItem               string
	InventorySpecialStockType    string
	InventoryStockType           string
	MaterialBaseUnit             string
	MatlWrhsStkQtyInMatlBaseUnit string
}{
	Material:                     "sap_material_stock_data.Material",
	Plant:                        "sap_material_stock_data.Plant",
	StorageLocation:              "sap_material_stock_data.StorageLocation",
	Batch:                        "sap_material_stock_data.Batch",
	Supplier:                     "sap_material_stock_data.Supplier",
	Customer:                     "sap_material_stock_data.Customer",
	WBSElementInternalID:         "sap_material_stock_data.WBSElementInternalID",
	SDDocument:                   "sap_material_stock_data.SDDocument",
	SDDocumentItem:               "sap_material_stock_data.SDDocumentItem",
	InventorySpecialStockType:    "sap_material_stock_data.InventorySpecialStockType",
	InventoryStockType:           "sap_material_stock_data.InventoryStockType",
	MaterialBaseUnit:             "sap_material_stock_data.MaterialBaseUnit",
	MatlWrhsStkQtyInMatlBaseUnit: "sap_material_stock_data.MatlWrhsStkQtyInMatlBaseUnit",
}

// Generated where

var SapMaterialStockDatumWhere = struct {
	Material                     whereHelperstring
	Plant                        whereHelperstring
	StorageLocation              whereHelperstring
	Batch                        whereHelperstring
	Supplier                     whereHelperstring
	Customer                     whereHelperstring
	WBSElementInternalID         whereHelperstring
	SDDocument                   whereHelperstring
	SDDocumentItem               whereHelperstring
	InventorySpecialStockType    whereHelperstring
	InventoryStockType           whereHelperstring
	MaterialBaseUnit             whereHelpernull_String
	MatlWrhsStkQtyInMatlBaseUnit whereHelpernull_String
}{
	Material:                     whereHelperstring{field: "`sap_material_stock_data`.`Material`"},
	Plant:                        whereHelperstring{field: "`sap_material_stock_data`.`Plant`"},
	StorageLocation:              whereHelperstring{field: "`sap_material_stock_data`.`StorageLocation`"},
	Batch:                        whereHelperstring{field: "`sap_material_stock_data`.`Batch`"},
	Supplier:                     whereHelperstring{field: "`sap_material_stock_data`.`Supplier`"},
	Customer:                     whereHelperstring{field: "`sap_material_stock_data`.`Customer`"},
	WBSElementInternalID:         whereHelperstring{field: "`sap_material_stock_data`.`WBSElementInternalID`"},
	SDDocument:                   whereHelperstring{field: "`sap_material_stock_data`.`SDDocument`"},
	SDDocumentItem:               whereHelperstring{field: "`sap_material_stock_data`.`SDDocumentItem`"},
	InventorySpecialStockType:    whereHelperstring{field: "`sap_material_stock_data`.`InventorySpecialStockType`"},
	InventoryStockType:           whereHelperstring{field: "`sap_material_stock_data`.`InventoryStockType`"},
	MaterialBaseUnit:             whereHelpernull_String{field: "`sap_material_stock_data`.`MaterialBaseUnit`"},
	MatlWrhsStkQtyInMatlBaseUnit: whereHelpernull_String{field: "`sap_material_stock_data`.`MatlWrhsStkQtyInMatlBaseUnit`"},
}

// SapMaterialStockDatumRels is where relationship names are stored.
var SapMaterialStockDatumRels = struct {
}{}

// sapMaterialStockDatumR is where relationships are stored.
type sapMaterialStockDatumR struct {
}

// NewStruct creates a new relationship struct
func (*sapMaterialStockDatumR) NewStruct() *sapMaterialStockDatumR {
	return &sapMaterialStockDatumR{}
}

// sapMaterialStockDatumL is where Load methods for each relationship are stored.
type sapMaterialStockDatumL struct{}

var (
	sapMaterialStockDatumAllColumns            = []string{"Material", "Plant", "StorageLocation", "Batch", "Supplier", "Customer", "WBSElementInternalID", "SDDocument", "SDDocumentItem", "InventorySpecialStockType", "InventoryStockType", "MaterialBaseUnit", "MatlWrhsStkQtyInMatlBaseUnit"}
	sapMaterialStockDatumColumnsWithoutDefault = []string{"Material", "Plant", "StorageLocation", "Batch", "Supplier", "Customer", "WBSElementInternalID", "SDDocument", "SDDocumentItem", "InventorySpecialStockType", "InventoryStockType", "MaterialBaseUnit", "MatlWrhsStkQtyInMatlBaseUnit"}
	sapMaterialStockDatumColumnsWithDefault    = []string{}
	sapMaterialStockDatumPrimaryKeyColumns     = []string{"Material", "Plant", "StorageLocation", "Batch", "Supplier", "Customer", "WBSElementInternalID", "SDDocument", "SDDocumentItem", "InventorySpecialStockType", "InventoryStockType"}
)

type (
	// SapMaterialStockDatumSlice is an alias for a slice of pointers to SapMaterialStockDatum.
	// This should almost always be used instead of []SapMaterialStockDatum.
	SapMaterialStockDatumSlice []*SapMaterialStockDatum
	// SapMaterialStockDatumHook is the signature for custom SapMaterialStockDatum hook methods
	SapMaterialStockDatumHook func(context.Context, boil.ContextExecutor, *SapMaterialStockDatum) error

	sapMaterialStockDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapMaterialStockDatumType                 = reflect.TypeOf(&SapMaterialStockDatum{})
	sapMaterialStockDatumMapping              = queries.MakeStructMapping(sapMaterialStockDatumType)
	sapMaterialStockDatumPrimaryKeyMapping, _ = queries.BindMapping(sapMaterialStockDatumType, sapMaterialStockDatumMapping, sapMaterialStockDatumPrimaryKeyColumns)
	sapMaterialStockDatumInsertCacheMut       sync.RWMutex
	sapMaterialStockDatumInsertCache          = make(map[string]insertCache)
	sapMaterialStockDatumUpdateCacheMut       sync.RWMutex
	sapMaterialStockDatumUpdateCache          = make(map[string]updateCache)
	sapMaterialStockDatumUpsertCacheMut       sync.RWMutex
	sapMaterialStockDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapMaterialStockDatumBeforeInsertHooks []SapMaterialStockDatumHook
var sapMaterialStockDatumBeforeUpdateHooks []SapMaterialStockDatumHook
var sapMaterialStockDatumBeforeDeleteHooks []SapMaterialStockDatumHook
var sapMaterialStockDatumBeforeUpsertHooks []SapMaterialStockDatumHook

var sapMaterialStockDatumAfterInsertHooks []SapMaterialStockDatumHook
var sapMaterialStockDatumAfterSelectHooks []SapMaterialStockDatumHook
var sapMaterialStockDatumAfterUpdateHooks []SapMaterialStockDatumHook
var sapMaterialStockDatumAfterDeleteHooks []SapMaterialStockDatumHook
var sapMaterialStockDatumAfterUpsertHooks []SapMaterialStockDatumHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapMaterialStockDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapMaterialStockDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapMaterialStockDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapMaterialStockDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapMaterialStockDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapMaterialStockDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapMaterialStockDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapMaterialStockDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapMaterialStockDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaterialStockDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapMaterialStockDatumHook registers your hook function for all future operations.
func AddSapMaterialStockDatumHook(hookPoint boil.HookPoint, sapMaterialStockDatumHook SapMaterialStockDatumHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		sapMaterialStockDatumBeforeInsertHooks = append(sapMaterialStockDatumBeforeInsertHooks, sapMaterialStockDatumHook)
	case boil.BeforeUpdateHook:
		sapMaterialStockDatumBeforeUpdateHooks = append(sapMaterialStockDatumBeforeUpdateHooks, sapMaterialStockDatumHook)
	case boil.BeforeDeleteHook:
		sapMaterialStockDatumBeforeDeleteHooks = append(sapMaterialStockDatumBeforeDeleteHooks, sapMaterialStockDatumHook)
	case boil.BeforeUpsertHook:
		sapMaterialStockDatumBeforeUpsertHooks = append(sapMaterialStockDatumBeforeUpsertHooks, sapMaterialStockDatumHook)
	case boil.AfterInsertHook:
		sapMaterialStockDatumAfterInsertHooks = append(sapMaterialStockDatumAfterInsertHooks, sapMaterialStockDatumHook)
	case boil.AfterSelectHook:
		sapMaterialStockDatumAfterSelectHooks = append(sapMaterialStockDatumAfterSelectHooks, sapMaterialStockDatumHook)
	case boil.AfterUpdateHook:
		sapMaterialStockDatumAfterUpdateHooks = append(sapMaterialStockDatumAfterUpdateHooks, sapMaterialStockDatumHook)
	case boil.AfterDeleteHook:
		sapMaterialStockDatumAfterDeleteHooks = append(sapMaterialStockDatumAfterDeleteHooks, sapMaterialStockDatumHook)
	case boil.AfterUpsertHook:
		sapMaterialStockDatumAfterUpsertHooks = append(sapMaterialStockDatumAfterUpsertHooks, sapMaterialStockDatumHook)
	}
}

// One returns a single sapMaterialStockDatum record from the query.
func (q sapMaterialStockDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapMaterialStockDatum, error) {
	o := &SapMaterialStockDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_material_stock_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapMaterialStockDatum records from the query.
func (q sapMaterialStockDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapMaterialStockDatumSlice, error) {
	var o []*SapMaterialStockDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapMaterialStockDatum slice")
	}

	if len(sapMaterialStockDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapMaterialStockDatum records in the query.
func (q sapMaterialStockDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_material_stock_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapMaterialStockDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_material_stock_data exists")
	}

	return count > 0, nil
}

// SapMaterialStockData retrieves all the records using an executor.
func SapMaterialStockData(mods ...qm.QueryMod) sapMaterialStockDatumQuery {
	mods = append(mods, qm.From("`sap_material_stock_data`"))
	return sapMaterialStockDatumQuery{NewQuery(mods...)}
}

// FindSapMaterialStockDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapMaterialStockDatum(ctx context.Context, exec boil.ContextExecutor, material string, plant string, storageLocation string, batch string, supplier string, customer string, wBSElementInternalID string, sDDocument string, sDDocumentItem string, inventorySpecialStockType string, inventoryStockType string, selectCols ...string) (*SapMaterialStockDatum, error) {
	sapMaterialStockDatumObj := &SapMaterialStockDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_material_stock_data` where `Material`=? AND `Plant`=? AND `StorageLocation`=? AND `Batch`=? AND `Supplier`=? AND `Customer`=? AND `WBSElementInternalID`=? AND `SDDocument`=? AND `SDDocumentItem`=? AND `InventorySpecialStockType`=? AND `InventoryStockType`=?", sel,
	)

	q := queries.Raw(query, material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)

	err := q.Bind(ctx, exec, sapMaterialStockDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_material_stock_data")
	}

	if err = sapMaterialStockDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapMaterialStockDatumObj, err
	}

	return sapMaterialStockDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapMaterialStockDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_material_stock_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapMaterialStockDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapMaterialStockDatumInsertCacheMut.RLock()
	cache, cached := sapMaterialStockDatumInsertCache[key]
	sapMaterialStockDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapMaterialStockDatumAllColumns,
			sapMaterialStockDatumColumnsWithDefault,
			sapMaterialStockDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapMaterialStockDatumType, sapMaterialStockDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapMaterialStockDatumType, sapMaterialStockDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_material_stock_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_material_stock_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_material_stock_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapMaterialStockDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_material_stock_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Material,
		o.Plant,
		o.StorageLocation,
		o.Batch,
		o.Supplier,
		o.Customer,
		o.WBSElementInternalID,
		o.SDDocument,
		o.SDDocumentItem,
		o.InventorySpecialStockType,
		o.InventoryStockType,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_material_stock_data")
	}

CacheNoHooks:
	if !cached {
		sapMaterialStockDatumInsertCacheMut.Lock()
		sapMaterialStockDatumInsertCache[key] = cache
		sapMaterialStockDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapMaterialStockDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapMaterialStockDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapMaterialStockDatumUpdateCacheMut.RLock()
	cache, cached := sapMaterialStockDatumUpdateCache[key]
	sapMaterialStockDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapMaterialStockDatumAllColumns,
			sapMaterialStockDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_material_stock_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_material_stock_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapMaterialStockDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapMaterialStockDatumType, sapMaterialStockDatumMapping, append(wl, sapMaterialStockDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_material_stock_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_material_stock_data")
	}

	if !cached {
		sapMaterialStockDatumUpdateCacheMut.Lock()
		sapMaterialStockDatumUpdateCache[key] = cache
		sapMaterialStockDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapMaterialStockDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_material_stock_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_material_stock_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapMaterialStockDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaterialStockDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_material_stock_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaterialStockDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapMaterialStockDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapMaterialStockDatum")
	}
	return rowsAff, nil
}

var mySQLSapMaterialStockDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapMaterialStockDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_material_stock_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapMaterialStockDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapMaterialStockDatumUniqueColumns, o)

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

	sapMaterialStockDatumUpsertCacheMut.RLock()
	cache, cached := sapMaterialStockDatumUpsertCache[key]
	sapMaterialStockDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapMaterialStockDatumAllColumns,
			sapMaterialStockDatumColumnsWithDefault,
			sapMaterialStockDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			sapMaterialStockDatumAllColumns,
			sapMaterialStockDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_material_stock_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_material_stock_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_material_stock_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapMaterialStockDatumType, sapMaterialStockDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapMaterialStockDatumType, sapMaterialStockDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_material_stock_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapMaterialStockDatumType, sapMaterialStockDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_material_stock_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_material_stock_data")
	}

CacheNoHooks:
	if !cached {
		sapMaterialStockDatumUpsertCacheMut.Lock()
		sapMaterialStockDatumUpsertCache[key] = cache
		sapMaterialStockDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapMaterialStockDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapMaterialStockDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapMaterialStockDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapMaterialStockDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_material_stock_data` WHERE `Material`=? AND `Plant`=? AND `StorageLocation`=? AND `Batch`=? AND `Supplier`=? AND `Customer`=? AND `WBSElementInternalID`=? AND `SDDocument`=? AND `SDDocumentItem`=? AND `InventorySpecialStockType`=? AND `InventoryStockType`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_material_stock_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_material_stock_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapMaterialStockDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapMaterialStockDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_material_stock_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_material_stock_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapMaterialStockDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapMaterialStockDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaterialStockDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_material_stock_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaterialStockDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapMaterialStockDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_material_stock_data")
	}

	if len(sapMaterialStockDatumAfterDeleteHooks) != 0 {
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
func (o *SapMaterialStockDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapMaterialStockDatum(ctx, exec, o.Material, o.Plant, o.StorageLocation, o.Batch, o.Supplier, o.Customer, o.WBSElementInternalID, o.SDDocument, o.SDDocumentItem, o.InventorySpecialStockType, o.InventoryStockType)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapMaterialStockDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapMaterialStockDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaterialStockDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_material_stock_data`.* FROM `sap_material_stock_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaterialStockDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapMaterialStockDatumSlice")
	}

	*o = slice

	return nil
}

// SapMaterialStockDatumExists checks if the SapMaterialStockDatum row exists.
func SapMaterialStockDatumExists(ctx context.Context, exec boil.ContextExecutor, material string, plant string, storageLocation string, batch string, supplier string, customer string, wBSElementInternalID string, sDDocument string, sDDocumentItem string, inventorySpecialStockType string, inventoryStockType string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_material_stock_data` where `Material`=? AND `Plant`=? AND `StorageLocation`=? AND `Batch`=? AND `Supplier`=? AND `Customer`=? AND `WBSElementInternalID`=? AND `SDDocument`=? AND `SDDocumentItem`=? AND `InventorySpecialStockType`=? AND `InventoryStockType`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)
	}
	row := exec.QueryRowContext(ctx, sql, material, plant, storageLocation, batch, supplier, customer, wBSElementInternalID, sDDocument, sDDocumentItem, inventorySpecialStockType, inventoryStockType)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_material_stock_data exists")
	}

	return exists, nil
}
