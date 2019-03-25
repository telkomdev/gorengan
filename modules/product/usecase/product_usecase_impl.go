package usecase

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/musobarlab/gorengan/modules/product/domain"
	"github.com/musobarlab/gorengan/modules/product/repository"
	"github.com/musobarlab/gorengan/modules/shared"
)

// ProductUsecaseImpl struct
type ProductUsecaseImpl struct {
	productRepositoryRead  repository.ProductRepository
	productRepositoryWrite repository.ProductRepository
}

// NewProductUsecaseImpl function
func NewProductUsecaseImpl(productRepositoryRead, productRepositoryWrite repository.ProductRepository) *ProductUsecaseImpl {
	return &ProductUsecaseImpl{productRepositoryRead: productRepositoryRead, productRepositoryWrite: productRepositoryWrite}
}

// CreateProduct function
func (u *ProductUsecaseImpl) CreateProduct(product *domain.Product) shared.Output {
	productOutput := u.productRepositoryRead.FindByID(product.ID)
	if productOutput.Err != nil && productOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output{Err: productOutput.Err}
	}

	if productOutput.Result != nil {
		productExist := productOutput.Result.(*domain.Product)

		if productExist != nil {
			return shared.Output{Err: fmt.Errorf("product with id %s already exist", product.ID)}
		}
	}

	err := product.Validate()

	if err != nil {
		return shared.Output{Err: err}
	}

	productSaveOutput := u.productRepositoryWrite.Save(product)
	if productSaveOutput.Err != nil {
		return shared.Output{Err: productSaveOutput.Err}
	}

	productSaved := productSaveOutput.Result.(*domain.Product)

	return shared.Output{Result: productSaved}
}

// GetAllProduct function
func (u *ProductUsecaseImpl) GetAllProduct(params *shared.Parameters) shared.Output {
	params.Page = 1

	if len(params.StrPage) > 0 {
		page, err := strconv.Atoi(params.StrPage)
		if err != nil {
			return shared.Output{Err: shared.NewErrorAllowNumericOnly("page")}
		}

		params.Page = page
	}

	params.Limit = 10
	if len(params.StrLimit) > 0 {
		limit, err := strconv.Atoi(params.StrLimit)
		if err != nil {
			return shared.Output{Err: shared.NewErrorAllowNumericOnly("limit")}
		}

		params.Limit = limit
	}

	params.Offset = (params.Page - 1) * params.Limit

	if len(params.OrderBy) > 0 {
		if !shared.StringInSlice(params.OrderBy, shared.AllowedSortFields) {
			return shared.Output{Err: fmt.Errorf(shared.ErrorParameterInvalid, "order by")}
		}
	} else {
		params.OrderBy = "created"
	}

	if len(params.Sort) > 0 {
		if !shared.StringInSlice(params.Sort, []string{"asc", "desc"}) {
			return shared.Output{Err: fmt.Errorf("parameter %s allow input asc and desc only", "sort")}

		}
	} else {
		params.Sort = "asc"
	}

	params.OrderBy = fmt.Sprintf(`"%s" %s`, params.OrderBy, params.Sort)

	productResult := u.productRepositoryRead.FindAll(params)
	if productResult.Err != nil {
		return shared.Output{Err: productResult.Err}
	}

	products := productResult.Result.(domain.Products)

	return shared.Output{Result: products}
}

// GetTotalProduct function
func (u *ProductUsecaseImpl) GetTotalProduct(params *shared.Parameters) shared.Output {
	productResult := u.productRepositoryRead.Count(params)
	if productResult.Err != nil {
		return shared.Output{Err: productResult.Err}
	}

	totalProduct := productResult.Result.(int)

	return shared.Output{Result: totalProduct}
}
