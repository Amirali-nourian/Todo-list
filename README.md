# 📝 Todo API - Go Clean Architecture

Simple **Todo API** built with **Go**, **Gin**, and **Swagger docs**, following the **Clean Architecture** pattern.

---

## 🚀 Run in 3 Seconds

```bash
# 1. Clone the repository
git clone https://github.com/mahdighadiriii/Todo-list-golang.git
cd Todo-list-golang

# 2. Run the server
go run cmd/main.go
```

Server runs at: [http://localhost:8080](http://localhost:8080)

---

## 📘 Swagger UI (API Docs)

Open in your browser:
👉 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Then click **“Try it out”** to test all endpoints!

---

## 🧩 Endpoints

| Method   | Endpoint             | Description             |
| -------- | -------------------- | ----------------------- |
| `POST`   | `/api/v1/todos`      | Create a todo           |
| `GET`    | `/api/v1/todos`      | List all todos          |
| `GET`    | `/api/v1/todos/{id}` | Get one todo            |
| `PUT`    | `/api/v1/todos/{id}` | Update (mark completed) |
| `DELETE` | `/api/v1/todos/{id}` | Delete a todo           |

---

## ⚙️ Tech Stack

* **Go + Gin**
* **Clean Architecture** (inspired by DDD)
* **In-memory DB** (easy to swap later)
* **Auto-generated Swagger Docs**

---

## Made with love by [Mahdi](https://github.com/https://github.com/mahdighadiriii)



### 🧾 Final Step

```bash
git add README.md
git commit -m "docs: add simple README"
git push origin develop
```
