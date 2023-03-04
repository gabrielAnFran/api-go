package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabrielAnFran/api-go/internal/dto"
	"github.com/gabrielAnFran/api-go/internal/entity"
	"github.com/gabrielAnFran/api-go/internal/infra/database"
	entitypkg "github.com/gabrielAnFran/api-go/pkg/entity"
	"github.com/go-chi/chi"
)

type ShoesHandler struct {
	ProductDB database.ShoesInterface
}

func NewProductHandler(db database.ShoesInterface) *ShoesHandler {
	return &ShoesHandler{
		ProductDB: db,
	}
}

// Create Shoes godoc
// @Summary      Create shoes
// @Description  Create shoes
// @Tags         shoes
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateShoesInput  true  "shoes request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /shoes [post]
// @Security ApiKeyAuth
func (h *ShoesHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var shoes dto.CreateShoesInput

	// Read from r.body and stores it in shoes
	err := json.NewDecoder(r.Body).Decode(&shoes)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(msg)
	}

	// Calls entity and instantiate a shoes object
	s, err := entity.NewShoes(shoes.Name, shoes.Brand, shoes.Price, shoes.Size)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(msg)
	}

	//
	err = h.ProductDB.Create(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

// GetShoes godoc
// @Summary      Get shoes
// @Description  Get shoes
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "shoes ID" Format(uuid)
// @Success      200  {object}  entity.Shoes
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /shoes/{id} [get]
// @Security ApiKeyAuth
func (h *ShoesHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Access the query params in the URL
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	shoes, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shoes)

}

// UpdateShoes godoc
// @Summary      Update shoes
// @Description  Update shoes
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "shoes ID" Format(uuid)
// @Param        request     body      dto.CreateShoesInput  true  "shoes request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /shoes/{id} [put]
// @Security ApiKeyAuth
func (h *ShoesHandler) UpdateShoes(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var shoes entity.Shoes

	err := json.NewDecoder(r.Body).Decode(&shoes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shoes.ID, err = entitypkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(&shoes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
