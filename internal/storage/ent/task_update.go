// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todolist/internal/storage/ent/predicate"
	"todolist/internal/storage/ent/task"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateAt sets the "update_at" field.
func (tu *TaskUpdate) SetUpdateAt(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdateAt(t)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TaskUpdate) SetTitle(s string) *TaskUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableTitle(s *string) *TaskUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TaskUpdate) SetDescription(s string) *TaskUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDescription(s *string) *TaskUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TaskUpdate) ClearDescription() *TaskUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetCompleted sets the "completed" field.
func (tu *TaskUpdate) SetCompleted(b bool) *TaskUpdate {
	tu.mutation.SetCompleted(b)
	return tu
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCompleted(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetCompleted(*b)
	}
	return tu
}

// SetDeleteAt sets the "delete_at" field.
func (tu *TaskUpdate) SetDeleteAt(t time.Time) *TaskUpdate {
	tu.mutation.SetDeleteAt(t)
	return tu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDeleteAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDeleteAt(*t)
	}
	return tu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (tu *TaskUpdate) ClearDeleteAt() *TaskUpdate {
	tu.mutation.ClearDeleteAt()
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() {
	if _, ok := tu.mutation.UpdateAt(); !ok {
		v := task.UpdateDefaultUpdateAt()
		tu.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateAt(); ok {
		_spec.SetField(task.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.Completed(); ok {
		_spec.SetField(task.FieldCompleted, field.TypeBool, value)
	}
	if value, ok := tu.mutation.DeleteAt(); ok {
		_spec.SetField(task.FieldDeleteAt, field.TypeTime, value)
	}
	if tu.mutation.DeleteAtCleared() {
		_spec.ClearField(task.FieldDeleteAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetUpdateAt sets the "update_at" field.
func (tuo *TaskUpdateOne) SetUpdateAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdateAt(t)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TaskUpdateOne) SetTitle(s string) *TaskUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTitle(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TaskUpdateOne) SetDescription(s string) *TaskUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDescription(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TaskUpdateOne) ClearDescription() *TaskUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetCompleted sets the "completed" field.
func (tuo *TaskUpdateOne) SetCompleted(b bool) *TaskUpdateOne {
	tuo.mutation.SetCompleted(b)
	return tuo
}

// SetNillableCompleted sets the "completed" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCompleted(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetCompleted(*b)
	}
	return tuo
}

// SetDeleteAt sets the "delete_at" field.
func (tuo *TaskUpdateOne) SetDeleteAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDeleteAt(t)
	return tuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDeleteAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDeleteAt(*t)
	}
	return tuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (tuo *TaskUpdateOne) ClearDeleteAt() *TaskUpdateOne {
	tuo.mutation.ClearDeleteAt()
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateAt(); !ok {
		v := task.UpdateDefaultUpdateAt()
		tuo.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdateAt(); ok {
		_spec.SetField(task.FieldUpdateAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(task.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.Completed(); ok {
		_spec.SetField(task.FieldCompleted, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.DeleteAt(); ok {
		_spec.SetField(task.FieldDeleteAt, field.TypeTime, value)
	}
	if tuo.mutation.DeleteAtCleared() {
		_spec.ClearField(task.FieldDeleteAt, field.TypeTime)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}