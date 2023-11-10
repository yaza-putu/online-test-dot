# Online Test DOT
![golang](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png)

## How to start

- install depedency
  ```bash
    make tidy
  ```
- copy environment
  ```bash
    make config
  ```
- generate key
  ```bash
  make key
  ```
- run migration
```bash
go run zoro.go migrate:up
```

- run seeder
```bash
go run zoro.go seed:up
```
- run dev mode
  ```bash
    make run
  ```
- build
  ```bash
  make build
  ```
  

# Default Login
email : user@mail.com

pass : Password1

## Zoro command
- make migration
  ```bash
    go run zoro.go make:migration file_name
  ```
- migration up
  ```bash
    go run zoro.go migrate:up
  ```
- migration down
  ```bash
    go run zoro.go migrate:down
  ```

# Alasan memakai service & repository pattern
Saya biasanya memakai pattern service dan repository bertujuan memisahkan business logic dengan query logic serta memfungsikan handler untuk fokus pada pengarahan lalulintas data
kalau tidak memakai pattern cendrung kita menulis logika programing pada satu file katakan handler kalau project kecil tidak masalah tetapi jika project besar ini akan sulit di maintenace karena kode program yang banyak dan bercampur

selain pattern di project ini saya menambahkan otomatisasi pembuatan migration dan seeder untuk memudahkan development


  

## Validation Unique With Struct Tag
- unique
```go
type v struct {
	Name string `validate:"unique=table_name:column_name"`
}
// ecample
type v struct {
Name string `validate:"unique=users:name"`
}
```
- unique with ignore
```go
type v struct {
Name string `validate:"unique=table_name:column_name:ignore_with_field_name"`
ID   string `validate:"required"`
}
// example
type v struct {
Name string `validate:"unique=users:name:ID"`
ID   string `validate:"required"`
}
```
## Stack 
- [Echo](https://echo.labstack.com)
- [Gorm](https://gorm.io)
- [Env](https://github.com/spf13/viper)

