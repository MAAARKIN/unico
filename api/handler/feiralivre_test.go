package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/domain"
	"github.com/MAAARKIN/unico/service"
	"github.com/stretchr/testify/mock"
)

var feira = domain.FeiraLivre{
	Id:         1,
	Long:       -46550164,
	Lat:        -23558733,
	SetCens:    355030885000091,
	Areap:      3550308005040,
	CodDist:    87,
	Distrito:   "VILA FORMOSA",
	CodSubPref: 26,
	SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
	Regiao5:    "Leste",
	Regiao8:    "Leste 1",
	NomeFeira:  "VILA FORMOSA",
	Registro:   "4041-0",
	Logradouro: "RUA MARAGOJIPE",
	Numero:     nil,
	Bairro:     nil,
	Referencia: nil,
}

var persistFeira = domain.FeiraLivrePersist{
	Long:       -46550164,
	Lat:        -23558733,
	SetCens:    355030885000091,
	Areap:      3550308005040,
	CodDist:    87,
	Distrito:   "VILA FORMOSA",
	CodSubPref: 26,
	SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
	Regiao5:    "Leste",
	Regiao8:    "Leste 1",
	NomeFeira:  "VILA FORMOSA",
	Registro:   "4041-0",
	Logradouro: "RUA MARAGOJIPE",
	Numero:     nil,
	Bairro:     nil,
	Referencia: nil,
}

var updateFeira = domain.FeiraLivreUpdate{
	Long:       -46550164,
	Lat:        -23558733,
	SetCens:    355030885000091,
	Areap:      3550308005040,
	CodDist:    87,
	Distrito:   "VILA FORMOSA",
	CodSubPref: 26,
	SubPrefe:   "ARICANDUVA-FORMOSA-CARRAO",
	Regiao5:    "Leste",
	Regiao8:    "Leste 1",
	NomeFeira:  "VILA FORMOSA",
	Logradouro: "RUA MARAGOJIPE",
	Numero:     nil,
	Bairro:     nil,
	Referencia: nil,
}

func TestGetAllFeiras(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	mocks := []domain.FeiraLivre{feira}

	feiraFilter := domain.FeiraFiltro{}
	serviceMock.On("GetAll", feiraFilter).Return(mocks, nil)

	req, err := http.NewRequest("GET", "/feiras", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.getAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":1,"long":-46550164,"lat":-23558733,"setCens":355030885000091,"areap":3550308005040,"codDist":87,"distrito":"VILA FORMOSA","codSubPref":26,"subPrefe":"ARICANDUVA-FORMOSA-CARRAO","regiao5":"Leste","regiao8":"Leste 1","nomeFeira":"VILA FORMOSA","logradouro":"RUA MARAGOJIPE","numero":null,"bairro":null,"referencia":null,"registro":"4041-0"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetAllFeiraWithoutRecord(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	mocks := []domain.FeiraLivre{}

	feiraFilter := domain.FeiraFiltro{}
	serviceMock.On("GetAll", feiraFilter).Return(mocks, nil)

	req, err := http.NewRequest("GET", "/feiras", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.getAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestGetAllFeiraWithError(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	feiraFilter := domain.FeiraFiltro{}
	serviceMock.On("GetAll", feiraFilter).Return(nil, errors.New("Mock error"))

	req, err := http.NewRequest("GET", "/feiras", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.getAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expected := `{"description":"Internal error, please report to unico Team"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateFeira(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	serviceMock.On("Create", persistFeira.ToDomain()).Return(uint64(1), nil)

	body, _ := json.Marshal(persistFeira)
	req, err := http.NewRequest("POST", "/feiras", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(feiraHandler.create)

	handlerFunc.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestUpdateFeira(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	serviceMock.On("Update", mock.Anything, updateFeira.ToDomain()).Return(&feira, nil)

	body, _ := json.Marshal(updateFeira)
	req, err := http.NewRequest("PUT", "/feiras/1", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.update)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetFeiraById(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	serviceMock.On("Get", mock.Anything).Return(&feira, nil)

	req, err := http.NewRequest("GET", "/feiras/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetWithoutFeira(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	serviceMock.On("Get", mock.Anything).Return(nil, db.ErrRecordNotFound)

	req, err := http.NewRequest("GET", "/feiras/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGetFeiraWithInternalError(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	serviceMock.On("Get", mock.Anything).Return(nil, errors.New("internal error"))

	req, err := http.NewRequest("GET", "/feiras/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.get)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestDeleteFeira(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	serviceMock.On("Delete", mock.Anything).Return(nil)

	req, err := http.NewRequest("DELETE", "/feiras/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.delete)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeleteFeiraWithInternalError(t *testing.T) {
	serviceMock := new(service.MockFeiraService)
	feiraHandler := NewFeiraHandler(serviceMock)

	serviceMock.On("Delete", mock.Anything).Return(errors.New("internal error"))

	req, err := http.NewRequest("DELETE", "/feiras/10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(feiraHandler.delete)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
