# Bối cảnh dự án

API quản lý văn thư nội bộ Vietlott. Backend Go (Gin + pgx), PostgreSQL
(hosting: Neon, connection string trong .env, KHÔNG hardcode).

## Nguyên tắc bắt buộc
- Dùng SQL thuần qua pgx, KHÔNG dùng ORM (yêu cầu audit/tuning cho dữ liệu
  nội bộ nhạy cảm).
- Query luôn dùng tham số hóa ($1, $2...), không nối chuỗi.
- Schema nguồn sự thật nằm ở migrations/00-08 — không tự đổi tên cột/bảng.
- Kiến trúc: handler → service → repository → database, giữ nguyên layer
  đang có trong internal/.

## Các bảng chính
- van_ban_den, van_ban_di_khong_ten_loai, van_ban_di_co_ten_loai
- Danh mục: don_vi, can_bo, loai_van_ban
- Lưu kho: thung, hop, ho_so_luu_tru + 3 view v_vi_tri_van_ban_*
- Mỗi bảng văn bản có ho_so_id (migration 08, backfill theo quy tắc khớp
  của view) → tra số hộp/số thùng; vi_tri được nhúng trong /van-ban/tim-kiem