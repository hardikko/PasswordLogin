package handlers

import (
	"encoding/json"
	"fmt"
	"learngo/helpers"
	"learngo/models"
	"learngo/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func OrgCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.Organization{}

	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	obj, err := services.OrgCreateService(ctx, req)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, obj)
}

func OrgHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := helpers.StringToInt64(idStr)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	obj, erEr := services.OrgService(ctx, id)
	if erEr != nil {
		render.JSON(w, r, erEr)
	}
	render.JSON(w, r, obj)
}

func OrgListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := services.OrgListService(ctx)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, result)
}

func OrgUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.Organization{}
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		render.JSON(w, r, err.Error())
	}
	defer r.Body.Close()

	obj, erEr := services.OrgUpdateService(ctx, id, req)
	if erEr != nil {
		render.JSON(w, r, erEr)
	}
	render.JSON(w, r, obj)

}
