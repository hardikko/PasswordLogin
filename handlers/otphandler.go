package handlers

import (
	"encoding/json"
	"fmt"
	"learngo/helpers"
	"learngo/models"
	"learngo/services"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func OtpCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.Otp{}

	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	obj, erEr := services.OtpCreateService(ctx, req)
	if erEr != nil {
		render.JSON(w, r, erEr)
	}
	render.JSON(w, r, obj)
}

func OtpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := helpers.StringToInt64(idStr)
	if err != nil {
		render.JSON(w, r, err.Error())
	}
	obj, erEr := services.OtpService(ctx, id)
	if erEr != nil {
		render.JSON(w, r, erEr)
	}
	render.JSON(w, r, obj)
}
