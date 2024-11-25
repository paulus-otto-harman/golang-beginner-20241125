package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"project/domain"
	"project/service"
	"project/util"
	"strconv"
)

type BestsellerHandler struct {
	BestsellerService *service.BestsellerService
	Logger            *zap.Logger
}

func InitBestsellerHandler(service service.Service, log *zap.Logger) BestsellerHandler {
	return BestsellerHandler{BestsellerService: service.Bestseller, Logger: log}
}

func (handler BestsellerHandler) All(w http.ResponseWriter, r *http.Request) {
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
		util.Response(w).Json(http.StatusBadRequest, "Invalid Number of items per page")
		return
	}

	totalItems, totalPages, items, err := handler.BestsellerService.Index(page, itemPerPage)
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
