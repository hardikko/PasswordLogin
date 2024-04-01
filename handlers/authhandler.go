package handlers

import (
	"encoding/json"
	"fmt"
	"learngo/models"
	"learngo/services"
	"net/http"

	"github.com/go-chi/render"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := models.Login{}

	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		err := fmt.Errorf("invalid json request")
		render.JSON(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	obj, err := services.LoginCreateService(ctx, req)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, obj)
}
