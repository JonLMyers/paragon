// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Target is the model entity for the Target schema.
type Target struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// MachineUUID holds the value of the "MachineUUID" field.
	MachineUUID string `json:"MachineUUID,omitempty"`
	// PrimaryIP holds the value of the "PrimaryIP" field.
	PrimaryIP string `json:"PrimaryIP,omitempty"`
	// PublicIP holds the value of the "PublicIP" field.
	PublicIP string `json:"PublicIP,omitempty"`
	// PrimaryMAC holds the value of the "PrimaryMAC" field.
	PrimaryMAC string `json:"PrimaryMAC,omitempty"`
	// Hostname holds the value of the "Hostname" field.
	Hostname string `json:"Hostname,omitempty"`
	// LastSeen holds the value of the "LastSeen" field.
	LastSeen int64 `json:"LastSeen,omitempty"`
}

// FromRows scans the sql response data into Target.
func (t *Target) FromRows(rows *sql.Rows) error {
	var vt struct {
		ID          int
		Name        sql.NullString
		MachineUUID sql.NullString
		PrimaryIP   sql.NullString
		PublicIP    sql.NullString
		PrimaryMAC  sql.NullString
		Hostname    sql.NullString
		LastSeen    sql.NullInt64
	}
	// the order here should be the same as in the `target.Columns`.
	if err := rows.Scan(
		&vt.ID,
		&vt.Name,
		&vt.MachineUUID,
		&vt.PrimaryIP,
		&vt.PublicIP,
		&vt.PrimaryMAC,
		&vt.Hostname,
		&vt.LastSeen,
	); err != nil {
		return err
	}
	t.ID = vt.ID
	t.Name = vt.Name.String
	t.MachineUUID = vt.MachineUUID.String
	t.PrimaryIP = vt.PrimaryIP.String
	t.PublicIP = vt.PublicIP.String
	t.PrimaryMAC = vt.PrimaryMAC.String
	t.Hostname = vt.Hostname.String
	t.LastSeen = vt.LastSeen.Int64
	return nil
}

// QueryTasks queries the tasks edge of the Target.
func (t *Target) QueryTasks() *TaskQuery {
	return (&TargetClient{t.config}).QueryTasks(t)
}

// Update returns a builder for updating this Target.
// Note that, you need to call Target.Unwrap() before calling this method, if this Target
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Target) Update() *TargetUpdateOne {
	return (&TargetClient{t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Target) Unwrap() *Target {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Target is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Target) String() string {
	var builder strings.Builder
	builder.WriteString("Target(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", Name=")
	builder.WriteString(t.Name)
	builder.WriteString(", MachineUUID=")
	builder.WriteString(t.MachineUUID)
	builder.WriteString(", PrimaryIP=")
	builder.WriteString(t.PrimaryIP)
	builder.WriteString(", PublicIP=")
	builder.WriteString(t.PublicIP)
	builder.WriteString(", PrimaryMAC=")
	builder.WriteString(t.PrimaryMAC)
	builder.WriteString(", Hostname=")
	builder.WriteString(t.Hostname)
	builder.WriteString(", LastSeen=")
	builder.WriteString(fmt.Sprintf("%v", t.LastSeen))
	builder.WriteByte(')')
	return builder.String()
}

// Targets is a parsable slice of Target.
type Targets []*Target

// FromRows scans the sql response data into Targets.
func (t *Targets) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vt := &Target{}
		if err := vt.FromRows(rows); err != nil {
			return err
		}
		*t = append(*t, vt)
	}
	return nil
}

func (t Targets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
