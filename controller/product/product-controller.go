package product

import (
	"strconv"
	"github.com/NetSinx/go-restful-API/model/domain"
	"github.com/NetSinx/go-restful-API/service/product"
	"github.com/NetSinx/go-restful-API/utils"
	"github.com/gin-gonic/gin"
)

type productcontroller struct {
	ProductService product.ProductService
}

func NewProductController(productService product.ProductService) *productcontroller {
	return &productcontroller{
        ProductService: productService,
    }
}

func (controller *productcontroller) Create(ctx *gin.Context) {
	var product domain.Product

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		utils.BadRequest(ctx)

        return
	}

	createProduct, _ := controller.ProductService.Create(product)

	utils.WriteResponse(ctx, createProduct)
}

func (controller *productcontroller) Update(ctx *gin.Context) {
	var updateProduct domain.Product
	
	getId := ctx.Param("productId")
	id, _ := strconv.Atoi(getId)

	err := ctx.ShouldBindJSON(&updateProduct)
	if err != nil {
		utils.BadRequest(ctx)

		return
	}

	responseProduct, err := controller.ProductService.Update(updateProduct, id)
	if err != nil {
		utils.NotFound(ctx)

		return
	}

	utils.WriteResponse(ctx, responseProduct)
}

func (controller *productcontroller) Delete(ctx *gin.Context) {
	var delProduct domain.Product

	getId := ctx.Param("productId")
	id, _ := strconv.Atoi(getId)

	err := controller.ProductService.Delete(delProduct, id)
	if err != nil {
		utils.NotFound(ctx)

		return
	}

	utils.ResponseDelete(ctx)
}

func (controller *productcontroller) FindById(ctx *gin.Context) {
	var product domain.Product

	getId := ctx.Param("productId")
	id, _ := strconv.Atoi(getId)

	findIdProduct, err := controller.ProductService.FindById(product, id)
	if err != nil {
		utils.NotFound(ctx)

		return
	}

	utils.WriteResponse(ctx, findIdProduct)
}

func (controller *productcontroller) FindAll(ctx *gin.Context) {
	var product []domain.Product

	findAllProduct, err := controller.ProductService.FindAll(product)
	if err != nil {
		utils.NotFound(ctx)

		return
	}

	utils.ResponseFindAll(ctx, findAllProduct)
}