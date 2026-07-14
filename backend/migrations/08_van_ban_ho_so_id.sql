-- ============================================================
-- 08. GẮN VĂN BẢN VÀO HỒ SƠ LƯU TRỮ (ho_so_id)
-- Thêm cột ho_so_id vào 3 bảng văn bản và backfill một lần theo
-- mô tả trong ho_so_luu_tru (cùng quy tắc khớp với các view
-- v_vi_tri_van_ban_* ở migration 01), để tra nhanh số hộp/số thùng.
--
-- Quy tắc chọn khi một văn bản khớp NHIỀU hồ sơ (khoảng chồng lấn):
--   1. Ưu tiên hồ sơ khớp theo KHOẢNG SỐ (tu_so/den_so) hơn khoảng ngày
--      (với văn bản đi; văn bản đến chỉ khớp theo ngày đến).
--   2. Trong các hồ sơ cùng mức, chọn hồ sơ có khoảng HẸP nhất
--      (mô tả cụ thể hơn thì đáng tin hơn).
--   3. Còn hòa thì lấy hồ sơ có so_ho_so nhỏ nhất (ổn định, tái lập được).
-- Văn bản không khớp hồ sơ nào (vd: 2024-2025 chưa có danh mục) giữ NULL.
-- ============================================================

ALTER TABLE van_ban_den
    ADD COLUMN IF NOT EXISTS ho_so_id INT REFERENCES ho_so_luu_tru(id);
ALTER TABLE van_ban_di_khong_ten_loai
    ADD COLUMN IF NOT EXISTS ho_so_id INT REFERENCES ho_so_luu_tru(id);
ALTER TABLE van_ban_di_co_ten_loai
    ADD COLUMN IF NOT EXISTS ho_so_id INT REFERENCES ho_so_luu_tru(id);

-- Văn bản đến: khớp theo ngày đến ∈ [tu_ngay, den_ngay] của tập CV_DEN.
UPDATE van_ban_den vb
SET ho_so_id = best.hs_id
FROM (
    SELECT DISTINCT ON (v.id) v.id AS vb_id, hs.id AS hs_id
    FROM van_ban_den v
    JOIN ho_so_luu_tru hs
      ON hs.loai_tap = 'CV_DEN'
     AND v.ngay_den BETWEEN hs.tu_ngay AND hs.den_ngay
    ORDER BY v.id, (hs.den_ngay - hs.tu_ngay), hs.so_ho_so
) best
WHERE vb.id = best.vb_id;

-- Văn bản đi không tên loại: ưu tiên khớp số đầu của so_ky_hieu ∈ [tu_so, den_so],
-- nếu hồ sơ không có khoảng số thì khớp theo khoảng ngày.
UPDATE van_ban_di_khong_ten_loai vb
SET ho_so_id = best.hs_id
FROM (
    SELECT DISTINCT ON (v.id) v.id AS vb_id, hs.id AS hs_id
    FROM van_ban_di_khong_ten_loai v
    JOIN ho_so_luu_tru hs
      ON hs.loai_tap = 'CV_DI_KHONG_TEN_LOAI'
     AND hs.nam = v.nam
     AND (
           (hs.tu_so IS NOT NULL
            AND (substring(v.so_ky_hieu FROM '^\d+'))::int BETWEEN hs.tu_so AND hs.den_so)
        OR (hs.tu_so IS NULL
            AND v.ngay_van_ban BETWEEN hs.tu_ngay AND hs.den_ngay)
     )
    ORDER BY v.id,
             (hs.tu_so IS NULL),                                        -- khớp số trước
             COALESCE(hs.den_so - hs.tu_so, hs.den_ngay - hs.tu_ngay),  -- khoảng hẹp nhất
             hs.so_ho_so
) best
WHERE vb.id = best.vb_id;

-- Văn bản đi có tên loại: cùng quy tắc với loại không tên loại.
UPDATE van_ban_di_co_ten_loai vb
SET ho_so_id = best.hs_id
FROM (
    SELECT DISTINCT ON (v.id) v.id AS vb_id, hs.id AS hs_id
    FROM van_ban_di_co_ten_loai v
    JOIN ho_so_luu_tru hs
      ON hs.loai_tap = 'CV_DI_CO_TEN_LOAI'
     AND hs.nam = v.nam
     AND (
           (hs.tu_so IS NOT NULL
            AND (substring(v.so_ky_hieu FROM '^\d+'))::int BETWEEN hs.tu_so AND hs.den_so)
        OR (hs.tu_so IS NULL
            AND v.ngay_van_ban BETWEEN hs.tu_ngay AND hs.den_ngay)
     )
    ORDER BY v.id,
             (hs.tu_so IS NULL),
             COALESCE(hs.den_so - hs.tu_so, hs.den_ngay - hs.tu_ngay),
             hs.so_ho_so
) best
WHERE vb.id = best.vb_id;

CREATE INDEX IF NOT EXISTS idx_vbd_ho_so    ON van_ban_den (ho_so_id);
CREATE INDEX IF NOT EXISTS idx_vbdktl_ho_so ON van_ban_di_khong_ten_loai (ho_so_id);
CREATE INDEX IF NOT EXISTS idx_vbdctl_ho_so ON van_ban_di_co_ten_loai (ho_so_id);
