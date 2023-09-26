// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Pyakz/buildbox-api/ent/generated/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldAccountID, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedBy, v))
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldClientID, v))
}

// ManagerID applies equality check predicate on the "manager_id" field. It's identical to ManagerIDEQ.
func ManagerID(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldManagerID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldLocation, v))
}

// Budget applies equality check predicate on the "budget" field. It's identical to BudgetEQ.
func Budget(v float64) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldBudget, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescription, v))
}

// Notes applies equality check predicate on the "notes" field. It's identical to NotesEQ.
func Notes(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldNotes, v))
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldEndDate, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUUID, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeleted, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldAccountID, vs...))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedBy, v))
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldClientID, v))
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldClientID, v))
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldClientID, vs...))
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldClientID, vs...))
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldClientID, v))
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldClientID, v))
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldClientID, v))
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldClientID, v))
}

// ClientIDIsNil applies the IsNil predicate on the "client_id" field.
func ClientIDIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldClientID))
}

// ClientIDNotNil applies the NotNil predicate on the "client_id" field.
func ClientIDNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldClientID))
}

// ManagerIDEQ applies the EQ predicate on the "manager_id" field.
func ManagerIDEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldManagerID, v))
}

// ManagerIDNEQ applies the NEQ predicate on the "manager_id" field.
func ManagerIDNEQ(v int) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldManagerID, v))
}

// ManagerIDIn applies the In predicate on the "manager_id" field.
func ManagerIDIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldManagerID, vs...))
}

// ManagerIDNotIn applies the NotIn predicate on the "manager_id" field.
func ManagerIDNotIn(vs ...int) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldManagerID, vs...))
}

// ManagerIDGT applies the GT predicate on the "manager_id" field.
func ManagerIDGT(v int) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldManagerID, v))
}

// ManagerIDGTE applies the GTE predicate on the "manager_id" field.
func ManagerIDGTE(v int) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldManagerID, v))
}

// ManagerIDLT applies the LT predicate on the "manager_id" field.
func ManagerIDLT(v int) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldManagerID, v))
}

// ManagerIDLTE applies the LTE predicate on the "manager_id" field.
func ManagerIDLTE(v int) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldManagerID, v))
}

// ManagerIDIsNil applies the IsNil predicate on the "manager_id" field.
func ManagerIDIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldManagerID))
}

// ManagerIDNotNil applies the NotNil predicate on the "manager_id" field.
func ManagerIDNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldManagerID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldName, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldStatus))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationIsNil applies the IsNil predicate on the "location" field.
func LocationIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldLocation))
}

// LocationNotNil applies the NotNil predicate on the "location" field.
func LocationNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldLocation))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldLocation, v))
}

// BudgetEQ applies the EQ predicate on the "budget" field.
func BudgetEQ(v float64) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldBudget, v))
}

// BudgetNEQ applies the NEQ predicate on the "budget" field.
func BudgetNEQ(v float64) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldBudget, v))
}

// BudgetIn applies the In predicate on the "budget" field.
func BudgetIn(vs ...float64) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldBudget, vs...))
}

// BudgetNotIn applies the NotIn predicate on the "budget" field.
func BudgetNotIn(vs ...float64) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldBudget, vs...))
}

// BudgetGT applies the GT predicate on the "budget" field.
func BudgetGT(v float64) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldBudget, v))
}

// BudgetGTE applies the GTE predicate on the "budget" field.
func BudgetGTE(v float64) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldBudget, v))
}

// BudgetLT applies the LT predicate on the "budget" field.
func BudgetLT(v float64) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldBudget, v))
}

// BudgetLTE applies the LTE predicate on the "budget" field.
func BudgetLTE(v float64) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldBudget, v))
}

// BudgetIsNil applies the IsNil predicate on the "budget" field.
func BudgetIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldBudget))
}

// BudgetNotNil applies the NotNil predicate on the "budget" field.
func BudgetNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldBudget))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldDescription, v))
}

// NotesEQ applies the EQ predicate on the "notes" field.
func NotesEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldNotes, v))
}

// NotesNEQ applies the NEQ predicate on the "notes" field.
func NotesNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldNotes, v))
}

// NotesIn applies the In predicate on the "notes" field.
func NotesIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldNotes, vs...))
}

// NotesNotIn applies the NotIn predicate on the "notes" field.
func NotesNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldNotes, vs...))
}

// NotesGT applies the GT predicate on the "notes" field.
func NotesGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldNotes, v))
}

// NotesGTE applies the GTE predicate on the "notes" field.
func NotesGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldNotes, v))
}

// NotesLT applies the LT predicate on the "notes" field.
func NotesLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldNotes, v))
}

// NotesLTE applies the LTE predicate on the "notes" field.
func NotesLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldNotes, v))
}

// NotesContains applies the Contains predicate on the "notes" field.
func NotesContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldNotes, v))
}

// NotesHasPrefix applies the HasPrefix predicate on the "notes" field.
func NotesHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldNotes, v))
}

// NotesHasSuffix applies the HasSuffix predicate on the "notes" field.
func NotesHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldNotes, v))
}

// NotesIsNil applies the IsNil predicate on the "notes" field.
func NotesIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldNotes))
}

// NotesNotNil applies the NotNil predicate on the "notes" field.
func NotesNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldNotes))
}

// NotesEqualFold applies the EqualFold predicate on the "notes" field.
func NotesEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldNotes, v))
}

// NotesContainsFold applies the ContainsFold predicate on the "notes" field.
func NotesContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldNotes, v))
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldStartDate, v))
}

// StartDateIsNil applies the IsNil predicate on the "start_date" field.
func StartDateIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldStartDate))
}

// StartDateNotNil applies the NotNil predicate on the "start_date" field.
func StartDateNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldStartDate))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldEndDate, v))
}

// EndDateIsNil applies the IsNil predicate on the "end_date" field.
func EndDateIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldEndDate))
}

// EndDateNotNil applies the NotNil predicate on the "end_date" field.
func EndDateNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldEndDate))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUUID, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDeleted, v))
}

// DeletedIsNil applies the IsNil predicate on the "deleted" field.
func DeletedIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDeleted))
}

// DeletedNotNil applies the NotNil predicate on the "deleted" field.
func DeletedNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDeleted))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedAt, v))
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(sql.NotPredicates(p))
}
