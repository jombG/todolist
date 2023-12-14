// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the task in the database.
	Table = "tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreateAt,
	FieldUpdateAt,
	FieldTitle,
	FieldDescription,
	FieldCompleted,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultCompleted holds the default value on creation for the "completed" field.
	DefaultCompleted bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateAt orders the results by the create_at field.
func ByCreateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateAt, opts...).ToFunc()
}

// ByUpdateAt orders the results by the update_at field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCompleted orders the results by the completed field.
func ByCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompleted, opts...).ToFunc()
}

// ByDeleteAt orders the results by the delete_at field.
func ByDeleteAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteAt, opts...).ToFunc()
}
