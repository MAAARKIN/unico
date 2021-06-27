package repository

import "github.com/MAAARKIN/unico/domain"

type FeiraStore interface {
	GetAll(op domain.FeiraFiltro) ([]domain.FeiraLivre, error)
	Get(id uint64) (*domain.FeiraLivre, error)
	GetByRegistro(registro string) (*domain.FeiraLivre, error)
	Create(item domain.FeiraLivre) (uint64, error)
	Update(id uint64, item domain.FeiraLivre) (*domain.FeiraLivre, error)
	Delete(id uint64) error
	DeleteByRegistro(registro string) error
}
