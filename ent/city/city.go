// Code generated by entc, DO NOT EDIT.

package city

import (
	"time"
)

const (
	// Label holds the string label denoting the city type in the database.
	Label = "city"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldX holds the string denoting the x field in the database.
	FieldX = "x"
	// FieldY holds the string denoting the y field in the database.
	FieldY = "y"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// FieldWoodProduction holds the string denoting the wood_production field in the database.
	FieldWoodProduction = "wood_production"
	// FieldStoneProduction holds the string denoting the stone_production field in the database.
	FieldStoneProduction = "stone_production"
	// FieldIronProduction holds the string denoting the iron_production field in the database.
	FieldIronProduction = "iron_production"
	// FieldFoodProduction holds the string denoting the food_production field in the database.
	FieldFoodProduction = "food_production"
	// FieldWoodStored holds the string denoting the wood_stored field in the database.
	FieldWoodStored = "wood_stored"
	// FieldStoneStored holds the string denoting the stone_stored field in the database.
	FieldStoneStored = "stone_stored"
	// FieldIronStored holds the string denoting the iron_stored field in the database.
	FieldIronStored = "iron_stored"
	// FieldFoodStored holds the string denoting the food_stored field in the database.
	FieldFoodStored = "food_stored"
	// FieldWoodLimit holds the string denoting the wood_limit field in the database.
	FieldWoodLimit = "wood_limit"
	// FieldStoneLimit holds the string denoting the stone_limit field in the database.
	FieldStoneLimit = "stone_limit"
	// FieldIronLimit holds the string denoting the iron_limit field in the database.
	FieldIronLimit = "iron_limit"
	// FieldFoodLimit holds the string denoting the food_limit field in the database.
	FieldFoodLimit = "food_limit"
	// FieldQueueStartedAt holds the string denoting the queue_started_at field in the database.
	FieldQueueStartedAt = "queue_started_at"
	// FieldQueueEndsAt holds the string denoting the queue_ends_at field in the database.
	FieldQueueEndsAt = "queue_ends_at"
	// FieldConstructionSpeed holds the string denoting the construction_speed field in the database.
	FieldConstructionSpeed = "construction_speed"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeConstructions holds the string denoting the constructions edge name in mutations.
	EdgeConstructions = "constructions"
	// EdgeQueue holds the string denoting the queue edge name in mutations.
	EdgeQueue = "queue"

	// Table holds the table name of the city in the database.
	Table = "cities"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "cities"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_cities"
	// ConstructionsTable is the table the holds the constructions relation/edge.
	ConstructionsTable = "constructions"
	// ConstructionsInverseTable is the table name for the Construction entity.
	// It exists in this package in order to avoid circular dependency with the "construction" package.
	ConstructionsInverseTable = "constructions"
	// ConstructionsColumn is the table column denoting the constructions relation/edge.
	ConstructionsColumn = "city_constructions"
	// QueueTable is the table the holds the queue relation/edge.
	QueueTable = "queue_items"
	// QueueInverseTable is the table name for the QueueItem entity.
	// It exists in this package in order to avoid circular dependency with the "queueitem" package.
	QueueInverseTable = "queue_items"
	// QueueColumn is the table column denoting the queue relation/edge.
	QueueColumn = "city_queue"
)

// Columns holds all SQL columns for city fields.
var Columns = []string{
	FieldID,
	FieldX,
	FieldY,
	FieldName,
	FieldPoints,
	FieldWoodProduction,
	FieldStoneProduction,
	FieldIronProduction,
	FieldFoodProduction,
	FieldWoodStored,
	FieldStoneStored,
	FieldIronStored,
	FieldFoodStored,
	FieldWoodLimit,
	FieldStoneLimit,
	FieldIronLimit,
	FieldFoodLimit,
	FieldQueueStartedAt,
	FieldQueueEndsAt,
	FieldConstructionSpeed,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the City type.
var ForeignKeys = []string{
	"user_cities",
}

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
	// DefaultX holds the default value on creation for the x field.
	DefaultX int
	// DefaultY holds the default value on creation for the y field.
	DefaultY int
	// DefaultName holds the default value on creation for the name field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultPoints holds the default value on creation for the points field.
	DefaultPoints int
	// DefaultWoodProduction holds the default value on creation for the wood_production field.
	DefaultWoodProduction float64
	// DefaultStoneProduction holds the default value on creation for the stone_production field.
	DefaultStoneProduction float64
	// DefaultIronProduction holds the default value on creation for the iron_production field.
	DefaultIronProduction float64
	// DefaultFoodProduction holds the default value on creation for the food_production field.
	DefaultFoodProduction float64
	// DefaultWoodStored holds the default value on creation for the wood_stored field.
	DefaultWoodStored float64
	// DefaultStoneStored holds the default value on creation for the stone_stored field.
	DefaultStoneStored float64
	// DefaultIronStored holds the default value on creation for the iron_stored field.
	DefaultIronStored float64
	// DefaultFoodStored holds the default value on creation for the food_stored field.
	DefaultFoodStored float64
	// DefaultWoodLimit holds the default value on creation for the wood_limit field.
	DefaultWoodLimit float64
	// DefaultStoneLimit holds the default value on creation for the stone_limit field.
	DefaultStoneLimit float64
	// DefaultIronLimit holds the default value on creation for the iron_limit field.
	DefaultIronLimit float64
	// DefaultFoodLimit holds the default value on creation for the food_limit field.
	DefaultFoodLimit float64
	// DefaultQueueStartedAt holds the default value on creation for the queue_started_at field.
	DefaultQueueStartedAt func() time.Time
	// DefaultQueueEndsAt holds the default value on creation for the queue_ends_at field.
	DefaultQueueEndsAt func() time.Time
	// DefaultConstructionSpeed holds the default value on creation for the construction_speed field.
	DefaultConstructionSpeed int
	// DefaultLastUpdated holds the default value on creation for the last_updated field.
	DefaultLastUpdated func() time.Time
)
