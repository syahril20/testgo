package handlers

import (
	"context"
	"net/http"
	dtoProduct "server/dto/product"
	dto "server/dto/result"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProductHandler(c *gin.Context) {
	productList, err := repositories.GetAllProduct(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    productList})
}

func CreateProductHandler(c *gin.Context) {
	var req dtoProduct.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.Name == "" || req.Content == "" || req.Image == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	productName, _ := repositories.GetProductByName(context.Background(), req.Name)
	if productName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now()

	Product := dtoProduct.CreateProductRequest{
		Id:         primitive.NewObjectID(),
		Name:       req.Name,
		Content:    req.Content,
		Image:      req.Image,
		SubProduct: []dtoProduct.CreateSubProductRequest{},
		CreatedAt:  currentTime,
		CreatedBy:  "System",
		UpdatedAt:  currentTime,
		UpdatedBy:  "System",
	}

	// Menambahkan Suburb ke database
	data, err := repositories.CreateProduct(context.Background(), Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    data})
}

func CreateSubProductHandler(c *gin.Context) {
	var req dtoProduct.CreateSubProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.IdProduct.Hex() == "" || req.Name == "" || req.Content == "" || req.Image == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(req.IdProduct.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Product ID"})
		return
	}

	product, _ := repositories.GetProductByID(context.Background(), objectId)
	if product == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Product Is Not Found"})
		return
	}

	subProducts, _ := repositories.GetSubProductByNameId(context.Background(), objectId, req.Name)
	if subProducts != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now()
	subProduct := dtoProduct.CreateSubProductRequest{
		Id:        primitive.NewObjectID(),
		IdProduct: objectId,
		Name:      req.Name,
		Content:   req.Content,
		Image:     "image/jpg",
		Addons:    []dtoProduct.CreateAddonsRequest{},
		CreatedAt: currentTime,
		CreatedBy: "System",
		UpdatedAt: currentTime,
		UpdatedBy: "System",
	}

	subProduct.Price = func() int32 {
		if req.Price != 0 {
			return int32(req.Price)
		}
		return int32(0)
	}()

	// Menambahkan Suburb ke database
	_, err = repositories.CreateSubProduct(context.Background(), objectId, subProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    subProduct})
}

func CreateAddonsHandler(c *gin.Context) {
	var req dtoProduct.CreateAddonsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}

	if req.IdSubProduct.Hex() == "" || req.Name == "" || req.Content == "" || req.Image == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Bad Request"})
		return
	}

	subProductID, err := primitive.ObjectIDFromHex(req.IdSubProduct.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Invalid Product ID"})
		return
	}

	subProducts, _ := repositories.GetSubProductById(context.Background(), subProductID)
	if subProducts == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Sub Product Not Found"})
		return
	}

	addonsName, _ := repositories.GetAddonsByName(context.Background(), req.Name)
	if addonsName != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Data Exist"})
		return
	}

	currentTime := time.Now()
	addons := dtoProduct.CreateAddonsRequest{
		Id:           primitive.NewObjectID(),
		IdSubProduct: subProductID,
		Name:         req.Name,
		Content:      req.Content,
		Image:        "image/jpg",
		CreatedAt:    currentTime,
		CreatedBy:    "System",
		UpdatedAt:    currentTime,
		UpdatedBy:    "System",
	}

	addons.Price = func() int32 {
		if req.Price != 0 {
			return int32(req.Price)
		}
		return int32(0)
	}()

	// Menambahkan Suburb ke database
	err = repositories.CreateAddons(context.Background(), subProductID, addons)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SuccessResult{
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    addons})
}

// func convertResponseProduct(product dtoProduct.CreateProductRequest) dtoProduct.ProductResponse {
// 	return dtoProduct.ProductResponse{
// 		Name:      product.Name,
// 		Content:   product.Content,
// 		DeletedAt: *product.DeletedAt,
// 		Image:     product.Image,
// 	}
// }
