package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"learngo/helpers"
	"learngo/models"
	"learngo/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func LeadCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.Lead{}

	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	obj, err := services.LeadCreateService(ctx, req)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, obj)
}

func LeadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := helpers.StringToInt64(idStr)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	obj, erEr := services.LeadService(ctx, id)
	if erEr != nil {
		render.JSON(w, r, erEr)
	}
	render.JSON(w, r, obj)
}

func LeadListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := services.LeadListService(ctx)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, result)
}

func LeadUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.Lead{}
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.JSON(w, r, err.Error())
	}
	defer r.Body.Close()

	obj, erEr := services.LeadUpdateService(ctx, id, req)
	if erEr != nil {
		render.JSON(w, r, erEr)
	}
	render.JSON(w, r, obj)

}
