-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE FEIRALIVRE(
   ID         SERIAL NOT NULL PRIMARY KEY,
   LONG       INTEGER  NOT NULL,
   LAT        INTEGER  NOT NULL,
   SETCENS    INTEGER  NOT NULL,
   AREAP      INTEGER  NOT NULL,
   CODDIST    INTEGER  NOT NULL,
   DISTRITO   VARCHAR(18) NOT NULL,
   CODSUBPREF INTEGER  NOT NULL,
   SUBPREFE   VARCHAR(25) NOT NULL,
   REGIAO5    VARCHAR(6) NOT NULL,
   REGIAO8    VARCHAR(7) NOT NULL,
   NOMEFEIRA VARCHAR(30) NOT NULL,
   REGISTRO   VARCHAR(6) NOT NULL UNIQUE,
   LOGRADOURO VARCHAR(34) NOT NULL,
   NUMERO     VARCHAR(11),
   BAIRRO     VARCHAR(20),
   REFERENCIA VARCHAR(30)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE FEIRALIVRE;