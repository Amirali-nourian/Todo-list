
# Todo API – Go Clean Architecture

A simple and clean Todo REST API built with **Go**, **Gin**, and **Swagger**, following principles of Clean Architecture.  
This project is intended as both a learning exercise and a portfolio-ready example of building maintainable HTTP services in Go.

---

## 🚀 Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/Amirali-nourian/Todo-list.git
cd Todo-list
````

### 2. Environment variables

Copy the example environment file:

```bash
cp .env.example .env
```

Then open `.env` and adjust values as needed.

---

## ▶️ Run locally (without Docker)

```bash
go mod tidy
go run ./cmd/main.go
```

Server default URL:

```
http://localhost:8080
```

---

## 🐳 Run with Docker / Docker Compose

```bash
docker-compose up --build
```

After containers start:

* API URL: `http://localhost:8080`
* Swagger UI: `http://localhost:8080/swagger/index.html`

---

## 📚 API Documentation (Swagger)

Interactive documentation is available here:

```
http://localhost:8080/swagger/index.html
```

Use **Try it out** to test endpoints directly.

---

## 🔗 API Endpoints

Base path: `/api/v1`

| Method | Endpoint    | Description      |
| ------ | ----------- | ---------------- |
| POST   | /todos      | Create a todo    |
| GET    | /todos      | Get all todos    |
| GET    | /todos/{id} | Get a todo by ID |
| PUT    | /todos/{id} | Update a todo    |
| DELETE | /todos/{id} | Delete a todo    |

---

## 🧱 Tech Stack

* Go
* Gin
* Clean Architecture
* Docker & Docker Compose
* Swagger / OpenAPI
* In-memory storage (can be replaced with real DB)

---

## 🗂 Project Structure

```
.
├── cmd/               # Application entrypoint
├── internal/          # Use cases, handlers, domain logic
├── docs/              # Swagger files
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── README.md
```

---

## ✅ Running Tests

```bash
go test ./...
```

---

## 🧠 Future Improvements

* Add PostgreSQL / MySQL
* Add JWT authentication
* Add pagination & filters
* Add CI/CD via GitHub Actions
* Better logging and error handling

---

## 👤 Author

Created and maintained by **Amirali Nourian**.

---

---


## Todo API – معماری پاک با Go

این پروژه یک API ساده برای مدیریت کارها (Todo) است که با **Go**، فریم‌ورک **Gin** و مستندسازی **Swagger** ساخته شده و ساختار آن بر پایهٔ اصول «معماری پاک» طراحی شده است.
هدف پروژه، هم یادگیری و هم ارائهٔ یک نمونه‌کار مناسب برای رزومه است.

---

## 🚀 شروع سریع

### ۱. کلون کردن مخزن

```bash
git clone https://github.com/Amirali-nourian/Todo-list.git
cd Todo-list
```

### ۲. ساخت فایل تنظیمات محیطی

```bash
cp .env.example .env
```

سپس مقادیر را تغییر دهید.

---

## ▶️ اجرای پروژه بدون Docker

```bash
go mod tidy
go run ./cmd/main.go
```

آدرس پیش‌فرض:

```
http://localhost:8080
```

---

## 🐳 اجرای پروژه با Docker / Docker Compose

```bash
docker-compose up --build
```

پس از اجرا:

* آدرس API: `http://localhost:8080`
* مستندات Swagger: `http://localhost:8080/swagger/index.html`

---

## 📚 مستندات API

```
http://localhost:8080/swagger/index.html
```

می‌توانید مستقیماً درخواست‌ها را تست کنید.

---

## 🔗 مسیرهای API

مسیر پایه: `/api/v1`

| متد    | مسیر        | توضیح             |
| ------ | ----------- | ----------------- |
| POST   | /todos      | ساخت یک کار       |
| GET    | /todos      | دریافت همهٔ کارها |
| GET    | /todos/{id} | دریافت یک کار     |
| PUT    | /todos/{id} | به‌روزرسانی کار   |
| DELETE | /todos/{id} | حذف کار           |

---

## 🧱 فناوری‌های استفاده‌شده

* زبان Go
* Gin
* معماری پاک
* Docker و Docker Compose
* Swagger
* حافظهٔ درون‌برنامه‌ای (قابل ارتقا به دیتابیس واقعی)

---

## 🗂 ساختار پروژه

```
.
├── cmd/               # نقطه ورود برنامه
├── internal/          # منطق برنامه و لایه‌ها
├── docs/              # فایل‌های Swagger
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── README.md
```

---

## ✅ تست‌ها

```bash
go test ./...
```

---

## 🧠 مسیر توسعه آینده

* اتصال به PostgreSQL یا MySQL
* افزودن JWT
* صفحه‌بندی و فیلتر
* افزودن CI در GitHub Actions
* بهبود لاگ‌گیری و مدیریت خطا

---

## 👤 سازنده

ساخته‌شده توسط **Amirali Nourian**.
