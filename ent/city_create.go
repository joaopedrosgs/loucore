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

// CityCreate is the builder for creating a City entity.
type CityCreate struct {
	config
	mutation *CityMutation
	hooks    []Hook
}

// SetX sets the x field.
func (cc *CityCreate) SetX(i int) *CityCreate {
	cc.mutation.SetX(i)
	return cc
}

// SetNillableX sets the x field if the given value is not nil.
func (cc *CityCreate) SetNillableX(i *int) *CityCreate {
	if i != nil {
		cc.SetX(*i)
	}
	return cc
}

// SetY sets the y field.
func (cc *CityCreate) SetY(i int) *CityCreate {
	cc.mutation.SetY(i)
	return cc
}

// SetNillableY sets the y field if the given value is not nil.
func (cc *CityCreate) SetNillableY(i *int) *CityCreate {
	if i != nil {
		cc.SetY(*i)
	}
	return cc
}

// SetName sets the name field.
func (cc *CityCreate) SetName(s string) *CityCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the name field if the given value is not nil.
func (cc *CityCreate) SetNillableName(s *string) *CityCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetPoints sets the points field.
func (cc *CityCreate) SetPoints(i int) *CityCreate {
	cc.mutation.SetPoints(i)
	return cc
}

// SetNillablePoints sets the points field if the given value is not nil.
func (cc *CityCreate) SetNillablePoints(i *int) *CityCreate {
	if i != nil {
		cc.SetPoints(*i)
	}
	return cc
}

// SetWoodProduction sets the wood_production field.
func (cc *CityCreate) SetWoodProduction(f float64) *CityCreate {
	cc.mutation.SetWoodProduction(f)
	return cc
}

// SetNillableWoodProduction sets the wood_production field if the given value is not nil.
func (cc *CityCreate) SetNillableWoodProduction(f *float64) *CityCreate {
	if f != nil {
		cc.SetWoodProduction(*f)
	}
	return cc
}

// SetStoneProduction sets the stone_production field.
func (cc *CityCreate) SetStoneProduction(f float64) *CityCreate {
	cc.mutation.SetStoneProduction(f)
	return cc
}

// SetNillableStoneProduction sets the stone_production field if the given value is not nil.
func (cc *CityCreate) SetNillableStoneProduction(f *float64) *CityCreate {
	if f != nil {
		cc.SetStoneProduction(*f)
	}
	return cc
}

// SetIronProduction sets the iron_production field.
func (cc *CityCreate) SetIronProduction(f float64) *CityCreate {
	cc.mutation.SetIronProduction(f)
	return cc
}

// SetNillableIronProduction sets the iron_production field if the given value is not nil.
func (cc *CityCreate) SetNillableIronProduction(f *float64) *CityCreate {
	if f != nil {
		cc.SetIronProduction(*f)
	}
	return cc
}

// SetFoodProduction sets the food_production field.
func (cc *CityCreate) SetFoodProduction(f float64) *CityCreate {
	cc.mutation.SetFoodProduction(f)
	return cc
}

// SetNillableFoodProduction sets the food_production field if the given value is not nil.
func (cc *CityCreate) SetNillableFoodProduction(f *float64) *CityCreate {
	if f != nil {
		cc.SetFoodProduction(*f)
	}
	return cc
}

// SetWoodStored sets the wood_stored field.
func (cc *CityCreate) SetWoodStored(f float64) *CityCreate {
	cc.mutation.SetWoodStored(f)
	return cc
}

// SetNillableWoodStored sets the wood_stored field if the given value is not nil.
func (cc *CityCreate) SetNillableWoodStored(f *float64) *CityCreate {
	if f != nil {
		cc.SetWoodStored(*f)
	}
	return cc
}

// SetStoneStored sets the stone_stored field.
func (cc *CityCreate) SetStoneStored(f float64) *CityCreate {
	cc.mutation.SetStoneStored(f)
	return cc
}

// SetNillableStoneStored sets the stone_stored field if the given value is not nil.
func (cc *CityCreate) SetNillableStoneStored(f *float64) *CityCreate {
	if f != nil {
		cc.SetStoneStored(*f)
	}
	return cc
}

// SetIronStored sets the iron_stored field.
func (cc *CityCreate) SetIronStored(f float64) *CityCreate {
	cc.mutation.SetIronStored(f)
	return cc
}

// SetNillableIronStored sets the iron_stored field if the given value is not nil.
func (cc *CityCreate) SetNillableIronStored(f *float64) *CityCreate {
	if f != nil {
		cc.SetIronStored(*f)
	}
	return cc
}

// SetFoodStored sets the food_stored field.
func (cc *CityCreate) SetFoodStored(f float64) *CityCreate {
	cc.mutation.SetFoodStored(f)
	return cc
}

// SetNillableFoodStored sets the food_stored field if the given value is not nil.
func (cc *CityCreate) SetNillableFoodStored(f *float64) *CityCreate {
	if f != nil {
		cc.SetFoodStored(*f)
	}
	return cc
}

// SetWoodLimit sets the wood_limit field.
func (cc *CityCreate) SetWoodLimit(f float64) *CityCreate {
	cc.mutation.SetWoodLimit(f)
	return cc
}

// SetNillableWoodLimit sets the wood_limit field if the given value is not nil.
func (cc *CityCreate) SetNillableWoodLimit(f *float64) *CityCreate {
	if f != nil {
		cc.SetWoodLimit(*f)
	}
	return cc
}

// SetStoneLimit sets the stone_limit field.
func (cc *CityCreate) SetStoneLimit(f float64) *CityCreate {
	cc.mutation.SetStoneLimit(f)
	return cc
}

// SetNillableStoneLimit sets the stone_limit field if the given value is not nil.
func (cc *CityCreate) SetNillableStoneLimit(f *float64) *CityCreate {
	if f != nil {
		cc.SetStoneLimit(*f)
	}
	return cc
}

// SetIronLimit sets the iron_limit field.
func (cc *CityCreate) SetIronLimit(f float64) *CityCreate {
	cc.mutation.SetIronLimit(f)
	return cc
}

// SetNillableIronLimit sets the iron_limit field if the given value is not nil.
func (cc *CityCreate) SetNillableIronLimit(f *float64) *CityCreate {
	if f != nil {
		cc.SetIronLimit(*f)
	}
	return cc
}

// SetFoodLimit sets the food_limit field.
func (cc *CityCreate) SetFoodLimit(f float64) *CityCreate {
	cc.mutation.SetFoodLimit(f)
	return cc
}

// SetNillableFoodLimit sets the food_limit field if the given value is not nil.
func (cc *CityCreate) SetNillableFoodLimit(f *float64) *CityCreate {
	if f != nil {
		cc.SetFoodLimit(*f)
	}
	return cc
}

// SetQueueEndsAt sets the queue_ends_at field.
func (cc *CityCreate) SetQueueEndsAt(t time.Time) *CityCreate {
	cc.mutation.SetQueueEndsAt(t)
	return cc
}

// SetNillableQueueEndsAt sets the queue_ends_at field if the given value is not nil.
func (cc *CityCreate) SetNillableQueueEndsAt(t *time.Time) *CityCreate {
	if t != nil {
		cc.SetQueueEndsAt(*t)
	}
	return cc
}

// SetConstructionSpeed sets the construction_speed field.
func (cc *CityCreate) SetConstructionSpeed(i int) *CityCreate {
	cc.mutation.SetConstructionSpeed(i)
	return cc
}

// SetNillableConstructionSpeed sets the construction_speed field if the given value is not nil.
func (cc *CityCreate) SetNillableConstructionSpeed(i *int) *CityCreate {
	if i != nil {
		cc.SetConstructionSpeed(*i)
	}
	return cc
}

// SetLastUpdated sets the last_updated field.
func (cc *CityCreate) SetLastUpdated(t time.Time) *CityCreate {
	cc.mutation.SetLastUpdated(t)
	return cc
}

// SetNillableLastUpdated sets the last_updated field if the given value is not nil.
func (cc *CityCreate) SetNillableLastUpdated(t *time.Time) *CityCreate {
	if t != nil {
		cc.SetLastUpdated(*t)
	}
	return cc
}

// SetOwnerID sets the owner edge to User by id.
func (cc *CityCreate) SetOwnerID(id int) *CityCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cc *CityCreate) SetNillableOwnerID(id *int) *CityCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the owner edge to User.
func (cc *CityCreate) SetOwner(u *User) *CityCreate {
	return cc.SetOwnerID(u.ID)
}

// AddConstructionIDs adds the constructions edge to Construction by ids.
func (cc *CityCreate) AddConstructionIDs(ids ...int) *CityCreate {
	cc.mutation.AddConstructionIDs(ids...)
	return cc
}

// AddConstructions adds the constructions edges to Construction.
func (cc *CityCreate) AddConstructions(c ...*Construction) *CityCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddConstructionIDs(ids...)
}

// AddQueueIDs adds the queue edge to QueueItem by ids.
func (cc *CityCreate) AddQueueIDs(ids ...int) *CityCreate {
	cc.mutation.AddQueueIDs(ids...)
	return cc
}

// AddQueue adds the queue edges to QueueItem.
func (cc *CityCreate) AddQueue(q ...*QueueItem) *CityCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return cc.AddQueueIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (cc *CityCreate) Mutation() *CityMutation {
	return cc.mutation
}

// Save creates the City in the database.
func (cc *CityCreate) Save(ctx context.Context) (*City, error) {
	var (
		err  error
		node *City
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CityMutation)
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
func (cc *CityCreate) SaveX(ctx context.Context) *City {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *CityCreate) defaults() {
	if _, ok := cc.mutation.X(); !ok {
		v := city.DefaultX
		cc.mutation.SetX(v)
	}
	if _, ok := cc.mutation.Y(); !ok {
		v := city.DefaultY
		cc.mutation.SetY(v)
	}
	if _, ok := cc.mutation.Name(); !ok {
		v := city.DefaultName
		cc.mutation.SetName(v)
	}
	if _, ok := cc.mutation.Points(); !ok {
		v := city.DefaultPoints
		cc.mutation.SetPoints(v)
	}
	if _, ok := cc.mutation.WoodProduction(); !ok {
		v := city.DefaultWoodProduction
		cc.mutation.SetWoodProduction(v)
	}
	if _, ok := cc.mutation.StoneProduction(); !ok {
		v := city.DefaultStoneProduction
		cc.mutation.SetStoneProduction(v)
	}
	if _, ok := cc.mutation.IronProduction(); !ok {
		v := city.DefaultIronProduction
		cc.mutation.SetIronProduction(v)
	}
	if _, ok := cc.mutation.FoodProduction(); !ok {
		v := city.DefaultFoodProduction
		cc.mutation.SetFoodProduction(v)
	}
	if _, ok := cc.mutation.WoodStored(); !ok {
		v := city.DefaultWoodStored
		cc.mutation.SetWoodStored(v)
	}
	if _, ok := cc.mutation.StoneStored(); !ok {
		v := city.DefaultStoneStored
		cc.mutation.SetStoneStored(v)
	}
	if _, ok := cc.mutation.IronStored(); !ok {
		v := city.DefaultIronStored
		cc.mutation.SetIronStored(v)
	}
	if _, ok := cc.mutation.FoodStored(); !ok {
		v := city.DefaultFoodStored
		cc.mutation.SetFoodStored(v)
	}
	if _, ok := cc.mutation.WoodLimit(); !ok {
		v := city.DefaultWoodLimit
		cc.mutation.SetWoodLimit(v)
	}
	if _, ok := cc.mutation.StoneLimit(); !ok {
		v := city.DefaultStoneLimit
		cc.mutation.SetStoneLimit(v)
	}
	if _, ok := cc.mutation.IronLimit(); !ok {
		v := city.DefaultIronLimit
		cc.mutation.SetIronLimit(v)
	}
	if _, ok := cc.mutation.FoodLimit(); !ok {
		v := city.DefaultFoodLimit
		cc.mutation.SetFoodLimit(v)
	}
	if _, ok := cc.mutation.QueueEndsAt(); !ok {
		v := city.DefaultQueueEndsAt()
		cc.mutation.SetQueueEndsAt(v)
	}
	if _, ok := cc.mutation.ConstructionSpeed(); !ok {
		v := city.DefaultConstructionSpeed
		cc.mutation.SetConstructionSpeed(v)
	}
	if _, ok := cc.mutation.LastUpdated(); !ok {
		v := city.DefaultLastUpdated()
		cc.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CityCreate) check() error {
	if _, ok := cc.mutation.X(); !ok {
		return &ValidationError{Name: "x", err: errors.New("ent: missing required field \"x\"")}
	}
	if _, ok := cc.mutation.Y(); !ok {
		return &ValidationError{Name: "y", err: errors.New("ent: missing required field \"y\"")}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := city.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New("ent: missing required field \"points\"")}
	}
	if _, ok := cc.mutation.WoodProduction(); !ok {
		return &ValidationError{Name: "wood_production", err: errors.New("ent: missing required field \"wood_production\"")}
	}
	if _, ok := cc.mutation.StoneProduction(); !ok {
		return &ValidationError{Name: "stone_production", err: errors.New("ent: missing required field \"stone_production\"")}
	}
	if _, ok := cc.mutation.IronProduction(); !ok {
		return &ValidationError{Name: "iron_production", err: errors.New("ent: missing required field \"iron_production\"")}
	}
	if _, ok := cc.mutation.FoodProduction(); !ok {
		return &ValidationError{Name: "food_production", err: errors.New("ent: missing required field \"food_production\"")}
	}
	if _, ok := cc.mutation.WoodStored(); !ok {
		return &ValidationError{Name: "wood_stored", err: errors.New("ent: missing required field \"wood_stored\"")}
	}
	if _, ok := cc.mutation.StoneStored(); !ok {
		return &ValidationError{Name: "stone_stored", err: errors.New("ent: missing required field \"stone_stored\"")}
	}
	if _, ok := cc.mutation.IronStored(); !ok {
		return &ValidationError{Name: "iron_stored", err: errors.New("ent: missing required field \"iron_stored\"")}
	}
	if _, ok := cc.mutation.FoodStored(); !ok {
		return &ValidationError{Name: "food_stored", err: errors.New("ent: missing required field \"food_stored\"")}
	}
	if _, ok := cc.mutation.WoodLimit(); !ok {
		return &ValidationError{Name: "wood_limit", err: errors.New("ent: missing required field \"wood_limit\"")}
	}
	if _, ok := cc.mutation.StoneLimit(); !ok {
		return &ValidationError{Name: "stone_limit", err: errors.New("ent: missing required field \"stone_limit\"")}
	}
	if _, ok := cc.mutation.IronLimit(); !ok {
		return &ValidationError{Name: "iron_limit", err: errors.New("ent: missing required field \"iron_limit\"")}
	}
	if _, ok := cc.mutation.FoodLimit(); !ok {
		return &ValidationError{Name: "food_limit", err: errors.New("ent: missing required field \"food_limit\"")}
	}
	if _, ok := cc.mutation.QueueEndsAt(); !ok {
		return &ValidationError{Name: "queue_ends_at", err: errors.New("ent: missing required field \"queue_ends_at\"")}
	}
	if _, ok := cc.mutation.ConstructionSpeed(); !ok {
		return &ValidationError{Name: "construction_speed", err: errors.New("ent: missing required field \"construction_speed\"")}
	}
	if _, ok := cc.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "last_updated", err: errors.New("ent: missing required field \"last_updated\"")}
	}
	return nil
}

func (cc *CityCreate) sqlSave(ctx context.Context) (*City, error) {
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

func (cc *CityCreate) createSpec() (*City, *sqlgraph.CreateSpec) {
	var (
		_node = &City{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: city.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.X(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldX,
		})
		_node.X = value
	}
	if value, ok := cc.mutation.Y(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldY,
		})
		_node.Y = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Points(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldPoints,
		})
		_node.Points = value
	}
	if value, ok := cc.mutation.WoodProduction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldWoodProduction,
		})
		_node.WoodProduction = value
	}
	if value, ok := cc.mutation.StoneProduction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldStoneProduction,
		})
		_node.StoneProduction = value
	}
	if value, ok := cc.mutation.IronProduction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldIronProduction,
		})
		_node.IronProduction = value
	}
	if value, ok := cc.mutation.FoodProduction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldFoodProduction,
		})
		_node.FoodProduction = value
	}
	if value, ok := cc.mutation.WoodStored(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldWoodStored,
		})
		_node.WoodStored = value
	}
	if value, ok := cc.mutation.StoneStored(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldStoneStored,
		})
		_node.StoneStored = value
	}
	if value, ok := cc.mutation.IronStored(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldIronStored,
		})
		_node.IronStored = value
	}
	if value, ok := cc.mutation.FoodStored(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldFoodStored,
		})
		_node.FoodStored = value
	}
	if value, ok := cc.mutation.WoodLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldWoodLimit,
		})
		_node.WoodLimit = value
	}
	if value, ok := cc.mutation.StoneLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldStoneLimit,
		})
		_node.StoneLimit = value
	}
	if value, ok := cc.mutation.IronLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldIronLimit,
		})
		_node.IronLimit = value
	}
	if value, ok := cc.mutation.FoodLimit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: city.FieldFoodLimit,
		})
		_node.FoodLimit = value
	}
	if value, ok := cc.mutation.QueueEndsAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: city.FieldQueueEndsAt,
		})
		_node.QueueEndsAt = value
	}
	if value, ok := cc.mutation.ConstructionSpeed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldConstructionSpeed,
		})
		_node.ConstructionSpeed = value
	}
	if value, ok := cc.mutation.LastUpdated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: city.FieldLastUpdated,
		})
		_node.LastUpdated = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   city.OwnerTable,
			Columns: []string{city.OwnerColumn},
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
	if nodes := cc.mutation.ConstructionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.ConstructionsTable,
			Columns: []string{city.ConstructionsColumn},
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
	if nodes := cc.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.QueueTable,
			Columns: []string{city.QueueColumn},
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
	return _node, _spec
}

// CityCreateBulk is the builder for creating a bulk of City entities.
type CityCreateBulk struct {
	config
	builders []*CityCreate
}

// Save creates the City entities in the database.
func (ccb *CityCreateBulk) Save(ctx context.Context) ([]*City, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*City, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CityMutation)
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
func (ccb *CityCreateBulk) SaveX(ctx context.Context) []*City {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
