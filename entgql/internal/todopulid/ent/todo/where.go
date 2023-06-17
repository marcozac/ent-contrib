// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package todo

import (
	"time"

	"github.com/marcozac/ent-contrib/entgql/internal/todopulid/ent/predicate"
	"github.com/marcozac/ent-contrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldPriority, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldText, v))
}

// Blob applies equality check predicate on the "blob" field. It's identical to BlobEQ.
func Blob(v []byte) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldBlob, v))
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCategoryID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldCreatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldStatus, vs...))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v int) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...int) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...int) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldPriority, vs...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v int) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldPriority, v))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v int) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldPriority, v))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v int) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldPriority, v))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v int) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldPriority, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldText, v))
}

// BlobEQ applies the EQ predicate on the "blob" field.
func BlobEQ(v []byte) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldBlob, v))
}

// BlobNEQ applies the NEQ predicate on the "blob" field.
func BlobNEQ(v []byte) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldBlob, v))
}

// BlobIn applies the In predicate on the "blob" field.
func BlobIn(vs ...[]byte) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldBlob, vs...))
}

// BlobNotIn applies the NotIn predicate on the "blob" field.
func BlobNotIn(vs ...[]byte) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldBlob, vs...))
}

// BlobGT applies the GT predicate on the "blob" field.
func BlobGT(v []byte) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldBlob, v))
}

// BlobGTE applies the GTE predicate on the "blob" field.
func BlobGTE(v []byte) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldBlob, v))
}

// BlobLT applies the LT predicate on the "blob" field.
func BlobLT(v []byte) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldBlob, v))
}

// BlobLTE applies the LTE predicate on the "blob" field.
func BlobLTE(v []byte) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldBlob, v))
}

// BlobIsNil applies the IsNil predicate on the "blob" field.
func BlobIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldBlob))
}

// BlobNotNil applies the NotNil predicate on the "blob" field.
func BlobNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldBlob))
}

// InitIsNil applies the IsNil predicate on the "init" field.
func InitIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldInit))
}

// InitNotNil applies the NotNil predicate on the "init" field.
func InitNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldInit))
}

// CustomIsNil applies the IsNil predicate on the "custom" field.
func CustomIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldCustom))
}

// CustomNotNil applies the NotNil predicate on the "custom" field.
func CustomNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldCustom))
}

// CustompIsNil applies the IsNil predicate on the "customp" field.
func CustompIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldCustomp))
}

// CustompNotNil applies the NotNil predicate on the "customp" field.
func CustompNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldCustomp))
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCategoryID, v))
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldCategoryID, v))
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldCategoryID, vs...))
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldCategoryID, vs...))
}

// CategoryIDGT applies the GT predicate on the "category_id" field.
func CategoryIDGT(v pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldCategoryID, v))
}

// CategoryIDGTE applies the GTE predicate on the "category_id" field.
func CategoryIDGTE(v pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldCategoryID, v))
}

// CategoryIDLT applies the LT predicate on the "category_id" field.
func CategoryIDLT(v pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldCategoryID, v))
}

// CategoryIDLTE applies the LTE predicate on the "category_id" field.
func CategoryIDLTE(v pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldCategoryID, v))
}

// CategoryIDContains applies the Contains predicate on the "category_id" field.
func CategoryIDContains(v pulid.ID) predicate.Todo {
	vc := string(v)
	return predicate.Todo(sql.FieldContains(FieldCategoryID, vc))
}

// CategoryIDHasPrefix applies the HasPrefix predicate on the "category_id" field.
func CategoryIDHasPrefix(v pulid.ID) predicate.Todo {
	vc := string(v)
	return predicate.Todo(sql.FieldHasPrefix(FieldCategoryID, vc))
}

// CategoryIDHasSuffix applies the HasSuffix predicate on the "category_id" field.
func CategoryIDHasSuffix(v pulid.ID) predicate.Todo {
	vc := string(v)
	return predicate.Todo(sql.FieldHasSuffix(FieldCategoryID, vc))
}

// CategoryIDIsNil applies the IsNil predicate on the "category_id" field.
func CategoryIDIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldCategoryID))
}

// CategoryIDNotNil applies the NotNil predicate on the "category_id" field.
func CategoryIDNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldCategoryID))
}

// CategoryIDEqualFold applies the EqualFold predicate on the "category_id" field.
func CategoryIDEqualFold(v pulid.ID) predicate.Todo {
	vc := string(v)
	return predicate.Todo(sql.FieldEqualFold(FieldCategoryID, vc))
}

// CategoryIDContainsFold applies the ContainsFold predicate on the "category_id" field.
func CategoryIDContainsFold(v pulid.ID) predicate.Todo {
	vc := string(v)
	return predicate.Todo(sql.FieldContainsFold(FieldCategoryID, vc))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSecret applies the HasEdge predicate on the "secret" edge.
func HasSecret() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SecretTable, SecretColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSecretWith applies the HasEdge predicate on the "secret" edge with a given conditions (other predicates).
func HasSecretWith(preds ...predicate.VerySecret) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newSecretStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		p(s.Not())
	})
}
