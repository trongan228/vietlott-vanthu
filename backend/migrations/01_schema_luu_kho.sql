-- ============================================================
-- MODULE LƯU KHO: THÙNG → HỘP → HỒ SƠ → VĂN BẢN
-- Chạy SAU khi đã tạo schema chính (van_ban_den, van_ban_di_*)
-- Mô hình vật lý: 1 THÙNG chứa nhiều HỘP; 1 HỘP chứa nhiều HỒ SƠ
-- (tập lưu); 1 HỒ SƠ mô tả một dải văn bản (theo khoảng ngày
-- hoặc khoảng số) → từ đó tra được văn bản nằm ở hộp/thùng nào.
-- ============================================================

-- 1. THÙNG (carton lưu kho, nguồn: file "Tra cứu lưu kho")
CREATE TABLE IF NOT EXISTS thung (
    id              SERIAL PRIMARY KEY,
    ma_thung        VARCHAR(20) UNIQUE NOT NULL,   -- vd: VTVP-0001
    so_serial       VARCHAR(20),                    -- số serial dán trên thùng (vd: 6711, 006952)
    dot_luu_kho     SMALLINT,                       -- đợt gửi kho: 1, 2, ...
    vi_tri_kho      VARCHAR(100),                   -- vị trí kệ/kho (bổ sung sau nếu có)
    ngay_nhap_kho   DATE,
    ghi_chu         TEXT,                           -- vd: "Hộp số 13;14;15 + HĐĐL 2018" (hộp vĩnh viễn kèm theo)
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- 2. HỘP
-- Lưu ý nghiệp vụ: hộp "có thời hạn" và hộp "vĩnh viễn" là HAI dãy số
-- riêng biệt (trong file lưu kho có cột riêng cho từng loại),
-- nên khóa duy nhất là (so_hop, loai_hop).
CREATE TABLE IF NOT EXISTS hop (
    id          SERIAL PRIMARY KEY,
    so_hop      INT NOT NULL,
    loai_hop    VARCHAR(15) NOT NULL DEFAULT 'thoi_han'
                CHECK (loai_hop IN ('thoi_han','vinh_vien')),
    thung_id    INT REFERENCES thung(id),           -- NULL = hộp chưa xếp vào thùng/chưa gửi kho
    ghi_chu     TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT uq_hop UNIQUE (so_hop, loai_hop)
);
CREATE INDEX IF NOT EXISTS idx_hop_thung ON hop (thung_id);

-- 3. HỒ SƠ LƯU TRỮ (tập lưu văn bản — nguồn: DANH MỤC TÀI LIỆU CÓ THỜI HẠN)
-- Mỗi hồ sơ mô tả một dải văn bản: loại tập + năm + khoảng ngày
-- và/hoặc khoảng số văn bản. Các cột loai_tap/nam/tu_ngay/den_ngay/
-- tu_so/den_so được bóc tách tự động từ tiêu đề để phục vụ tra cứu;
-- tieu_de luôn giữ nguyên văn làm nguồn sự thật.
CREATE TABLE IF NOT EXISTS ho_so_luu_tru (
    id                      SERIAL PRIMARY KEY,
    so_ho_so                INT UNIQUE NOT NULL,     -- số hồ sơ, đánh liên tục toàn phông (1..1752..)
    hop_id                  INT NOT NULL REFERENCES hop(id),
    tieu_de                 TEXT NOT NULL,           -- nguyên văn từ danh mục
    loai_tap                VARCHAR(30) NOT NULL,    -- xem CHECK bên dưới
    nam                     SMALLINT,                -- năm của tập lưu (nếu bóc tách được)
    tu_ngay                 DATE,                    -- khoảng ngày văn bản trong tập
    den_ngay                DATE,
    tu_so                   INT,                     -- khoảng số văn bản trong tập (với CV đi)
    den_so                  INT,
    so_to                   INT,                     -- cột "Số tờ" (nguồn hầu hết để trống)
    thoi_han_bao_quan_nam   INT,                     -- 5/10/20/50/70; NULL nếu vĩnh viễn
    vinh_vien               BOOLEAN NOT NULL DEFAULT FALSE,
    ghi_chu                 TEXT,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT ck_loai_tap CHECK (loai_tap IN (
        'CV_DEN',                  -- tập lưu công văn đến
        'CV_DI_CO_TEN_LOAI',       -- tập lưu công văn đi có tên loại
        'CV_DI_KHONG_TEN_LOAI',    -- tập lưu công văn đi không tên loại
        'CV_TO_BAN_HD',            -- công văn Tổ, Ban, Hội đồng
        'CV_DOAN_THE',             -- công văn tổ chức đoàn thể
        'CV_CHI_NHANH',            -- tập lưu công văn các chi nhánh
        'CV_NOI_BO',               -- công văn đi nội bộ các phòng
        'HDLD',                    -- hợp đồng lao động & phụ lục
        'HD_LAY_SO',               -- hợp đồng lấy số
        'KHAC'                     -- biên bản, BCTC, giấy giới thiệu, sao y...
    ))
);
CREATE INDEX IF NOT EXISTS idx_hstl_hop      ON ho_so_luu_tru (hop_id);
CREATE INDEX IF NOT EXISTS idx_hstl_loai_nam ON ho_so_luu_tru (loai_tap, nam);
CREATE INDEX IF NOT EXISTS idx_hstl_ngay     ON ho_so_luu_tru (tu_ngay, den_ngay);
CREATE INDEX IF NOT EXISTS idx_hstl_so       ON ho_so_luu_tru (tu_so, den_so);

-- ============================================================
-- 4. VIEW ĐỊNH VỊ: văn bản → hồ sơ → hộp → thùng
-- Dùng để trả lời câu hỏi "văn bản X đang nằm ở hộp nào, thùng nào".
-- Nguyên tắc khớp:
--   * CV đến      : theo NGÀY ĐẾN nằm trong [tu_ngay, den_ngay]
--   * CV đi (2 loại): ưu tiên khớp theo KHOẢNG SỐ (số đầu của Số ký
--     hiệu ∈ [tu_so, den_so]); nếu hồ sơ không có khoảng số thì khớp
--     theo khoảng ngày.
-- ============================================================

CREATE OR REPLACE VIEW v_vi_tri_van_ban_den AS
SELECT vb.id AS van_ban_id, vb.nam, vb.so_den, vb.ngay_den, vb.trich_yeu,
       hs.so_ho_so, hs.tieu_de AS ho_so_tieu_de,
       h.so_hop, t.ma_thung, t.so_serial
FROM van_ban_den vb
JOIN ho_so_luu_tru hs
  ON hs.loai_tap = 'CV_DEN'
 AND vb.ngay_den BETWEEN hs.tu_ngay AND hs.den_ngay
JOIN hop h   ON h.id = hs.hop_id
LEFT JOIN thung t ON t.id = h.thung_id;

CREATE OR REPLACE VIEW v_vi_tri_van_ban_di_co_ten_loai AS
SELECT vb.id AS van_ban_id, vb.nam, vb.so_ky_hieu, vb.ngay_van_ban, vb.trich_yeu,
       hs.so_ho_so, hs.tieu_de AS ho_so_tieu_de,
       h.so_hop, t.ma_thung, t.so_serial
FROM van_ban_di_co_ten_loai vb
JOIN ho_so_luu_tru hs
  ON hs.loai_tap = 'CV_DI_CO_TEN_LOAI'
 AND hs.nam = vb.nam
 AND (
       (hs.tu_so IS NOT NULL
        AND (substring(vb.so_ky_hieu FROM '^\d+'))::int BETWEEN hs.tu_so AND hs.den_so)
    OR (hs.tu_so IS NULL
        AND vb.ngay_van_ban BETWEEN hs.tu_ngay AND hs.den_ngay)
 )
JOIN hop h   ON h.id = hs.hop_id
LEFT JOIN thung t ON t.id = h.thung_id;

CREATE OR REPLACE VIEW v_vi_tri_van_ban_di_khong_ten_loai AS
SELECT vb.id AS van_ban_id, vb.nam, vb.so_ky_hieu, vb.ngay_van_ban, vb.trich_yeu,
       hs.so_ho_so, hs.tieu_de AS ho_so_tieu_de,
       h.so_hop, t.ma_thung, t.so_serial
FROM van_ban_di_khong_ten_loai vb
JOIN ho_so_luu_tru hs
  ON hs.loai_tap = 'CV_DI_KHONG_TEN_LOAI'
 AND hs.nam = vb.nam
 AND (
       (hs.tu_so IS NOT NULL
        AND (substring(vb.so_ky_hieu FROM '^\d+'))::int BETWEEN hs.tu_so AND hs.den_so)
    OR (hs.tu_so IS NULL
        AND vb.ngay_van_ban BETWEEN hs.tu_ngay AND hs.den_ngay)
 )
JOIN hop h   ON h.id = hs.hop_id
LEFT JOIN thung t ON t.id = h.thung_id;

-- Ví dụ tra cứu: văn bản đến số 150 năm 2022 nằm ở đâu?
-- SELECT * FROM v_vi_tri_van_ban_den WHERE nam = 2022 AND so_den = 150;
