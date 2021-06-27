# Getting Started

To setup the project you will need

- Docker & Docker-compose

```shellscript
git clone https://github.com/MAAARKIN/unico.git

cd unico

docker-compose up -d
```

The application will be available on port :8080, to access the services available from the api, you can access the url: http://localhost:8080/swagger/index.html

# Run coverage test

To run the coverage test you will need

```shellscript
go test ./... -coverprofile=coverage.out

//to see the coverage
go tool cover -func=coverage.out

//to see the coverage in html file
go tool cover -html=coverage.out
```