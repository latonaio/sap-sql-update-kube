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

// SapWarehouseResourceHeaderDatum is an object representing the database table.
type SapWarehouseResourceHeaderDatum struct {
	EWMWarehouse                  string      `boil:"EWMWarehouse" json:"EWMWarehouse" toml:"EWMWarehouse" yaml:"EWMWarehouse"`
	EWMResource                   string      `boil:"EWMResource" json:"EWMResource" toml:"EWMResource" yaml:"EWMResource"`
	UserName                      null.String `boil:"UserName" json:"UserName,omitempty" toml:"UserName" yaml:"UserName,omitempty"`
	EWMResourceGroup              null.String `boil:"EWMResourceGroup" json:"EWMResourceGroup,omitempty" toml:"EWMResourceGroup" yaml:"EWMResourceGroup,omitempty"`
	EWMResourceType               null.String `boil:"EWMResourceType" json:"EWMResourceType,omitempty" toml:"EWMResourceType" yaml:"EWMResourceType,omitempty"`
	WarehouseOrderQueue           null.String `boil:"WarehouseOrderQueue" json:"WarehouseOrderQueue,omitempty" toml:"WarehouseOrderQueue" yaml:"WarehouseOrderQueue,omitempty"`
	EWMCurrentQueue               null.String `boil:"EWMCurrentQueue" json:"EWMCurrentQueue,omitempty" toml:"EWMCurrentQueue" yaml:"EWMCurrentQueue,omitempty"`
	EWMStorTypeOfLastWhseTaskConf null.String `boil:"EWMStorTypeOfLastWhseTaskConf" json:"EWMStorTypeOfLastWhseTaskConf,omitempty" toml:"EWMStorTypeOfLastWhseTaskConf" yaml:"EWMStorTypeOfLastWhseTaskConf,omitempty"`
	EWMResourceLogonDateTime      null.String `boil:"EWMResourceLogonDateTime" json:"EWMResourceLogonDateTime,omitempty" toml:"EWMResourceLogonDateTime" yaml:"EWMResourceLogonDateTime,omitempty"`
	EWMRsceIsLoggedOntoByCurUser  null.Bool   `boil:"EWMRsceIsLoggedOntoByCurUser" json:"EWMRsceIsLoggedOntoByCurUser,omitempty" toml:"EWMRsceIsLoggedOntoByCurUser" yaml:"EWMRsceIsLoggedOntoByCurUser,omitempty"`

	R *sapWarehouseResourceHeaderDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapWarehouseResourceHeaderDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapWarehouseResourceHeaderDatumColumns = struct {
	EWMWarehouse                  string
	EWMResource                   string
	UserName                      string
	EWMResourceGroup              string
	EWMResourceType               string
	WarehouseOrderQueue           string
	EWMCurrentQueue               string
	EWMStorTypeOfLastWhseTaskConf string
	EWMResourceLogonDateTime      string
	EWMRsceIsLoggedOntoByCurUser  string
}{
	EWMWarehouse:                  "EWMWarehouse",
	EWMResource:                   "EWMResource",
	UserName:                      "UserName",
	EWMResourceGroup:              "EWMResourceGroup",
	EWMResourceType:               "EWMResourceType",
	WarehouseOrderQueue:           "WarehouseOrderQueue",
	EWMCurrentQueue:               "EWMCurrentQueue",
	EWMStorTypeOfLastWhseTaskConf: "EWMStorTypeOfLastWhseTaskConf",
	EWMResourceLogonDateTime:      "EWMResourceLogonDateTime",
	EWMRsceIsLoggedOntoByCurUser:  "EWMRsceIsLoggedOntoByCurUser",
}

var SapWarehouseResourceHeaderDatumTableColumns = struct {
	EWMWarehouse                  string
	EWMResource                   string
	UserName                      string
	EWMResourceGroup              string
	EWMResourceType               string
	WarehouseOrderQueue           string
	EWMCurrentQueue               string
	EWMStorTypeOfLastWhseTaskConf string
	EWMResourceLogonDateTime      string
	EWMRsceIsLoggedOntoByCurUser  string
}{
	EWMWarehouse:                  "sap_warehouse_resource_header_data.EWMWarehouse",
	EWMResource:                   "sap_warehouse_resource_header_data.EWMResource",
	UserName:                      "sap_warehouse_resource_header_data.UserName",
	EWMResourceGroup:              "sap_warehouse_resource_header_data.EWMResourceGroup",
	EWMResourceType:               "sap_warehouse_resource_header_data.EWMResourceType",
	WarehouseOrderQueue:           "sap_warehouse_resource_header_data.WarehouseOrderQueue",
	EWMCurrentQueue:               "sap_warehouse_resource_header_data.EWMCurrentQueue",
	EWMStorTypeOfLastWhseTaskConf: "sap_warehouse_resource_header_data.EWMStorTypeOfLastWhseTaskConf",
	EWMResourceLogonDateTime:      "sap_warehouse_resource_header_data.EWMResourceLogonDateTime",
	EWMRsceIsLoggedOntoByCurUser:  "sap_warehouse_resource_header_data.EWMRsceIsLoggedOntoByCurUser",
}

// Generated where

var SapWarehouseResourceHeaderDatumWhere = struct {
	EWMWarehouse                  whereHelperstring
	EWMResource                   whereHelperstring
	UserName                      whereHelpernull_String
	EWMResourceGroup              whereHelpernull_String
	EWMResourceType               whereHelpernull_String
	WarehouseOrderQueue           whereHelpernull_String
	EWMCurrentQueue               whereHelpernull_String
	EWMStorTypeOfLastWhseTaskConf whereHelpernull_String
	EWMResourceLogonDateTime      whereHelpernull_String
	EWMRsceIsLoggedOntoByCurUser  whereHelpernull_Bool
}{
	EWMWarehouse:                  whereHelperstring{field: "`sap_warehouse_resource_header_data`.`EWMWarehouse`"},
	EWMResource:                   whereHelperstring{field: "`sap_warehouse_resource_header_data`.`EWMResource`"},
	UserName:                      whereHelpernull_String{field: "`sap_warehouse_resource_header_data`.`UserName`"},
	EWMResourceGroup:              whereHelpernull_String{field: "`sap_warehouse_resource_header_data`.`EWMResourceGroup`"},
	EWMResourceType:               whereHelpernull_String{field: "`sap_warehouse_resource_header_data`.`EWMResourceType`"},
	WarehouseOrderQueue:           whereHelpernull_String{field: "`sap_warehouse_resource_header_data`.`WarehouseOrderQueue`"},
	EWMCurrentQueue:               whereHelpernull_String{field: "`sap_warehouse_resource_header_data`.`EWMCurrentQueue`"},
	EWMStorTypeOfLastWhseTaskConf: whereHelpernull_String{field: "`sap_warehouse_resource_header_data`.`EWMStorTypeOfLastWhseTaskConf`"},
	EWMResourceLogonDateTime:      whereHelpernull_String{field: "`sap_warehouse_resource_header_data`.`EWMResourceLogonDateTime`"},
	EWMRsceIsLoggedOntoByCurUser:  whereHelpernull_Bool{field: "`sap_warehouse_resource_header_data`.`EWMRsceIsLoggedOntoByCurUser`"},
}

// SapWarehouseResourceHeaderDatumRels is where relationship names are stored.
var SapWarehouseResourceHeaderDatumRels = struct {
}{}

// sapWarehouseResourceHeaderDatumR is where relationships are stored.
type sapWarehouseResourceHeaderDatumR struct {
}

// NewStruct creates a new relationship struct
func (*sapWarehouseResourceHeaderDatumR) NewStruct() *sapWarehouseResourceHeaderDatumR {
	return &sapWarehouseResourceHeaderDatumR{}
}

// sapWarehouseResourceHeaderDatumL is where Load methods for each relationship are stored.
type sapWarehouseResourceHeaderDatumL struct{}

var (
	sapWarehouseResourceHeaderDatumAllColumns            = []string{"EWMWarehouse", "EWMResource", "UserName", "EWMResourceGroup", "EWMResourceType", "WarehouseOrderQueue", "EWMCurrentQueue", "EWMStorTypeOfLastWhseTaskConf", "EWMResourceLogonDateTime", "EWMRsceIsLoggedOntoByCurUser"}
	sapWarehouseResourceHeaderDatumColumnsWithoutDefault = []string{"EWMWarehouse", "EWMResource", "UserName", "EWMResourceGroup", "EWMResourceType", "WarehouseOrderQueue", "EWMCurrentQueue", "EWMStorTypeOfLastWhseTaskConf", "EWMResourceLogonDateTime", "EWMRsceIsLoggedOntoByCurUser"}
	sapWarehouseResourceHeaderDatumColumnsWithDefault    = []string{}
	sapWarehouseResourceHeaderDatumPrimaryKeyColumns     = []string{"EWMWarehouse", "EWMResource"}
	sapWarehouseResourceHeaderDatumGeneratedColumns      = []string{}
)

type (
	// SapWarehouseResourceHeaderDatumSlice is an alias for a slice of pointers to SapWarehouseResourceHeaderDatum.
	// This should almost always be used instead of []SapWarehouseResourceHeaderDatum.
	SapWarehouseResourceHeaderDatumSlice []*SapWarehouseResourceHeaderDatum
	// SapWarehouseResourceHeaderDatumHook is the signature for custom SapWarehouseResourceHeaderDatum hook methods
	SapWarehouseResourceHeaderDatumHook func(context.Context, boil.ContextExecutor, *SapWarehouseResourceHeaderDatum) error

	sapWarehouseResourceHeaderDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapWarehouseResourceHeaderDatumType                 = reflect.TypeOf(&SapWarehouseResourceHeaderDatum{})
	sapWarehouseResourceHeaderDatumMapping              = queries.MakeStructMapping(sapWarehouseResourceHeaderDatumType)
	sapWarehouseResourceHeaderDatumPrimaryKeyMapping, _ = queries.BindMapping(sapWarehouseResourceHeaderDatumType, sapWarehouseResourceHeaderDatumMapping, sapWarehouseResourceHeaderDatumPrimaryKeyColumns)
	sapWarehouseResourceHeaderDatumInsertCacheMut       sync.RWMutex
	sapWarehouseResourceHeaderDatumInsertCache          = make(map[string]insertCache)
	sapWarehouseResourceHeaderDatumUpdateCacheMut       sync.RWMutex
	sapWarehouseResourceHeaderDatumUpdateCache          = make(map[string]updateCache)
	sapWarehouseResourceHeaderDatumUpsertCacheMut       sync.RWMutex
	sapWarehouseResourceHeaderDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapWarehouseResourceHeaderDatumAfterSelectHooks []SapWarehouseResourceHeaderDatumHook

var sapWarehouseResourceHeaderDatumBeforeInsertHooks []SapWarehouseResourceHeaderDatumHook
var sapWarehouseResourceHeaderDatumAfterInsertHooks []SapWarehouseResourceHeaderDatumHook

var sapWarehouseResourceHeaderDatumBeforeUpdateHooks []SapWarehouseResourceHeaderDatumHook
var sapWarehouseResourceHeaderDatumAfterUpdateHooks []SapWarehouseResourceHeaderDatumHook

var sapWarehouseResourceHeaderDatumBeforeDeleteHooks []SapWarehouseResourceHeaderDatumHook
var sapWarehouseResourceHeaderDatumAfterDeleteHooks []SapWarehouseResourceHeaderDatumHook

var sapWarehouseResourceHeaderDatumBeforeUpsertHooks []SapWarehouseResourceHeaderDatumHook
var sapWarehouseResourceHeaderDatumAfterUpsertHooks []SapWarehouseResourceHeaderDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapWarehouseResourceHeaderDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapWarehouseResourceHeaderDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapWarehouseResourceHeaderDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapWarehouseResourceHeaderDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapWarehouseResourceHeaderDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapWarehouseResourceHeaderDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapWarehouseResourceHeaderDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapWarehouseResourceHeaderDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapWarehouseResourceHeaderDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWarehouseResourceHeaderDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapWarehouseResourceHeaderDatumHook registers your hook function for all future operations.
func AddSapWarehouseResourceHeaderDatumHook(hookPoint boil.HookPoint, sapWarehouseResourceHeaderDatumHook SapWarehouseResourceHeaderDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sapWarehouseResourceHeaderDatumAfterSelectHooks = append(sapWarehouseResourceHeaderDatumAfterSelectHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.BeforeInsertHook:
		sapWarehouseResourceHeaderDatumBeforeInsertHooks = append(sapWarehouseResourceHeaderDatumBeforeInsertHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.AfterInsertHook:
		sapWarehouseResourceHeaderDatumAfterInsertHooks = append(sapWarehouseResourceHeaderDatumAfterInsertHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.BeforeUpdateHook:
		sapWarehouseResourceHeaderDatumBeforeUpdateHooks = append(sapWarehouseResourceHeaderDatumBeforeUpdateHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.AfterUpdateHook:
		sapWarehouseResourceHeaderDatumAfterUpdateHooks = append(sapWarehouseResourceHeaderDatumAfterUpdateHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.BeforeDeleteHook:
		sapWarehouseResourceHeaderDatumBeforeDeleteHooks = append(sapWarehouseResourceHeaderDatumBeforeDeleteHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.AfterDeleteHook:
		sapWarehouseResourceHeaderDatumAfterDeleteHooks = append(sapWarehouseResourceHeaderDatumAfterDeleteHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.BeforeUpsertHook:
		sapWarehouseResourceHeaderDatumBeforeUpsertHooks = append(sapWarehouseResourceHeaderDatumBeforeUpsertHooks, sapWarehouseResourceHeaderDatumHook)
	case boil.AfterUpsertHook:
		sapWarehouseResourceHeaderDatumAfterUpsertHooks = append(sapWarehouseResourceHeaderDatumAfterUpsertHooks, sapWarehouseResourceHeaderDatumHook)
	}
}

// One returns a single sapWarehouseResourceHeaderDatum record from the query.
func (q sapWarehouseResourceHeaderDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapWarehouseResourceHeaderDatum, error) {
	o := &SapWarehouseResourceHeaderDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_warehouse_resource_header_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapWarehouseResourceHeaderDatum records from the query.
func (q sapWarehouseResourceHeaderDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapWarehouseResourceHeaderDatumSlice, error) {
	var o []*SapWarehouseResourceHeaderDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapWarehouseResourceHeaderDatum slice")
	}

	if len(sapWarehouseResourceHeaderDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapWarehouseResourceHeaderDatum records in the query.
func (q sapWarehouseResourceHeaderDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_warehouse_resource_header_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapWarehouseResourceHeaderDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_warehouse_resource_header_data exists")
	}

	return count > 0, nil
}

// SapWarehouseResourceHeaderData retrieves all the records using an executor.
func SapWarehouseResourceHeaderData(mods ...qm.QueryMod) sapWarehouseResourceHeaderDatumQuery {
	mods = append(mods, qm.From("`sap_warehouse_resource_header_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`sap_warehouse_resource_header_data`.*"})
	}

	return sapWarehouseResourceHeaderDatumQuery{q}
}

// FindSapWarehouseResourceHeaderDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapWarehouseResourceHeaderDatum(ctx context.Context, exec boil.ContextExecutor, eWMWarehouse string, eWMResource string, selectCols ...string) (*SapWarehouseResourceHeaderDatum, error) {
	sapWarehouseResourceHeaderDatumObj := &SapWarehouseResourceHeaderDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_warehouse_resource_header_data` where `EWMWarehouse`=? AND `EWMResource`=?", sel,
	)

	q := queries.Raw(query, eWMWarehouse, eWMResource)

	err := q.Bind(ctx, exec, sapWarehouseResourceHeaderDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_warehouse_resource_header_data")
	}

	if err = sapWarehouseResourceHeaderDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapWarehouseResourceHeaderDatumObj, err
	}

	return sapWarehouseResourceHeaderDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapWarehouseResourceHeaderDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_warehouse_resource_header_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapWarehouseResourceHeaderDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapWarehouseResourceHeaderDatumInsertCacheMut.RLock()
	cache, cached := sapWarehouseResourceHeaderDatumInsertCache[key]
	sapWarehouseResourceHeaderDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapWarehouseResourceHeaderDatumAllColumns,
			sapWarehouseResourceHeaderDatumColumnsWithDefault,
			sapWarehouseResourceHeaderDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapWarehouseResourceHeaderDatumType, sapWarehouseResourceHeaderDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapWarehouseResourceHeaderDatumType, sapWarehouseResourceHeaderDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_warehouse_resource_header_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_warehouse_resource_header_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_warehouse_resource_header_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapWarehouseResourceHeaderDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_warehouse_resource_header_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.EWMWarehouse,
		o.EWMResource,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_warehouse_resource_header_data")
	}

CacheNoHooks:
	if !cached {
		sapWarehouseResourceHeaderDatumInsertCacheMut.Lock()
		sapWarehouseResourceHeaderDatumInsertCache[key] = cache
		sapWarehouseResourceHeaderDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapWarehouseResourceHeaderDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapWarehouseResourceHeaderDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapWarehouseResourceHeaderDatumUpdateCacheMut.RLock()
	cache, cached := sapWarehouseResourceHeaderDatumUpdateCache[key]
	sapWarehouseResourceHeaderDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapWarehouseResourceHeaderDatumAllColumns,
			sapWarehouseResourceHeaderDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_warehouse_resource_header_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_warehouse_resource_header_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapWarehouseResourceHeaderDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapWarehouseResourceHeaderDatumType, sapWarehouseResourceHeaderDatumMapping, append(wl, sapWarehouseResourceHeaderDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_warehouse_resource_header_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_warehouse_resource_header_data")
	}

	if !cached {
		sapWarehouseResourceHeaderDatumUpdateCacheMut.Lock()
		sapWarehouseResourceHeaderDatumUpdateCache[key] = cache
		sapWarehouseResourceHeaderDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapWarehouseResourceHeaderDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_warehouse_resource_header_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_warehouse_resource_header_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapWarehouseResourceHeaderDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapWarehouseResourceHeaderDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_warehouse_resource_header_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapWarehouseResourceHeaderDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapWarehouseResourceHeaderDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapWarehouseResourceHeaderDatum")
	}
	return rowsAff, nil
}

var mySQLSapWarehouseResourceHeaderDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapWarehouseResourceHeaderDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_warehouse_resource_header_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapWarehouseResourceHeaderDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapWarehouseResourceHeaderDatumUniqueColumns, o)

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

	sapWarehouseResourceHeaderDatumUpsertCacheMut.RLock()
	cache, cached := sapWarehouseResourceHeaderDatumUpsertCache[key]
	sapWarehouseResourceHeaderDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapWarehouseResourceHeaderDatumAllColumns,
			sapWarehouseResourceHeaderDatumColumnsWithDefault,
			sapWarehouseResourceHeaderDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sapWarehouseResourceHeaderDatumAllColumns,
			sapWarehouseResourceHeaderDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_warehouse_resource_header_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_warehouse_resource_header_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_warehouse_resource_header_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapWarehouseResourceHeaderDatumType, sapWarehouseResourceHeaderDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapWarehouseResourceHeaderDatumType, sapWarehouseResourceHeaderDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_warehouse_resource_header_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapWarehouseResourceHeaderDatumType, sapWarehouseResourceHeaderDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_warehouse_resource_header_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_warehouse_resource_header_data")
	}

CacheNoHooks:
	if !cached {
		sapWarehouseResourceHeaderDatumUpsertCacheMut.Lock()
		sapWarehouseResourceHeaderDatumUpsertCache[key] = cache
		sapWarehouseResourceHeaderDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapWarehouseResourceHeaderDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapWarehouseResourceHeaderDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapWarehouseResourceHeaderDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapWarehouseResourceHeaderDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_warehouse_resource_header_data` WHERE `EWMWarehouse`=? AND `EWMResource`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_warehouse_resource_header_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_warehouse_resource_header_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapWarehouseResourceHeaderDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapWarehouseResourceHeaderDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_warehouse_resource_header_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_warehouse_resource_header_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapWarehouseResourceHeaderDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapWarehouseResourceHeaderDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapWarehouseResourceHeaderDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_warehouse_resource_header_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapWarehouseResourceHeaderDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapWarehouseResourceHeaderDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_warehouse_resource_header_data")
	}

	if len(sapWarehouseResourceHeaderDatumAfterDeleteHooks) != 0 {
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
func (o *SapWarehouseResourceHeaderDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapWarehouseResourceHeaderDatum(ctx, exec, o.EWMWarehouse, o.EWMResource)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapWarehouseResourceHeaderDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapWarehouseResourceHeaderDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapWarehouseResourceHeaderDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_warehouse_resource_header_data`.* FROM `sap_warehouse_resource_header_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapWarehouseResourceHeaderDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapWarehouseResourceHeaderDatumSlice")
	}

	*o = slice

	return nil
}

// SapWarehouseResourceHeaderDatumExists checks if the SapWarehouseResourceHeaderDatum row exists.
func SapWarehouseResourceHeaderDatumExists(ctx context.Context, exec boil.ContextExecutor, eWMWarehouse string, eWMResource string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_warehouse_resource_header_data` where `EWMWarehouse`=? AND `EWMResource`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, eWMWarehouse, eWMResource)
	}
	row := exec.QueryRowContext(ctx, sql, eWMWarehouse, eWMResource)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_warehouse_resource_header_data exists")
	}

	return exists, nil
}
