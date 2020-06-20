// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/predicate"
	"github.com/joaopedrosgs/loucore/ent/queueitem"
	"github.com/joaopedrosgs/loucore/ent/user"
)

// QueueItemUpdate is the builder for updating QueueItem entities.
type QueueItemUpdate struct {
	config
	hooks      []Hook
	mutation   *QueueItemMutation
	predicates []predicate.QueueItem
}

// Where adds a new predicate for the builder.
func (qiu *QueueItemUpdate) Where(ps ...predicate.QueueItem) *QueueItemUpdate {
	qiu.predicates = append(qiu.predicates, ps...)
	return qiu
}

// SetStartAt sets the start_at field.
func (qiu *QueueItemUpdate) SetStartAt(t time.Time) *QueueItemUpdate {
	qiu.mutation.SetStartAt(t)
	return qiu
}

// SetDuration sets the duration field.
func (qiu *QueueItemUpdate) SetDuration(i int) *QueueItemUpdate {
	qiu.mutation.ResetDuration()
	qiu.mutation.SetDuration(i)
	return qiu
}

// AddDuration adds i to duration.
func (qiu *QueueItemUpdate) AddDuration(i int) *QueueItemUpdate {
	qiu.mutation.AddDuration(i)
	return qiu
}

// SetCompletion sets the completion field.
func (qiu *QueueItemUpdate) SetCompletion(t time.Time) *QueueItemUpdate {
	qiu.mutation.SetCompletion(t)
	return qiu
}

// SetAction sets the action field.
func (qiu *QueueItemUpdate) SetAction(i int) *QueueItemUpdate {
	qiu.mutation.ResetAction()
	qiu.mutation.SetAction(i)
	return qiu
}

// AddAction adds i to action.
func (qiu *QueueItemUpdate) AddAction(i int) *QueueItemUpdate {
	qiu.mutation.AddAction(i)
	return qiu
}

// SetOrder sets the order field.
func (qiu *QueueItemUpdate) SetOrder(i int) *QueueItemUpdate {
	qiu.mutation.ResetOrder()
	qiu.mutation.SetOrder(i)
	return qiu
}

// AddOrder adds i to order.
func (qiu *QueueItemUpdate) AddOrder(i int) *QueueItemUpdate {
	qiu.mutation.AddOrder(i)
	return qiu
}

// SetOwnerID sets the owner edge to User by id.
func (qiu *QueueItemUpdate) SetOwnerID(id int) *QueueItemUpdate {
	qiu.mutation.SetOwnerID(id)
	return qiu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (qiu *QueueItemUpdate) SetNillableOwnerID(id *int) *QueueItemUpdate {
	if id != nil {
		qiu = qiu.SetOwnerID(*id)
	}
	return qiu
}

// SetOwner sets the owner edge to User.
func (qiu *QueueItemUpdate) SetOwner(u *User) *QueueItemUpdate {
	return qiu.SetOwnerID(u.ID)
}

// SetCityID sets the city edge to City by id.
func (qiu *QueueItemUpdate) SetCityID(id int) *QueueItemUpdate {
	qiu.mutation.SetCityID(id)
	return qiu
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (qiu *QueueItemUpdate) SetNillableCityID(id *int) *QueueItemUpdate {
	if id != nil {
		qiu = qiu.SetCityID(*id)
	}
	return qiu
}

// SetCity sets the city edge to City.
func (qiu *QueueItemUpdate) SetCity(c *City) *QueueItemUpdate {
	return qiu.SetCityID(c.ID)
}

// SetConstructionID sets the construction edge to Construction by id.
func (qiu *QueueItemUpdate) SetConstructionID(id int) *QueueItemUpdate {
	qiu.mutation.SetConstructionID(id)
	return qiu
}

// SetNillableConstructionID sets the construction edge to Construction by id if the given value is not nil.
func (qiu *QueueItemUpdate) SetNillableConstructionID(id *int) *QueueItemUpdate {
	if id != nil {
		qiu = qiu.SetConstructionID(*id)
	}
	return qiu
}

// SetConstruction sets the construction edge to Construction.
func (qiu *QueueItemUpdate) SetConstruction(c *Construction) *QueueItemUpdate {
	return qiu.SetConstructionID(c.ID)
}

// ClearOwner clears the owner edge to User.
func (qiu *QueueItemUpdate) ClearOwner() *QueueItemUpdate {
	qiu.mutation.ClearOwner()
	return qiu
}

// ClearCity clears the city edge to City.
func (qiu *QueueItemUpdate) ClearCity() *QueueItemUpdate {
	qiu.mutation.ClearCity()
	return qiu
}

// ClearConstruction clears the construction edge to Construction.
func (qiu *QueueItemUpdate) ClearConstruction() *QueueItemUpdate {
	qiu.mutation.ClearConstruction()
	return qiu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (qiu *QueueItemUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(qiu.hooks) == 0 {
		affected, err = qiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qiu.mutation = mutation
			affected, err = qiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qiu.hooks) - 1; i >= 0; i-- {
			mut = qiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qiu *QueueItemUpdate) SaveX(ctx context.Context) int {
	affected, err := qiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qiu *QueueItemUpdate) Exec(ctx context.Context) error {
	_, err := qiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qiu *QueueItemUpdate) ExecX(ctx context.Context) {
	if err := qiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qiu *QueueItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queueitem.Table,
			Columns: queueitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queueitem.FieldID,
			},
		},
	}
	if ps := qiu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qiu.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queueitem.FieldStartAt,
		})
	}
	if value, ok := qiu.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldDuration,
		})
	}
	if value, ok := qiu.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldDuration,
		})
	}
	if value, ok := qiu.mutation.Completion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queueitem.FieldCompletion,
		})
	}
	if value, ok := qiu.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldAction,
		})
	}
	if value, ok := qiu.mutation.AddedAction(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldAction,
		})
	}
	if value, ok := qiu.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldOrder,
		})
	}
	if value, ok := qiu.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldOrder,
		})
	}
	if qiu.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qiu.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qiu.mutation.CityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qiu.mutation.CityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qiu.mutation.ConstructionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qiu.mutation.ConstructionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{queueitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// QueueItemUpdateOne is the builder for updating a single QueueItem entity.
type QueueItemUpdateOne struct {
	config
	hooks    []Hook
	mutation *QueueItemMutation
}

// SetStartAt sets the start_at field.
func (qiuo *QueueItemUpdateOne) SetStartAt(t time.Time) *QueueItemUpdateOne {
	qiuo.mutation.SetStartAt(t)
	return qiuo
}

// SetDuration sets the duration field.
func (qiuo *QueueItemUpdateOne) SetDuration(i int) *QueueItemUpdateOne {
	qiuo.mutation.ResetDuration()
	qiuo.mutation.SetDuration(i)
	return qiuo
}

// AddDuration adds i to duration.
func (qiuo *QueueItemUpdateOne) AddDuration(i int) *QueueItemUpdateOne {
	qiuo.mutation.AddDuration(i)
	return qiuo
}

// SetCompletion sets the completion field.
func (qiuo *QueueItemUpdateOne) SetCompletion(t time.Time) *QueueItemUpdateOne {
	qiuo.mutation.SetCompletion(t)
	return qiuo
}

// SetAction sets the action field.
func (qiuo *QueueItemUpdateOne) SetAction(i int) *QueueItemUpdateOne {
	qiuo.mutation.ResetAction()
	qiuo.mutation.SetAction(i)
	return qiuo
}

// AddAction adds i to action.
func (qiuo *QueueItemUpdateOne) AddAction(i int) *QueueItemUpdateOne {
	qiuo.mutation.AddAction(i)
	return qiuo
}

// SetOrder sets the order field.
func (qiuo *QueueItemUpdateOne) SetOrder(i int) *QueueItemUpdateOne {
	qiuo.mutation.ResetOrder()
	qiuo.mutation.SetOrder(i)
	return qiuo
}

// AddOrder adds i to order.
func (qiuo *QueueItemUpdateOne) AddOrder(i int) *QueueItemUpdateOne {
	qiuo.mutation.AddOrder(i)
	return qiuo
}

// SetOwnerID sets the owner edge to User by id.
func (qiuo *QueueItemUpdateOne) SetOwnerID(id int) *QueueItemUpdateOne {
	qiuo.mutation.SetOwnerID(id)
	return qiuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (qiuo *QueueItemUpdateOne) SetNillableOwnerID(id *int) *QueueItemUpdateOne {
	if id != nil {
		qiuo = qiuo.SetOwnerID(*id)
	}
	return qiuo
}

// SetOwner sets the owner edge to User.
func (qiuo *QueueItemUpdateOne) SetOwner(u *User) *QueueItemUpdateOne {
	return qiuo.SetOwnerID(u.ID)
}

// SetCityID sets the city edge to City by id.
func (qiuo *QueueItemUpdateOne) SetCityID(id int) *QueueItemUpdateOne {
	qiuo.mutation.SetCityID(id)
	return qiuo
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (qiuo *QueueItemUpdateOne) SetNillableCityID(id *int) *QueueItemUpdateOne {
	if id != nil {
		qiuo = qiuo.SetCityID(*id)
	}
	return qiuo
}

// SetCity sets the city edge to City.
func (qiuo *QueueItemUpdateOne) SetCity(c *City) *QueueItemUpdateOne {
	return qiuo.SetCityID(c.ID)
}

// SetConstructionID sets the construction edge to Construction by id.
func (qiuo *QueueItemUpdateOne) SetConstructionID(id int) *QueueItemUpdateOne {
	qiuo.mutation.SetConstructionID(id)
	return qiuo
}

// SetNillableConstructionID sets the construction edge to Construction by id if the given value is not nil.
func (qiuo *QueueItemUpdateOne) SetNillableConstructionID(id *int) *QueueItemUpdateOne {
	if id != nil {
		qiuo = qiuo.SetConstructionID(*id)
	}
	return qiuo
}

// SetConstruction sets the construction edge to Construction.
func (qiuo *QueueItemUpdateOne) SetConstruction(c *Construction) *QueueItemUpdateOne {
	return qiuo.SetConstructionID(c.ID)
}

// ClearOwner clears the owner edge to User.
func (qiuo *QueueItemUpdateOne) ClearOwner() *QueueItemUpdateOne {
	qiuo.mutation.ClearOwner()
	return qiuo
}

// ClearCity clears the city edge to City.
func (qiuo *QueueItemUpdateOne) ClearCity() *QueueItemUpdateOne {
	qiuo.mutation.ClearCity()
	return qiuo
}

// ClearConstruction clears the construction edge to Construction.
func (qiuo *QueueItemUpdateOne) ClearConstruction() *QueueItemUpdateOne {
	qiuo.mutation.ClearConstruction()
	return qiuo
}

// Save executes the query and returns the updated entity.
func (qiuo *QueueItemUpdateOne) Save(ctx context.Context) (*QueueItem, error) {

	var (
		err  error
		node *QueueItem
	)
	if len(qiuo.hooks) == 0 {
		node, err = qiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QueueItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qiuo.mutation = mutation
			node, err = qiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qiuo.hooks) - 1; i >= 0; i-- {
			mut = qiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (qiuo *QueueItemUpdateOne) SaveX(ctx context.Context) *QueueItem {
	qi, err := qiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return qi
}

// Exec executes the query on the entity.
func (qiuo *QueueItemUpdateOne) Exec(ctx context.Context) error {
	_, err := qiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qiuo *QueueItemUpdateOne) ExecX(ctx context.Context) {
	if err := qiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qiuo *QueueItemUpdateOne) sqlSave(ctx context.Context) (qi *QueueItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   queueitem.Table,
			Columns: queueitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: queueitem.FieldID,
			},
		},
	}
	id, ok := qiuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing QueueItem.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := qiuo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queueitem.FieldStartAt,
		})
	}
	if value, ok := qiuo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldDuration,
		})
	}
	if value, ok := qiuo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldDuration,
		})
	}
	if value, ok := qiuo.mutation.Completion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: queueitem.FieldCompletion,
		})
	}
	if value, ok := qiuo.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldAction,
		})
	}
	if value, ok := qiuo.mutation.AddedAction(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldAction,
		})
	}
	if value, ok := qiuo.mutation.Order(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldOrder,
		})
	}
	if value, ok := qiuo.mutation.AddedOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: queueitem.FieldOrder,
		})
	}
	if qiuo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qiuo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qiuo.mutation.CityCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qiuo.mutation.CityIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qiuo.mutation.ConstructionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qiuo.mutation.ConstructionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	qi = &QueueItem{config: qiuo.config}
	_spec.Assign = qi.assignValues
	_spec.ScanValues = qi.scanValues()
	if err = sqlgraph.UpdateNode(ctx, qiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{queueitem.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return qi, nil
}
