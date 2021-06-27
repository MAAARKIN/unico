package service

import (
	"errors"
	"testing"

	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/domain"
	"github.com/MAAARKIN/unico/repository"
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

func TestGetAllFeiras(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	feirasMocked := []domain.FeiraLivre{feira}
	feiraFilter := domain.FeiraFiltro{
		Distrito:  "distrito",
		Regiao5:   "regiao5",
		NomeFeira: "nomefeira",
		Bairro:    "bairro",
	}
	storeMock.On("GetAll", feiraFilter).Return(feirasMocked, nil)

	feiras, err := service.GetAll(feiraFilter)

	if err != nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}

	if len(feiras) != 1 {
		t.Errorf("service returned wrong data: got %v want %v", len(feiras), 1)
	}
}

func TestGetFeira(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	storeMock.On("Get", uint64(1)).Return(&feira, nil)

	res, err := service.Get(uint64(1))

	if err != nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}

	if res.Registro != feira.Registro && res.Id != feira.Id {
		t.Errorf("service returned wrong data: got %v want %v", *res, feira)
	}
}

func TestCreateFeira(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	codeToRecover := uint64(5)

	storeMock.On("GetByRegistro", feira.Registro).Return(nil, db.ErrRecordNotFound)
	storeMock.On("Create", feira).Return(codeToRecover, nil)

	res, err := service.Create(feira)

	if err != nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}

	if res != codeToRecover {
		t.Errorf("service returned wrong data: got %v want %v", res, codeToRecover)
	}
}

func TestCreateFeiraWithInternalError(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	storeMock.On("GetByRegistro", feira.Registro).Return(nil, errors.New("internal error"))

	res, err := service.Create(feira)

	if err == nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}

	if res != 0 {
		t.Errorf("service returned wrong data: got %v want %v", res, 0)
	}
}

func TestCreateFeiraWithDuplicatedRegistro(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	storeMock.On("GetByRegistro", feira.Registro).Return(&feira, nil)

	res, err := service.Create(feira)

	if err != domain.ErrFeiraWithRegistroAlreadyExist {
		t.Errorf("service returned wrong error: got %v want %v", err, domain.ErrFeiraWithRegistroAlreadyExist)
	}

	if res != 0 {
		t.Errorf("service returned wrong data: got %v want %v", res, 0)
	}
}

func TestUpdateFeiraWithAnotherRegistro(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	var feiraToUpdate = domain.FeiraLivre{
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
		Registro:   "4041-1",
		Logradouro: "RUA MARAGOJIPE",
		Numero:     nil,
		Bairro:     nil,
		Referencia: nil,
	}

	storeMock.On("Update", uint64(1), feiraToUpdate).Return(&feiraToUpdate, nil)

	res, err := service.Update(uint64(1), feiraToUpdate)

	if err != nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}

	if res.Registro != feiraToUpdate.Registro {
		t.Errorf("service returned wrong data: got %v want %v", res.Registro, feiraToUpdate.Registro)
	}
}

func TestUpdateFeira(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	var feiraToUpdate = domain.FeiraLivre{
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
		Registro:   "4041-1",
		Logradouro: "RUA MARAGOJIPE",
		Numero:     nil,
		Bairro:     nil,
		Referencia: nil,
	}

	storeMock.On("GetByRegistro", feiraToUpdate.Registro).Return(&feira, nil)
	storeMock.On("Update", uint64(1), feiraToUpdate).Return(&feiraToUpdate, nil)

	res, err := service.Update(uint64(1), feiraToUpdate)

	if err != nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}

	if res.Registro != feiraToUpdate.Registro {
		t.Errorf("service returned wrong data: got %v want %v", res.Registro, feiraToUpdate.Registro)
	}
}

func TestDeleteFeira(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	codeToDelete := uint64(5)
	storeMock.On("Delete", codeToDelete).Return(nil)

	err := service.Delete(codeToDelete)

	if err != nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}
}

func TestDeleteFeiraByRegistro(t *testing.T) {
	storeMock := new(repository.MockFeiraStore)
	service := NewFeiraService(storeMock)

	registroToBeDelete := "4041-0"
	storeMock.On("DeleteByRegistro", registroToBeDelete).Return(nil)

	err := service.DeleteByRegistro(registroToBeDelete)

	if err != nil {
		t.Errorf("service returned wrong error: got %v want %v", err, nil)
	}
}
