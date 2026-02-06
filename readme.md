# Manual CRUD básico en Go

## Requisitos

- Go 1.22+
- MySQL corriendo
- Archivo `.env` configurado

---

## 1. Clonar proyecto

```bash
git clone <repo>
cd <repo>
```

---

## 2. Instalar dependencias

El proyecto ya incluye `go.mod` y `go.sum`.

```bash
go mod tidy
```

---

## 3. Configurar variables de entorno

Crear archivo `.env` en la raíz:

```
JWT_EXPIRATION_HOURS=24
COOKIE_DOMAIN=localhost
COOKIE_SECURE=false
COOKIE_SAME_SITE=Lax
DB_DSN=user:pass@tcp(localhost:3306)/testdb?parseTime=true
```

---

## 4. Ejecutar

```bash
go run main.go
```

---

## 5. Build binario

```bash
go build -o app
./app
```

---

## 6. Ejecutar con Docker

```bash
docker build -t go-crud .
docker run -p 8080:8080 go-crud
```

---

## Comandos útiles

Actualizar dependencias:

```bash
go mod tidy
```

Actualizar módulos:

```bash
go get -u ./...
```
