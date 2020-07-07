// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/queueitem"
	"github.com/joaopedrosgs/loucore/ent/user"
)

// QueueItem is the model entity for the QueueItem schema.
type QueueItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt time.Time `json:"start_at,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// Completion holds the value of the "completion" field.
	Completion time.Time `json:"completion,omitempty"`
	// Action holds the value of the "action" field.
	Action int `json:"action,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QueueItemQuery when eager-loading is set.
	Edges              QueueItemEdges `json:"edges"`
	city_queue         *int
	construction_queue *int
	user_queue         *int
}

// QueueItemEdges holds the relations/edges for other nodes in the graph.
type QueueItemEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User
	// City holds the value of the city edge.
	City *City
	// Construction holds the value of the construction edge.
	Construction *Construction
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QueueItemEdges) OwnerOrErr() (*User, error) {
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

// CityOrErr returns the City value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QueueItemEdges) CityOrErr() (*City, error) {
	if e.loadedTypes[1] {
		if e.City == nil {
			// The edge city was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: city.Label}
		}
		return e.City, nil
	}
	return nil, &NotLoadedError{edge: "city"}
}

// ConstructionOrErr returns the Construction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QueueItemEdges) ConstructionOrErr() (*Construction, error) {
	if e.loadedTypes[2] {
		if e.Construction == nil {
			// The edge construction was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: construction.Label}
		}
		return e.Construction, nil
	}
	return nil, &NotLoadedError{edge: "construction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*QueueItem) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // start_at
		&sql.NullInt64{}, // duration
		&sql.NullTime{},  // completion
		&sql.NullInt64{}, // action
		&sql.NullInt64{}, // order
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*QueueItem) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // city_queue
		&sql.NullInt64{}, // construction_queue
		&sql.NullInt64{}, // user_queue
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the QueueItem fields.
func (qi *QueueItem) assignValues(values ...interface{}) error {
	if m, n := len(values), len(queueitem.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	qi.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field start_at", values[0])
	} else if value.Valid {
		qi.StartAt = value.Time
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field duration", values[1])
	} else if value.Valid {
		qi.Duration = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field completion", values[2])
	} else if value.Valid {
		qi.Completion = value.Time
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field action", values[3])
	} else if value.Valid {
		qi.Action = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field order", values[4])
	} else if value.Valid {
		qi.Order = int(value.Int64)
	}
	values = values[5:]
	if len(values) == len(queueitem.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field city_queue", value)
		} else if value.Valid {
			qi.city_queue = new(int)
			*qi.city_queue = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field construction_queue", value)
		} else if value.Valid {
			qi.construction_queue = new(int)
			*qi.construction_queue = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_queue", value)
		} else if value.Valid {
			qi.user_queue = new(int)
			*qi.user_queue = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the QueueItem.
func (qi *QueueItem) QueryOwner() *UserQuery {
	return (&QueueItemClient{config: qi.config}).QueryOwner(qi)
}

// QueryCity queries the city edge of the QueueItem.
func (qi *QueueItem) QueryCity() *CityQuery {
	return (&QueueItemClient{config: qi.config}).QueryCity(qi)
}

// QueryConstruction queries the construction edge of the QueueItem.
func (qi *QueueItem) QueryConstruction() *ConstructionQuery {
	return (&QueueItemClient{config: qi.config}).QueryConstruction(qi)
}

// Update returns a builder for updating this QueueItem.
// Note that, you need to call QueueItem.Unwrap() before calling this method, if this QueueItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (qi *QueueItem) Update() *QueueItemUpdateOne {
	return (&QueueItemClient{config: qi.config}).UpdateOne(qi)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (qi *QueueItem) Unwrap() *QueueItem {
	tx, ok := qi.config.driver.(*txDriver)
	if !ok {
		panic("ent: QueueItem is not a transactional entity")
	}
	qi.config.driver = tx.drv
	return qi
}

// String implements the fmt.Stringer.
func (qi *QueueItem) String() string {
	var builder strings.Builder
	builder.WriteString("QueueItem(")
	builder.WriteString(fmt.Sprintf("id=%v", qi.ID))
	builder.WriteString(", start_at=")
	builder.WriteString(qi.StartAt.Format(time.ANSIC))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", qi.Duration))
	builder.WriteString(", completion=")
	builder.WriteString(qi.Completion.Format(time.ANSIC))
	builder.WriteString(", action=")
	builder.WriteString(fmt.Sprintf("%v", qi.Action))
	builder.WriteString(", order=")
	builder.WriteString(fmt.Sprintf("%v", qi.Order))
	builder.WriteByte(')')
	return builder.String()
}

// QueueItems is a parsable slice of QueueItem.
type QueueItems []*QueueItem

func (qi QueueItems) config(cfg config) {
	for _i := range qi {
		qi[_i].config = cfg
	}
}