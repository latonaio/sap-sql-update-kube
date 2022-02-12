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

// SapWorkCenterDatum is an object representing the database table.
type SapWorkCenterDatum struct {
	WorkCenterInternalID         string      `boil:"WorkCenterInternalID" json:"WorkCenterInternalID" toml:"WorkCenterInternalID" yaml:"WorkCenterInternalID"`
	WorkCenterTypeCode           string      `boil:"WorkCenterTypeCode" json:"WorkCenterTypeCode" toml:"WorkCenterTypeCode" yaml:"WorkCenterTypeCode"`
	WorkCenter                   null.String `boil:"WorkCenter" json:"WorkCenter,omitempty" toml:"WorkCenter" yaml:"WorkCenter,omitempty"`
	WorkCenterDesc               null.String `boil:"WorkCenter_desc" json:"WorkCenter_desc,omitempty" toml:"WorkCenter_desc" yaml:"WorkCenter_desc,omitempty"`
	Plant                        null.String `boil:"Plant" json:"Plant,omitempty" toml:"Plant" yaml:"Plant,omitempty"`
	WorkCenterCategoryCode       null.String `boil:"WorkCenterCategoryCode" json:"WorkCenterCategoryCode,omitempty" toml:"WorkCenterCategoryCode" yaml:"WorkCenterCategoryCode,omitempty"`
	WorkCenterResponsible        null.String `boil:"WorkCenterResponsible" json:"WorkCenterResponsible,omitempty" toml:"WorkCenterResponsible" yaml:"WorkCenterResponsible,omitempty"`
	SupplyArea                   null.String `boil:"SupplyArea" json:"SupplyArea,omitempty" toml:"SupplyArea" yaml:"SupplyArea,omitempty"`
	WorkCenterUsage              null.String `boil:"WorkCenterUsage" json:"WorkCenterUsage,omitempty" toml:"WorkCenterUsage" yaml:"WorkCenterUsage,omitempty"`
	MatlCompIsMarkedForBackflush null.Bool   `boil:"MatlCompIsMarkedForBackflush" json:"MatlCompIsMarkedForBackflush,omitempty" toml:"MatlCompIsMarkedForBackflush" yaml:"MatlCompIsMarkedForBackflush,omitempty"`
	WorkCenterLocation           null.String `boil:"WorkCenterLocation" json:"WorkCenterLocation,omitempty" toml:"WorkCenterLocation" yaml:"WorkCenterLocation,omitempty"`
	CapacityInternalID           null.String `boil:"CapacityInternalID" json:"CapacityInternalID,omitempty" toml:"CapacityInternalID" yaml:"CapacityInternalID,omitempty"`
	CapacityCategoryCode         null.String `boil:"CapacityCategoryCode" json:"CapacityCategoryCode,omitempty" toml:"CapacityCategoryCode" yaml:"CapacityCategoryCode,omitempty"`
	ValidityStartDate            null.String `boil:"ValidityStartDate" json:"ValidityStartDate,omitempty" toml:"ValidityStartDate" yaml:"ValidityStartDate,omitempty"`
	ValidityEndDate              null.String `boil:"ValidityEndDate" json:"ValidityEndDate,omitempty" toml:"ValidityEndDate" yaml:"ValidityEndDate,omitempty"`
	WorkCenterIsToBeDeleted      null.Bool   `boil:"WorkCenterIsToBeDeleted" json:"WorkCenterIsToBeDeleted,omitempty" toml:"WorkCenterIsToBeDeleted" yaml:"WorkCenterIsToBeDeleted,omitempty"`

	R *sapWorkCenterDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapWorkCenterDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapWorkCenterDatumColumns = struct {
	WorkCenterInternalID         string
	WorkCenterTypeCode           string
	WorkCenter                   string
	WorkCenterDesc               string
	Plant                        string
	WorkCenterCategoryCode       string
	WorkCenterResponsible        string
	SupplyArea                   string
	WorkCenterUsage              string
	MatlCompIsMarkedForBackflush string
	WorkCenterLocation           string
	CapacityInternalID           string
	CapacityCategoryCode         string
	ValidityStartDate            string
	ValidityEndDate              string
	WorkCenterIsToBeDeleted      string
}{
	WorkCenterInternalID:         "WorkCenterInternalID",
	WorkCenterTypeCode:           "WorkCenterTypeCode",
	WorkCenter:                   "WorkCenter",
	WorkCenterDesc:               "WorkCenter_desc",
	Plant:                        "Plant",
	WorkCenterCategoryCode:       "WorkCenterCategoryCode",
	WorkCenterResponsible:        "WorkCenterResponsible",
	SupplyArea:                   "SupplyArea",
	WorkCenterUsage:              "WorkCenterUsage",
	MatlCompIsMarkedForBackflush: "MatlCompIsMarkedForBackflush",
	WorkCenterLocation:           "WorkCenterLocation",
	CapacityInternalID:           "CapacityInternalID",
	CapacityCategoryCode:         "CapacityCategoryCode",
	ValidityStartDate:            "ValidityStartDate",
	ValidityEndDate:              "ValidityEndDate",
	WorkCenterIsToBeDeleted:      "WorkCenterIsToBeDeleted",
}

var SapWorkCenterDatumTableColumns = struct {
	WorkCenterInternalID         string
	WorkCenterTypeCode           string
	WorkCenter                   string
	WorkCenterDesc               string
	Plant                        string
	WorkCenterCategoryCode       string
	WorkCenterResponsible        string
	SupplyArea                   string
	WorkCenterUsage              string
	MatlCompIsMarkedForBackflush string
	WorkCenterLocation           string
	CapacityInternalID           string
	CapacityCategoryCode         string
	ValidityStartDate            string
	ValidityEndDate              string
	WorkCenterIsToBeDeleted      string
}{
	WorkCenterInternalID:         "sap_work_center_data.WorkCenterInternalID",
	WorkCenterTypeCode:           "sap_work_center_data.WorkCenterTypeCode",
	WorkCenter:                   "sap_work_center_data.WorkCenter",
	WorkCenterDesc:               "sap_work_center_data.WorkCenter_desc",
	Plant:                        "sap_work_center_data.Plant",
	WorkCenterCategoryCode:       "sap_work_center_data.WorkCenterCategoryCode",
	WorkCenterResponsible:        "sap_work_center_data.WorkCenterResponsible",
	SupplyArea:                   "sap_work_center_data.SupplyArea",
	WorkCenterUsage:              "sap_work_center_data.WorkCenterUsage",
	MatlCompIsMarkedForBackflush: "sap_work_center_data.MatlCompIsMarkedForBackflush",
	WorkCenterLocation:           "sap_work_center_data.WorkCenterLocation",
	CapacityInternalID:           "sap_work_center_data.CapacityInternalID",
	CapacityCategoryCode:         "sap_work_center_data.CapacityCategoryCode",
	ValidityStartDate:            "sap_work_center_data.ValidityStartDate",
	ValidityEndDate:              "sap_work_center_data.ValidityEndDate",
	WorkCenterIsToBeDeleted:      "sap_work_center_data.WorkCenterIsToBeDeleted",
}

// Generated where

var SapWorkCenterDatumWhere = struct {
	WorkCenterInternalID         whereHelperstring
	WorkCenterTypeCode           whereHelperstring
	WorkCenter                   whereHelpernull_String
	WorkCenterDesc               whereHelpernull_String
	Plant                        whereHelpernull_String
	WorkCenterCategoryCode       whereHelpernull_String
	WorkCenterResponsible        whereHelpernull_String
	SupplyArea                   whereHelpernull_String
	WorkCenterUsage              whereHelpernull_String
	MatlCompIsMarkedForBackflush whereHelpernull_Bool
	WorkCenterLocation           whereHelpernull_String
	CapacityInternalID           whereHelpernull_String
	CapacityCategoryCode         whereHelpernull_String
	ValidityStartDate            whereHelpernull_String
	ValidityEndDate              whereHelpernull_String
	WorkCenterIsToBeDeleted      whereHelpernull_Bool
}{
	WorkCenterInternalID:         whereHelperstring{field: "`sap_work_center_data`.`WorkCenterInternalID`"},
	WorkCenterTypeCode:           whereHelperstring{field: "`sap_work_center_data`.`WorkCenterTypeCode`"},
	WorkCenter:                   whereHelpernull_String{field: "`sap_work_center_data`.`WorkCenter`"},
	WorkCenterDesc:               whereHelpernull_String{field: "`sap_work_center_data`.`WorkCenter_desc`"},
	Plant:                        whereHelpernull_String{field: "`sap_work_center_data`.`Plant`"},
	WorkCenterCategoryCode:       whereHelpernull_String{field: "`sap_work_center_data`.`WorkCenterCategoryCode`"},
	WorkCenterResponsible:        whereHelpernull_String{field: "`sap_work_center_data`.`WorkCenterResponsible`"},
	SupplyArea:                   whereHelpernull_String{field: "`sap_work_center_data`.`SupplyArea`"},
	WorkCenterUsage:              whereHelpernull_String{field: "`sap_work_center_data`.`WorkCenterUsage`"},
	MatlCompIsMarkedForBackflush: whereHelpernull_Bool{field: "`sap_work_center_data`.`MatlCompIsMarkedForBackflush`"},
	WorkCenterLocation:           whereHelpernull_String{field: "`sap_work_center_data`.`WorkCenterLocation`"},
	CapacityInternalID:           whereHelpernull_String{field: "`sap_work_center_data`.`CapacityInternalID`"},
	CapacityCategoryCode:         whereHelpernull_String{field: "`sap_work_center_data`.`CapacityCategoryCode`"},
	ValidityStartDate:            whereHelpernull_String{field: "`sap_work_center_data`.`ValidityStartDate`"},
	ValidityEndDate:              whereHelpernull_String{field: "`sap_work_center_data`.`ValidityEndDate`"},
	WorkCenterIsToBeDeleted:      whereHelpernull_Bool{field: "`sap_work_center_data`.`WorkCenterIsToBeDeleted`"},
}

// SapWorkCenterDatumRels is where relationship names are stored.
var SapWorkCenterDatumRels = struct {
}{}

// sapWorkCenterDatumR is where relationships are stored.
type sapWorkCenterDatumR struct {
}

// NewStruct creates a new relationship struct
func (*sapWorkCenterDatumR) NewStruct() *sapWorkCenterDatumR {
	return &sapWorkCenterDatumR{}
}

// sapWorkCenterDatumL is where Load methods for each relationship are stored.
type sapWorkCenterDatumL struct{}

var (
	sapWorkCenterDatumAllColumns            = []string{"WorkCenterInternalID", "WorkCenterTypeCode", "WorkCenter", "WorkCenter_desc", "Plant", "WorkCenterCategoryCode", "WorkCenterResponsible", "SupplyArea", "WorkCenterUsage", "MatlCompIsMarkedForBackflush", "WorkCenterLocation", "CapacityInternalID", "CapacityCategoryCode", "ValidityStartDate", "ValidityEndDate", "WorkCenterIsToBeDeleted"}
	sapWorkCenterDatumColumnsWithoutDefault = []string{"WorkCenterInternalID", "WorkCenterTypeCode", "WorkCenter", "WorkCenter_desc", "Plant", "WorkCenterCategoryCode", "WorkCenterResponsible", "SupplyArea", "WorkCenterUsage", "MatlCompIsMarkedForBackflush", "WorkCenterLocation", "CapacityInternalID", "CapacityCategoryCode", "ValidityStartDate", "ValidityEndDate", "WorkCenterIsToBeDeleted"}
	sapWorkCenterDatumColumnsWithDefault    = []string{}
	sapWorkCenterDatumPrimaryKeyColumns     = []string{"WorkCenterInternalID", "WorkCenterTypeCode"}
	sapWorkCenterDatumGeneratedColumns      = []string{}
)

type (
	// SapWorkCenterDatumSlice is an alias for a slice of pointers to SapWorkCenterDatum.
	// This should almost always be used instead of []SapWorkCenterDatum.
	SapWorkCenterDatumSlice []*SapWorkCenterDatum
	// SapWorkCenterDatumHook is the signature for custom SapWorkCenterDatum hook methods
	SapWorkCenterDatumHook func(context.Context, boil.ContextExecutor, *SapWorkCenterDatum) error

	sapWorkCenterDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapWorkCenterDatumType                 = reflect.TypeOf(&SapWorkCenterDatum{})
	sapWorkCenterDatumMapping              = queries.MakeStructMapping(sapWorkCenterDatumType)
	sapWorkCenterDatumPrimaryKeyMapping, _ = queries.BindMapping(sapWorkCenterDatumType, sapWorkCenterDatumMapping, sapWorkCenterDatumPrimaryKeyColumns)
	sapWorkCenterDatumInsertCacheMut       sync.RWMutex
	sapWorkCenterDatumInsertCache          = make(map[string]insertCache)
	sapWorkCenterDatumUpdateCacheMut       sync.RWMutex
	sapWorkCenterDatumUpdateCache          = make(map[string]updateCache)
	sapWorkCenterDatumUpsertCacheMut       sync.RWMutex
	sapWorkCenterDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapWorkCenterDatumAfterSelectHooks []SapWorkCenterDatumHook

var sapWorkCenterDatumBeforeInsertHooks []SapWorkCenterDatumHook
var sapWorkCenterDatumAfterInsertHooks []SapWorkCenterDatumHook

var sapWorkCenterDatumBeforeUpdateHooks []SapWorkCenterDatumHook
var sapWorkCenterDatumAfterUpdateHooks []SapWorkCenterDatumHook

var sapWorkCenterDatumBeforeDeleteHooks []SapWorkCenterDatumHook
var sapWorkCenterDatumAfterDeleteHooks []SapWorkCenterDatumHook

var sapWorkCenterDatumBeforeUpsertHooks []SapWorkCenterDatumHook
var sapWorkCenterDatumAfterUpsertHooks []SapWorkCenterDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapWorkCenterDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapWorkCenterDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapWorkCenterDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapWorkCenterDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapWorkCenterDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapWorkCenterDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapWorkCenterDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapWorkCenterDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapWorkCenterDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapWorkCenterDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapWorkCenterDatumHook registers your hook function for all future operations.
func AddSapWorkCenterDatumHook(hookPoint boil.HookPoint, sapWorkCenterDatumHook SapWorkCenterDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sapWorkCenterDatumAfterSelectHooks = append(sapWorkCenterDatumAfterSelectHooks, sapWorkCenterDatumHook)
	case boil.BeforeInsertHook:
		sapWorkCenterDatumBeforeInsertHooks = append(sapWorkCenterDatumBeforeInsertHooks, sapWorkCenterDatumHook)
	case boil.AfterInsertHook:
		sapWorkCenterDatumAfterInsertHooks = append(sapWorkCenterDatumAfterInsertHooks, sapWorkCenterDatumHook)
	case boil.BeforeUpdateHook:
		sapWorkCenterDatumBeforeUpdateHooks = append(sapWorkCenterDatumBeforeUpdateHooks, sapWorkCenterDatumHook)
	case boil.AfterUpdateHook:
		sapWorkCenterDatumAfterUpdateHooks = append(sapWorkCenterDatumAfterUpdateHooks, sapWorkCenterDatumHook)
	case boil.BeforeDeleteHook:
		sapWorkCenterDatumBeforeDeleteHooks = append(sapWorkCenterDatumBeforeDeleteHooks, sapWorkCenterDatumHook)
	case boil.AfterDeleteHook:
		sapWorkCenterDatumAfterDeleteHooks = append(sapWorkCenterDatumAfterDeleteHooks, sapWorkCenterDatumHook)
	case boil.BeforeUpsertHook:
		sapWorkCenterDatumBeforeUpsertHooks = append(sapWorkCenterDatumBeforeUpsertHooks, sapWorkCenterDatumHook)
	case boil.AfterUpsertHook:
		sapWorkCenterDatumAfterUpsertHooks = append(sapWorkCenterDatumAfterUpsertHooks, sapWorkCenterDatumHook)
	}
}

// One returns a single sapWorkCenterDatum record from the query.
func (q sapWorkCenterDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapWorkCenterDatum, error) {
	o := &SapWorkCenterDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_work_center_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapWorkCenterDatum records from the query.
func (q sapWorkCenterDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapWorkCenterDatumSlice, error) {
	var o []*SapWorkCenterDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapWorkCenterDatum slice")
	}

	if len(sapWorkCenterDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapWorkCenterDatum records in the query.
func (q sapWorkCenterDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_work_center_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapWorkCenterDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_work_center_data exists")
	}

	return count > 0, nil
}

// SapWorkCenterData retrieves all the records using an executor.
func SapWorkCenterData(mods ...qm.QueryMod) sapWorkCenterDatumQuery {
	mods = append(mods, qm.From("`sap_work_center_data`"))
	return sapWorkCenterDatumQuery{NewQuery(mods...)}
}

// FindSapWorkCenterDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapWorkCenterDatum(ctx context.Context, exec boil.ContextExecutor, workCenterInternalID string, workCenterTypeCode string, selectCols ...string) (*SapWorkCenterDatum, error) {
	sapWorkCenterDatumObj := &SapWorkCenterDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_work_center_data` where `WorkCenterInternalID`=? AND `WorkCenterTypeCode`=?", sel,
	)

	q := queries.Raw(query, workCenterInternalID, workCenterTypeCode)

	err := q.Bind(ctx, exec, sapWorkCenterDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_work_center_data")
	}

	if err = sapWorkCenterDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapWorkCenterDatumObj, err
	}

	return sapWorkCenterDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapWorkCenterDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_work_center_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapWorkCenterDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapWorkCenterDatumInsertCacheMut.RLock()
	cache, cached := sapWorkCenterDatumInsertCache[key]
	sapWorkCenterDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapWorkCenterDatumAllColumns,
			sapWorkCenterDatumColumnsWithDefault,
			sapWorkCenterDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapWorkCenterDatumType, sapWorkCenterDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapWorkCenterDatumType, sapWorkCenterDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_work_center_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_work_center_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_work_center_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapWorkCenterDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_work_center_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.WorkCenterInternalID,
		o.WorkCenterTypeCode,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_work_center_data")
	}

CacheNoHooks:
	if !cached {
		sapWorkCenterDatumInsertCacheMut.Lock()
		sapWorkCenterDatumInsertCache[key] = cache
		sapWorkCenterDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapWorkCenterDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapWorkCenterDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapWorkCenterDatumUpdateCacheMut.RLock()
	cache, cached := sapWorkCenterDatumUpdateCache[key]
	sapWorkCenterDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapWorkCenterDatumAllColumns,
			sapWorkCenterDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_work_center_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_work_center_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapWorkCenterDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapWorkCenterDatumType, sapWorkCenterDatumMapping, append(wl, sapWorkCenterDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_work_center_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_work_center_data")
	}

	if !cached {
		sapWorkCenterDatumUpdateCacheMut.Lock()
		sapWorkCenterDatumUpdateCache[key] = cache
		sapWorkCenterDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapWorkCenterDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_work_center_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_work_center_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapWorkCenterDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapWorkCenterDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_work_center_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapWorkCenterDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapWorkCenterDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapWorkCenterDatum")
	}
	return rowsAff, nil
}

var mySQLSapWorkCenterDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapWorkCenterDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_work_center_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapWorkCenterDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapWorkCenterDatumUniqueColumns, o)

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

	sapWorkCenterDatumUpsertCacheMut.RLock()
	cache, cached := sapWorkCenterDatumUpsertCache[key]
	sapWorkCenterDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapWorkCenterDatumAllColumns,
			sapWorkCenterDatumColumnsWithDefault,
			sapWorkCenterDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sapWorkCenterDatumAllColumns,
			sapWorkCenterDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_work_center_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_work_center_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_work_center_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapWorkCenterDatumType, sapWorkCenterDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapWorkCenterDatumType, sapWorkCenterDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_work_center_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapWorkCenterDatumType, sapWorkCenterDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_work_center_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_work_center_data")
	}

CacheNoHooks:
	if !cached {
		sapWorkCenterDatumUpsertCacheMut.Lock()
		sapWorkCenterDatumUpsertCache[key] = cache
		sapWorkCenterDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapWorkCenterDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapWorkCenterDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapWorkCenterDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapWorkCenterDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_work_center_data` WHERE `WorkCenterInternalID`=? AND `WorkCenterTypeCode`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_work_center_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_work_center_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapWorkCenterDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapWorkCenterDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_work_center_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_work_center_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapWorkCenterDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapWorkCenterDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapWorkCenterDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_work_center_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapWorkCenterDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapWorkCenterDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_work_center_data")
	}

	if len(sapWorkCenterDatumAfterDeleteHooks) != 0 {
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
func (o *SapWorkCenterDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapWorkCenterDatum(ctx, exec, o.WorkCenterInternalID, o.WorkCenterTypeCode)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapWorkCenterDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapWorkCenterDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapWorkCenterDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_work_center_data`.* FROM `sap_work_center_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapWorkCenterDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapWorkCenterDatumSlice")
	}

	*o = slice

	return nil
}

// SapWorkCenterDatumExists checks if the SapWorkCenterDatum row exists.
func SapWorkCenterDatumExists(ctx context.Context, exec boil.ContextExecutor, workCenterInternalID string, workCenterTypeCode string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_work_center_data` where `WorkCenterInternalID`=? AND `WorkCenterTypeCode`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, workCenterInternalID, workCenterTypeCode)
	}
	row := exec.QueryRowContext(ctx, sql, workCenterInternalID, workCenterTypeCode)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_work_center_data exists")
	}

	return exists, nil
}
