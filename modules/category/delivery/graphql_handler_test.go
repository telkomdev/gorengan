package delivery

import (
	"testing"

	"github.com/musobarlab/gorengan/modules/category/graphql/schema"
	usecaseMock "github.com/musobarlab/gorengan/modules/category/usecase/mock"
	"golang.org/x/net/context"
)

func TestGraphQLHandler(t *testing.T) {

	t.Run("should return success test mutation create category", func(t *testing.T) {
		categoryUsecaseMock := usecaseMock.NewCategoryUsecaseMock()

		handler := &GraphQLCategoryHandler{
			CategoryUsecase: categoryUsecaseMock,
		}

		ctx := context.Background()

		categoryInputArgs := &CategoryInputArgs{
			Category: schema.CategorySchemaInput{
				ID:   "1",
				Name: "Music",
			},
		}

		categoryCreated, err := handler.CreateCategory(ctx, categoryInputArgs)

		if err != nil {
			t.Error("create category mutation should return success")
		}

		if categoryCreated == nil {
			t.Error("category created should not be nil")
		}

	})

}
