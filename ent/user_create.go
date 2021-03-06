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

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetPasswordHash sets the password_hash field.
func (uc *UserCreate) SetPasswordHash(s string) *UserCreate {
	uc.mutation.SetPasswordHash(s)
	return uc
}

// SetGold sets the gold field.
func (uc *UserCreate) SetGold(i int) *UserCreate {
	uc.mutation.SetGold(i)
	return uc
}

// SetNillableGold sets the gold field if the given value is not nil.
func (uc *UserCreate) SetNillableGold(i *int) *UserCreate {
	if i != nil {
		uc.SetGold(*i)
	}
	return uc
}

// SetDiamonds sets the diamonds field.
func (uc *UserCreate) SetDiamonds(i int) *UserCreate {
	uc.mutation.SetDiamonds(i)
	return uc
}

// SetNillableDiamonds sets the diamonds field if the given value is not nil.
func (uc *UserCreate) SetNillableDiamonds(i *int) *UserCreate {
	if i != nil {
		uc.SetDiamonds(*i)
	}
	return uc
}

// SetDarkwood sets the darkwood field.
func (uc *UserCreate) SetDarkwood(i int) *UserCreate {
	uc.mutation.SetDarkwood(i)
	return uc
}

// SetNillableDarkwood sets the darkwood field if the given value is not nil.
func (uc *UserCreate) SetNillableDarkwood(i *int) *UserCreate {
	if i != nil {
		uc.SetDarkwood(*i)
	}
	return uc
}

// SetRunestone sets the runestone field.
func (uc *UserCreate) SetRunestone(i int) *UserCreate {
	uc.mutation.SetRunestone(i)
	return uc
}

// SetNillableRunestone sets the runestone field if the given value is not nil.
func (uc *UserCreate) SetNillableRunestone(i *int) *UserCreate {
	if i != nil {
		uc.SetRunestone(*i)
	}
	return uc
}

// SetVeritium sets the veritium field.
func (uc *UserCreate) SetVeritium(i int) *UserCreate {
	uc.mutation.SetVeritium(i)
	return uc
}

// SetNillableVeritium sets the veritium field if the given value is not nil.
func (uc *UserCreate) SetNillableVeritium(i *int) *UserCreate {
	if i != nil {
		uc.SetVeritium(*i)
	}
	return uc
}

// SetTrueseed sets the trueseed field.
func (uc *UserCreate) SetTrueseed(i int) *UserCreate {
	uc.mutation.SetTrueseed(i)
	return uc
}

// SetNillableTrueseed sets the trueseed field if the given value is not nil.
func (uc *UserCreate) SetNillableTrueseed(i *int) *UserCreate {
	if i != nil {
		uc.SetTrueseed(*i)
	}
	return uc
}

// SetRank sets the rank field.
func (uc *UserCreate) SetRank(i int) *UserCreate {
	uc.mutation.SetRank(i)
	return uc
}

// SetNillableRank sets the rank field if the given value is not nil.
func (uc *UserCreate) SetNillableRank(i *int) *UserCreate {
	if i != nil {
		uc.SetRank(*i)
	}
	return uc
}

// SetAllianceRank sets the alliance_rank field.
func (uc *UserCreate) SetAllianceRank(i int) *UserCreate {
	uc.mutation.SetAllianceRank(i)
	return uc
}

// SetNillableAllianceRank sets the alliance_rank field if the given value is not nil.
func (uc *UserCreate) SetNillableAllianceRank(i *int) *UserCreate {
	if i != nil {
		uc.SetAllianceRank(*i)
	}
	return uc
}

// SetLastUpdated sets the last_updated field.
func (uc *UserCreate) SetLastUpdated(t time.Time) *UserCreate {
	uc.mutation.SetLastUpdated(t)
	return uc
}

// SetNillableLastUpdated sets the last_updated field if the given value is not nil.
func (uc *UserCreate) SetNillableLastUpdated(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastUpdated(*t)
	}
	return uc
}

// AddCityIDs adds the cities edge to City by ids.
func (uc *UserCreate) AddCityIDs(ids ...int) *UserCreate {
	uc.mutation.AddCityIDs(ids...)
	return uc
}

// AddCities adds the cities edges to City.
func (uc *UserCreate) AddCities(c ...*City) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCityIDs(ids...)
}

// AddQueueIDs adds the queue edge to QueueItem by ids.
func (uc *UserCreate) AddQueueIDs(ids ...int) *UserCreate {
	uc.mutation.AddQueueIDs(ids...)
	return uc
}

// AddQueue adds the queue edges to QueueItem.
func (uc *UserCreate) AddQueue(q ...*QueueItem) *UserCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return uc.AddQueueIDs(ids...)
}

// AddConstructionIDs adds the constructions edge to Construction by ids.
func (uc *UserCreate) AddConstructionIDs(ids ...int) *UserCreate {
	uc.mutation.AddConstructionIDs(ids...)
	return uc
}

// AddConstructions adds the constructions edges to Construction.
func (uc *UserCreate) AddConstructions(c ...*Construction) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddConstructionIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Gold(); !ok {
		v := user.DefaultGold
		uc.mutation.SetGold(v)
	}
	if _, ok := uc.mutation.Diamonds(); !ok {
		v := user.DefaultDiamonds
		uc.mutation.SetDiamonds(v)
	}
	if _, ok := uc.mutation.Darkwood(); !ok {
		v := user.DefaultDarkwood
		uc.mutation.SetDarkwood(v)
	}
	if _, ok := uc.mutation.Runestone(); !ok {
		v := user.DefaultRunestone
		uc.mutation.SetRunestone(v)
	}
	if _, ok := uc.mutation.Veritium(); !ok {
		v := user.DefaultVeritium
		uc.mutation.SetVeritium(v)
	}
	if _, ok := uc.mutation.Trueseed(); !ok {
		v := user.DefaultTrueseed
		uc.mutation.SetTrueseed(v)
	}
	if _, ok := uc.mutation.Rank(); !ok {
		v := user.DefaultRank
		uc.mutation.SetRank(v)
	}
	if _, ok := uc.mutation.AllianceRank(); !ok {
		v := user.DefaultAllianceRank
		uc.mutation.SetAllianceRank(v)
	}
	if _, ok := uc.mutation.LastUpdated(); !ok {
		v := user.DefaultLastUpdated()
		uc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := uc.mutation.PasswordHash(); !ok {
		return &ValidationError{Name: "password_hash", err: errors.New("ent: missing required field \"password_hash\"")}
	}
	if v, ok := uc.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf("ent: validator failed for field \"password_hash\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Gold(); !ok {
		return &ValidationError{Name: "gold", err: errors.New("ent: missing required field \"gold\"")}
	}
	if _, ok := uc.mutation.Diamonds(); !ok {
		return &ValidationError{Name: "diamonds", err: errors.New("ent: missing required field \"diamonds\"")}
	}
	if _, ok := uc.mutation.Darkwood(); !ok {
		return &ValidationError{Name: "darkwood", err: errors.New("ent: missing required field \"darkwood\"")}
	}
	if _, ok := uc.mutation.Runestone(); !ok {
		return &ValidationError{Name: "runestone", err: errors.New("ent: missing required field \"runestone\"")}
	}
	if _, ok := uc.mutation.Veritium(); !ok {
		return &ValidationError{Name: "veritium", err: errors.New("ent: missing required field \"veritium\"")}
	}
	if _, ok := uc.mutation.Trueseed(); !ok {
		return &ValidationError{Name: "trueseed", err: errors.New("ent: missing required field \"trueseed\"")}
	}
	if _, ok := uc.mutation.Rank(); !ok {
		return &ValidationError{Name: "rank", err: errors.New("ent: missing required field \"rank\"")}
	}
	if _, ok := uc.mutation.AllianceRank(); !ok {
		return &ValidationError{Name: "alliance_rank", err: errors.New("ent: missing required field \"alliance_rank\"")}
	}
	if _, ok := uc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "last_updated", err: errors.New("ent: missing required field \"last_updated\"")}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		_node.Name = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := uc.mutation.PasswordHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPasswordHash,
		})
		_node.PasswordHash = value
	}
	if value, ok := uc.mutation.Gold(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldGold,
		})
		_node.Gold = value
	}
	if value, ok := uc.mutation.Diamonds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldDiamonds,
		})
		_node.Diamonds = value
	}
	if value, ok := uc.mutation.Darkwood(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldDarkwood,
		})
		_node.Darkwood = value
	}
	if value, ok := uc.mutation.Runestone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldRunestone,
		})
		_node.Runestone = value
	}
	if value, ok := uc.mutation.Veritium(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldVeritium,
		})
		_node.Veritium = value
	}
	if value, ok := uc.mutation.Trueseed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldTrueseed,
		})
		_node.Trueseed = value
	}
	if value, ok := uc.mutation.Rank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldRank,
		})
		_node.Rank = value
	}
	if value, ok := uc.mutation.AllianceRank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAllianceRank,
		})
		_node.AllianceRank = value
	}
	if value, ok := uc.mutation.LastUpdated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldLastUpdated,
		})
		_node.LastUpdated = value
	}
	if nodes := uc.mutation.CitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CitiesTable,
			Columns: []string{user.CitiesColumn},
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
	if nodes := uc.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.QueueTable,
			Columns: []string{user.QueueColumn},
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
	if nodes := uc.mutation.ConstructionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ConstructionsTable,
			Columns: []string{user.ConstructionsColumn},
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

// UserCreateBulk is the builder for creating a bulk of User entities.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
