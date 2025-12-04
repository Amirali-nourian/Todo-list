# Todo API – Go Clean Architecture

A simple and clean Todo REST API built with **Go**, **Gin**, **PostgreSQL**, and **Swagger**, following principles of Clean Architecture.  
This project includes a modern **Web UI** and is intended as both a learning exercise and a portfolio-ready example of building maintainable HTTP services in Go.

---

## 🚀 Quick Start

### 1. Clone the repository

```bash
git clone [https://github.com/Amirali-nourian/Todo-list.git](https://github.com/Amirali-nourian/Todo-list.git)
cd Todo-list
````

### 2\. Environment variables

Copy the example environment file:

```bash
cp .env.example .env
```

Then open `.env` and adjust values as needed.

-----

## ▶️ Run locally (without Docker)

Ensure you have a running PostgreSQL instance and update `.env` with your database credentials.

```bash
go mod tidy
go run ./cmd/main.go
```

Server default URL:

```
http://localhost:8080
```

-----

## 🐳 Run with Docker / Docker Compose (Recommended)

```bash
docker-compose up --build
```

After containers start:

  * **Web UI (Task Manager):** `http://localhost:8080`
  * **Swagger UI:** `http://localhost:8080/swagger/index.html`

-----

## 📚 API Documentation (Swagger)

Interactive documentation is available here:

```
http://localhost:8080/swagger/index.html
```

Use **Try it out** to test endpoints directly.

-----

## 🔗 API Endpoints

Base path: `/api/v1`

| Method | Endpoint    | Description      |
| ------ | ----------- | ---------------- |
| POST   | /todos      | Create a todo    |
| GET    | /todos      | Get all todos    |
| GET    | /todos/{id} | Get a todo by ID |
| PUT    | /todos/{id} | Update a todo    |
| DELETE | /todos/{id} | Delete a todo    |

-----

## 🧱 Tech Stack

  * Go
  * Gin
  * Clean Architecture
  * Docker & Docker Compose
  * Swagger / OpenAPI
  * **PostgreSQL** (Database)
  * **GORM** (ORM)
  * **HTML/CSS/JS** (Web UI)

-----

## 🗂 Project Structure

```
.
├── cmd/               # Application entrypoint
├── internal/          # Use cases, handlers, domain logic
├── web/               # Frontend files (HTML/CSS/JS)
├── docs/              # Swagger files
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── README.md
```

-----

## ✅ Running Tests

```bash
go test ./...
```

-----

## 🧠 Future Improvements

  * Add JWT authentication
  * Add pagination & filters
  * Add CI/CD via GitHub Actions
  * Better logging and error handling

-----

## 👤 Author

Created and maintained by **Amirali Nourian**.

-----

-----

## Todo API – معماری پاک با Go

این پروژه یک API مدیریت کارها (Todo) است که با **Go**، فریم‌ورک **Gin**، دیتابیس **PostgreSQL** و مستندسازی **Swagger** ساخته شده و دارای یک **رابط کاربری (Web UI)** مدرن است. ساختار پروژه بر پایهٔ اصول «معماری پاک» طراحی شده است.

-----

## 🚀 شروع سریع

### ۱. کلون کردن مخزن

```bash
git clone [https://github.com/Amirali-nourian/Todo-list.git](https://github.com/Amirali-nourian/Todo-list.git)
cd Todo-list
```

### ۲. ساخت فایل تنظیمات محیطی

```bash
cp .env.example .env
```

سپس مقادیر را تغییر دهید.

-----

## ▶️ اجرای پروژه بدون Docker

ابتدا مطمئن شوید سرویس PostgreSQL در حال اجراست و اطلاعات اتصال را در فایل `.env` وارد کرده‌اید.

```bash
go mod tidy
go run ./cmd/main.go
```

آدرس پیش‌فرض:

```
http://localhost:8080
```

-----

## 🐳 اجرای پروژه با Docker / Docker Compose (پیشنهادی)

```bash
docker-compose up --build
```

پس از اجرا:

  * **رابط کاربری (مدیریت کارها):** `http://localhost:8080`
  * **مستندات Swagger:** `http://localhost:8080/swagger/index.html`

-----

## 📚 مستندات API

```
http://localhost:8080/swagger/index.html
```

می‌توانید مستقیماً درخواست‌ها را تست کنید.

-----

## 🔗 مسیرهای API

مسیر پایه: `/api/v1`

| متد    | مسیر        | توضیح             |
| ------ | ----------- | ----------------- |
| POST   | /todos      | ساخت یک کار       |
| GET    | /todos      | دریافت همهٔ کارها |
| GET    | /todos/{id} | دریافت یک کار     |
| PUT    | /todos/{id} | به‌روزرسانی کار   |
| DELETE | /todos/{id} | حذف کار           |

-----

## 🧱 فناوری‌های استفاده‌شده

  * زبان Go
  * Gin
  * معماری پاک
  * Docker و Docker Compose
  * Swagger
  * **PostgreSQL** (دیتابیس)
  * **GORM** (ORM)
  * **HTML/CSS/JS** (رابط کاربری وب)

-----

## 🗂 ساختار پروژه

```
.
├── cmd/               # نقطه ورود برنامه
├── internal/          # منطق برنامه و لایه‌ها
├── web/               # فایل‌های فرانت‌اند
├── docs/              # فایل‌های Swagger
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── README.md
```

-----

## ✅ تست‌ها

```bash
go test ./...
```

-----

## 🧠 مسیر توسعه آینده

  * افزودن احراز هویت با JWT
  * صفحه‌بندی (Pagination) و فیلتر
  * افزودن CI در GitHub Actions
  * بهبود لاگ‌گیری و مدیریت خطا

-----

## 👤 سازنده

ساخته‌شده توسط **Amirali Nourian**.

```
