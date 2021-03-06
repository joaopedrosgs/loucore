// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// CitiesColumns holds the columns for the "cities" table.
	CitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "x", Type: field.TypeInt},
		{Name: "y", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "New city"},
		{Name: "points", Type: field.TypeInt, Default: 3},
		{Name: "wood_production", Type: field.TypeFloat64, Default: 5},
		{Name: "stone_production", Type: field.TypeFloat64},
		{Name: "iron_production", Type: field.TypeFloat64},
		{Name: "food_production", Type: field.TypeFloat64},
		{Name: "wood_stored", Type: field.TypeFloat64, Default: 300},
		{Name: "stone_stored", Type: field.TypeFloat64},
		{Name: "iron_stored", Type: field.TypeFloat64},
		{Name: "food_stored", Type: field.TypeFloat64},
		{Name: "wood_limit", Type: field.TypeFloat64, Default: 300},
		{Name: "stone_limit", Type: field.TypeFloat64},
		{Name: "iron_limit", Type: field.TypeFloat64},
		{Name: "food_limit", Type: field.TypeFloat64},
		{Name: "queue_started_at", Type: field.TypeTime},
		{Name: "queue_ends_at", Type: field.TypeTime},
		{Name: "construction_speed", Type: field.TypeInt, Default: 1},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "user_cities", Type: field.TypeInt, Nullable: true},
	}
	// CitiesTable holds the schema information for the "cities" table.
	CitiesTable = &schema.Table{
		Name:       "cities",
		Columns:    CitiesColumns,
		PrimaryKey: []*schema.Column{CitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "cities_users_cities",
				Columns: []*schema.Column{CitiesColumns[21]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "city_x_y",
				Unique:  true,
				Columns: []*schema.Column{CitiesColumns[1], CitiesColumns[2]},
			},
		},
	}
	// ConstructionsColumns holds the columns for the "constructions" table.
	ConstructionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "x", Type: field.TypeInt},
		{Name: "y", Type: field.TypeInt},
		{Name: "raw_production", Type: field.TypeFloat64},
		{Name: "production", Type: field.TypeFloat64},
		{Name: "type", Type: field.TypeInt},
		{Name: "level", Type: field.TypeInt},
		{Name: "modifier", Type: field.TypeFloat64, Default: 1},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "need_refresh", Type: field.TypeBool, Default: true},
		{Name: "city_constructions", Type: field.TypeInt, Nullable: true},
		{Name: "user_constructions", Type: field.TypeInt, Nullable: true},
	}
	// ConstructionsTable holds the schema information for the "constructions" table.
	ConstructionsTable = &schema.Table{
		Name:       "constructions",
		Columns:    ConstructionsColumns,
		PrimaryKey: []*schema.Column{ConstructionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "constructions_cities_constructions",
				Columns: []*schema.Column{ConstructionsColumns[10]},

				RefColumns: []*schema.Column{CitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "constructions_users_constructions",
				Columns: []*schema.Column{ConstructionsColumns[11]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "construction_x_y_city_constructions",
				Unique:  true,
				Columns: []*schema.Column{ConstructionsColumns[1], ConstructionsColumns[2], ConstructionsColumns[10]},
			},
		},
	}
	// QueueItemsColumns holds the columns for the "queue_items" table.
	QueueItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "duration", Type: field.TypeInt},
		{Name: "action", Type: field.TypeInt},
		{Name: "position", Type: field.TypeInt},
		{Name: "city_queue", Type: field.TypeInt, Nullable: true},
		{Name: "construction_queue", Type: field.TypeInt, Nullable: true},
		{Name: "user_queue", Type: field.TypeInt, Nullable: true},
	}
	// QueueItemsTable holds the schema information for the "queue_items" table.
	QueueItemsTable = &schema.Table{
		Name:       "queue_items",
		Columns:    QueueItemsColumns,
		PrimaryKey: []*schema.Column{QueueItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "queue_items_cities_queue",
				Columns: []*schema.Column{QueueItemsColumns[4]},

				RefColumns: []*schema.Column{CitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "queue_items_constructions_queue",
				Columns: []*schema.Column{QueueItemsColumns[5]},

				RefColumns: []*schema.Column{ConstructionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "queue_items_users_queue",
				Columns: []*schema.Column{QueueItemsColumns[6]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "gold", Type: field.TypeInt},
		{Name: "diamonds", Type: field.TypeInt},
		{Name: "darkwood", Type: field.TypeInt},
		{Name: "runestone", Type: field.TypeInt},
		{Name: "veritium", Type: field.TypeInt},
		{Name: "trueseed", Type: field.TypeInt},
		{Name: "rank", Type: field.TypeInt},
		{Name: "alliance_rank", Type: field.TypeInt},
		{Name: "last_updated", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ConstructionAffectsColumns holds the columns for the "construction_affects" table.
	ConstructionAffectsColumns = []*schema.Column{
		{Name: "construction_id", Type: field.TypeInt},
		{Name: "affected_by_id", Type: field.TypeInt},
	}
	// ConstructionAffectsTable holds the schema information for the "construction_affects" table.
	ConstructionAffectsTable = &schema.Table{
		Name:       "construction_affects",
		Columns:    ConstructionAffectsColumns,
		PrimaryKey: []*schema.Column{ConstructionAffectsColumns[0], ConstructionAffectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "construction_affects_construction_id",
				Columns: []*schema.Column{ConstructionAffectsColumns[0]},

				RefColumns: []*schema.Column{ConstructionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "construction_affects_affected_by_id",
				Columns: []*schema.Column{ConstructionAffectsColumns[1]},

				RefColumns: []*schema.Column{ConstructionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CitiesTable,
		ConstructionsTable,
		QueueItemsTable,
		UsersTable,
		ConstructionAffectsTable,
	}
)

func init() {
	CitiesTable.ForeignKeys[0].RefTable = UsersTable
	ConstructionsTable.ForeignKeys[0].RefTable = CitiesTable
	ConstructionsTable.ForeignKeys[1].RefTable = UsersTable
	QueueItemsTable.ForeignKeys[0].RefTable = CitiesTable
	QueueItemsTable.ForeignKeys[1].RefTable = ConstructionsTable
	QueueItemsTable.ForeignKeys[2].RefTable = UsersTable
	ConstructionAffectsTable.ForeignKeys[0].RefTable = ConstructionsTable
	ConstructionAffectsTable.ForeignKeys[1].RefTable = ConstructionsTable
}
