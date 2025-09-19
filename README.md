# Base Code PKMB — Golang Microservices (Admin, Users) + PostgreSQL + Gateway

Base Code PKMB adalah starter kit microservices berbasis **Golang** dengan **PostgreSQL** dan sebuah **Gateway** sederhana.
Terdiri dari tiga service:
- `services/admin` — service admin (REST)
- `services/users` — service users (REST)
- `gateway` — reverse proxy sederhana yang meneruskan request ke admin & users
- `db/migrations` — skrip SQL inisialisasi

---

## Struktur Direktori
```
project-root/
├─ .env
├─ docker-compose.yml
├─ db/
│  └─ migrations/
├─ gateway/
│  ├─ Dockerfile
│  └─ main.go
├─ services/
│  ├─ admin/
│  │  ├─ Dockerfile
│  │  ├─ go.mod
│  │  └─ main.go
│  └─ users/
│     ├─ Dockerfile
│     ├─ go.mod
│     └─ main.go
└─ .gitignore
```

---

## Menjalankan dengan Docker Compose

1. **Buat file `.env`** (jangan commit file ini):
```env
POSTGRES_USER=appuser
POSTGRES_PASSWORD=apppassword
POSTGRES_DB=appdb
POSTGRES_PORT=5432
DATABASE_URL=postgres://appuser:apppassword@postgres:5432/appdb?sslmode=disable
GATEWAY_PORT=8080
ADMIN_PORT=8001
USERS_PORT=8002
```

2. **Build & Jalankan**
```bash
docker compose build --no-cache
docker compose up -d
```

3. **Cek Health**
- Gateway: `http://localhost:8080/health`
- Admin via gateway: `http://localhost:8080/admin/...`
- Users via gateway: `http://localhost:8080/users/...`

---

## Debugging
```bash
docker compose ps
docker compose logs --tail=200 --follow admin users gateway postgres
```
Jika perlu reset DB (hapus semua data):
```bash
docker compose down -v
docker compose up -d --build
```

---

## Development Tips
- Jangan commit `.env`.
- Gunakan GitHub Secrets untuk CI/CD.
- Pertimbangkan menambahkan healthcheck di `docker-compose.yml` agar gateway menunggu service healthy.

---

## Contributing
1. Fork → buat branch `feat/your-feature`
2. Commit dan buat pull request

---

## License
MIT
