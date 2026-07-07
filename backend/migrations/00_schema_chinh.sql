-- ============================================================
-- 00_schema_chinh.sql — Schema chính: Văn bản đến / Văn bản đi
-- Trích từ tài liệu "thiet_ke_csdl_van_thu.md", mục 3
-- Chạy đầu tiên, trước 01_schema_luu_kho.sql
-- ============================================================

-- ============================================================
-- 0. TIỆN ÍCH
-- ============================================================
CREATE EXTENSION IF NOT EXISTS unaccent;   -- tìm kiếm tiếng Việt không dấu
CREATE EXTENSION IF NOT EXISTS pg_trgm;    -- tìm kiếm gần đúng (LIKE %...%) có index

-- ============================================================
-- 1. BẢNG DANH MỤC DÙNG CHUNG
-- ============================================================

CREATE TABLE don_vi (
    id              SERIAL PRIMARY KEY,
    ma_don_vi       VARCHAR(30) UNIQUE,
    ten_don_vi      VARCHAR(255) NOT NULL,
    loai_don_vi     VARCHAR(20) NOT NULL
                    CHECK (loai_don_vi IN ('noi_bo','chi_nhanh','doi_tac','to_chuc')),
    dia_chi         VARCHAR(255),
    ghi_chu         TEXT,
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE can_bo (
    id              SERIAL PRIMARY KEY,
    ho_ten          VARCHAR(150) NOT NULL,
    chuc_danh       VARCHAR(100),
    don_vi_id       INT REFERENCES don_vi(id),
    so_dien_thoai   VARCHAR(20),
    email           VARCHAR(150),
    is_van_thu      BOOLEAN NOT NULL DEFAULT FALSE,
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    ghi_chu         TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE loai_van_ban (
    id                      SERIAL PRIMARY KEY,
    ma_loai                 VARCHAR(20) UNIQUE NOT NULL,
    ten_loai                VARCHAR(150) NOT NULL,
    thoi_han_bao_quan_nam   INT,
    ghi_chu                 TEXT
);

INSERT INTO loai_van_ban (ma_loai, ten_loai, thoi_han_bao_quan_nam) VALUES
    ('QĐ',  'Quyết định',   20),
    ('TB',  'Thông báo',    20),
    ('TTr', 'Tờ trình',     20),
    ('BC',  'Báo cáo',      20),
    ('UQ',  'Ủy quyền',     20),
    ('CT',  'Chỉ thị',      20),
    ('GM',  'Giấy mời',     20),
    ('CV',  'Công văn',     20);

-- ============================================================
-- 2. VĂN BẢN ĐẾN
-- ============================================================
CREATE TABLE van_ban_den (
    id                  BIGSERIAL PRIMARY KEY,
    nam                 SMALLINT NOT NULL,
    so_den              INT NOT NULL,
    ngay_den            DATE NOT NULL,
    noi_gui_id          INT REFERENCES don_vi(id),
    noi_gui_text        VARCHAR(255) NOT NULL,
    so_ky_hieu          VARCHAR(100),
    ngay_van_ban        DATE,
    trich_yeu           TEXT NOT NULL,
    don_vi_xu_ly_id     INT REFERENCES don_vi(id),
    don_vi_nhan_text    VARCHAR(255),
    ky_nhan             VARCHAR(150),
    ghi_chu             TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT uq_van_ban_den_nam_so UNIQUE (nam, so_den)
);

CREATE INDEX idx_vbd_ngay_den    ON van_ban_den (ngay_den);
CREATE INDEX idx_vbd_ngay_vb     ON van_ban_den (ngay_van_ban);
CREATE INDEX idx_vbd_noi_gui     ON van_ban_den (noi_gui_id);
CREATE INDEX idx_vbd_trich_yeu_trgm ON van_ban_den USING gin (trich_yeu gin_trgm_ops);

-- ============================================================
-- 3. VĂN BẢN ĐI — KHÔNG TÊN LOẠI
-- ============================================================
CREATE TABLE van_ban_di_khong_ten_loai (
    id              BIGSERIAL PRIMARY KEY,
    nam             SMALLINT NOT NULL,
    so_ky_hieu      VARCHAR(100) NOT NULL,
    ngay_van_ban    DATE NOT NULL,
    trich_yeu       TEXT NOT NULL,
    nguoi_ky_id     INT REFERENCES can_bo(id),
    nguoi_ky_text   VARCHAR(150),
    noi_nhan_text   VARCHAR(500),
    so_luong_ban    INT,
    ghi_chu         TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT uq_vbdktl_nam_soky UNIQUE (nam, so_ky_hieu)
);

CREATE INDEX idx_vbdktl_ngay_vb ON van_ban_di_khong_ten_loai (ngay_van_ban);
CREATE INDEX idx_vbdktl_nguoiky ON van_ban_di_khong_ten_loai (nguoi_ky_id);
CREATE INDEX idx_vbdktl_trich_yeu_trgm ON van_ban_di_khong_ten_loai USING gin (trich_yeu gin_trgm_ops);

-- ============================================================
-- 4. VĂN BẢN ĐI — CÓ TÊN LOẠI
-- ============================================================
CREATE TABLE van_ban_di_co_ten_loai (
    id              BIGSERIAL PRIMARY KEY,
    nam             SMALLINT NOT NULL,
    loai_van_ban_id INT REFERENCES loai_van_ban(id),
    so_ky_hieu      VARCHAR(100) NOT NULL,
    ngay_van_ban    DATE NOT NULL,
    trich_yeu       TEXT NOT NULL,
    nguoi_ky_id     INT REFERENCES can_bo(id),
    nguoi_ky_text   VARCHAR(150),
    noi_nhan_text   VARCHAR(500),
    so_luong_ban    INT,
    ghi_chu         TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT uq_vbdctl_nam_soky UNIQUE (nam, so_ky_hieu)
);

CREATE INDEX idx_vbdctl_ngay_vb   ON van_ban_di_co_ten_loai (ngay_van_ban);
CREATE INDEX idx_vbdctl_loai      ON van_ban_di_co_ten_loai (loai_van_ban_id);
CREATE INDEX idx_vbdctl_nguoiky   ON van_ban_di_co_ten_loai (nguoi_ky_id);
CREATE INDEX idx_vbdctl_trich_yeu_trgm ON van_ban_di_co_ten_loai USING gin (trich_yeu gin_trgm_ops);

-- ============================================================
-- 5. TRIGGER cập nhật updated_at
-- ============================================================
CREATE OR REPLACE FUNCTION trg_set_updated_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_vbd  BEFORE UPDATE ON van_ban_den
    FOR EACH ROW EXECUTE FUNCTION trg_set_updated_at();
CREATE TRIGGER set_updated_at_vbdktl BEFORE UPDATE ON van_ban_di_khong_ten_loai
    FOR EACH ROW EXECUTE FUNCTION trg_set_updated_at();
CREATE TRIGGER set_updated_at_vbdctl BEFORE UPDATE ON van_ban_di_co_ten_loai
    FOR EACH ROW EXECUTE FUNCTION trg_set_updated_at();
