// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/queueitem"
	"github.com/joaopedrosgs/loucore/ent/user"
)

// QueueItemCreate is the builder for creating a QueueItem entity.
type QueueItemCreate struct {
	config
	mutation *QueueItemMutation
	hooks    []Hook
}

// SetStartAt sets the start_at field.
func (qic *QueueItemCreate) SetStartAt(t time.Time) *QueueItemCreate {
	qic.mutation.SetStartAt(t)
	return qic
}

// SetDuration sets the duration field.
func (qic *QueueItemCreate) SetDuration(i int) *QueueItemCreate {
	qic.mutation.SetDuration(i)
	return qic
}

// SetCompletion sets the completion field.
func (qic *QueueItemCreate) SetCompletion(t time.Time) *QueueItemCreate {
	qic.mutation.SetCompletion(t)
	return qic
}

// SetAction sets the action field.
func (qic *QueueItemCreate) SetAction(i int) *QueueItemCreate {
	qic.mutation.SetAction(i)
	return qic
}

// SetOrder sets the order field.
func (qic *QueueItemCreate) SetOrder(i int) *QueueItemCreate {
	qic.mutation.SetOrder(i)
	return qic
}

// SetOwnerID sets the owner edge to User by id.
func (qic *QueueItemCreate) SetOwnerID(id int) *QueueItemCreate {
	qic.mutation.SetOwnerID(id)
	return qic
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (qic *QueueItemCreate) SetNillableOwnerID(id *int) *QueueItemCreate {
	if id != nil {
		qic = qic.SetOwnerID(*id)
	}
	return qic
}

// SetOwner sets the owner edge to User.
func (qic *QueueItemCreate) SetOwner(u *User) *QueueItemCreate {
	return qic.SetOwnerID(u.ID)
}

// SetCityID sets the city edge to City by id.
func (qic *QueueItemCreate) SetCityID(id int) *QueueItemCreate {
	qic.mutation.SetCityID(id)
	return qic
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (qic *QueueItemCreate) SetNillableCityID(id *int) *QueueItemCreate {
	if id != nil {
		qic = qic.SetCityID(*id)
	}
	return qic
}

// SetCity sets the city edge to City.
func (qic *QueueItemCreate) SetCity(c *City) *QueueItemCreate {
	return qic.SetCityID(c.ID)
}

// SetConstructionID sets the construction edge to Construction by id.
func (qic *QueueItemCreate) SetConstructionID(id int) *QueueItemCreate {
	qic.mutation.SetConstructionID(id)
	return qic
}

// SetNillableConstructionID sets the construction edge to Construction by id if the given value is not nil.
func (qic *QueueItemCreate) SetNillableConstructionID(id *int) *QueueItemCreate {
	if id != nil {
		qic = qic.SetConstructionID(*id)
	}
	return qic
}

// SetConstruction sets the construction edge to Construction.
func (qic *QueueItemCreate) SetConstruction(c *Construction) *QueueItemCreate {
	return qic.SetConstructionID(c.ID)
}

// Save creates the QueueItem in the database.
func (qic *QueueItemCreate) Save(ctx context.Context) (*QueueItem, error) {
	if _, ok := qic.mutation.StartAt(); !ok {
		return nil, errors.New("ent: missing required field \"start_at\"")
	}
	if _, ok := qic.mutation.Duration(); !ok {
		return nil, errors.New("ent: missing required field \"duration\"")
	}
	if _, ok := qic.mutation.Completion(); !ok {
		return nil, errors.New("ent: missing required field \"completion\"")
	}
	if _, ok := qic.mutation.Action(); !ok {
		return nil, errors.New("ent: missing required field \"action\"")
	}
	if _, ok := qic.mutation.Order(); !ok {
		return nil, errors.New("ent: missing required field \"order\"")
	}
	var (
		err  error
		node *QueueItem
	)
	if len(qic.hooks) == 0 {
		node, err = qic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qic.mutation = mutation
			node, err = qic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qic.hooks) - 1; i >= 0; i-- {
			mut = qic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qic *QueueItemCreate) SaveX(ctx context.Context) *QueueItem {
	v, err := qic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (qic *QueueItemCreate) sqlSave(ctx context.Context) (*QueueItem, error) {
	var (
		qi    = &QueueItem{config: qic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: queueitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queueitem.FieldID,
			},
		}
	)
	if value, ok := qic.mutation.StartAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queueitem.FieldStartAt,
		})
		qi.StartAt = value
	}
	if value, ok := qic.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldDuration,
		})
		qi.Duration = value
	}
	if value, ok := qic.mutation.Completion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queueitem.FieldCompletion,
		})
		qi.Completion = value
	}
	if value, ok := qic.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldAction,
		})
		qi.Action = value
	}
	if value, ok := qic.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldOrder,
		})
		qi.Order = value
	}
	if nodes := qic.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queueitem.OwnerTable,
			Columns: []string{queueitem.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qic.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queueitem.CityTable,
			Columns: []string{queueitem.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qic.mutation.ConstructionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   queueitem.ConstructionTable,
			Columns: []string{queueitem.ConstructionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: construction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, qic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	qi.ID = int(id)
	return qi, nil
}
