// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/user"
)

// City is the model entity for the City schema.
type City struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// X holds the value of the "x" field.
	X int `json:"x,omitempty"`
	// Y holds the value of the "y" field.
	Y int `json:"y,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Points holds the value of the "points" field.
	Points int `json:"points,omitempty"`
	// WoodProduction holds the value of the "wood_production" field.
	WoodProduction float64 `json:"wood_production,omitempty"`
	// StoneProduction holds the value of the "stone_production" field.
	StoneProduction float64 `json:"stone_production,omitempty"`
	// IronProduction holds the value of the "iron_production" field.
	IronProduction float64 `json:"iron_production,omitempty"`
	// FoodProduction holds the value of the "food_production" field.
	FoodProduction float64 `json:"food_production,omitempty"`
	// WoodStored holds the value of the "wood_stored" field.
	WoodStored float64 `json:"wood_stored,omitempty"`
	// StoneStored holds the value of the "stone_stored" field.
	StoneStored float64 `json:"stone_stored,omitempty"`
	// IronStored holds the value of the "iron_stored" field.
	IronStored float64 `json:"iron_stored,omitempty"`
	// FoodStored holds the value of the "food_stored" field.
	FoodStored float64 `json:"food_stored,omitempty"`
	// WoodLimit holds the value of the "wood_limit" field.
	WoodLimit float64 `json:"wood_limit,omitempty"`
	// StoneLimit holds the value of the "stone_limit" field.
	StoneLimit float64 `json:"stone_limit,omitempty"`
	// IronLimit holds the value of the "iron_limit" field.
	IronLimit float64 `json:"iron_limit,omitempty"`
	// FoodLimit holds the value of the "food_limit" field.
	FoodLimit float64 `json:"food_limit,omitempty"`
	// QueueEndsAt holds the value of the "queue_ends_at" field.
	QueueEndsAt time.Time `json:"queue_ends_at,omitempty"`
	// ConstructionSpeed holds the value of the "construction_speed" field.
	ConstructionSpeed int `json:"construction_speed,omitempty"`
	// LastUpdated holds the value of the "last_updated" field.
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CityQuery when eager-loading is set.
	Edges       CityEdges `json:"edges"`
	user_cities *int
}

// CityEdges holds the relations/edges for other nodes in the graph.
type CityEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// Constructions holds the value of the constructions edge.
	Constructions []*Construction
	// Queue holds the value of the queue edge.
	Queue []*QueueItem
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CityEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// ConstructionsOrErr returns the Constructions value or an error if the edge
// was not loaded in eager-loading.
func (e CityEdges) ConstructionsOrErr() ([]*Construction, error) {
	if e.loadedTypes[1] {
		return e.Constructions, nil
	}
	return nil, &NotLoadedError{edge: "constructions"}
}

// QueueOrErr returns the Queue value or an error if the edge
// was not loaded in eager-loading.
func (e CityEdges) QueueOrErr() ([]*QueueItem, error) {
	if e.loadedTypes[2] {
		return e.Queue, nil
	}
	return nil, &NotLoadedError{edge: "queue"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*City) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullInt64{},   // x
		&sql.NullInt64{},   // y
		&sql.NullString{},  // name
		&sql.NullInt64{},   // points
		&sql.NullFloat64{}, // wood_production
		&sql.NullFloat64{}, // stone_production
		&sql.NullFloat64{}, // iron_production
		&sql.NullFloat64{}, // food_production
		&sql.NullFloat64{}, // wood_stored
		&sql.NullFloat64{}, // stone_stored
		&sql.NullFloat64{}, // iron_stored
		&sql.NullFloat64{}, // food_stored
		&sql.NullFloat64{}, // wood_limit
		&sql.NullFloat64{}, // stone_limit
		&sql.NullFloat64{}, // iron_limit
		&sql.NullFloat64{}, // food_limit
		&sql.NullTime{},    // queue_ends_at
		&sql.NullInt64{},   // construction_speed
		&sql.NullTime{},    // last_updated
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*City) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_cities
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the City fields.
func (c *City) assignValues(values ...interface{}) error {
	if m, n := len(values), len(city.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field x", values[0])
	} else if value.Valid {
		c.X = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field y", values[1])
	} else if value.Valid {
		c.Y = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		c.Name = value.String
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field points", values[3])
	} else if value.Valid {
		c.Points = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field wood_production", values[4])
	} else if value.Valid {
		c.WoodProduction = value.Float64
	}
	if value, ok := values[5].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field stone_production", values[5])
	} else if value.Valid {
		c.StoneProduction = value.Float64
	}
	if value, ok := values[6].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field iron_production", values[6])
	} else if value.Valid {
		c.IronProduction = value.Float64
	}
	if value, ok := values[7].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field food_production", values[7])
	} else if value.Valid {
		c.FoodProduction = value.Float64
	}
	if value, ok := values[8].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field wood_stored", values[8])
	} else if value.Valid {
		c.WoodStored = value.Float64
	}
	if value, ok := values[9].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field stone_stored", values[9])
	} else if value.Valid {
		c.StoneStored = value.Float64
	}
	if value, ok := values[10].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field iron_stored", values[10])
	} else if value.Valid {
		c.IronStored = value.Float64
	}
	if value, ok := values[11].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field food_stored", values[11])
	} else if value.Valid {
		c.FoodStored = value.Float64
	}
	if value, ok := values[12].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field wood_limit", values[12])
	} else if value.Valid {
		c.WoodLimit = value.Float64
	}
	if value, ok := values[13].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field stone_limit", values[13])
	} else if value.Valid {
		c.StoneLimit = value.Float64
	}
	if value, ok := values[14].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field iron_limit", values[14])
	} else if value.Valid {
		c.IronLimit = value.Float64
	}
	if value, ok := values[15].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field food_limit", values[15])
	} else if value.Valid {
		c.FoodLimit = value.Float64
	}
	if value, ok := values[16].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field queue_ends_at", values[16])
	} else if value.Valid {
		c.QueueEndsAt = value.Time
	}
	if value, ok := values[17].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field construction_speed", values[17])
	} else if value.Valid {
		c.ConstructionSpeed = int(value.Int64)
	}
	if value, ok := values[18].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field last_updated", values[18])
	} else if value.Valid {
		c.LastUpdated = value.Time
	}
	values = values[19:]
	if len(values) == len(city.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_cities", value)
		} else if value.Valid {
			c.user_cities = new(int)
			*c.user_cities = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the City.
func (c *City) QueryOwner() *UserQuery {
	return (&CityClient{config: c.config}).QueryOwner(c)
}

// QueryConstructions queries the constructions edge of the City.
func (c *City) QueryConstructions() *ConstructionQuery {
	return (&CityClient{config: c.config}).QueryConstructions(c)
}

// QueryQueue queries the queue edge of the City.
func (c *City) QueryQueue() *QueueItemQuery {
	return (&CityClient{config: c.config}).QueryQueue(c)
}

// Update returns a builder for updating this City.
// Note that, you need to call City.Unwrap() before calling this method, if this City
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *City) Update() *CityUpdateOne {
	return (&CityClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *City) Unwrap() *City {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: City is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *City) String() string {
	var builder strings.Builder
	builder.WriteString("City(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", x=")
	builder.WriteString(fmt.Sprintf("%v", c.X))
	builder.WriteString(", y=")
	builder.WriteString(fmt.Sprintf("%v", c.Y))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", points=")
	builder.WriteString(fmt.Sprintf("%v", c.Points))
	builder.WriteString(", wood_production=")
	builder.WriteString(fmt.Sprintf("%v", c.WoodProduction))
	builder.WriteString(", stone_production=")
	builder.WriteString(fmt.Sprintf("%v", c.StoneProduction))
	builder.WriteString(", iron_production=")
	builder.WriteString(fmt.Sprintf("%v", c.IronProduction))
	builder.WriteString(", food_production=")
	builder.WriteString(fmt.Sprintf("%v", c.FoodProduction))
	builder.WriteString(", wood_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.WoodStored))
	builder.WriteString(", stone_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.StoneStored))
	builder.WriteString(", iron_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.IronStored))
	builder.WriteString(", food_stored=")
	builder.WriteString(fmt.Sprintf("%v", c.FoodStored))
	builder.WriteString(", wood_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.WoodLimit))
	builder.WriteString(", stone_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.StoneLimit))
	builder.WriteString(", iron_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.IronLimit))
	builder.WriteString(", food_limit=")
	builder.WriteString(fmt.Sprintf("%v", c.FoodLimit))
	builder.WriteString(", queue_ends_at=")
	builder.WriteString(c.QueueEndsAt.Format(time.ANSIC))
	builder.WriteString(", construction_speed=")
	builder.WriteString(fmt.Sprintf("%v", c.ConstructionSpeed))
	builder.WriteString(", last_updated=")
	builder.WriteString(c.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Cities is a parsable slice of City.
type Cities []*City

func (c Cities) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
