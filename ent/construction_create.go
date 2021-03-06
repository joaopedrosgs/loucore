// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/queueitem"
	"github.com/joaopedrosgs/loucore/ent/user"
)

// ConstructionCreate is the builder for creating a Construction entity.
type ConstructionCreate struct {
	config
	mutation *ConstructionMutation
	hooks    []Hook
}

// SetX sets the x field.
func (cc *ConstructionCreate) SetX(i int) *ConstructionCreate {
	cc.mutation.SetX(i)
	return cc
}

// SetY sets the y field.
func (cc *ConstructionCreate) SetY(i int) *ConstructionCreate {
	cc.mutation.SetY(i)
	return cc
}

// SetRawProduction sets the raw_production field.
func (cc *ConstructionCreate) SetRawProduction(f float64) *ConstructionCreate {
	cc.mutation.SetRawProduction(f)
	return cc
}

// SetNillableRawProduction sets the raw_production field if the given value is not nil.
func (cc *ConstructionCreate) SetNillableRawProduction(f *float64) *ConstructionCreate {
	if f != nil {
		cc.SetRawProduction(*f)
	}
	return cc
}

// SetProduction sets the production field.
func (cc *ConstructionCreate) SetProduction(f float64) *ConstructionCreate {
	cc.mutation.SetProduction(f)
	return cc
}

// SetNillableProduction sets the production field if the given value is not nil.
func (cc *ConstructionCreate) SetNillableProduction(f *float64) *ConstructionCreate {
	if f != nil {
		cc.SetProduction(*f)
	}
	return cc
}

// SetType sets the type field.
func (cc *ConstructionCreate) SetType(i int) *ConstructionCreate {
	cc.mutation.SetType(i)
	return cc
}

// SetNillableType sets the type field if the given value is not nil.
func (cc *ConstructionCreate) SetNillableType(i *int) *ConstructionCreate {
	if i != nil {
		cc.SetType(*i)
	}
	return cc
}

// SetLevel sets the level field.
func (cc *ConstructionCreate) SetLevel(i int) *ConstructionCreate {
	cc.mutation.SetLevel(i)
	return cc
}

// SetNillableLevel sets the level field if the given value is not nil.
func (cc *ConstructionCreate) SetNillableLevel(i *int) *ConstructionCreate {
	if i != nil {
		cc.SetLevel(*i)
	}
	return cc
}

// SetModifier sets the modifier field.
func (cc *ConstructionCreate) SetModifier(f float64) *ConstructionCreate {
	cc.mutation.SetModifier(f)
	return cc
}

// SetNillableModifier sets the modifier field if the given value is not nil.
func (cc *ConstructionCreate) SetNillableModifier(f *float64) *ConstructionCreate {
	if f != nil {
		cc.SetModifier(*f)
	}
	return cc
}

// SetLastUpdated sets the last_updated field.
func (cc *ConstructionCreate) SetLastUpdated(t time.Time) *ConstructionCreate {
	cc.mutation.SetLastUpdated(t)
	return cc
}

// SetNillableLastUpdated sets the last_updated field if the given value is not nil.
func (cc *ConstructionCreate) SetNillableLastUpdated(t *time.Time) *ConstructionCreate {
	if t != nil {
		cc.SetLastUpdated(*t)
	}
	return cc
}

// SetNeedRefresh sets the need_refresh field.
func (cc *ConstructionCreate) SetNeedRefresh(b bool) *ConstructionCreate {
	cc.mutation.SetNeedRefresh(b)
	return cc
}

// SetNillableNeedRefresh sets the need_refresh field if the given value is not nil.
func (cc *ConstructionCreate) SetNillableNeedRefresh(b *bool) *ConstructionCreate {
	if b != nil {
		cc.SetNeedRefresh(*b)
	}
	return cc
}

// SetCityID sets the city edge to City by id.
func (cc *ConstructionCreate) SetCityID(id int) *ConstructionCreate {
	cc.mutation.SetCityID(id)
	return cc
}

// SetCity sets the city edge to City.
func (cc *ConstructionCreate) SetCity(c *City) *ConstructionCreate {
	return cc.SetCityID(c.ID)
}

// SetOwnerID sets the owner edge to User by id.
func (cc *ConstructionCreate) SetOwnerID(id int) *ConstructionCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cc *ConstructionCreate) SetNillableOwnerID(id *int) *ConstructionCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the owner edge to User.
func (cc *ConstructionCreate) SetOwner(u *User) *ConstructionCreate {
	return cc.SetOwnerID(u.ID)
}

// AddQueueIDs adds the queue edge to QueueItem by ids.
func (cc *ConstructionCreate) AddQueueIDs(ids ...int) *ConstructionCreate {
	cc.mutation.AddQueueIDs(ids...)
	return cc
}

// AddQueue adds the queue edges to QueueItem.
func (cc *ConstructionCreate) AddQueue(q ...*QueueItem) *ConstructionCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return cc.AddQueueIDs(ids...)
}

// AddAffectIDs adds the affects edge to Construction by ids.
func (cc *ConstructionCreate) AddAffectIDs(ids ...int) *ConstructionCreate {
	cc.mutation.AddAffectIDs(ids...)
	return cc
}

// AddAffects adds the affects edges to Construction.
func (cc *ConstructionCreate) AddAffects(c ...*Construction) *ConstructionCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddAffectIDs(ids...)
}

// AddAffectedByIDs adds the affected_by edge to Construction by ids.
func (cc *ConstructionCreate) AddAffectedByIDs(ids ...int) *ConstructionCreate {
	cc.mutation.AddAffectedByIDs(ids...)
	return cc
}

// AddAffectedBy adds the affected_by edges to Construction.
func (cc *ConstructionCreate) AddAffectedBy(c ...*Construction) *ConstructionCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddAffectedByIDs(ids...)
}

// Mutation returns the ConstructionMutation object of the builder.
func (cc *ConstructionCreate) Mutation() *ConstructionMutation {
	return cc.mutation
}

// Save creates the Construction in the database.
func (cc *ConstructionCreate) Save(ctx context.Context) (*Construction, error) {
	var (
		err  error
		node *Construction
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConstructionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConstructionCreate) SaveX(ctx context.Context) *Construction {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *ConstructionCreate) defaults() {
	if _, ok := cc.mutation.RawProduction(); !ok {
		v := construction.DefaultRawProduction
		cc.mutation.SetRawProduction(v)
	}
	if _, ok := cc.mutation.Production(); !ok {
		v := construction.DefaultProduction
		cc.mutation.SetProduction(v)
	}
	if _, ok := cc.mutation.GetType(); !ok {
		v := construction.DefaultType
		cc.mutation.SetType(v)
	}
	if _, ok := cc.mutation.Level(); !ok {
		v := construction.DefaultLevel
		cc.mutation.SetLevel(v)
	}
	if _, ok := cc.mutation.Modifier(); !ok {
		v := construction.DefaultModifier
		cc.mutation.SetModifier(v)
	}
	if _, ok := cc.mutation.LastUpdated(); !ok {
		v := construction.DefaultLastUpdated()
		cc.mutation.SetLastUpdated(v)
	}
	if _, ok := cc.mutation.NeedRefresh(); !ok {
		v := construction.DefaultNeedRefresh
		cc.mutation.SetNeedRefresh(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConstructionCreate) check() error {
	if _, ok := cc.mutation.X(); !ok {
		return &ValidationError{Name: "x", err: errors.New("ent: missing required field \"x\"")}
	}
	if _, ok := cc.mutation.Y(); !ok {
		return &ValidationError{Name: "y", err: errors.New("ent: missing required field \"y\"")}
	}
	if _, ok := cc.mutation.RawProduction(); !ok {
		return &ValidationError{Name: "raw_production", err: errors.New("ent: missing required field \"raw_production\"")}
	}
	if _, ok := cc.mutation.Production(); !ok {
		return &ValidationError{Name: "production", err: errors.New("ent: missing required field \"production\"")}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if _, ok := cc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New("ent: missing required field \"level\"")}
	}
	if _, ok := cc.mutation.Modifier(); !ok {
		return &ValidationError{Name: "modifier", err: errors.New("ent: missing required field \"modifier\"")}
	}
	if _, ok := cc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "last_updated", err: errors.New("ent: missing required field \"last_updated\"")}
	}
	if _, ok := cc.mutation.NeedRefresh(); !ok {
		return &ValidationError{Name: "need_refresh", err: errors.New("ent: missing required field \"need_refresh\"")}
	}
	if _, ok := cc.mutation.CityID(); !ok {
		return &ValidationError{Name: "city", err: errors.New("ent: missing required edge \"city\"")}
	}
	return nil
}

func (cc *ConstructionCreate) sqlSave(ctx context.Context) (*Construction, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ConstructionCreate) createSpec() (*Construction, *sqlgraph.CreateSpec) {
	var (
		_node = &Construction{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: construction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: construction.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.X(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: construction.FieldX,
		})
		_node.X = value
	}
	if value, ok := cc.mutation.Y(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: construction.FieldY,
		})
		_node.Y = value
	}
	if value, ok := cc.mutation.RawProduction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: construction.FieldRawProduction,
		})
		_node.RawProduction = value
	}
	if value, ok := cc.mutation.Production(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: construction.FieldProduction,
		})
		_node.Production = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: construction.FieldType,
		})
		_node.Type = value
	}
	if value, ok := cc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: construction.FieldLevel,
		})
		_node.Level = value
	}
	if value, ok := cc.mutation.Modifier(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: construction.FieldModifier,
		})
		_node.Modifier = value
	}
	if value, ok := cc.mutation.LastUpdated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: construction.FieldLastUpdated,
		})
		_node.LastUpdated = value
	}
	if value, ok := cc.mutation.NeedRefresh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: construction.FieldNeedRefresh,
		})
		_node.NeedRefresh = value
	}
	if nodes := cc.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   construction.CityTable,
			Columns: []string{construction.CityColumn},
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
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   construction.OwnerTable,
			Columns: []string{construction.OwnerColumn},
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
	if nodes := cc.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   construction.QueueTable,
			Columns: []string{construction.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queueitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AffectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   construction.AffectsTable,
			Columns: construction.AffectsPrimaryKey,
			Bidi:    true,
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
	if nodes := cc.mutation.AffectedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   construction.AffectedByTable,
			Columns: construction.AffectedByPrimaryKey,
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

// ConstructionCreateBulk is the builder for creating a bulk of Construction entities.
type ConstructionCreateBulk struct {
	config
	builders []*ConstructionCreate
}

// Save creates the Construction entities in the database.
func (ccb *ConstructionCreateBulk) Save(ctx context.Context) ([]*Construction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Construction, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConstructionMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ccb *ConstructionCreateBulk) SaveX(ctx context.Context) []*Construction {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
