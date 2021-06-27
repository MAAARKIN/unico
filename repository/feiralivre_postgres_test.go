package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MAAARKIN/unico/db"
	"github.com/MAAARKIN/unico/domain"
	"github.com/jmoiron/sqlx"
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
	SubPref:    "ARICANDUVA-FORMOSA-CARRAO",
	Regiao5:    "Leste",
	Regiao8:    "Leste 1",
	NomeFeira:  "VILA FORMOSA",
	Registro:   "4041-0",
	Logradouro: "RUA MARAGOJIPE",
	Numero:     "S/N",
	Bairro:     "VL FORMOSA",
	Referencia: "TV RUA PRETORIA",
}

func TestGetAllFeiras(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	feiraFilter := domain.FeiraFiltro{}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows(
		[]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subpref", "regiao5", "regiao8", "nomefeira", "registro", "logradouro", "numero", "bairro", "referencia"},
	).AddRow(1, -46550164, -23558733, 355030885000091, 3550308005040, 87, "VILA FORMOSA", 26, "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA")

	mock.ExpectQuery("^SELECT (.+) FROM feiralivre").WillReturnRows(rows)

	repository := NewFeiraStorePostgres(sqlxDB)

	content, err := repository.GetAll(feiraFilter)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}

	if len(content) != 1 {
		t.Errorf("we expect 1 record but was: %d", len(content))
	}
}

func TestGetAllFeirasWithError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	feiraFilter := domain.FeiraFiltro{}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectQuery("^SELECT (.+) FROM feiralivre").WillReturnError(errors.New("internal error"))

	repository := NewFeiraStorePostgres(sqlxDB)

	content, err := repository.GetAll(feiraFilter)
	if err == nil {
		t.Errorf("error was expected while query data: %s", err)
	}

	if content != nil {
		t.Errorf("we not expect record but was: %v", content)
	}
}

func TestGetAllFeirasWithoutRecord(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	feiraFilter := domain.FeiraFiltro{}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows(nil)
	mock.ExpectQuery("^SELECT (.+) FROM feiralivre").WillReturnRows(rows)

	repository := NewFeiraStorePostgres(sqlxDB)

	content, err := repository.GetAll(feiraFilter)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}

	if len(content) != 0 {
		t.Errorf("we expect 1 record but was: %d", len(content))
	}
}

func TestGetAllFeirasWithFilter(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	feiraFilter := domain.FeiraFiltro{
		Distrito:  "distrito",
		Regiao5:   "regiao5",
		NomeFeira: "nomefeira",
		Bairro:    "bairro",
	}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	rows := sqlmock.NewRows(
		[]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subpref", "regiao5", "regiao8", "nomefeira", "registro", "logradouro", "numero", "bairro", "referencia"},
	).AddRow(1, -46550164, -23558733, 355030885000091, 3550308005040, 87, "VILA FORMOSA", 26, "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA")

	mock.ExpectQuery("^SELECT (.+) FROM feiralivre WHERE").WillReturnRows(rows)

	repository := NewFeiraStorePostgres(sqlxDB)

	content, err := repository.GetAll(feiraFilter)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}

	if len(content) != 1 {
		t.Errorf("we expect 1 record but was: %d", len(content))
	}
}

func TestDeleteFeiraWithId(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeDeleted := uint64(3)

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM feiralivre where id=$1`)).
		WithArgs(codeToBeDeleted).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewFeiraStorePostgres(sqlxDB)

	err = repository.Delete(codeToBeDeleted)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}
}

func TestDeleteFeiraWithError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeDeleted := uint64(3)

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM feiralivre where id=$1`)).
		WithArgs(codeToBeDeleted).
		WillReturnError(errors.New("internal error"))

	repository := NewFeiraStorePostgres(sqlxDB)

	err = repository.Delete(codeToBeDeleted)
	if err == nil {
		t.Errorf("error was expected while query data")
	}
}

func TestDeleteFeiraByRegistro(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	registroToBeDeleted := "4041-0"

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM feiralivre where registro=$1`)).
		WithArgs(registroToBeDeleted).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewFeiraStorePostgres(sqlxDB)

	err = repository.DeleteByRegistro(registroToBeDeleted)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}
}

func TestDeleteFeiraByRegistroWithError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	registroToBeDeleted := "4041-0"

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM feiralivre where registro=$1`)).
		WithArgs(registroToBeDeleted).
		WillReturnError(errors.New("internal error"))

	repository := NewFeiraStorePostgres(sqlxDB)

	err = repository.DeleteByRegistro(registroToBeDeleted)
	if err == nil {
		t.Errorf("error was expected while query data")
	}
}

func TestUpdateFeira(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeUpdated := uint64(1)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE feiralivre`)).
		WithArgs(feira.Long, feira.Lat, feira.SetCens, feira.Areap, feira.CodDist, feira.Distrito, feira.CodSubPref, feira.SubPref, feira.Regiao5, feira.Regiao8, feira.NomeFeira, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia, codeToBeUpdated).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewFeiraStorePostgres(sqlxDB)

	result, err := repository.Update(codeToBeUpdated, feira)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}

	if result.Long != feira.Long {
		t.Errorf("service returned wrong data: got %v want %v", result.Long, feira.Long)
	}
}

func TestUpdateFeiraWithError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeUpdated := uint64(1)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE feiralivre`)).
		WithArgs(feira.Long, feira.Lat, feira.SetCens, feira.Areap, feira.CodDist, feira.Distrito, feira.CodSubPref, feira.SubPref, feira.Regiao5, feira.Regiao8, feira.NomeFeira, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia, codeToBeUpdated).
		WillReturnError(errors.New("internal error"))

	repository := NewFeiraStorePostgres(sqlxDB)

	_, err = repository.Update(codeToBeUpdated, feira)
	if err == nil {
		t.Errorf("error was expected while query data")
	}
}

func TestUpdateFeiraWithErrorRowsAffected(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeUpdated := uint64(1)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE feiralivre`)).
		WithArgs(feira.Long, feira.Lat, feira.SetCens, feira.Areap, feira.CodDist, feira.Distrito, feira.CodSubPref, feira.SubPref, feira.Regiao5, feira.Regiao8, feira.NomeFeira, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia, codeToBeUpdated).
		WillReturnResult(sqlmock.NewErrorResult(errors.New("internal error")))

	repository := NewFeiraStorePostgres(sqlxDB)

	_, err = repository.Update(codeToBeUpdated, feira)
	if err == nil {
		t.Errorf("error was expected while query data")
	}
}

func TestUpdateFeiraWithoutRowsAffected(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeUpdated := uint64(1)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE feiralivre`)).
		WithArgs(feira.Long, feira.Lat, feira.SetCens, feira.Areap, feira.CodDist, feira.Distrito, feira.CodSubPref, feira.SubPref, feira.Regiao5, feira.Regiao8, feira.NomeFeira, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia, codeToBeUpdated).
		WillReturnResult(sqlmock.NewResult(0, 0))

	repository := NewFeiraStorePostgres(sqlxDB)

	_, err = repository.Update(codeToBeUpdated, feira)
	if err != db.ErrRecordNotFound {
		t.Errorf("error was expected while query data: got %v want %v", err, db.ErrRecordNotFound)
	}
}

func TestCreateFeira(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO feiralivre`)).
		WithArgs(feira.Long, feira.Lat, feira.SetCens, feira.Areap, feira.CodDist, feira.Distrito, feira.CodSubPref, feira.SubPref, feira.Regiao5, feira.Regiao8, feira.NomeFeira, feira.Registro, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repository := NewFeiraStorePostgres(sqlxDB)

	result, err := repository.Create(feira)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}

	if result != 1 {
		t.Errorf("service returned wrong data: got %v want %v", result, 1)
	}
}

func TestCreateFeiraWithError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO feiralivre`)).
		WillReturnError(errors.New("internal error"))

	repository := NewFeiraStorePostgres(sqlxDB)

	result, err := repository.Create(feira)
	if err == nil {
		t.Errorf("error was expected while query data: %s", err)
	}

	if result != 0 {
		t.Errorf("service returned wrong data: got %v want %v", result, 0)
	}
}

func TestGetFeira(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeRecovered := uint64(10)
	rows := sqlmock.NewRows(
		[]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subpref", "regiao5", "regiao8", "nomefeira", "registro", "logradouro", "numero", "bairro", "referencia"},
	).AddRow(10, -46550164, -23558733, 355030885000091, 3550308005040, 87, "VILA FORMOSA", 26, "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM feiralivre WHERE id=$1")).
		WithArgs(codeToBeRecovered).
		WillReturnRows(rows)

	repository := NewFeiraStorePostgres(sqlxDB)

	content, err := repository.Get(codeToBeRecovered)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}

	if content.Id != codeToBeRecovered {
		t.Errorf("we expect 1 record but was: %d", codeToBeRecovered)
	}
}

func TestGetFeiraWithNoRows(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	codeToBeRecovered := uint64(10)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM feiralivre WHERE id=$1")).
		WithArgs(codeToBeRecovered).
		WillReturnError(sql.ErrNoRows)

	repository := NewFeiraStorePostgres(sqlxDB)

	_, err = repository.Get(codeToBeRecovered)
	if !errors.Is(err, db.ErrRecordNotFound) {
		t.Errorf("error was expected while query data: %s", db.ErrRecordNotFound)
	}
}

func TestGetFeiraWithInternalError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	var internalError = errors.New("internal error")

	codeToBeRecovered := uint64(10)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM feiralivre WHERE id=$1")).
		WithArgs(codeToBeRecovered).
		WillReturnError(internalError)

	repository := NewFeiraStorePostgres(sqlxDB)

	_, err = repository.Get(codeToBeRecovered)
	if !errors.Is(err, internalError) {
		t.Errorf("error was expected while query data: got %v want %v", err, internalError)
	}
}

func TestGetFeiraByRegistro(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	registroToBeRecovered := "4041-0"
	rows := sqlmock.NewRows(
		[]string{"id", "long", "lat", "setcens", "areap", "coddist", "distrito", "codsubpref", "subpref", "regiao5", "regiao8", "nomefeira", "registro", "logradouro", "numero", "bairro", "referencia"},
	).AddRow(10, -46550164, -23558733, 355030885000091, 3550308005040, 87, "VILA FORMOSA", 26, "ARICANDUVA-FORMOSA-CARRAO", "Leste", "Leste 1", "VILA FORMOSA", "4041-0", "RUA MARAGOJIPE", "S/N", "VL FORMOSA", "TV RUA PRETORIA")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM feiralivre WHERE registro=$1")).
		WithArgs(registroToBeRecovered).
		WillReturnRows(rows)

	repository := NewFeiraStorePostgres(sqlxDB)

	content, err := repository.GetByRegistro(registroToBeRecovered)
	if err != nil {
		t.Errorf("error was not expected while query data: %s", err)
	}

	if content.Registro != registroToBeRecovered {
		t.Errorf("we expect 1 record but was: %s", registroToBeRecovered)
	}
}

func TestGetFeiraByRegistroWithNoRows(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	registroToBeRecovered := "4041-0"
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM feiralivre WHERE registro=$1")).
		WithArgs(registroToBeRecovered).
		WillReturnError(sql.ErrNoRows)

	repository := NewFeiraStorePostgres(sqlxDB)

	_, err = repository.GetByRegistro(registroToBeRecovered)
	if !errors.Is(err, db.ErrRecordNotFound) {
		t.Errorf("error was expected while query data: %s", db.ErrRecordNotFound)
	}
}

func TestGetFeiraByRegistroWithInternalError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	internalError := errors.New("internal error")
	registroToBeRecovered := "4041-0"
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM feiralivre WHERE registro=$1")).
		WithArgs(registroToBeRecovered).
		WillReturnError(internalError)

	repository := NewFeiraStorePostgres(sqlxDB)

	_, err = repository.GetByRegistro(registroToBeRecovered)
	if !errors.Is(err, internalError) {
		t.Errorf("error was expected while query data: %s", internalError)
	}
}
