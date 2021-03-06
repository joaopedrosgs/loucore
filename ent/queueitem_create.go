// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
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

// SetDuration sets the duration field.
func (qic *QueueItemCreate) SetDuration(i int) *QueueItemCreate {
	qic.mutation.SetDuration(i)
	return qic
}

// SetAction sets the action field.
func (qic *QueueItemCreate) SetAction(i int) *QueueItemCreate {
	qic.mutation.SetAction(i)
	return qic
}

// SetPosition sets the position field.
func (qic *QueueItemCreate) SetPosition(i int) *QueueItemCreate {
	qic.mutation.SetPosition(i)
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

// Mutation returns the QueueItemMutation object of the builder.
func (qic *QueueItemCreate) Mutation() *QueueItemMutation {
	return qic.mutation
}

// Save creates the QueueItem in the database.
func (qic *QueueItemCreate) Save(ctx context.Context) (*QueueItem, error) {
	var (
		err  error
		node *QueueItem
	)
	if len(qic.hooks) == 0 {
		if err = qic.check(); err != nil {
			return nil, err
		}
		node, err = qic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qic.check(); err != nil {
				return nil, err
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

// check runs all checks and user-defined validators on the builder.
func (qic *QueueItemCreate) check() error {
	if _, ok := qic.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New("ent: missing required field \"duration\"")}
	}
	if _, ok := qic.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New("ent: missing required field \"action\"")}
	}
	if _, ok := qic.mutation.Position(); !ok {
		return &ValidationError{Name: "position", err: errors.New("ent: missing required field \"position\"")}
	}
	return nil
}

func (qic *QueueItemCreate) sqlSave(ctx context.Context) (*QueueItem, error) {
	_node, _spec := qic.createSpec()
	if err := sqlgraph.CreateNode(ctx, qic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (qic *QueueItemCreate) createSpec() (*QueueItem, *sqlgraph.CreateSpec) {
	var (
		_node = &QueueItem{config: qic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: queueitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queueitem.FieldID,
			},
		}
	)
	if value, ok := qic.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := qic.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldAction,
		})
		_node.Action = value
	}
	if value, ok := qic.mutation.Position(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldPosition,
		})
		_node.Position = value
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
	return _node, _spec
}

// QueueItemCreateBulk is the builder for creating a bulk of QueueItem entities.
type QueueItemCreateBulk struct {
	config
	builders []*QueueItemCreate
}

// Save creates the QueueItem entities in the database.
func (qicb *QueueItemCreateBulk) Save(ctx context.Context) ([]*QueueItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qicb.builders))
	nodes := make([]*QueueItem, len(qicb.builders))
	mutators := make([]Mutator, len(qicb.builders))
	for i := range qicb.builders {
		func(i int, root context.Context) {
			builder := qicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QueueItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, qicb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qicb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, qicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (qicb *QueueItemCreateBulk) SaveX(ctx context.Context) []*QueueItem {
	v, err := qicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
