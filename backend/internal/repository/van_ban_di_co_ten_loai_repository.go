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

type VanBanDiCoTenLoaiRepository struct {
	db *pgxpool.Pool
}

func NewVanBanDiCoTenLoaiRepository(db *pgxpool.Pool) *VanBanDiCoTenLoaiRepository {
	return &VanBanDiCoTenLoaiRepository{db: db}
}

const vanBanDiCoTenLoaiColumns = `id, nam, loai_van_ban_id, so_ky_hieu, ngay_van_ban, trich_yeu,
	nguoi_ky_id, nguoi_ky_text, noi_nhan_text, so_luong_ban, ghi_chu, created_at, updated_at`

func scanVanBanDiCoTenLoai(row pgx.Row) (*model.VanBanDiCoTenLoai, error) {
	var (
		m          model.VanBanDiCoTenLoai
		ngayVanBan time.Time
	)

	err := row.Scan(
		&m.ID, &m.Nam, &m.LoaiVanBanID, &m.SoKyHieu, &ngayVanBan, &m.TrichYeu,
		&m.NguoiKyID, &m.NguoiKyText, &m.NoiNhanText, &m.SoLuongBan, &m.GhiChu,
		&m.CreatedAt, &m.UpdatedAt,
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

func (r *VanBanDiCoTenLoaiRepository) Create(ctx context.Context, in *model.VanBanDiCoTenLoaiInput) (*model.VanBanDiCoTenLoai, error) {
	query := `INSERT INTO van_ban_di_co_ten_loai
		(nam, loai_van_ban_id, so_ky_hieu, ngay_van_ban, trich_yeu, nguoi_ky_id, nguoi_ky_text, noi_nhan_text, so_luong_ban, ghi_chu)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING ` + vanBanDiCoTenLoaiColumns

	row := r.db.QueryRow(ctx, query,
		in.Nam, in.LoaiVanBanID, in.SoKyHieu, in.NgayVanBan.Time, in.TrichYeu, in.NguoiKyID,
		in.NguoiKyText, in.NoiNhanText, in.SoLuongBan, in.GhiChu,
	)
	return scanVanBanDiCoTenLoai(row)
}

func (r *VanBanDiCoTenLoaiRepository) GetByID(ctx context.Context, id int64) (*model.VanBanDiCoTenLoai, error) {
	query := `SELECT ` + vanBanDiCoTenLoaiColumns + ` FROM van_ban_di_co_ten_loai WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	return scanVanBanDiCoTenLoai(row)
}

func (r *VanBanDiCoTenLoaiRepository) Update(ctx context.Context, id int64, in *model.VanBanDiCoTenLoaiInput) (*model.VanBanDiCoTenLoai, error) {
	query := `UPDATE van_ban_di_co_ten_loai SET
		nam = $1, loai_van_ban_id = $2, so_ky_hieu = $3, ngay_van_ban = $4, trich_yeu = $5,
		nguoi_ky_id = $6, nguoi_ky_text = $7, noi_nhan_text = $8, so_luong_ban = $9, ghi_chu = $10
		WHERE id = $11
		RETURNING ` + vanBanDiCoTenLoaiColumns

	row := r.db.QueryRow(ctx, query,
		in.Nam, in.LoaiVanBanID, in.SoKyHieu, in.NgayVanBan.Time, in.TrichYeu, in.NguoiKyID,
		in.NguoiKyText, in.NoiNhanText, in.SoLuongBan, in.GhiChu, id,
	)
	return scanVanBanDiCoTenLoai(row)
}

func (r *VanBanDiCoTenLoaiRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM van_ban_di_co_ten_loai WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *VanBanDiCoTenLoaiRepository) List(ctx context.Context, f model.VanBanDiCoTenLoaiFilter) ([]model.VanBanDiCoTenLoai, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.Nam != nil {
		where = append(where, fmt.Sprintf("nam = $%d", argN))
		args = append(args, *f.Nam)
		argN++
	}
	if f.LoaiVanBanID != nil {
		where = append(where, fmt.Sprintf("loai_van_ban_id = $%d", argN))
		args = append(args, *f.LoaiVanBanID)
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
		where = append(where, fmt.Sprintf("trich_yeu ILIKE $%d", argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	orderBy, err := buildOrderBy(vanBanDiSortColumns, f.SortBy, f.SortDir, "nam DESC, ngay_van_ban DESC")
	if err != nil {
		return nil, 0, err
	}

	var total int64
	countQuery := `SELECT count(*) FROM van_ban_di_co_ten_loai WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + vanBanDiCoTenLoaiColumns + ` FROM van_ban_di_co_ten_loai WHERE ` + whereClause +
		" ORDER BY " + orderBy +
		fmt.Sprintf(" LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.VanBanDiCoTenLoai, 0)
	for rows.Next() {
		m, err := scanVanBanDiCoTenLoai(rows)
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
