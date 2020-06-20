// Code generated by entc, DO NOT EDIT.

package queueitem

const (
	// Label holds the string label denoting the queueitem type in the database.
	Label = "queue_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldCompletion holds the string denoting the completion field in the database.
	FieldCompletion = "completion"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeCity holds the string denoting the city edge name in mutations.
	EdgeCity = "city"
	// EdgeConstruction holds the string denoting the construction edge name in mutations.
	EdgeConstruction = "construction"

	// Table holds the table name of the queueitem in the database.
	Table = "queue_items"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "queue_items"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_queue"
	// CityTable is the table the holds the city relation/edge.
	CityTable = "queue_items"
	// CityInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CityInverseTable = "cities"
	// CityColumn is the table column denoting the city relation/edge.
	CityColumn = "city_queue"
	// ConstructionTable is the table the holds the construction relation/edge.
	ConstructionTable = "queue_items"
	// ConstructionInverseTable is the table name for the Construction entity.
	// It exists in this package in order to avoid circular dependency with the "construction" package.
	ConstructionInverseTable = "constructions"
	// ConstructionColumn is the table column denoting the construction relation/edge.
	ConstructionColumn = "construction_queue"
)

// Columns holds all SQL columns for queueitem fields.
var Columns = []string{
	FieldID,
	FieldStartAt,
	FieldDuration,
	FieldCompletion,
	FieldAction,
	FieldOrder,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the QueueItem type.
var ForeignKeys = []string{
	"city_queue",
	"construction_queue",
	"user_queue",
}
