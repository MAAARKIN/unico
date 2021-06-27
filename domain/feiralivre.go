package domain

import (
	"errors"

	mapper "gopkg.in/jeevatkm/go-model.v1"
)

var ErrFeiraWithRegistroAlreadyExist = errors.New("Feiralivre com esse registro já existe")

type FeiraLivre struct {
	Id         uint64  `json:"id"`
	Long       int64   `json:"long"`
	Lat        int64   `json:"lat"`
	SetCens    int64   `json:"setCens"`
	Areap      int64   `json:"areap"`
	CodDist    uint64  `json:"codDist"`
	Distrito   string  `json:"distrito"`
	CodSubPref uint64  `json:"codSubPref"`
	SubPrefe   string  `json:"subPrefe"`
	Regiao5    string  `json:"regiao5"`
	Regiao8    string  `json:"regiao8"`
	NomeFeira  string  `json:"nomeFeira"`
	Logradouro string  `json:"logradouro"`
	Numero     *string `json:"numero"`
	Bairro     *string `json:"bairro"`
	Referencia *string `json:"referencia"`
	Registro   string  `json:"registro"`
}

type FeiraFiltro struct {
	Distrito  string `json:"distrito" model:"distrito,omitempty"`
	Regiao5   string `json:"regiao5" model:"regiao5,omitempty"`
	NomeFeira string `json:"nomeFeira" model:"nomefeira,omitempty"`
	Bairro    string `json:"bairro" model:"bairro,omitempty"`
}

func (f *FeiraFiltro) ToMap() map[string]interface{} {
	result, err := mapper.Map(f)
	if err != nil {
		return map[string]interface{}{}
	}
	return result
}

type FeiraLivrePersist struct {
	Long       int64   `json:"long"`
	Lat        int64   `json:"lat"`
	SetCens    int64   `json:"setCens"`
	Areap      int64   `json:"areap"`
	CodDist    uint64  `json:"codDist"`
	Distrito   string  `json:"distrito"`
	CodSubPref uint64  `json:"codSubPref"`
	SubPrefe   string  `json:"subPrefe"`
	Regiao5    string  `json:"regiao5"`
	Regiao8    string  `json:"regiao8"`
	NomeFeira  string  `json:"nomeFeira"`
	Registro   string  `json:"registro"`
	Logradouro string  `json:"logradouro"`
	Numero     *string `json:"numero"`
	Bairro     *string `json:"bairro"`
	Referencia *string `json:"referencia"`
}

func (f *FeiraLivrePersist) ToDomain() FeiraLivre {
	return FeiraLivre{
		Long:       f.Long,
		Lat:        f.Lat,
		SetCens:    f.SetCens,
		Areap:      f.Areap,
		CodDist:    f.CodDist,
		Distrito:   f.Distrito,
		CodSubPref: f.CodSubPref,
		SubPrefe:   f.SubPrefe,
		Regiao5:    f.Regiao5,
		Regiao8:    f.Regiao8,
		NomeFeira:  f.NomeFeira,
		Registro:   f.Registro,
		Logradouro: f.Logradouro,
		Numero:     f.Numero,
		Bairro:     f.Bairro,
		Referencia: f.Referencia,
	}
}

type FeiraLivreUpdate struct {
	Long       int64   `json:"long"`
	Lat        int64   `json:"lat"`
	SetCens    int64   `json:"setCens"`
	Areap      int64   `json:"areap"`
	CodDist    uint64  `json:"codDist"`
	Distrito   string  `json:"distrito"`
	CodSubPref uint64  `json:"codSubPref"`
	SubPrefe   string  `json:"subPrefe"`
	Regiao5    string  `json:"regiao5"`
	Regiao8    string  `json:"regiao8"`
	NomeFeira  string  `json:"nomeFeira"`
	Logradouro string  `json:"logradouro"`
	Numero     *string `json:"numero"`
	Bairro     *string `json:"bairro"`
	Referencia *string `json:"referencia"`
}

func (f *FeiraLivreUpdate) ToDomain() FeiraLivre {
	return FeiraLivre{
		Long:       f.Long,
		Lat:        f.Lat,
		SetCens:    f.SetCens,
		Areap:      f.Areap,
		CodDist:    f.CodDist,
		Distrito:   f.Distrito,
		CodSubPref: f.CodSubPref,
		SubPrefe:   f.SubPrefe,
		Regiao5:    f.Regiao5,
		Regiao8:    f.Regiao8,
		NomeFeira:  f.NomeFeira,
		Logradouro: f.Logradouro,
		Numero:     f.Numero,
		Bairro:     f.Bairro,
		Referencia: f.Referencia,
	}
}
