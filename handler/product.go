package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"log"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
	"strconv"
)

type ProductHandler struct {
	ProductService *service.ProductService
	Logger         *zap.Logger
}

func InitProductHandler(service service.Service, log *zap.Logger) ProductHandler {
	return ProductHandler{ProductService: service.Product, Logger: log}
}

func (handler ProductHandler) All(w http.ResponseWriter, r *http.Request) {
	const ItemPerPage = 6

	page := 1
	var err error
	if p := r.URL.Query().Get("page"); p != "" {
		page, err = strconv.Atoi(p)
	}

	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "Invalid Page Number")
		return
	}

	itemPerPage := ItemPerPage
	if i := r.URL.Query().Get("i"); i != "" {
		itemPerPage, err = strconv.Atoi(i)
	}

	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "Invalid Item Per Page")
		return
	}

	productName := r.URL.Query().Get("q")
	categoryId := 0
	if category := chi.URLParam(r, "category_id"); category != "" {
		categoryId, err = strconv.Atoi(category)
	}

	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "Invalid Category ID")
	}

	totalItems, totalPages, items, err := handler.ProductService.Index(page, itemPerPage, productName, categoryId, r.Header.Get("token"))
	if err != nil {
		util.Response(w).Json(http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(domain.DataPage{
		StatusCode: http.StatusOK,
		Page:       page,
		Limit:      itemPerPage,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       items,
	})
}

func (handler ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		util.Response(w).Json(http.StatusBadRequest, "Invalid Product ID")
		return
	}

	product := domain.ProductDetail{Id: id}
	err = handler.ProductService.Get(&product, r.Header.Get("token"))
	if err != nil {
		log.Println(err)
		util.Response(w).Json(http.StatusBadRequest, "Failed to retrieve Product")
		return
	}
	util.Response(w).Json(http.StatusOK, "Product Found", product)
}
