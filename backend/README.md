# Vietlott Văn Thư — Backend

API quản lý văn thư nội bộ Vietlott (công văn đến/đi, danh mục, tra cứu lưu kho vật lý).
Xây dựng bằng **Go + Gin + pgx**, dữ liệu lưu trên **PostgreSQL** (hosting Neon).
Frontend React (thư mục `../frontend`) gọi vào API này qua REST.

## Công nghệ chính

- **Go 1.25** — ngôn ngữ chính.
- **Gin (`gin-gonic/gin`)** — HTTP router & middleware (logger, recovery).
- **pgx v5 (`jackc/pgx/v5` + `pgxpool`)** — driver PostgreSQL, dùng connection pool.
  **KHÔNG dùng ORM** — toàn bộ truy vấn là SQL thuần, tham số hóa (`$1, $2...`) để phục vụ
  audit/tuning cho dữ liệu nội bộ nhạy cảm.
- **godotenv (`joho/godotenv`)** — nạp cấu hình từ file `.env` khi chạy local.

## Cách chạy

```bash
go mod download
go build ./...            # kiểm tra biên dịch
go run ./cmd/server       # chạy server, mặc định tại http://localhost:8080
```

Cấu hình qua file [.env](.env) (hoặc biến môi trường hệ thống):

| Biến           | Bắt buộc | Mô tả                                                        |
| -------------- | -------- | ------------------------------------------------------------ |
| `DATABASE_URL` | ✅       | Connection string PostgreSQL. **KHÔNG hardcode** trong code. |
| `PORT`         | ❌       | Cổng HTTP, mặc định `8080`.                                  |

Khi khởi động, server ping thử database (timeout 5s) — nếu kết nối thất bại thì dừng ngay
(xem [internal/db/db.go](internal/db/db.go)). Health check: `GET /healthz` → `{"status":"ok"}`.

## Cấu trúc thư mục

```
backend/
├── cmd/
│   └── server/
│       └── main.go              # entry point: load config → tạo pool → wire router → Run
├── internal/
│   ├── config/config.go         # đọc DATABASE_URL, PORT từ .env / env
│   ├── db/db.go                 # tạo pgxpool.Pool + ping kiểm tra kết nối
│   ├── router/router.go         # khởi tạo Gin, wire 3 layer cho mọi tài nguyên, group /api/v1
│   ├── model/                   # struct DTO/Input/Filter + phân trang, kiểu ngày tùy biến
│   │   ├── pagination.go        #   ListResult[T], Pagination, chuẩn hóa page/page_size
│   │   ├── date.go              #   kiểu ngày parse/format theo format của hệ thống
│   │   └── *.go                 #   van_ban_den, van_ban_di_*, don_vi, can_bo, thung, hop...
│   ├── repository/              # lớp truy cập DB — SQL thuần qua pgx, không ORM
│   │   ├── errors.go            #   ErrNotFound, ErrInvalidInput (lỗi domain)
│   │   ├── sort.go              #   whitelist cột sort (chống SQL injection ở ORDER BY)
│   │   └── *_repository.go      #   1 file / bảng: CRUD + List có filter/phân trang
│   ├── service/                 # lớp nghiệp vụ — validate, điều phối repository
│   │   └── *_service.go
│   ├── handler/                 # lớp HTTP — bind request, gọi service, trả JSON
│   │   ├── params.go            #   helper đọc path/query param (int16/int32, paging, bool)
│   │   └── *_handler.go         #   mỗi handler tự Register route group của mình
│   └── httpx/error.go           # ánh xạ lỗi domain / lỗi Postgres → HTTP status phù hợp
├── migrations/                  # SQL nguồn sự thật của schema + dữ liệu seed
│   ├── 00_schema_chinh.sql      #   bảng nghiệp vụ + danh mục
│   ├── 01_schema_luu_kho.sql    #   bảng lưu kho + 3 view v_vi_tri_van_ban_*
│   └── 02–07_data_*.sql         #   dữ liệu mẫu (thùng, hộp, hồ sơ, văn bản đến/đi)
├── postman/                     # collection + environment để test API thủ công
├── go.mod / go.sum
└── .env                         # DATABASE_URL, PORT (không commit giá trị thật)
```

## Kiến trúc: 4 lớp

Luồng một request đi qua các lớp cố định (không được phá vỡ):

```
HTTP request → handler → service → repository → database
```

1. **handler** (`internal/handler`) — bind & validate cú pháp request (JSON body, path/query param),
   gọi service, chuyển kết quả thành JSON. Mỗi handler tự `Register(rg)` route group của nó
   trong [router.go](internal/router/router.go).
2. **service** (`internal/service`) — logic nghiệp vụ & validate ngữ nghĩa, điều phối repository.
3. **repository** (`internal/repository`) — SQL thuần qua pgx, tham số hóa hoàn toàn. Cột dùng để
   `ORDER BY` được whitelist trong [sort.go](internal/repository/sort.go) để chống SQL injection.
4. **httpx** ([error.go](internal/httpx/error.go)) — điểm ánh xạ lỗi tập trung: `ErrNotFound` → 404,
   `ErrInvalidInput` → 400, lỗi Postgres `23505/23503/23514` → 409/400 với thông báo tiếng Việt.

Wiring đi từ trong ra ngoài: `repository.New...(pool)` → `service.New...(repo)` → `handler.New...(svc)`.

## Các nhóm endpoint (prefix `/api/v1`)

Mỗi tài nguyên dưới đây có đủ CRUD: `POST /`, `GET /`, `GET /:id`, `PUT /:id`, `DELETE /:id`.

**Nghiệp vụ — văn bản:**
- `van-ban-den` — văn bản đến
- `van-ban-di-khong-ten-loai` — văn bản đi không có tên loại
- `van-ban-di-co-ten-loai` — văn bản đi có tên loại

**Danh mục:**
- `loai-van-ban`, `don-vi`, `can-bo`

**Lưu kho vật lý:**
- `thung`, `hop`, `ho-so-luu-tru`

**Tra cứu (chỉ đọc):**
- `GET /vi-tri/van-ban-den?nam=&so_den=` — vị trí lưu kho (văn bản → hồ sơ → hộp → thùng),
  đọc từ 3 view `v_vi_tri_van_ban_*`.
- `GET /vi-tri/van-ban-di-co-ten-loai?nam=&so_ky_hieu=`
- `GET /vi-tri/van-ban-di-khong-ten-loai?nam=&so_ky_hieu=`
- `GET /van-ban/tim-kiem?q=&so=&nam=&limit=20` — tìm đồng thời trên cả 3 bảng văn bản theo
  trích yếu (`q`) và/hoặc số văn bản (`so`); cần ít nhất một trong `q` hoặc `so`.

### Danh sách & phân trang

Các endpoint `GET /` (List) nhận query param dùng chung: `page` (mặc định 1), `page_size`
(mặc định 20, tối đa 100), `sort_by`, `sort_dir`, và `q` để tìm kiếm; kèm các filter riêng theo
bảng (ví dụ văn bản đến: `nam`, `so_den`, `so_ky_hieu`, `noi_gui_id`, `don_vi_xu_ly_id`).
Response có dạng:

```json
{
  "items": [ /* ... */ ],
  "pagination": { "page": 1, "page_size": 20, "total": 123 }
}
```

## Cơ sở dữ liệu & migrations

Schema **nguồn sự thật** nằm ở thư mục `migrations/` (đánh số `00`→`07`) — không tự đổi tên
cột/bảng ngoài các file này.

- `00_schema_chinh.sql`, `01_schema_luu_kho.sql` — DDL: bảng nghiệp vụ, danh mục, lưu kho
  và 3 view `v_vi_tri_van_ban_den` / `_di_co_ten_loai` / `_di_khong_ten_loai`.
- `02`–`07` — dữ liệu seed (thùng, hộp, hồ sơ lưu trữ, và ba loại văn bản).

Bảng chính: `van_ban_den`, `van_ban_di_khong_ten_loai`, `van_ban_di_co_ten_loai`;
danh mục `don_vi`, `can_bo`, `loai_van_ban`; lưu kho `thung`, `hop`, `ho_so_luu_tru`.

## Nguyên tắc bắt buộc

- Dùng SQL thuần qua pgx, **KHÔNG dùng ORM** (yêu cầu audit/tuning cho dữ liệu nhạy cảm).
- Query luôn **tham số hóa** (`$1, $2...`), không nối chuỗi; cột `ORDER BY` phải qua whitelist.
- Schema nguồn sự thật ở `migrations/00-07` — không tự đổi tên cột/bảng.
- Giữ nguyên kiến trúc **handler → service → repository → database** đang có trong `internal/`.
- Connection string **KHÔNG hardcode** — luôn lấy từ `DATABASE_URL`.

## Kiểm thử API

Import [postman/Vietlott-VanThu-Backend.postman_collection.json](postman/Vietlott-VanThu-Backend.postman_collection.json)
cùng environment [postman/Vietlott-VanThu-Local.postman_environment.json](postman/Vietlott-VanThu-Local.postman_environment.json)
vào Postman để gọi thử toàn bộ endpoint ở môi trường local.
