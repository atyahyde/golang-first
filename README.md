# Dokumentasi Instalasi

## Prasyarat

Pastikan perangkat sudah memiliki:

- Docker Engine
- Docker Compose
- Go 1.27 atau versi yang kompatibel jika ingin menjalankan aplikasi tanpa container
- Git, jika project diambil dari repository

## Instalasi dengan Docker Compose

### 1. Masuk ke folder project

```bash
cd /home/yurein/Documents/golang/assesment
```

### 2. Buat file `.env`

Docker Compose menggunakan file `.env` pada service aplikasi. Buat file tersebut di root project:

```env
DB_HOST=mysql
DB_PORT=3306
DB_USER=my_user
DB_PASSWORD=my_password
DB_NAME=my_db
REDIS_HOST=redis
REDIS_PORT=6379
```

Nilai database di atas harus sesuai dengan konfigurasi pada `docker-compose.yml`.

### 3. Build dan jalankan service

```bash
docker compose up --build -d
```

Jika Docker membutuhkan hak akses administrator:

```bash
sudo docker compose up --build -d
```

Perintah ini menjalankan:

- API Go pada port `3002`
- MySQL pada port `3306`
- Redis pada port `6379`

Service aplikasi menunggu MySQL dan Redis berstatus healthy sebelum dijalankan.

### 4. Periksa status container

```bash
docker compose ps
```

Lihat log aplikasi jika diperlukan:

```bash
docker compose logs -f app
```

### 5. Menghentikan service

```bash
docker compose down
```

Data MySQL dan Redis tetap tersimpan di Docker volume. Untuk menghapus container sekaligus seluruh data volume:

```bash
docker compose down -v
```

## Menjalankan Tanpa Container

Gunakan cara ini jika MySQL dan Redis sudah berjalan secara lokal.

### 1. Pastikan database tersedia

Buat database dan user MySQL yang sesuai, atau jalankan MySQL melalui Docker saja:

```bash
docker compose up -d mysql redis
```

### 2. Atur environment variable

Buat file `.env` di root project:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=my_user
DB_PASSWORD=my_password
DB_NAME=my_db
```

### 3. Download dependency dan jalankan aplikasi

```bash
go mod download
go run .
```

API tersedia di `http://localhost:3002`.

Saat aplikasi mulai, GORM menjalankan `AutoMigrate` untuk membuat atau memperbarui table `events` berdasarkan model `models.Event`.

## API Endpoint

### Membuat event

```http
POST http://localhost:3002/api/events
Content-Type: application/json
```

Contoh body:

```json
{
  "name": "Tech Conference",
  "description": "Conference teknologi tahunan",
  "location": "Jakarta",
  "datetime": "2026-10-01T09:00:00Z"
}
```

Field `name`, `description`, `location`, dan `datetime` wajib diisi. `user_id` saat ini diisi otomatis dengan nilai dummy `1` oleh aplikasi.

Contoh menggunakan cURL:

```bash
curl -X POST http://localhost:3002/api/events \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Tech Conference",
    "description": "Conference teknologi tahunan",
    "location": "Jakarta",
    "datetime": "2026-10-01T09:00:00Z"
  }'
```

### Mengambil semua event

```bash
curl http://localhost:3002/api/events
```

## Migrasi Database

Untuk menambahkan model baru, tambahkan model tersebut ke `AutoMigrate` di `main.go`:

```go
if err := config.DB.AutoMigrate(
    &models.Event{},
    &models.User{},
); err != nil {
    log.Fatalf("Gagal melakukan migrasi: %v", err)
}
```

Setelah perubahan, rebuild dan restart aplikasi:

```bash
docker compose up --build -d
```

`AutoMigrate` dapat membuat table baru dan menambahkan kolom baru. Proses ini tidak otomatis menghapus table atau kolom yang sudah ada.

## Troubleshooting

### Tidak dapat terhubung ke MySQL

Periksa status service dan log MySQL:

```bash
docker compose ps
docker compose logs mysql
```

Saat aplikasi berjalan di Docker, gunakan `DB_HOST=mysql`. Saat aplikasi dijalankan langsung dengan `go run .`, gunakan `DB_HOST=127.0.0.1`.

### Port sudah digunakan

Periksa proses yang menggunakan port `3002`, `3306`, atau `6379`, lalu hentikan proses tersebut atau ubah port pada `docker-compose.yml`.

### Data event tidak tersimpan

Periksa response API. Jika insert gagal, endpoint akan mengembalikan status HTTP `500` beserta pesan error database. Pastikan request menggunakan JSON valid dan seluruh field wajib sudah diisi.
