// Code generated by entc, DO NOT EDIT.

package construction

import (
	"time"
)

const (
	// Label holds the string label denoting the construction type in the database.
	Label = "construction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldX holds the string denoting the x field in the database.
	FieldX = "x"
	// FieldY holds the string denoting the y field in the database.
	FieldY = "y"
	// FieldRawProduction holds the string denoting the raw_production field in the database.
	FieldRawProduction = "raw_production"
	// FieldProduction holds the string denoting the production field in the database.
	FieldProduction = "production"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldModifier holds the string denoting the modifier field in the database.
	FieldModifier = "modifier"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"
	// FieldNeedRefresh holds the string denoting the need_refresh field in the database.
	FieldNeedRefresh = "need_refresh"

	// EdgeCity holds the string denoting the city edge name in mutations.
	EdgeCity = "city"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeQueue holds the string denoting the queue edge name in mutations.
	EdgeQueue = "queue"
	// EdgeAffects holds the string denoting the affects edge name in mutations.
	EdgeAffects = "affects"
	// EdgeAffectedBy holds the string denoting the affected_by edge name in mutations.
	EdgeAffectedBy = "affected_by"

	// Table holds the table name of the construction in the database.
	Table = "constructions"
	// CityTable is the table the holds the city relation/edge.
	CityTable = "constructions"
	// CityInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CityInverseTable = "cities"
	// CityColumn is the table column denoting the city relation/edge.
	CityColumn = "city_constructions"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "constructions"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_constructions"
	// QueueTable is the table the holds the queue relation/edge.
	QueueTable = "queue_items"
	// QueueInverseTable is the table name for the QueueItem entity.
	// It exists in this package in order to avoid circular dependency with the "queueitem" package.
	QueueInverseTable = "queue_items"
	// QueueColumn is the table column denoting the queue relation/edge.
	QueueColumn = "construction_queue"
	// AffectsTable is the table the holds the affects relation/edge. The primary key declared below.
	AffectsTable = "construction_affects"
	// AffectedByTable is the table the holds the affected_by relation/edge. The primary key declared below.
	AffectedByTable = "construction_affects"
)

// Columns holds all SQL columns for construction fields.
var Columns = []string{
	FieldID,
	FieldX,
	FieldY,
	FieldRawProduction,
	FieldProduction,
	FieldType,
	FieldLevel,
	FieldModifier,
	FieldLastUpdated,
	FieldNeedRefresh,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Construction type.
var ForeignKeys = []string{
	"city_constructions",
	"user_constructions",
}

var (
	// AffectsPrimaryKey and AffectsColumn2 are the table columns denoting the
	// primary key for the affects relation (M2M).
	AffectsPrimaryKey = []string{"construction_id", "affected_by_id"}
	// AffectedByPrimaryKey and AffectedByColumn2 are the table columns denoting the
	// primary key for the affected_by relation (M2M).
	AffectedByPrimaryKey = []string{"construction_id", "affected_by_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRawProduction holds the default value on creation for the raw_production field.
	DefaultRawProduction float64
	// DefaultProduction holds the default value on creation for the production field.
	DefaultProduction float64
	// DefaultType holds the default value on creation for the type field.
	DefaultType int
	// DefaultLevel holds the default value on creation for the level field.
	DefaultLevel int
	// DefaultModifier holds the default value on creation for the modifier field.
	DefaultModifier float64
	// DefaultLastUpdated holds the default value on creation for the last_updated field.
	DefaultLastUpdated func() time.Time
	// DefaultNeedRefresh holds the default value on creation for the need_refresh field.
	DefaultNeedRefresh bool
)
