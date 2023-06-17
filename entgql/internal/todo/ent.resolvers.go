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

package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marcozac/ent-contrib/entgql"
	"github.com/marcozac/ent-contrib/entgql/internal/todo/ent"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) BillProducts(ctx context.Context) ([]*ent.BillProduct, error) {
	return r.client.BillProduct.Query().All(ctx)
}

func (r *queryResolver) Categories(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*ent.CategoryOrder, where *ent.CategoryWhereInput) (*ent.CategoryConnection, error) {
	return r.client.Category.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithCategoryOrder(orderBy),
			ent.WithCategoryFilter(where.Filter),
		)
}

func (r *queryResolver) Groups(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.client.Group.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGroupFilter(where.Filter),
		)
}

func (r *queryResolver) OneToMany(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.OneToManyOrder, where *ent.OneToManyWhereInput) (*ent.OneToManyConnection, error) {
	return r.client.OneToMany.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithOneToManyFilter(where.Filter),
		)
}

func (r *queryResolver) Todos(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}

func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserFilter(where.Filter),
		)
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// CreateCategoryInput returns CreateCategoryInputResolver implementation.
func (r *Resolver) CreateCategoryInput() CreateCategoryInputResolver {
	return &createCategoryInputResolver{r}
}

// TodoWhereInput returns TodoWhereInputResolver implementation.
func (r *Resolver) TodoWhereInput() TodoWhereInputResolver { return &todoWhereInputResolver{r} }

type (
	categoryResolver            struct{ *Resolver }
	queryResolver               struct{ *Resolver }
	todoResolver                struct{ *Resolver }
	createCategoryInputResolver struct{ *Resolver }
	todoWhereInputResolver      struct{ *Resolver }
)
