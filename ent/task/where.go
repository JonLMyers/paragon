// Code generated by entc, DO NOT EDIT.

package task

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// QueueTime applies equality check predicate on the "QueueTime" field. It's identical to QueueTimeEQ.
func QueueTime(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldQueueTime), v))
		},
	)
}

// ClaimTime applies equality check predicate on the "ClaimTime" field. It's identical to ClaimTimeEQ.
func ClaimTime(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldClaimTime), v))
		},
	)
}

// ExecStartTime applies equality check predicate on the "ExecStartTime" field. It's identical to ExecStartTimeEQ.
func ExecStartTime(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExecStartTime), v))
		},
	)
}

// ExecStopTime applies equality check predicate on the "ExecStopTime" field. It's identical to ExecStopTimeEQ.
func ExecStopTime(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExecStopTime), v))
		},
	)
}

// Content applies equality check predicate on the "Content" field. It's identical to ContentEQ.
func Content(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldContent), v))
		},
	)
}

// Error applies equality check predicate on the "Error" field. It's identical to ErrorEQ.
func Error(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldError), v))
		},
	)
}

// SessionID applies equality check predicate on the "SessionID" field. It's identical to SessionIDEQ.
func SessionID(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSessionID), v))
		},
	)
}

// QueueTimeEQ applies the EQ predicate on the "QueueTime" field.
func QueueTimeEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldQueueTime), v))
		},
	)
}

// QueueTimeNEQ applies the NEQ predicate on the "QueueTime" field.
func QueueTimeNEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldQueueTime), v))
		},
	)
}

// QueueTimeIn applies the In predicate on the "QueueTime" field.
func QueueTimeIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldQueueTime), v...))
		},
	)
}

// QueueTimeNotIn applies the NotIn predicate on the "QueueTime" field.
func QueueTimeNotIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldQueueTime), v...))
		},
	)
}

// QueueTimeGT applies the GT predicate on the "QueueTime" field.
func QueueTimeGT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldQueueTime), v))
		},
	)
}

// QueueTimeGTE applies the GTE predicate on the "QueueTime" field.
func QueueTimeGTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldQueueTime), v))
		},
	)
}

// QueueTimeLT applies the LT predicate on the "QueueTime" field.
func QueueTimeLT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldQueueTime), v))
		},
	)
}

// QueueTimeLTE applies the LTE predicate on the "QueueTime" field.
func QueueTimeLTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldQueueTime), v))
		},
	)
}

// ClaimTimeEQ applies the EQ predicate on the "ClaimTime" field.
func ClaimTimeEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldClaimTime), v))
		},
	)
}

// ClaimTimeNEQ applies the NEQ predicate on the "ClaimTime" field.
func ClaimTimeNEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldClaimTime), v))
		},
	)
}

// ClaimTimeIn applies the In predicate on the "ClaimTime" field.
func ClaimTimeIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldClaimTime), v...))
		},
	)
}

// ClaimTimeNotIn applies the NotIn predicate on the "ClaimTime" field.
func ClaimTimeNotIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldClaimTime), v...))
		},
	)
}

// ClaimTimeGT applies the GT predicate on the "ClaimTime" field.
func ClaimTimeGT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldClaimTime), v))
		},
	)
}

// ClaimTimeGTE applies the GTE predicate on the "ClaimTime" field.
func ClaimTimeGTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldClaimTime), v))
		},
	)
}

// ClaimTimeLT applies the LT predicate on the "ClaimTime" field.
func ClaimTimeLT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldClaimTime), v))
		},
	)
}

// ClaimTimeLTE applies the LTE predicate on the "ClaimTime" field.
func ClaimTimeLTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldClaimTime), v))
		},
	)
}

// ClaimTimeIsNil applies the IsNil predicate on the "ClaimTime" field.
func ClaimTimeIsNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldClaimTime)))
		},
	)
}

// ClaimTimeNotNil applies the NotNil predicate on the "ClaimTime" field.
func ClaimTimeNotNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldClaimTime)))
		},
	)
}

// ExecStartTimeEQ applies the EQ predicate on the "ExecStartTime" field.
func ExecStartTimeEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExecStartTime), v))
		},
	)
}

// ExecStartTimeNEQ applies the NEQ predicate on the "ExecStartTime" field.
func ExecStartTimeNEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldExecStartTime), v))
		},
	)
}

// ExecStartTimeIn applies the In predicate on the "ExecStartTime" field.
func ExecStartTimeIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldExecStartTime), v...))
		},
	)
}

// ExecStartTimeNotIn applies the NotIn predicate on the "ExecStartTime" field.
func ExecStartTimeNotIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldExecStartTime), v...))
		},
	)
}

// ExecStartTimeGT applies the GT predicate on the "ExecStartTime" field.
func ExecStartTimeGT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldExecStartTime), v))
		},
	)
}

// ExecStartTimeGTE applies the GTE predicate on the "ExecStartTime" field.
func ExecStartTimeGTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldExecStartTime), v))
		},
	)
}

// ExecStartTimeLT applies the LT predicate on the "ExecStartTime" field.
func ExecStartTimeLT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldExecStartTime), v))
		},
	)
}

// ExecStartTimeLTE applies the LTE predicate on the "ExecStartTime" field.
func ExecStartTimeLTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldExecStartTime), v))
		},
	)
}

// ExecStartTimeIsNil applies the IsNil predicate on the "ExecStartTime" field.
func ExecStartTimeIsNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldExecStartTime)))
		},
	)
}

// ExecStartTimeNotNil applies the NotNil predicate on the "ExecStartTime" field.
func ExecStartTimeNotNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldExecStartTime)))
		},
	)
}

// ExecStopTimeEQ applies the EQ predicate on the "ExecStopTime" field.
func ExecStopTimeEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExecStopTime), v))
		},
	)
}

// ExecStopTimeNEQ applies the NEQ predicate on the "ExecStopTime" field.
func ExecStopTimeNEQ(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldExecStopTime), v))
		},
	)
}

// ExecStopTimeIn applies the In predicate on the "ExecStopTime" field.
func ExecStopTimeIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldExecStopTime), v...))
		},
	)
}

// ExecStopTimeNotIn applies the NotIn predicate on the "ExecStopTime" field.
func ExecStopTimeNotIn(vs ...int64) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldExecStopTime), v...))
		},
	)
}

// ExecStopTimeGT applies the GT predicate on the "ExecStopTime" field.
func ExecStopTimeGT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldExecStopTime), v))
		},
	)
}

// ExecStopTimeGTE applies the GTE predicate on the "ExecStopTime" field.
func ExecStopTimeGTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldExecStopTime), v))
		},
	)
}

// ExecStopTimeLT applies the LT predicate on the "ExecStopTime" field.
func ExecStopTimeLT(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldExecStopTime), v))
		},
	)
}

// ExecStopTimeLTE applies the LTE predicate on the "ExecStopTime" field.
func ExecStopTimeLTE(v int64) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldExecStopTime), v))
		},
	)
}

// ExecStopTimeIsNil applies the IsNil predicate on the "ExecStopTime" field.
func ExecStopTimeIsNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldExecStopTime)))
		},
	)
}

// ExecStopTimeNotNil applies the NotNil predicate on the "ExecStopTime" field.
func ExecStopTimeNotNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldExecStopTime)))
		},
	)
}

// ContentEQ applies the EQ predicate on the "Content" field.
func ContentEQ(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldContent), v))
		},
	)
}

// ContentNEQ applies the NEQ predicate on the "Content" field.
func ContentNEQ(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldContent), v))
		},
	)
}

// ContentIn applies the In predicate on the "Content" field.
func ContentIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldContent), v...))
		},
	)
}

// ContentNotIn applies the NotIn predicate on the "Content" field.
func ContentNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldContent), v...))
		},
	)
}

// ContentGT applies the GT predicate on the "Content" field.
func ContentGT(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldContent), v))
		},
	)
}

// ContentGTE applies the GTE predicate on the "Content" field.
func ContentGTE(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldContent), v))
		},
	)
}

// ContentLT applies the LT predicate on the "Content" field.
func ContentLT(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldContent), v))
		},
	)
}

// ContentLTE applies the LTE predicate on the "Content" field.
func ContentLTE(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldContent), v))
		},
	)
}

// ContentContains applies the Contains predicate on the "Content" field.
func ContentContains(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldContent), v))
		},
	)
}

// ContentHasPrefix applies the HasPrefix predicate on the "Content" field.
func ContentHasPrefix(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldContent), v))
		},
	)
}

// ContentHasSuffix applies the HasSuffix predicate on the "Content" field.
func ContentHasSuffix(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldContent), v))
		},
	)
}

// ContentEqualFold applies the EqualFold predicate on the "Content" field.
func ContentEqualFold(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldContent), v))
		},
	)
}

// ContentContainsFold applies the ContainsFold predicate on the "Content" field.
func ContentContainsFold(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldContent), v))
		},
	)
}

// OutputIsNil applies the IsNil predicate on the "Output" field.
func OutputIsNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldOutput)))
		},
	)
}

// OutputNotNil applies the NotNil predicate on the "Output" field.
func OutputNotNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldOutput)))
		},
	)
}

// ErrorEQ applies the EQ predicate on the "Error" field.
func ErrorEQ(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldError), v))
		},
	)
}

// ErrorNEQ applies the NEQ predicate on the "Error" field.
func ErrorNEQ(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldError), v))
		},
	)
}

// ErrorIn applies the In predicate on the "Error" field.
func ErrorIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldError), v...))
		},
	)
}

// ErrorNotIn applies the NotIn predicate on the "Error" field.
func ErrorNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldError), v...))
		},
	)
}

// ErrorGT applies the GT predicate on the "Error" field.
func ErrorGT(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldError), v))
		},
	)
}

// ErrorGTE applies the GTE predicate on the "Error" field.
func ErrorGTE(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldError), v))
		},
	)
}

// ErrorLT applies the LT predicate on the "Error" field.
func ErrorLT(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldError), v))
		},
	)
}

// ErrorLTE applies the LTE predicate on the "Error" field.
func ErrorLTE(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldError), v))
		},
	)
}

// ErrorContains applies the Contains predicate on the "Error" field.
func ErrorContains(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldError), v))
		},
	)
}

// ErrorHasPrefix applies the HasPrefix predicate on the "Error" field.
func ErrorHasPrefix(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldError), v))
		},
	)
}

// ErrorHasSuffix applies the HasSuffix predicate on the "Error" field.
func ErrorHasSuffix(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldError), v))
		},
	)
}

// ErrorIsNil applies the IsNil predicate on the "Error" field.
func ErrorIsNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldError)))
		},
	)
}

// ErrorNotNil applies the NotNil predicate on the "Error" field.
func ErrorNotNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldError)))
		},
	)
}

// ErrorEqualFold applies the EqualFold predicate on the "Error" field.
func ErrorEqualFold(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldError), v))
		},
	)
}

// ErrorContainsFold applies the ContainsFold predicate on the "Error" field.
func ErrorContainsFold(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldError), v))
		},
	)
}

// SessionIDEQ applies the EQ predicate on the "SessionID" field.
func SessionIDEQ(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDNEQ applies the NEQ predicate on the "SessionID" field.
func SessionIDNEQ(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDIn applies the In predicate on the "SessionID" field.
func SessionIDIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldSessionID), v...))
		},
	)
}

// SessionIDNotIn applies the NotIn predicate on the "SessionID" field.
func SessionIDNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldSessionID), v...))
		},
	)
}

// SessionIDGT applies the GT predicate on the "SessionID" field.
func SessionIDGT(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDGTE applies the GTE predicate on the "SessionID" field.
func SessionIDGTE(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDLT applies the LT predicate on the "SessionID" field.
func SessionIDLT(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDLTE applies the LTE predicate on the "SessionID" field.
func SessionIDLTE(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDContains applies the Contains predicate on the "SessionID" field.
func SessionIDContains(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDHasPrefix applies the HasPrefix predicate on the "SessionID" field.
func SessionIDHasPrefix(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDHasSuffix applies the HasSuffix predicate on the "SessionID" field.
func SessionIDHasSuffix(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDIsNil applies the IsNil predicate on the "SessionID" field.
func SessionIDIsNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldSessionID)))
		},
	)
}

// SessionIDNotNil applies the NotNil predicate on the "SessionID" field.
func SessionIDNotNil() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldSessionID)))
		},
	)
}

// SessionIDEqualFold applies the EqualFold predicate on the "SessionID" field.
func SessionIDEqualFold(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldSessionID), v))
		},
	)
}

// SessionIDContainsFold applies the ContainsFold predicate on the "SessionID" field.
func SessionIDContainsFold(v string) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldSessionID), v))
		},
	)
}

// HasTarget applies the HasEdge predicate on the "target" edge.
func HasTarget() predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(TargetColumn)))
		},
	)
}

// HasTargetWith applies the HasEdge predicate on the "target" edge with a given conditions (other predicates).
func HasTargetWith(preds ...predicate.Target) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FieldID).From(sql.Table(TargetInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(TargetColumn), t2))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			for _, p := range predicates {
				p(s)
			}
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			for i, p := range predicates {
				if i > 0 {
					s.Or()
				}
				p(s)
			}
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
