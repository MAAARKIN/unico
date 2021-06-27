package service

import (
	"errors"

	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/domain"
	"github.com/MAAARKIN/unico/repository"
)

type FeiraService interface {
	GetAll(op domain.FeiraFiltro) ([]domain.FeiraLivre, error)
	Get(id uint64) (*domain.FeiraLivre, error)
	GetByRegistro(registro string) (*domain.FeiraLivre, error)
	Create(item domain.FeiraLivre) (uint64, error)
	Update(id uint64, item domain.FeiraLivre) (*domain.FeiraLivre, error)
	Delete(id uint64) error
	DeleteByRegistro(registro string) error
}

type feiraServiceImpl struct {
	store repository.FeiraStore
}

func NewFeiraService(store repository.FeiraStore) FeiraService {
	return &feiraServiceImpl{
		store: store,
	}
}

func (f *feiraServiceImpl) GetAll(op domain.FeiraFiltro) ([]domain.FeiraLivre, error) {
	return f.store.GetAll(op)
}

func (f *feiraServiceImpl) Get(id uint64) (*domain.FeiraLivre, error) {
	return f.store.Get(id)
}

func (f *feiraServiceImpl) GetByRegistro(registro string) (*domain.FeiraLivre, error) {
	return f.store.GetByRegistro(registro)
}

func (f *feiraServiceImpl) Create(item domain.FeiraLivre) (uint64, error) {
	_, err := f.GetByRegistro(item.Registro)

	if errors.Is(err, db.ErrRecordNotFound) {
		return f.store.Create(item)
	}

	if err != nil {
		return 0, err
	}

	return 0, domain.ErrFeiraWithRegistroAlreadyExist
}

func (f *feiraServiceImpl) Update(id uint64, item domain.FeiraLivre) (*domain.FeiraLivre, error) {
	return f.store.Update(id, item)
}

func (f *feiraServiceImpl) Delete(id uint64) error {
	return f.store.Delete(id)
}

func (f *feiraServiceImpl) DeleteByRegistro(registro string) error {
	return f.store.DeleteByRegistro(registro)
}
