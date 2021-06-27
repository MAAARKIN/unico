package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type feiraStorePostgres struct {
	db *sqlx.DB
}

func NewFeiraStorePostgres(db *sqlx.DB) FeiraStore {
	return &feiraStorePostgres{db}
}

func addWhereFromOptions(query string, op map[string]interface{}) (string, []interface{}) {
	var values []interface{}
	var where []string
	for k, v := range op {
		if v != nil && v != "" {
			values = append(values, v)
			where = append(where, fmt.Sprintf(`%s=%s`, k, "$"+strconv.Itoa(len(values))))
		}
	}

	query += " WHERE " + strings.Join(where, " AND ")
	return query, values
}

func (f *feiraStorePostgres) GetAll(op domain.FeiraFiltro) ([]domain.FeiraLivre, error) {

	query := "SELECT * FROM feiralivre"

	filterMap := op.ToMap()
	filterValues := []interface{}{}
	hasOptions := len(filterMap) > 0
	if hasOptions {
		query, filterValues = addWhereFromOptions(query, filterMap)
	}

	var result []domain.FeiraLivre
	var err error

	if hasOptions {
		err = f.db.Select(&result, query, filterValues...)
	} else {
		err = f.db.Select(&result, query)
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result == nil {
		return []domain.FeiraLivre{}, nil
	}
	return result, nil
}

func (f *feiraStorePostgres) Get(id uint64) (*domain.FeiraLivre, error) {
	var result domain.FeiraLivre
	if err := f.db.Get(&result, "SELECT * FROM feiralivre WHERE id=$1", id); err != nil {
		if err == sql.ErrNoRows {
			return nil, db.ErrRecordNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &result, nil
}

func (f *feiraStorePostgres) GetByRegistro(registro string) (*domain.FeiraLivre, error) {
	var result domain.FeiraLivre
	if err := f.db.Get(&result, "SELECT * FROM feiralivre WHERE registro=$1", registro); err != nil {
		if err == sql.ErrNoRows {
			return nil, db.ErrRecordNotFound
		}
		return nil, errors.WithStack(err)
	}
	return &result, nil
}

func (f *feiraStorePostgres) Create(item domain.FeiraLivre) (uint64, error) {
	var id uint64
	query := `
		INSERT INTO feiralivre(LONG,LAT,SETCENS,AREAP,CODDIST,DISTRITO,CODSUBPREF,SUBPREFE,REGIAO5,REGIAO8,NOMEFEIRA,REGISTRO,LOGRADOURO,NUMERO,BAIRRO,REFERENCIA) 
		VALUES 
			($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16) 
		RETURNING id
	`
	err := f.db.QueryRow(
		query,
		item.Long,
		item.Lat,
		item.SetCens,
		item.Areap,
		item.CodDist,
		item.Distrito,
		item.CodSubPref,
		item.SubPrefe,
		item.Regiao5,
		item.Regiao8,
		item.NomeFeira,
		item.Registro,
		item.Logradouro,
		item.Numero,
		item.Bairro,
		item.Referencia,
	).Scan(&id)

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return id, nil
}

func (f *feiraStorePostgres) Update(id uint64, item domain.FeiraLivre) (*domain.FeiraLivre, error) {
	item.Id = id
	query := `
		UPDATE feiralivre 
		SET 
			LONG=$1,
			LAT=$2,
			SETCENS=$3,
			AREAP=$4,
			CODDIST=$5,
			DISTRITO=$6,
			CODSUBPREF=$7,
			SUBPREFE=$8,
			REGIAO5=$9,
			REGIAO8=$10,
			NOMEFEIRA=$11,
			LOGRADOURO=$12,
			NUMERO=$13,
			BAIRRO=$14,
			REFERENCIA=$15
		WHERE id=$16
	`
	res, err := f.db.Exec(
		query,
		item.Long,
		item.Lat,
		item.SetCens,
		item.Areap,
		item.CodDist,
		item.Distrito,
		item.CodSubPref,
		item.SubPrefe,
		item.Regiao5,
		item.Regiao8,
		item.NomeFeira,
		item.Logradouro,
		item.Numero,
		item.Bairro,
		item.Referencia,
		id,
	)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if affect == 0 {
		return nil, db.ErrRecordNotFound
	}

	return &item, nil
}

func (f *feiraStorePostgres) Delete(id uint64) error {
	_, err := f.db.Exec("DELETE FROM feiralivre where id=$1", id)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (f *feiraStorePostgres) DeleteByRegistro(registro string) error {
	_, err := f.db.Exec("DELETE FROM feiralivre where registro=$1", registro)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
