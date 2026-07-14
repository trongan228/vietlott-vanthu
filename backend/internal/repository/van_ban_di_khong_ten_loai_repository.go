package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"vanthu-backend/internal/model"
)

type VanBanDiKhongTenLoaiRepository struct {
	db *pgxpool.Pool
}

func NewVanBanDiKhongTenLoaiRepository(db *pgxpool.Pool) *VanBanDiKhongTenLoaiRepository {
	return &VanBanDiKhongTenLoaiRepository{db: db}
}

const vanBanDiKhongTenLoaiColumns = `id, nam, so_ky_hieu, ngay_van_ban, trich_yeu, nguoi_ky_id,
	nguoi_ky_text, noi_nhan_text, so_luong_ban, ghi_chu, ho_so_id, created_at, updated_at`

func scanVanBanDiKhongTenLoai(row pgx.Row) (*model.VanBanDiKhongTenLoai, error) {
	var (
		m          model.VanBanDiKhongTenLoai
		ngayVanBan time.Time
	)

	err := row.Scan(
		&m.ID, &m.Nam, &m.SoKyHieu, &ngayVanBan, &m.TrichYeu, &m.NguoiKyID,
		&m.NguoiKyText, &m.NoiNhanText, &m.SoLuongBan, &m.GhiChu, &m.HoSoID, &m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	m.NgayVanBan = model.DateFromTime(ngayVanBan)
	return &m, nil
}

func (r *VanBanDiKhongTenLoaiRepository) Create(ctx context.Context, in *model.VanBanDiKhongTenLoaiInput) (*model.VanBanDiKhongTenLoai, error) {
	query := `INSERT INTO van_ban_di_khong_ten_loai
		(nam, so_ky_hieu, ngay_van_ban, trich_yeu, nguoi_ky_id, nguoi_ky_text, noi_nhan_text, so_luong_ban, ghi_chu)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING ` + vanBanDiKhongTenLoaiColumns

	row := r.db.QueryRow(ctx, query,
		in.Nam, in.SoKyHieu, in.NgayVanBan.Time, in.TrichYeu, in.NguoiKyID, in.NguoiKyText,
		in.NoiNhanText, in.SoLuongBan, in.GhiChu,
	)
	return scanVanBanDiKhongTenLoai(row)
}

func (r *VanBanDiKhongTenLoaiRepository) GetByID(ctx context.Context, id int64) (*model.VanBanDiKhongTenLoai, error) {
	query := `SELECT ` + vanBanDiKhongTenLoaiColumns + ` FROM van_ban_di_khong_ten_loai WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	return scanVanBanDiKhongTenLoai(row)
}

func (r *VanBanDiKhongTenLoaiRepository) Update(ctx context.Context, id int64, in *model.VanBanDiKhongTenLoaiInput) (*model.VanBanDiKhongTenLoai, error) {
	query := `UPDATE van_ban_di_khong_ten_loai SET
		nam = $1, so_ky_hieu = $2, ngay_van_ban = $3, trich_yeu = $4, nguoi_ky_id = $5,
		nguoi_ky_text = $6, noi_nhan_text = $7, so_luong_ban = $8, ghi_chu = $9
		WHERE id = $10
		RETURNING ` + vanBanDiKhongTenLoaiColumns

	row := r.db.QueryRow(ctx, query,
		in.Nam, in.SoKyHieu, in.NgayVanBan.Time, in.TrichYeu, in.NguoiKyID, in.NguoiKyText,
		in.NoiNhanText, in.SoLuongBan, in.GhiChu, id,
	)
	return scanVanBanDiKhongTenLoai(row)
}

func (r *VanBanDiKhongTenLoaiRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM van_ban_di_khong_ten_loai WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// vanBanDiSortColumns là whitelist cột được phép sắp xếp cho 2 bảng văn bản đi.
var vanBanDiSortColumns = map[string]string{
	"so_ky_hieu":   "so_ky_hieu",
	"ngay_van_ban": "ngay_van_ban",
	"nam":          "nam",
	"created_at":   "created_at",
}

func (r *VanBanDiKhongTenLoaiRepository) List(ctx context.Context, f model.VanBanDiKhongTenLoaiFilter) ([]model.VanBanDiKhongTenLoai, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.Nam != nil {
		where = append(where, fmt.Sprintf("nam = $%d", argN))
		args = append(args, *f.Nam)
		argN++
	}
	if f.SoKyHieu != nil {
		where = append(where, fmt.Sprintf("so_ky_hieu ILIKE $%d", argN))
		args = append(args, "%"+*f.SoKyHieu+"%")
		argN++
	}
	if f.NguoiKyID != nil {
		where = append(where, fmt.Sprintf("nguoi_ky_id = $%d", argN))
		args = append(args, *f.NguoiKyID)
		argN++
	}
	if f.Search != "" {
		// q khớp gần đúng theo tên/trích yếu hoặc số văn bản.
		where = append(where, fmt.Sprintf("(trich_yeu ILIKE $%d OR so_ky_hieu ILIKE $%d)", argN, argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	orderBy, err := buildOrderBy(vanBanDiSortColumns, f.SortBy, f.SortDir, "nam DESC, ngay_van_ban DESC")
	if err != nil {
		return nil, 0, err
	}

	var total int64
	countQuery := `SELECT count(*) FROM van_ban_di_khong_ten_loai WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + vanBanDiKhongTenLoaiColumns + ` FROM van_ban_di_khong_ten_loai WHERE ` + whereClause +
		" ORDER BY " + orderBy +
		fmt.Sprintf(" LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.VanBanDiKhongTenLoai, 0)
	for rows.Next() {
		m, err := scanVanBanDiKhongTenLoai(rows)
		if err != nil {
			return nil, 0, err
		}
		items = append(items, *m)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
