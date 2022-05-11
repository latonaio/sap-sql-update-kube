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

// SapBatchMasterRecord is an object representing the database table.
type SapBatchMasterRecord struct {
	Material                  string      `boil:"Material" json:"Material" toml:"Material" yaml:"Material"`
	BatchIdentifyingPlant     string      `boil:"BatchIdentifyingPlant" json:"BatchIdentifyingPlant" toml:"BatchIdentifyingPlant" yaml:"BatchIdentifyingPlant"`
	Batch                     string      `boil:"Batch" json:"Batch" toml:"Batch" yaml:"Batch"`
	Supplier                  null.String `boil:"Supplier" json:"Supplier,omitempty" toml:"Supplier" yaml:"Supplier,omitempty"`
	BatchBySupplier           null.String `boil:"BatchBySupplier" json:"BatchBySupplier,omitempty" toml:"BatchBySupplier" yaml:"BatchBySupplier,omitempty"`
	CountryOfOrigin           null.String `boil:"CountryOfOrigin" json:"CountryOfOrigin,omitempty" toml:"CountryOfOrigin" yaml:"CountryOfOrigin,omitempty"`
	RegionOfOrigin            null.String `boil:"RegionOfOrigin" json:"RegionOfOrigin,omitempty" toml:"RegionOfOrigin" yaml:"RegionOfOrigin,omitempty"`
	MatlBatchAvailabilityDate null.String `boil:"MatlBatchAvailabilityDate" json:"MatlBatchAvailabilityDate,omitempty" toml:"MatlBatchAvailabilityDate" yaml:"MatlBatchAvailabilityDate,omitempty"`
	ShelfLifeExpirationDate   null.String `boil:"ShelfLifeExpirationDate" json:"ShelfLifeExpirationDate,omitempty" toml:"ShelfLifeExpirationDate" yaml:"ShelfLifeExpirationDate,omitempty"`
	ManufactureDate           null.String `boil:"ManufactureDate" json:"ManufactureDate,omitempty" toml:"ManufactureDate" yaml:"ManufactureDate,omitempty"`
	CreationDateTime          null.String `boil:"CreationDateTime" json:"CreationDateTime,omitempty" toml:"CreationDateTime" yaml:"CreationDateTime,omitempty"`
	LastChangeDateTime        null.String `boil:"LastChangeDateTime" json:"LastChangeDateTime,omitempty" toml:"LastChangeDateTime" yaml:"LastChangeDateTime,omitempty"`
	BatchIsMarkedForDeletion  null.Bool   `boil:"BatchIsMarkedForDeletion" json:"BatchIsMarkedForDeletion,omitempty" toml:"BatchIsMarkedForDeletion" yaml:"BatchIsMarkedForDeletion,omitempty"`

	R *sapBatchMasterRecordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapBatchMasterRecordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapBatchMasterRecordColumns = struct {
	Material                  string
	BatchIdentifyingPlant     string
	Batch                     string
	Supplier                  string
	BatchBySupplier           string
	CountryOfOrigin           string
	RegionOfOrigin            string
	MatlBatchAvailabilityDate string
	ShelfLifeExpirationDate   string
	ManufactureDate           string
	CreationDateTime          string
	LastChangeDateTime        string
	BatchIsMarkedForDeletion  string
}{
	Material:                  "Material",
	BatchIdentifyingPlant:     "BatchIdentifyingPlant",
	Batch:                     "Batch",
	Supplier:                  "Supplier",
	BatchBySupplier:           "BatchBySupplier",
	CountryOfOrigin:           "CountryOfOrigin",
	RegionOfOrigin:            "RegionOfOrigin",
	MatlBatchAvailabilityDate: "MatlBatchAvailabilityDate",
	ShelfLifeExpirationDate:   "ShelfLifeExpirationDate",
	ManufactureDate:           "ManufactureDate",
	CreationDateTime:          "CreationDateTime",
	LastChangeDateTime:        "LastChangeDateTime",
	BatchIsMarkedForDeletion:  "BatchIsMarkedForDeletion",
}

var SapBatchMasterRecordTableColumns = struct {
	Material                  string
	BatchIdentifyingPlant     string
	Batch                     string
	Supplier                  string
	BatchBySupplier           string
	CountryOfOrigin           string
	RegionOfOrigin            string
	MatlBatchAvailabilityDate string
	ShelfLifeExpirationDate   string
	ManufactureDate           string
	CreationDateTime          string
	LastChangeDateTime        string
	BatchIsMarkedForDeletion  string
}{
	Material:                  "sap_batch_master_record.Material",
	BatchIdentifyingPlant:     "sap_batch_master_record.BatchIdentifyingPlant",
	Batch:                     "sap_batch_master_record.Batch",
	Supplier:                  "sap_batch_master_record.Supplier",
	BatchBySupplier:           "sap_batch_master_record.BatchBySupplier",
	CountryOfOrigin:           "sap_batch_master_record.CountryOfOrigin",
	RegionOfOrigin:            "sap_batch_master_record.RegionOfOrigin",
	MatlBatchAvailabilityDate: "sap_batch_master_record.MatlBatchAvailabilityDate",
	ShelfLifeExpirationDate:   "sap_batch_master_record.ShelfLifeExpirationDate",
	ManufactureDate:           "sap_batch_master_record.ManufactureDate",
	CreationDateTime:          "sap_batch_master_record.CreationDateTime",
	LastChangeDateTime:        "sap_batch_master_record.LastChangeDateTime",
	BatchIsMarkedForDeletion:  "sap_batch_master_record.BatchIsMarkedForDeletion",
}

// Generated where

var SapBatchMasterRecordWhere = struct {
	Material                  whereHelperstring
	BatchIdentifyingPlant     whereHelperstring
	Batch                     whereHelperstring
	Supplier                  whereHelpernull_String
	BatchBySupplier           whereHelpernull_String
	CountryOfOrigin           whereHelpernull_String
	RegionOfOrigin            whereHelpernull_String
	MatlBatchAvailabilityDate whereHelpernull_String
	ShelfLifeExpirationDate   whereHelpernull_String
	ManufactureDate           whereHelpernull_String
	CreationDateTime          whereHelpernull_String
	LastChangeDateTime        whereHelpernull_String
	BatchIsMarkedForDeletion  whereHelpernull_Bool
}{
	Material:                  whereHelperstring{field: "`sap_batch_master_record`.`Material`"},
	BatchIdentifyingPlant:     whereHelperstring{field: "`sap_batch_master_record`.`BatchIdentifyingPlant`"},
	Batch:                     whereHelperstring{field: "`sap_batch_master_record`.`Batch`"},
	Supplier:                  whereHelpernull_String{field: "`sap_batch_master_record`.`Supplier`"},
	BatchBySupplier:           whereHelpernull_String{field: "`sap_batch_master_record`.`BatchBySupplier`"},
	CountryOfOrigin:           whereHelpernull_String{field: "`sap_batch_master_record`.`CountryOfOrigin`"},
	RegionOfOrigin:            whereHelpernull_String{field: "`sap_batch_master_record`.`RegionOfOrigin`"},
	MatlBatchAvailabilityDate: whereHelpernull_String{field: "`sap_batch_master_record`.`MatlBatchAvailabilityDate`"},
	ShelfLifeExpirationDate:   whereHelpernull_String{field: "`sap_batch_master_record`.`ShelfLifeExpirationDate`"},
	ManufactureDate:           whereHelpernull_String{field: "`sap_batch_master_record`.`ManufactureDate`"},
	CreationDateTime:          whereHelpernull_String{field: "`sap_batch_master_record`.`CreationDateTime`"},
	LastChangeDateTime:        whereHelpernull_String{field: "`sap_batch_master_record`.`LastChangeDateTime`"},
	BatchIsMarkedForDeletion:  whereHelpernull_Bool{field: "`sap_batch_master_record`.`BatchIsMarkedForDeletion`"},
}

// SapBatchMasterRecordRels is where relationship names are stored.
var SapBatchMasterRecordRels = struct {
}{}

// sapBatchMasterRecordR is where relationships are stored.
type sapBatchMasterRecordR struct {
}

// NewStruct creates a new relationship struct
func (*sapBatchMasterRecordR) NewStruct() *sapBatchMasterRecordR {
	return &sapBatchMasterRecordR{}
}

// sapBatchMasterRecordL is where Load methods for each relationship are stored.
type sapBatchMasterRecordL struct{}

var (
	sapBatchMasterRecordAllColumns            = []string{"Material", "BatchIdentifyingPlant", "Batch", "Supplier", "BatchBySupplier", "CountryOfOrigin", "RegionOfOrigin", "MatlBatchAvailabilityDate", "ShelfLifeExpirationDate", "ManufactureDate", "CreationDateTime", "LastChangeDateTime", "BatchIsMarkedForDeletion"}
	sapBatchMasterRecordColumnsWithoutDefault = []string{"Material", "BatchIdentifyingPlant", "Batch", "Supplier", "BatchBySupplier", "CountryOfOrigin", "RegionOfOrigin", "MatlBatchAvailabilityDate", "ShelfLifeExpirationDate", "ManufactureDate", "CreationDateTime", "LastChangeDateTime", "BatchIsMarkedForDeletion"}
	sapBatchMasterRecordColumnsWithDefault    = []string{}
	sapBatchMasterRecordPrimaryKeyColumns     = []string{"Material", "BatchIdentifyingPlant", "Batch"}
	sapBatchMasterRecordGeneratedColumns      = []string{}
)

type (
	// SapBatchMasterRecordSlice is an alias for a slice of pointers to SapBatchMasterRecord.
	// This should almost always be used instead of []SapBatchMasterRecord.
	SapBatchMasterRecordSlice []*SapBatchMasterRecord
	// SapBatchMasterRecordHook is the signature for custom SapBatchMasterRecord hook methods
	SapBatchMasterRecordHook func(context.Context, boil.ContextExecutor, *SapBatchMasterRecord) error

	sapBatchMasterRecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapBatchMasterRecordType                 = reflect.TypeOf(&SapBatchMasterRecord{})
	sapBatchMasterRecordMapping              = queries.MakeStructMapping(sapBatchMasterRecordType)
	sapBatchMasterRecordPrimaryKeyMapping, _ = queries.BindMapping(sapBatchMasterRecordType, sapBatchMasterRecordMapping, sapBatchMasterRecordPrimaryKeyColumns)
	sapBatchMasterRecordInsertCacheMut       sync.RWMutex
	sapBatchMasterRecordInsertCache          = make(map[string]insertCache)
	sapBatchMasterRecordUpdateCacheMut       sync.RWMutex
	sapBatchMasterRecordUpdateCache          = make(map[string]updateCache)
	sapBatchMasterRecordUpsertCacheMut       sync.RWMutex
	sapBatchMasterRecordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapBatchMasterRecordAfterSelectHooks []SapBatchMasterRecordHook

var sapBatchMasterRecordBeforeInsertHooks []SapBatchMasterRecordHook
var sapBatchMasterRecordAfterInsertHooks []SapBatchMasterRecordHook

var sapBatchMasterRecordBeforeUpdateHooks []SapBatchMasterRecordHook
var sapBatchMasterRecordAfterUpdateHooks []SapBatchMasterRecordHook

var sapBatchMasterRecordBeforeDeleteHooks []SapBatchMasterRecordHook
var sapBatchMasterRecordAfterDeleteHooks []SapBatchMasterRecordHook

var sapBatchMasterRecordBeforeUpsertHooks []SapBatchMasterRecordHook
var sapBatchMasterRecordAfterUpsertHooks []SapBatchMasterRecordHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapBatchMasterRecord) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapBatchMasterRecord) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapBatchMasterRecord) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapBatchMasterRecord) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapBatchMasterRecord) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapBatchMasterRecord) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapBatchMasterRecord) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapBatchMasterRecord) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapBatchMasterRecord) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapBatchMasterRecordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapBatchMasterRecordHook registers your hook function for all future operations.
func AddSapBatchMasterRecordHook(hookPoint boil.HookPoint, sapBatchMasterRecordHook SapBatchMasterRecordHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sapBatchMasterRecordAfterSelectHooks = append(sapBatchMasterRecordAfterSelectHooks, sapBatchMasterRecordHook)
	case boil.BeforeInsertHook:
		sapBatchMasterRecordBeforeInsertHooks = append(sapBatchMasterRecordBeforeInsertHooks, sapBatchMasterRecordHook)
	case boil.AfterInsertHook:
		sapBatchMasterRecordAfterInsertHooks = append(sapBatchMasterRecordAfterInsertHooks, sapBatchMasterRecordHook)
	case boil.BeforeUpdateHook:
		sapBatchMasterRecordBeforeUpdateHooks = append(sapBatchMasterRecordBeforeUpdateHooks, sapBatchMasterRecordHook)
	case boil.AfterUpdateHook:
		sapBatchMasterRecordAfterUpdateHooks = append(sapBatchMasterRecordAfterUpdateHooks, sapBatchMasterRecordHook)
	case boil.BeforeDeleteHook:
		sapBatchMasterRecordBeforeDeleteHooks = append(sapBatchMasterRecordBeforeDeleteHooks, sapBatchMasterRecordHook)
	case boil.AfterDeleteHook:
		sapBatchMasterRecordAfterDeleteHooks = append(sapBatchMasterRecordAfterDeleteHooks, sapBatchMasterRecordHook)
	case boil.BeforeUpsertHook:
		sapBatchMasterRecordBeforeUpsertHooks = append(sapBatchMasterRecordBeforeUpsertHooks, sapBatchMasterRecordHook)
	case boil.AfterUpsertHook:
		sapBatchMasterRecordAfterUpsertHooks = append(sapBatchMasterRecordAfterUpsertHooks, sapBatchMasterRecordHook)
	}
}

// One returns a single sapBatchMasterRecord record from the query.
func (q sapBatchMasterRecordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapBatchMasterRecord, error) {
	o := &SapBatchMasterRecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_batch_master_record")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapBatchMasterRecord records from the query.
func (q sapBatchMasterRecordQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapBatchMasterRecordSlice, error) {
	var o []*SapBatchMasterRecord

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapBatchMasterRecord slice")
	}

	if len(sapBatchMasterRecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapBatchMasterRecord records in the query.
func (q sapBatchMasterRecordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_batch_master_record rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapBatchMasterRecordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_batch_master_record exists")
	}

	return count > 0, nil
}

// SapBatchMasterRecords retrieves all the records using an executor.
func SapBatchMasterRecords(mods ...qm.QueryMod) sapBatchMasterRecordQuery {
	mods = append(mods, qm.From("`sap_batch_master_record`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`sap_batch_master_record`.*"})
	}

	return sapBatchMasterRecordQuery{q}
}

// FindSapBatchMasterRecord retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapBatchMasterRecord(ctx context.Context, exec boil.ContextExecutor, material string, batchIdentifyingPlant string, batch string, selectCols ...string) (*SapBatchMasterRecord, error) {
	sapBatchMasterRecordObj := &SapBatchMasterRecord{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_batch_master_record` where `Material`=? AND `BatchIdentifyingPlant`=? AND `Batch`=?", sel,
	)

	q := queries.Raw(query, material, batchIdentifyingPlant, batch)

	err := q.Bind(ctx, exec, sapBatchMasterRecordObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_batch_master_record")
	}

	if err = sapBatchMasterRecordObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapBatchMasterRecordObj, err
	}

	return sapBatchMasterRecordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapBatchMasterRecord) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_batch_master_record provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapBatchMasterRecordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapBatchMasterRecordInsertCacheMut.RLock()
	cache, cached := sapBatchMasterRecordInsertCache[key]
	sapBatchMasterRecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapBatchMasterRecordAllColumns,
			sapBatchMasterRecordColumnsWithDefault,
			sapBatchMasterRecordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapBatchMasterRecordType, sapBatchMasterRecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapBatchMasterRecordType, sapBatchMasterRecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_batch_master_record` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_batch_master_record` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_batch_master_record` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapBatchMasterRecordPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_batch_master_record")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Material,
		o.BatchIdentifyingPlant,
		o.Batch,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_batch_master_record")
	}

CacheNoHooks:
	if !cached {
		sapBatchMasterRecordInsertCacheMut.Lock()
		sapBatchMasterRecordInsertCache[key] = cache
		sapBatchMasterRecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapBatchMasterRecord.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapBatchMasterRecord) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapBatchMasterRecordUpdateCacheMut.RLock()
	cache, cached := sapBatchMasterRecordUpdateCache[key]
	sapBatchMasterRecordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapBatchMasterRecordAllColumns,
			sapBatchMasterRecordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_batch_master_record, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_batch_master_record` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapBatchMasterRecordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapBatchMasterRecordType, sapBatchMasterRecordMapping, append(wl, sapBatchMasterRecordPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_batch_master_record row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_batch_master_record")
	}

	if !cached {
		sapBatchMasterRecordUpdateCacheMut.Lock()
		sapBatchMasterRecordUpdateCache[key] = cache
		sapBatchMasterRecordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapBatchMasterRecordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_batch_master_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_batch_master_record")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapBatchMasterRecordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapBatchMasterRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_batch_master_record` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapBatchMasterRecordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapBatchMasterRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapBatchMasterRecord")
	}
	return rowsAff, nil
}

var mySQLSapBatchMasterRecordUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapBatchMasterRecord) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_batch_master_record provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapBatchMasterRecordColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapBatchMasterRecordUniqueColumns, o)

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

	sapBatchMasterRecordUpsertCacheMut.RLock()
	cache, cached := sapBatchMasterRecordUpsertCache[key]
	sapBatchMasterRecordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapBatchMasterRecordAllColumns,
			sapBatchMasterRecordColumnsWithDefault,
			sapBatchMasterRecordColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sapBatchMasterRecordAllColumns,
			sapBatchMasterRecordPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_batch_master_record, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_batch_master_record`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_batch_master_record` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapBatchMasterRecordType, sapBatchMasterRecordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapBatchMasterRecordType, sapBatchMasterRecordMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_batch_master_record")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapBatchMasterRecordType, sapBatchMasterRecordMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_batch_master_record")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_batch_master_record")
	}

CacheNoHooks:
	if !cached {
		sapBatchMasterRecordUpsertCacheMut.Lock()
		sapBatchMasterRecordUpsertCache[key] = cache
		sapBatchMasterRecordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapBatchMasterRecord record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapBatchMasterRecord) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapBatchMasterRecord provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapBatchMasterRecordPrimaryKeyMapping)
	sql := "DELETE FROM `sap_batch_master_record` WHERE `Material`=? AND `BatchIdentifyingPlant`=? AND `Batch`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_batch_master_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_batch_master_record")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapBatchMasterRecordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapBatchMasterRecordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_batch_master_record")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_batch_master_record")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapBatchMasterRecordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapBatchMasterRecordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapBatchMasterRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_batch_master_record` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapBatchMasterRecordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapBatchMasterRecord slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_batch_master_record")
	}

	if len(sapBatchMasterRecordAfterDeleteHooks) != 0 {
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
func (o *SapBatchMasterRecord) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapBatchMasterRecord(ctx, exec, o.Material, o.BatchIdentifyingPlant, o.Batch)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapBatchMasterRecordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapBatchMasterRecordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapBatchMasterRecordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_batch_master_record`.* FROM `sap_batch_master_record` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapBatchMasterRecordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapBatchMasterRecordSlice")
	}

	*o = slice

	return nil
}

// SapBatchMasterRecordExists checks if the SapBatchMasterRecord row exists.
func SapBatchMasterRecordExists(ctx context.Context, exec boil.ContextExecutor, material string, batchIdentifyingPlant string, batch string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_batch_master_record` where `Material`=? AND `BatchIdentifyingPlant`=? AND `Batch`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, material, batchIdentifyingPlant, batch)
	}
	row := exec.QueryRowContext(ctx, sql, material, batchIdentifyingPlant, batch)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_batch_master_record exists")
	}

	return exists, nil
}
