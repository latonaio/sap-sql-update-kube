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

// SapMaintenanceItemCallObjectsDatum is an object representing the database table.
type SapMaintenanceItemCallObjectsDatum struct {
	MaintenancePlan              string      `boil:"MaintenancePlan" json:"MaintenancePlan" toml:"MaintenancePlan" yaml:"MaintenancePlan"`
	MaintenanceItem              string      `boil:"MaintenanceItem" json:"MaintenanceItem" toml:"MaintenanceItem" yaml:"MaintenanceItem"`
	MaintenancePlanCallNumber    int         `boil:"MaintenancePlanCallNumber" json:"MaintenancePlanCallNumber" toml:"MaintenancePlanCallNumber" yaml:"MaintenancePlanCallNumber"`
	MaintenanceOrder             null.String `boil:"MaintenanceOrder" json:"MaintenanceOrder,omitempty" toml:"MaintenanceOrder" yaml:"MaintenanceOrder,omitempty"`
	MaintenanceNotification      null.String `boil:"MaintenanceNotification" json:"MaintenanceNotification,omitempty" toml:"MaintenanceNotification" yaml:"MaintenanceNotification,omitempty"`
	ServiceOrder                 null.String `boil:"ServiceOrder" json:"ServiceOrder,omitempty" toml:"ServiceOrder" yaml:"ServiceOrder,omitempty"`
	MaintCallHorizonIsNotReached null.Bool   `boil:"MaintCallHorizonIsNotReached" json:"MaintCallHorizonIsNotReached,omitempty" toml:"MaintCallHorizonIsNotReached" yaml:"MaintCallHorizonIsNotReached,omitempty"`
	SchedulingStatus             null.String `boil:"SchedulingStatus" json:"SchedulingStatus,omitempty" toml:"SchedulingStatus" yaml:"SchedulingStatus,omitempty"`
	PlannedStartDate             null.String `boil:"PlannedStartDate" json:"PlannedStartDate,omitempty" toml:"PlannedStartDate" yaml:"PlannedStartDate,omitempty"`

	R *sapMaintenanceItemCallObjectsDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapMaintenanceItemCallObjectsDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapMaintenanceItemCallObjectsDatumColumns = struct {
	MaintenancePlan              string
	MaintenanceItem              string
	MaintenancePlanCallNumber    string
	MaintenanceOrder             string
	MaintenanceNotification      string
	ServiceOrder                 string
	MaintCallHorizonIsNotReached string
	SchedulingStatus             string
	PlannedStartDate             string
}{
	MaintenancePlan:              "MaintenancePlan",
	MaintenanceItem:              "MaintenanceItem",
	MaintenancePlanCallNumber:    "MaintenancePlanCallNumber",
	MaintenanceOrder:             "MaintenanceOrder",
	MaintenanceNotification:      "MaintenanceNotification",
	ServiceOrder:                 "ServiceOrder",
	MaintCallHorizonIsNotReached: "MaintCallHorizonIsNotReached",
	SchedulingStatus:             "SchedulingStatus",
	PlannedStartDate:             "PlannedStartDate",
}

var SapMaintenanceItemCallObjectsDatumTableColumns = struct {
	MaintenancePlan              string
	MaintenanceItem              string
	MaintenancePlanCallNumber    string
	MaintenanceOrder             string
	MaintenanceNotification      string
	ServiceOrder                 string
	MaintCallHorizonIsNotReached string
	SchedulingStatus             string
	PlannedStartDate             string
}{
	MaintenancePlan:              "sap_maintenance_item_call_objects_data.MaintenancePlan",
	MaintenanceItem:              "sap_maintenance_item_call_objects_data.MaintenanceItem",
	MaintenancePlanCallNumber:    "sap_maintenance_item_call_objects_data.MaintenancePlanCallNumber",
	MaintenanceOrder:             "sap_maintenance_item_call_objects_data.MaintenanceOrder",
	MaintenanceNotification:      "sap_maintenance_item_call_objects_data.MaintenanceNotification",
	ServiceOrder:                 "sap_maintenance_item_call_objects_data.ServiceOrder",
	MaintCallHorizonIsNotReached: "sap_maintenance_item_call_objects_data.MaintCallHorizonIsNotReached",
	SchedulingStatus:             "sap_maintenance_item_call_objects_data.SchedulingStatus",
	PlannedStartDate:             "sap_maintenance_item_call_objects_data.PlannedStartDate",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var SapMaintenanceItemCallObjectsDatumWhere = struct {
	MaintenancePlan              whereHelperstring
	MaintenanceItem              whereHelperstring
	MaintenancePlanCallNumber    whereHelperint
	MaintenanceOrder             whereHelpernull_String
	MaintenanceNotification      whereHelpernull_String
	ServiceOrder                 whereHelpernull_String
	MaintCallHorizonIsNotReached whereHelpernull_Bool
	SchedulingStatus             whereHelpernull_String
	PlannedStartDate             whereHelpernull_String
}{
	MaintenancePlan:              whereHelperstring{field: "`sap_maintenance_item_call_objects_data`.`MaintenancePlan`"},
	MaintenanceItem:              whereHelperstring{field: "`sap_maintenance_item_call_objects_data`.`MaintenanceItem`"},
	MaintenancePlanCallNumber:    whereHelperint{field: "`sap_maintenance_item_call_objects_data`.`MaintenancePlanCallNumber`"},
	MaintenanceOrder:             whereHelpernull_String{field: "`sap_maintenance_item_call_objects_data`.`MaintenanceOrder`"},
	MaintenanceNotification:      whereHelpernull_String{field: "`sap_maintenance_item_call_objects_data`.`MaintenanceNotification`"},
	ServiceOrder:                 whereHelpernull_String{field: "`sap_maintenance_item_call_objects_data`.`ServiceOrder`"},
	MaintCallHorizonIsNotReached: whereHelpernull_Bool{field: "`sap_maintenance_item_call_objects_data`.`MaintCallHorizonIsNotReached`"},
	SchedulingStatus:             whereHelpernull_String{field: "`sap_maintenance_item_call_objects_data`.`SchedulingStatus`"},
	PlannedStartDate:             whereHelpernull_String{field: "`sap_maintenance_item_call_objects_data`.`PlannedStartDate`"},
}

// SapMaintenanceItemCallObjectsDatumRels is where relationship names are stored.
var SapMaintenanceItemCallObjectsDatumRels = struct {
}{}

// sapMaintenanceItemCallObjectsDatumR is where relationships are stored.
type sapMaintenanceItemCallObjectsDatumR struct {
}

// NewStruct creates a new relationship struct
func (*sapMaintenanceItemCallObjectsDatumR) NewStruct() *sapMaintenanceItemCallObjectsDatumR {
	return &sapMaintenanceItemCallObjectsDatumR{}
}

// sapMaintenanceItemCallObjectsDatumL is where Load methods for each relationship are stored.
type sapMaintenanceItemCallObjectsDatumL struct{}

var (
	sapMaintenanceItemCallObjectsDatumAllColumns            = []string{"MaintenancePlan", "MaintenanceItem", "MaintenancePlanCallNumber", "MaintenanceOrder", "MaintenanceNotification", "ServiceOrder", "MaintCallHorizonIsNotReached", "SchedulingStatus", "PlannedStartDate"}
	sapMaintenanceItemCallObjectsDatumColumnsWithoutDefault = []string{"MaintenancePlan", "MaintenanceItem", "MaintenancePlanCallNumber", "MaintenanceOrder", "MaintenanceNotification", "ServiceOrder", "MaintCallHorizonIsNotReached", "SchedulingStatus", "PlannedStartDate"}
	sapMaintenanceItemCallObjectsDatumColumnsWithDefault    = []string{}
	sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns     = []string{"MaintenancePlan", "MaintenanceItem", "MaintenancePlanCallNumber"}
	sapMaintenanceItemCallObjectsDatumGeneratedColumns      = []string{}
)

type (
	// SapMaintenanceItemCallObjectsDatumSlice is an alias for a slice of pointers to SapMaintenanceItemCallObjectsDatum.
	// This should almost always be used instead of []SapMaintenanceItemCallObjectsDatum.
	SapMaintenanceItemCallObjectsDatumSlice []*SapMaintenanceItemCallObjectsDatum
	// SapMaintenanceItemCallObjectsDatumHook is the signature for custom SapMaintenanceItemCallObjectsDatum hook methods
	SapMaintenanceItemCallObjectsDatumHook func(context.Context, boil.ContextExecutor, *SapMaintenanceItemCallObjectsDatum) error

	sapMaintenanceItemCallObjectsDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapMaintenanceItemCallObjectsDatumType                 = reflect.TypeOf(&SapMaintenanceItemCallObjectsDatum{})
	sapMaintenanceItemCallObjectsDatumMapping              = queries.MakeStructMapping(sapMaintenanceItemCallObjectsDatumType)
	sapMaintenanceItemCallObjectsDatumPrimaryKeyMapping, _ = queries.BindMapping(sapMaintenanceItemCallObjectsDatumType, sapMaintenanceItemCallObjectsDatumMapping, sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns)
	sapMaintenanceItemCallObjectsDatumInsertCacheMut       sync.RWMutex
	sapMaintenanceItemCallObjectsDatumInsertCache          = make(map[string]insertCache)
	sapMaintenanceItemCallObjectsDatumUpdateCacheMut       sync.RWMutex
	sapMaintenanceItemCallObjectsDatumUpdateCache          = make(map[string]updateCache)
	sapMaintenanceItemCallObjectsDatumUpsertCacheMut       sync.RWMutex
	sapMaintenanceItemCallObjectsDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapMaintenanceItemCallObjectsDatumAfterSelectHooks []SapMaintenanceItemCallObjectsDatumHook

var sapMaintenanceItemCallObjectsDatumBeforeInsertHooks []SapMaintenanceItemCallObjectsDatumHook
var sapMaintenanceItemCallObjectsDatumAfterInsertHooks []SapMaintenanceItemCallObjectsDatumHook

var sapMaintenanceItemCallObjectsDatumBeforeUpdateHooks []SapMaintenanceItemCallObjectsDatumHook
var sapMaintenanceItemCallObjectsDatumAfterUpdateHooks []SapMaintenanceItemCallObjectsDatumHook

var sapMaintenanceItemCallObjectsDatumBeforeDeleteHooks []SapMaintenanceItemCallObjectsDatumHook
var sapMaintenanceItemCallObjectsDatumAfterDeleteHooks []SapMaintenanceItemCallObjectsDatumHook

var sapMaintenanceItemCallObjectsDatumBeforeUpsertHooks []SapMaintenanceItemCallObjectsDatumHook
var sapMaintenanceItemCallObjectsDatumAfterUpsertHooks []SapMaintenanceItemCallObjectsDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapMaintenanceItemCallObjectsDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapMaintenanceItemCallObjectsDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapMaintenanceItemCallObjectsDatumHook registers your hook function for all future operations.
func AddSapMaintenanceItemCallObjectsDatumHook(hookPoint boil.HookPoint, sapMaintenanceItemCallObjectsDatumHook SapMaintenanceItemCallObjectsDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sapMaintenanceItemCallObjectsDatumAfterSelectHooks = append(sapMaintenanceItemCallObjectsDatumAfterSelectHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.BeforeInsertHook:
		sapMaintenanceItemCallObjectsDatumBeforeInsertHooks = append(sapMaintenanceItemCallObjectsDatumBeforeInsertHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.AfterInsertHook:
		sapMaintenanceItemCallObjectsDatumAfterInsertHooks = append(sapMaintenanceItemCallObjectsDatumAfterInsertHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.BeforeUpdateHook:
		sapMaintenanceItemCallObjectsDatumBeforeUpdateHooks = append(sapMaintenanceItemCallObjectsDatumBeforeUpdateHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.AfterUpdateHook:
		sapMaintenanceItemCallObjectsDatumAfterUpdateHooks = append(sapMaintenanceItemCallObjectsDatumAfterUpdateHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.BeforeDeleteHook:
		sapMaintenanceItemCallObjectsDatumBeforeDeleteHooks = append(sapMaintenanceItemCallObjectsDatumBeforeDeleteHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.AfterDeleteHook:
		sapMaintenanceItemCallObjectsDatumAfterDeleteHooks = append(sapMaintenanceItemCallObjectsDatumAfterDeleteHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.BeforeUpsertHook:
		sapMaintenanceItemCallObjectsDatumBeforeUpsertHooks = append(sapMaintenanceItemCallObjectsDatumBeforeUpsertHooks, sapMaintenanceItemCallObjectsDatumHook)
	case boil.AfterUpsertHook:
		sapMaintenanceItemCallObjectsDatumAfterUpsertHooks = append(sapMaintenanceItemCallObjectsDatumAfterUpsertHooks, sapMaintenanceItemCallObjectsDatumHook)
	}
}

// One returns a single sapMaintenanceItemCallObjectsDatum record from the query.
func (q sapMaintenanceItemCallObjectsDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapMaintenanceItemCallObjectsDatum, error) {
	o := &SapMaintenanceItemCallObjectsDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_maintenance_item_call_objects_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapMaintenanceItemCallObjectsDatum records from the query.
func (q sapMaintenanceItemCallObjectsDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapMaintenanceItemCallObjectsDatumSlice, error) {
	var o []*SapMaintenanceItemCallObjectsDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapMaintenanceItemCallObjectsDatum slice")
	}

	if len(sapMaintenanceItemCallObjectsDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapMaintenanceItemCallObjectsDatum records in the query.
func (q sapMaintenanceItemCallObjectsDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_maintenance_item_call_objects_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapMaintenanceItemCallObjectsDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_maintenance_item_call_objects_data exists")
	}

	return count > 0, nil
}

// SapMaintenanceItemCallObjectsData retrieves all the records using an executor.
func SapMaintenanceItemCallObjectsData(mods ...qm.QueryMod) sapMaintenanceItemCallObjectsDatumQuery {
	mods = append(mods, qm.From("`sap_maintenance_item_call_objects_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`sap_maintenance_item_call_objects_data`.*"})
	}

	return sapMaintenanceItemCallObjectsDatumQuery{q}
}

// FindSapMaintenanceItemCallObjectsDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapMaintenanceItemCallObjectsDatum(ctx context.Context, exec boil.ContextExecutor, maintenancePlan string, maintenanceItem string, maintenancePlanCallNumber int, selectCols ...string) (*SapMaintenanceItemCallObjectsDatum, error) {
	sapMaintenanceItemCallObjectsDatumObj := &SapMaintenanceItemCallObjectsDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_maintenance_item_call_objects_data` where `MaintenancePlan`=? AND `MaintenanceItem`=? AND `MaintenancePlanCallNumber`=?", sel,
	)

	q := queries.Raw(query, maintenancePlan, maintenanceItem, maintenancePlanCallNumber)

	err := q.Bind(ctx, exec, sapMaintenanceItemCallObjectsDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_maintenance_item_call_objects_data")
	}

	if err = sapMaintenanceItemCallObjectsDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapMaintenanceItemCallObjectsDatumObj, err
	}

	return sapMaintenanceItemCallObjectsDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapMaintenanceItemCallObjectsDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_maintenance_item_call_objects_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapMaintenanceItemCallObjectsDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapMaintenanceItemCallObjectsDatumInsertCacheMut.RLock()
	cache, cached := sapMaintenanceItemCallObjectsDatumInsertCache[key]
	sapMaintenanceItemCallObjectsDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapMaintenanceItemCallObjectsDatumAllColumns,
			sapMaintenanceItemCallObjectsDatumColumnsWithDefault,
			sapMaintenanceItemCallObjectsDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapMaintenanceItemCallObjectsDatumType, sapMaintenanceItemCallObjectsDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapMaintenanceItemCallObjectsDatumType, sapMaintenanceItemCallObjectsDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_maintenance_item_call_objects_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_maintenance_item_call_objects_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_maintenance_item_call_objects_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_maintenance_item_call_objects_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.MaintenancePlan,
		o.MaintenanceItem,
		o.MaintenancePlanCallNumber,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_maintenance_item_call_objects_data")
	}

CacheNoHooks:
	if !cached {
		sapMaintenanceItemCallObjectsDatumInsertCacheMut.Lock()
		sapMaintenanceItemCallObjectsDatumInsertCache[key] = cache
		sapMaintenanceItemCallObjectsDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapMaintenanceItemCallObjectsDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapMaintenanceItemCallObjectsDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapMaintenanceItemCallObjectsDatumUpdateCacheMut.RLock()
	cache, cached := sapMaintenanceItemCallObjectsDatumUpdateCache[key]
	sapMaintenanceItemCallObjectsDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapMaintenanceItemCallObjectsDatumAllColumns,
			sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_maintenance_item_call_objects_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_maintenance_item_call_objects_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapMaintenanceItemCallObjectsDatumType, sapMaintenanceItemCallObjectsDatumMapping, append(wl, sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_maintenance_item_call_objects_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_maintenance_item_call_objects_data")
	}

	if !cached {
		sapMaintenanceItemCallObjectsDatumUpdateCacheMut.Lock()
		sapMaintenanceItemCallObjectsDatumUpdateCache[key] = cache
		sapMaintenanceItemCallObjectsDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapMaintenanceItemCallObjectsDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_maintenance_item_call_objects_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_maintenance_item_call_objects_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapMaintenanceItemCallObjectsDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaintenanceItemCallObjectsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_maintenance_item_call_objects_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapMaintenanceItemCallObjectsDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapMaintenanceItemCallObjectsDatum")
	}
	return rowsAff, nil
}

var mySQLSapMaintenanceItemCallObjectsDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapMaintenanceItemCallObjectsDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_maintenance_item_call_objects_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapMaintenanceItemCallObjectsDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapMaintenanceItemCallObjectsDatumUniqueColumns, o)

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

	sapMaintenanceItemCallObjectsDatumUpsertCacheMut.RLock()
	cache, cached := sapMaintenanceItemCallObjectsDatumUpsertCache[key]
	sapMaintenanceItemCallObjectsDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapMaintenanceItemCallObjectsDatumAllColumns,
			sapMaintenanceItemCallObjectsDatumColumnsWithDefault,
			sapMaintenanceItemCallObjectsDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sapMaintenanceItemCallObjectsDatumAllColumns,
			sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_maintenance_item_call_objects_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_maintenance_item_call_objects_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_maintenance_item_call_objects_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapMaintenanceItemCallObjectsDatumType, sapMaintenanceItemCallObjectsDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapMaintenanceItemCallObjectsDatumType, sapMaintenanceItemCallObjectsDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_maintenance_item_call_objects_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapMaintenanceItemCallObjectsDatumType, sapMaintenanceItemCallObjectsDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_maintenance_item_call_objects_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_maintenance_item_call_objects_data")
	}

CacheNoHooks:
	if !cached {
		sapMaintenanceItemCallObjectsDatumUpsertCacheMut.Lock()
		sapMaintenanceItemCallObjectsDatumUpsertCache[key] = cache
		sapMaintenanceItemCallObjectsDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapMaintenanceItemCallObjectsDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapMaintenanceItemCallObjectsDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapMaintenanceItemCallObjectsDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapMaintenanceItemCallObjectsDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_maintenance_item_call_objects_data` WHERE `MaintenancePlan`=? AND `MaintenanceItem`=? AND `MaintenancePlanCallNumber`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_maintenance_item_call_objects_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_maintenance_item_call_objects_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapMaintenanceItemCallObjectsDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapMaintenanceItemCallObjectsDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_maintenance_item_call_objects_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_maintenance_item_call_objects_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapMaintenanceItemCallObjectsDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapMaintenanceItemCallObjectsDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaintenanceItemCallObjectsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_maintenance_item_call_objects_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapMaintenanceItemCallObjectsDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_maintenance_item_call_objects_data")
	}

	if len(sapMaintenanceItemCallObjectsDatumAfterDeleteHooks) != 0 {
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
func (o *SapMaintenanceItemCallObjectsDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapMaintenanceItemCallObjectsDatum(ctx, exec, o.MaintenancePlan, o.MaintenanceItem, o.MaintenancePlanCallNumber)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapMaintenanceItemCallObjectsDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapMaintenanceItemCallObjectsDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapMaintenanceItemCallObjectsDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_maintenance_item_call_objects_data`.* FROM `sap_maintenance_item_call_objects_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapMaintenanceItemCallObjectsDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapMaintenanceItemCallObjectsDatumSlice")
	}

	*o = slice

	return nil
}

// SapMaintenanceItemCallObjectsDatumExists checks if the SapMaintenanceItemCallObjectsDatum row exists.
func SapMaintenanceItemCallObjectsDatumExists(ctx context.Context, exec boil.ContextExecutor, maintenancePlan string, maintenanceItem string, maintenancePlanCallNumber int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_maintenance_item_call_objects_data` where `MaintenancePlan`=? AND `MaintenanceItem`=? AND `MaintenancePlanCallNumber`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, maintenancePlan, maintenanceItem, maintenancePlanCallNumber)
	}
	row := exec.QueryRowContext(ctx, sql, maintenancePlan, maintenanceItem, maintenancePlanCallNumber)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_maintenance_item_call_objects_data exists")
	}

	return exists, nil
}
