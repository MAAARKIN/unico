package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MAAARKIN/unico/api/helper"
	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/domain"
	"github.com/MAAARKIN/unico/service"
	"github.com/go-chi/chi/v5"
)

type FeiraHandler struct {
	service service.FeiraService
}

func NewFeiraHandler(service service.FeiraService) FeiraHandler {
	return FeiraHandler{service}
}

func (f FeiraHandler) Route(r chi.Router) {
	r.Get("/", f.getAll)
	r.Post("/", f.create)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", f.get)
		r.Put("/", f.update)
		r.Delete("/", f.delete)
	})
	r.Route("/registro/{registro}", func(r chi.Router) {
		r.Delete("/", f.deleteByRegistro)
	})
}

func getFilterQuery(r *http.Request) domain.FeiraFiltro {
	return domain.FeiraFiltro{
		Distrito:  r.URL.Query().Get("distrito"),
		Regiao5:   r.URL.Query().Get("regiao5"),
		NomeFeira: r.URL.Query().Get("nomefeira"),
		Bairro:    r.URL.Query().Get("bairro"),
	}
}

// @Title getAllFeiras
// @Tags Feiras
// @Summary List feiras
// @Description Displays all feiras registered on the API.
// @Param distrito query string false "Name of the distrito"
// @Param regiao5 query string false "Name of regiao5"
// @Param nomefeira query string false "Name of feira"
// @Param bairro query string false  "Name of bairro"
// @Success 200 {object} []domain.FeiraLivre
// @Success 204 "Feiras not found"
// @Accept json
// @Router /feiras [get]
func (f FeiraHandler) getAll(w http.ResponseWriter, r *http.Request) {
	feiraFilter := getFilterQuery(r)
	if results, err := f.service.GetAll(feiraFilter); err != nil {
		helper.HandleError(w, err)
	} else {
		if results != nil && len(results) > 0 {
			helper.JsonResponse(w, results, http.StatusOK)
		} else {
			helper.NoContent(w)
		}
	}
}

// @Title createFeira
// @Tags Feiras
// @Summary Create a feira
// @Description Creates a new feira
// @Param content body domain.FeiraLivrePersist true "Object for persisting the feira"
// @Success 201 {object} domain.FeiraLivre
// @Failure 400 "Bad request"
// @Accept json
// @Router /feiras [post]
func (f FeiraHandler) create(w http.ResponseWriter, r *http.Request) {
	dto := domain.FeiraLivrePersist{}
	if err := helper.BindJson(r, &dto); err != nil {
		helper.HandleError(w, err)
		return
	}

	entity := dto.ToDomain()

	if _, err := f.service.Create(entity); err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			helper.CustomError(w, err, http.StatusBadRequest)
		} else {
			helper.HandleError(w, err)
		}
	} else {
		helper.Response(w, http.StatusCreated)
	}
}

// @Title updateFeira
// @Tags Feiras
// @Summary Update a feira
// @Description Updates a feira
// @Param id path string true "The identifier for the feira"
// @Param content body domain.FeiraLivreUpdate true "Object for updating a feira"
// @Success 200 "Ok"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Accept json
// @Router /feiras/{id} [put]
func (f FeiraHandler) update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	teste := chi.URLParam(r, "id")
	fmt.Println(teste)
	dto := domain.FeiraLivreUpdate{}
	if err := helper.BindJson(r, &dto); err != nil {
		helper.HandleError(w, err)
		return
	}

	entity := dto.ToDomain()
	fmt.Println(id)
	if _, err := f.service.Update(id, entity); err != nil {
		helper.HandleError(w, err)
	} else {
		helper.Response(w, http.StatusOK)
	}
}

// @Title deleteFeiraById
// @Tags Feiras
// @Summary Delete a feira
// @Description Delete a feira by id
// @Param id path string true "The identifier for the feira"
// @Success 200 "Deleted"
// @Failure 204 "Feira not found"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Accept json
// @Router /feiras/{id} [delete]
func (f FeiraHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	if err := f.service.Delete(id); err != nil {
		helper.HandleError(w, err)
	} else {
		helper.Response(w, http.StatusOK)
	}
}

// @Title getFeira
// @Tags Feiras
// @Summary Get a feira
// @Description Get a feira
// @Param id path string true "The identifier for the feira"
// @Success 200 {object} domain.FeiraLivre
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Accept json
// @Router /feiras/{id} [get]
func (f FeiraHandler) get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	if location, err := f.service.Get(id); err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			helper.CustomError(w, err, http.StatusBadRequest)
		} else {
			helper.HandleError(w, err)
		}
	} else {
		helper.JsonResponse(w, location, http.StatusOK)
	}
}

// @Title deleteFeiraByRegistro
// @Tags Feiras
// @Summary Delete a feira
// @Description Delete a feira by id
// @Param id path string true "The identifier for the feira"
// @Success 200 "Deleted"
// @Failure 204 "Feira not found"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Accept json
// @Router /feiras/registro/{registro} [delete]
func (f FeiraHandler) deleteByRegistro(w http.ResponseWriter, r *http.Request) {
	registro := chi.URLParam(r, "registro")

	if err := f.service.DeleteByRegistro(registro); err != nil {
		helper.HandleError(w, err)
	} else {
		helper.Response(w, http.StatusOK)
	}
}
