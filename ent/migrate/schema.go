// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/kcarretto/paragon/ent/task"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// TargetsColumns holds the columns for the "targets" table.
	TargetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "machine_uuid", Type: field.TypeString, Unique: true},
		{Name: "primary_ip", Type: field.TypeString, Nullable: true},
		{Name: "public_ip", Type: field.TypeString, Nullable: true},
		{Name: "primary_mac", Type: field.TypeString, Nullable: true},
		{Name: "hostname", Type: field.TypeString, Nullable: true},
		{Name: "last_seen", Type: field.TypeInt64, Nullable: true},
	}
	// TargetsTable holds the schema information for the "targets" table.
	TargetsTable = &schema.Table{
		Name:        "targets",
		Columns:     TargetsColumns,
		PrimaryKey:  []*schema.Column{TargetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "queue_time", Type: field.TypeInt64, Default: task.DefaultQueueTime},
		{Name: "claim_time", Type: field.TypeInt64, Nullable: true},
		{Name: "exec_start_time", Type: field.TypeInt64, Nullable: true},
		{Name: "exec_stop_time", Type: field.TypeInt64, Nullable: true},
		{Name: "content", Type: field.TypeString},
		{Name: "output", Type: field.TypeJSON, Nullable: true},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "session_id", Type: field.TypeString, Nullable: true},
		{Name: "target_id", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "tasks_targets_tasks",
				Columns: []*schema.Column{TasksColumns[9]},

				RefColumns: []*schema.Column{TargetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TargetsTable,
		TasksTable,
	}
)

func init() {
	TasksTable.ForeignKeys[0].RefTable = TargetsTable
}
