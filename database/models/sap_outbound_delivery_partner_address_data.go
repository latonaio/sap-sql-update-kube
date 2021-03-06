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

// SapOutboundDeliveryPartnerAddressDatum is an object representing the database table.
type SapOutboundDeliveryPartnerAddressDatum struct {
	AddressID              string      `boil:"AddressID" json:"AddressID" toml:"AddressID" yaml:"AddressID"`
	BusinessPartnerName1   null.String `boil:"BusinessPartnerName1" json:"BusinessPartnerName1,omitempty" toml:"BusinessPartnerName1" yaml:"BusinessPartnerName1,omitempty"`
	Country                null.String `boil:"Country" json:"Country,omitempty" toml:"Country" yaml:"Country,omitempty"`
	Region                 null.String `boil:"Region" json:"Region,omitempty" toml:"Region" yaml:"Region,omitempty"`
	StreetName             null.String `boil:"StreetName" json:"StreetName,omitempty" toml:"StreetName" yaml:"StreetName,omitempty"`
	CityName               null.String `boil:"CityName" json:"CityName,omitempty" toml:"CityName" yaml:"CityName,omitempty"`
	PostalCode             null.String `boil:"PostalCode" json:"PostalCode,omitempty" toml:"PostalCode" yaml:"PostalCode,omitempty"`
	CorrespondenceLanguage null.String `boil:"CorrespondenceLanguage" json:"CorrespondenceLanguage,omitempty" toml:"CorrespondenceLanguage" yaml:"CorrespondenceLanguage,omitempty"`
	FaxNumber              null.String `boil:"FaxNumber" json:"FaxNumber,omitempty" toml:"FaxNumber" yaml:"FaxNumber,omitempty"`
	PhoneNumber            null.String `boil:"PhoneNumber" json:"PhoneNumber,omitempty" toml:"PhoneNumber" yaml:"PhoneNumber,omitempty"`
	SDDocument             null.String `boil:"SDDocument" json:"SDDocument,omitempty" toml:"SDDocument" yaml:"SDDocument,omitempty"`
	SDDocumentItem         null.String `boil:"SDDocumentItem" json:"SDDocumentItem,omitempty" toml:"SDDocumentItem" yaml:"SDDocumentItem,omitempty"`

	R *sapOutboundDeliveryPartnerAddressDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sapOutboundDeliveryPartnerAddressDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SapOutboundDeliveryPartnerAddressDatumColumns = struct {
	AddressID              string
	BusinessPartnerName1   string
	Country                string
	Region                 string
	StreetName             string
	CityName               string
	PostalCode             string
	CorrespondenceLanguage string
	FaxNumber              string
	PhoneNumber            string
	SDDocument             string
	SDDocumentItem         string
}{
	AddressID:              "AddressID",
	BusinessPartnerName1:   "BusinessPartnerName1",
	Country:                "Country",
	Region:                 "Region",
	StreetName:             "StreetName",
	CityName:               "CityName",
	PostalCode:             "PostalCode",
	CorrespondenceLanguage: "CorrespondenceLanguage",
	FaxNumber:              "FaxNumber",
	PhoneNumber:            "PhoneNumber",
	SDDocument:             "SDDocument",
	SDDocumentItem:         "SDDocumentItem",
}

var SapOutboundDeliveryPartnerAddressDatumTableColumns = struct {
	AddressID              string
	BusinessPartnerName1   string
	Country                string
	Region                 string
	StreetName             string
	CityName               string
	PostalCode             string
	CorrespondenceLanguage string
	FaxNumber              string
	PhoneNumber            string
	SDDocument             string
	SDDocumentItem         string
}{
	AddressID:              "sap_outbound_delivery_partner_address_data.AddressID",
	BusinessPartnerName1:   "sap_outbound_delivery_partner_address_data.BusinessPartnerName1",
	Country:                "sap_outbound_delivery_partner_address_data.Country",
	Region:                 "sap_outbound_delivery_partner_address_data.Region",
	StreetName:             "sap_outbound_delivery_partner_address_data.StreetName",
	CityName:               "sap_outbound_delivery_partner_address_data.CityName",
	PostalCode:             "sap_outbound_delivery_partner_address_data.PostalCode",
	CorrespondenceLanguage: "sap_outbound_delivery_partner_address_data.CorrespondenceLanguage",
	FaxNumber:              "sap_outbound_delivery_partner_address_data.FaxNumber",
	PhoneNumber:            "sap_outbound_delivery_partner_address_data.PhoneNumber",
	SDDocument:             "sap_outbound_delivery_partner_address_data.SDDocument",
	SDDocumentItem:         "sap_outbound_delivery_partner_address_data.SDDocumentItem",
}

// Generated where

var SapOutboundDeliveryPartnerAddressDatumWhere = struct {
	AddressID              whereHelperstring
	BusinessPartnerName1   whereHelpernull_String
	Country                whereHelpernull_String
	Region                 whereHelpernull_String
	StreetName             whereHelpernull_String
	CityName               whereHelpernull_String
	PostalCode             whereHelpernull_String
	CorrespondenceLanguage whereHelpernull_String
	FaxNumber              whereHelpernull_String
	PhoneNumber            whereHelpernull_String
	SDDocument             whereHelpernull_String
	SDDocumentItem         whereHelpernull_String
}{
	AddressID:              whereHelperstring{field: "`sap_outbound_delivery_partner_address_data`.`AddressID`"},
	BusinessPartnerName1:   whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`BusinessPartnerName1`"},
	Country:                whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`Country`"},
	Region:                 whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`Region`"},
	StreetName:             whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`StreetName`"},
	CityName:               whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`CityName`"},
	PostalCode:             whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`PostalCode`"},
	CorrespondenceLanguage: whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`CorrespondenceLanguage`"},
	FaxNumber:              whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`FaxNumber`"},
	PhoneNumber:            whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`PhoneNumber`"},
	SDDocument:             whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`SDDocument`"},
	SDDocumentItem:         whereHelpernull_String{field: "`sap_outbound_delivery_partner_address_data`.`SDDocumentItem`"},
}

// SapOutboundDeliveryPartnerAddressDatumRels is where relationship names are stored.
var SapOutboundDeliveryPartnerAddressDatumRels = struct {
}{}

// sapOutboundDeliveryPartnerAddressDatumR is where relationships are stored.
type sapOutboundDeliveryPartnerAddressDatumR struct {
}

// NewStruct creates a new relationship struct
func (*sapOutboundDeliveryPartnerAddressDatumR) NewStruct() *sapOutboundDeliveryPartnerAddressDatumR {
	return &sapOutboundDeliveryPartnerAddressDatumR{}
}

// sapOutboundDeliveryPartnerAddressDatumL is where Load methods for each relationship are stored.
type sapOutboundDeliveryPartnerAddressDatumL struct{}

var (
	sapOutboundDeliveryPartnerAddressDatumAllColumns            = []string{"AddressID", "BusinessPartnerName1", "Country", "Region", "StreetName", "CityName", "PostalCode", "CorrespondenceLanguage", "FaxNumber", "PhoneNumber", "SDDocument", "SDDocumentItem"}
	sapOutboundDeliveryPartnerAddressDatumColumnsWithoutDefault = []string{"AddressID", "BusinessPartnerName1", "Country", "Region", "StreetName", "CityName", "PostalCode", "CorrespondenceLanguage", "FaxNumber", "PhoneNumber", "SDDocument", "SDDocumentItem"}
	sapOutboundDeliveryPartnerAddressDatumColumnsWithDefault    = []string{}
	sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns     = []string{"AddressID"}
	sapOutboundDeliveryPartnerAddressDatumGeneratedColumns      = []string{}
)

type (
	// SapOutboundDeliveryPartnerAddressDatumSlice is an alias for a slice of pointers to SapOutboundDeliveryPartnerAddressDatum.
	// This should almost always be used instead of []SapOutboundDeliveryPartnerAddressDatum.
	SapOutboundDeliveryPartnerAddressDatumSlice []*SapOutboundDeliveryPartnerAddressDatum
	// SapOutboundDeliveryPartnerAddressDatumHook is the signature for custom SapOutboundDeliveryPartnerAddressDatum hook methods
	SapOutboundDeliveryPartnerAddressDatumHook func(context.Context, boil.ContextExecutor, *SapOutboundDeliveryPartnerAddressDatum) error

	sapOutboundDeliveryPartnerAddressDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sapOutboundDeliveryPartnerAddressDatumType                 = reflect.TypeOf(&SapOutboundDeliveryPartnerAddressDatum{})
	sapOutboundDeliveryPartnerAddressDatumMapping              = queries.MakeStructMapping(sapOutboundDeliveryPartnerAddressDatumType)
	sapOutboundDeliveryPartnerAddressDatumPrimaryKeyMapping, _ = queries.BindMapping(sapOutboundDeliveryPartnerAddressDatumType, sapOutboundDeliveryPartnerAddressDatumMapping, sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns)
	sapOutboundDeliveryPartnerAddressDatumInsertCacheMut       sync.RWMutex
	sapOutboundDeliveryPartnerAddressDatumInsertCache          = make(map[string]insertCache)
	sapOutboundDeliveryPartnerAddressDatumUpdateCacheMut       sync.RWMutex
	sapOutboundDeliveryPartnerAddressDatumUpdateCache          = make(map[string]updateCache)
	sapOutboundDeliveryPartnerAddressDatumUpsertCacheMut       sync.RWMutex
	sapOutboundDeliveryPartnerAddressDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var sapOutboundDeliveryPartnerAddressDatumAfterSelectHooks []SapOutboundDeliveryPartnerAddressDatumHook

var sapOutboundDeliveryPartnerAddressDatumBeforeInsertHooks []SapOutboundDeliveryPartnerAddressDatumHook
var sapOutboundDeliveryPartnerAddressDatumAfterInsertHooks []SapOutboundDeliveryPartnerAddressDatumHook

var sapOutboundDeliveryPartnerAddressDatumBeforeUpdateHooks []SapOutboundDeliveryPartnerAddressDatumHook
var sapOutboundDeliveryPartnerAddressDatumAfterUpdateHooks []SapOutboundDeliveryPartnerAddressDatumHook

var sapOutboundDeliveryPartnerAddressDatumBeforeDeleteHooks []SapOutboundDeliveryPartnerAddressDatumHook
var sapOutboundDeliveryPartnerAddressDatumAfterDeleteHooks []SapOutboundDeliveryPartnerAddressDatumHook

var sapOutboundDeliveryPartnerAddressDatumBeforeUpsertHooks []SapOutboundDeliveryPartnerAddressDatumHook
var sapOutboundDeliveryPartnerAddressDatumAfterUpsertHooks []SapOutboundDeliveryPartnerAddressDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SapOutboundDeliveryPartnerAddressDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range sapOutboundDeliveryPartnerAddressDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSapOutboundDeliveryPartnerAddressDatumHook registers your hook function for all future operations.
func AddSapOutboundDeliveryPartnerAddressDatumHook(hookPoint boil.HookPoint, sapOutboundDeliveryPartnerAddressDatumHook SapOutboundDeliveryPartnerAddressDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		sapOutboundDeliveryPartnerAddressDatumAfterSelectHooks = append(sapOutboundDeliveryPartnerAddressDatumAfterSelectHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.BeforeInsertHook:
		sapOutboundDeliveryPartnerAddressDatumBeforeInsertHooks = append(sapOutboundDeliveryPartnerAddressDatumBeforeInsertHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.AfterInsertHook:
		sapOutboundDeliveryPartnerAddressDatumAfterInsertHooks = append(sapOutboundDeliveryPartnerAddressDatumAfterInsertHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.BeforeUpdateHook:
		sapOutboundDeliveryPartnerAddressDatumBeforeUpdateHooks = append(sapOutboundDeliveryPartnerAddressDatumBeforeUpdateHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.AfterUpdateHook:
		sapOutboundDeliveryPartnerAddressDatumAfterUpdateHooks = append(sapOutboundDeliveryPartnerAddressDatumAfterUpdateHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.BeforeDeleteHook:
		sapOutboundDeliveryPartnerAddressDatumBeforeDeleteHooks = append(sapOutboundDeliveryPartnerAddressDatumBeforeDeleteHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.AfterDeleteHook:
		sapOutboundDeliveryPartnerAddressDatumAfterDeleteHooks = append(sapOutboundDeliveryPartnerAddressDatumAfterDeleteHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.BeforeUpsertHook:
		sapOutboundDeliveryPartnerAddressDatumBeforeUpsertHooks = append(sapOutboundDeliveryPartnerAddressDatumBeforeUpsertHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	case boil.AfterUpsertHook:
		sapOutboundDeliveryPartnerAddressDatumAfterUpsertHooks = append(sapOutboundDeliveryPartnerAddressDatumAfterUpsertHooks, sapOutboundDeliveryPartnerAddressDatumHook)
	}
}

// One returns a single sapOutboundDeliveryPartnerAddressDatum record from the query.
func (q sapOutboundDeliveryPartnerAddressDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SapOutboundDeliveryPartnerAddressDatum, error) {
	o := &SapOutboundDeliveryPartnerAddressDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sap_outbound_delivery_partner_address_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SapOutboundDeliveryPartnerAddressDatum records from the query.
func (q sapOutboundDeliveryPartnerAddressDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (SapOutboundDeliveryPartnerAddressDatumSlice, error) {
	var o []*SapOutboundDeliveryPartnerAddressDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SapOutboundDeliveryPartnerAddressDatum slice")
	}

	if len(sapOutboundDeliveryPartnerAddressDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SapOutboundDeliveryPartnerAddressDatum records in the query.
func (q sapOutboundDeliveryPartnerAddressDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sap_outbound_delivery_partner_address_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sapOutboundDeliveryPartnerAddressDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sap_outbound_delivery_partner_address_data exists")
	}

	return count > 0, nil
}

// SapOutboundDeliveryPartnerAddressData retrieves all the records using an executor.
func SapOutboundDeliveryPartnerAddressData(mods ...qm.QueryMod) sapOutboundDeliveryPartnerAddressDatumQuery {
	mods = append(mods, qm.From("`sap_outbound_delivery_partner_address_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`sap_outbound_delivery_partner_address_data`.*"})
	}

	return sapOutboundDeliveryPartnerAddressDatumQuery{q}
}

// FindSapOutboundDeliveryPartnerAddressDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSapOutboundDeliveryPartnerAddressDatum(ctx context.Context, exec boil.ContextExecutor, addressID string, selectCols ...string) (*SapOutboundDeliveryPartnerAddressDatum, error) {
	sapOutboundDeliveryPartnerAddressDatumObj := &SapOutboundDeliveryPartnerAddressDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sap_outbound_delivery_partner_address_data` where `AddressID`=?", sel,
	)

	q := queries.Raw(query, addressID)

	err := q.Bind(ctx, exec, sapOutboundDeliveryPartnerAddressDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sap_outbound_delivery_partner_address_data")
	}

	if err = sapOutboundDeliveryPartnerAddressDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return sapOutboundDeliveryPartnerAddressDatumObj, err
	}

	return sapOutboundDeliveryPartnerAddressDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SapOutboundDeliveryPartnerAddressDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_outbound_delivery_partner_address_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapOutboundDeliveryPartnerAddressDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sapOutboundDeliveryPartnerAddressDatumInsertCacheMut.RLock()
	cache, cached := sapOutboundDeliveryPartnerAddressDatumInsertCache[key]
	sapOutboundDeliveryPartnerAddressDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sapOutboundDeliveryPartnerAddressDatumAllColumns,
			sapOutboundDeliveryPartnerAddressDatumColumnsWithDefault,
			sapOutboundDeliveryPartnerAddressDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sapOutboundDeliveryPartnerAddressDatumType, sapOutboundDeliveryPartnerAddressDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sapOutboundDeliveryPartnerAddressDatumType, sapOutboundDeliveryPartnerAddressDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sap_outbound_delivery_partner_address_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sap_outbound_delivery_partner_address_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sap_outbound_delivery_partner_address_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into sap_outbound_delivery_partner_address_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.AddressID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_outbound_delivery_partner_address_data")
	}

CacheNoHooks:
	if !cached {
		sapOutboundDeliveryPartnerAddressDatumInsertCacheMut.Lock()
		sapOutboundDeliveryPartnerAddressDatumInsertCache[key] = cache
		sapOutboundDeliveryPartnerAddressDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SapOutboundDeliveryPartnerAddressDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SapOutboundDeliveryPartnerAddressDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	sapOutboundDeliveryPartnerAddressDatumUpdateCacheMut.RLock()
	cache, cached := sapOutboundDeliveryPartnerAddressDatumUpdateCache[key]
	sapOutboundDeliveryPartnerAddressDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sapOutboundDeliveryPartnerAddressDatumAllColumns,
			sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sap_outbound_delivery_partner_address_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sap_outbound_delivery_partner_address_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sapOutboundDeliveryPartnerAddressDatumType, sapOutboundDeliveryPartnerAddressDatumMapping, append(wl, sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sap_outbound_delivery_partner_address_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sap_outbound_delivery_partner_address_data")
	}

	if !cached {
		sapOutboundDeliveryPartnerAddressDatumUpdateCacheMut.Lock()
		sapOutboundDeliveryPartnerAddressDatumUpdateCache[key] = cache
		sapOutboundDeliveryPartnerAddressDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q sapOutboundDeliveryPartnerAddressDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sap_outbound_delivery_partner_address_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sap_outbound_delivery_partner_address_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SapOutboundDeliveryPartnerAddressDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapOutboundDeliveryPartnerAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sap_outbound_delivery_partner_address_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sapOutboundDeliveryPartnerAddressDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sapOutboundDeliveryPartnerAddressDatum")
	}
	return rowsAff, nil
}

var mySQLSapOutboundDeliveryPartnerAddressDatumUniqueColumns = []string{
	"AddressID",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SapOutboundDeliveryPartnerAddressDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sap_outbound_delivery_partner_address_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(sapOutboundDeliveryPartnerAddressDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSapOutboundDeliveryPartnerAddressDatumUniqueColumns, o)

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

	sapOutboundDeliveryPartnerAddressDatumUpsertCacheMut.RLock()
	cache, cached := sapOutboundDeliveryPartnerAddressDatumUpsertCache[key]
	sapOutboundDeliveryPartnerAddressDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sapOutboundDeliveryPartnerAddressDatumAllColumns,
			sapOutboundDeliveryPartnerAddressDatumColumnsWithDefault,
			sapOutboundDeliveryPartnerAddressDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			sapOutboundDeliveryPartnerAddressDatumAllColumns,
			sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sap_outbound_delivery_partner_address_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sap_outbound_delivery_partner_address_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sap_outbound_delivery_partner_address_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sapOutboundDeliveryPartnerAddressDatumType, sapOutboundDeliveryPartnerAddressDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sapOutboundDeliveryPartnerAddressDatumType, sapOutboundDeliveryPartnerAddressDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for sap_outbound_delivery_partner_address_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sapOutboundDeliveryPartnerAddressDatumType, sapOutboundDeliveryPartnerAddressDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sap_outbound_delivery_partner_address_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sap_outbound_delivery_partner_address_data")
	}

CacheNoHooks:
	if !cached {
		sapOutboundDeliveryPartnerAddressDatumUpsertCacheMut.Lock()
		sapOutboundDeliveryPartnerAddressDatumUpsertCache[key] = cache
		sapOutboundDeliveryPartnerAddressDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SapOutboundDeliveryPartnerAddressDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SapOutboundDeliveryPartnerAddressDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SapOutboundDeliveryPartnerAddressDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sapOutboundDeliveryPartnerAddressDatumPrimaryKeyMapping)
	sql := "DELETE FROM `sap_outbound_delivery_partner_address_data` WHERE `AddressID`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sap_outbound_delivery_partner_address_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sap_outbound_delivery_partner_address_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sapOutboundDeliveryPartnerAddressDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sapOutboundDeliveryPartnerAddressDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sap_outbound_delivery_partner_address_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_outbound_delivery_partner_address_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SapOutboundDeliveryPartnerAddressDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(sapOutboundDeliveryPartnerAddressDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapOutboundDeliveryPartnerAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sap_outbound_delivery_partner_address_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sapOutboundDeliveryPartnerAddressDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sap_outbound_delivery_partner_address_data")
	}

	if len(sapOutboundDeliveryPartnerAddressDatumAfterDeleteHooks) != 0 {
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
func (o *SapOutboundDeliveryPartnerAddressDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSapOutboundDeliveryPartnerAddressDatum(ctx, exec, o.AddressID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SapOutboundDeliveryPartnerAddressDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SapOutboundDeliveryPartnerAddressDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sapOutboundDeliveryPartnerAddressDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sap_outbound_delivery_partner_address_data`.* FROM `sap_outbound_delivery_partner_address_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sapOutboundDeliveryPartnerAddressDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SapOutboundDeliveryPartnerAddressDatumSlice")
	}

	*o = slice

	return nil
}

// SapOutboundDeliveryPartnerAddressDatumExists checks if the SapOutboundDeliveryPartnerAddressDatum row exists.
func SapOutboundDeliveryPartnerAddressDatumExists(ctx context.Context, exec boil.ContextExecutor, addressID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sap_outbound_delivery_partner_address_data` where `AddressID`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, addressID)
	}
	row := exec.QueryRowContext(ctx, sql, addressID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sap_outbound_delivery_partner_address_data exists")
	}

	return exists, nil
}
