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

type VanBanDenRepository struct {
	db *pgxpool.Pool
}

func NewVanBanDenRepository(db *pgxpool.Pool) *VanBanDenRepository {
	return &VanBanDenRepository{db: db}
}

const vanBanDenColumns = `id, nam, so_den, ngay_den, noi_gui_id, noi_gui_text, so_ky_hieu,
	ngay_van_ban, trich_yeu, don_vi_xu_ly_id, don_vi_nhan_text, ky_nhan, ghi_chu,
	created_at, updated_at`

func scanVanBanDen(row pgx.Row) (*model.VanBanDen, error) {
	var (
		m          model.VanBanDen
		ngayDen    time.Time
		ngayVanBan *time.Time
	)

	err := row.Scan(
		&m.ID, &m.Nam, &m.SoDen, &ngayDen, &m.NoiGuiID, &m.NoiGuiText, &m.SoKyHieu,
		&ngayVanBan, &m.TrichYeu, &m.DonViXuLyID, &m.DonViNhanText, &m.KyNhan, &m.GhiChu,
		&m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	m.NgayDen = model.DateFromTime(ngayDen)
	m.NgayVanBan = model.DatePtrFromTimePtr(ngayVanBan)
	return &m, nil
}

func (r *VanBanDenRepository) Create(ctx context.Context, in *model.VanBanDenInput) (*model.VanBanDen, error) {
	query := `INSERT INTO van_ban_den
		(nam, so_den, ngay_den, noi_gui_id, noi_gui_text, so_ky_hieu, ngay_van_ban,
		 trich_yeu, don_vi_xu_ly_id, don_vi_nhan_text, ky_nhan, ghi_chu)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		RETURNING ` + vanBanDenColumns

	row := r.db.QueryRow(ctx, query,
		in.Nam, in.SoDen, in.NgayDen.Time, in.NoiGuiID, in.NoiGuiText, in.SoKyHieu, in.NgayVanBan.TimePtr(),
		in.TrichYeu, in.DonViXuLyID, in.DonViNhanText, in.KyNhan, in.GhiChu,
	)
	return scanVanBanDen(row)
}

func (r *VanBanDenRepository) GetByID(ctx context.Context, id int64) (*model.VanBanDen, error) {
	query := `SELECT ` + vanBanDenColumns + ` FROM van_ban_den WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)
	return scanVanBanDen(row)
}

func (r *VanBanDenRepository) Update(ctx context.Context, id int64, in *model.VanBanDenInput) (*model.VanBanDen, error) {
	query := `UPDATE van_ban_den SET
		nam = $1, so_den = $2, ngay_den = $3, noi_gui_id = $4, noi_gui_text = $5,
		so_ky_hieu = $6, ngay_van_ban = $7, trich_yeu = $8, don_vi_xu_ly_id = $9,
		don_vi_nhan_text = $10, ky_nhan = $11, ghi_chu = $12
		WHERE id = $13
		RETURNING ` + vanBanDenColumns

	row := r.db.QueryRow(ctx, query,
		in.Nam, in.SoDen, in.NgayDen.Time, in.NoiGuiID, in.NoiGuiText, in.SoKyHieu, in.NgayVanBan.TimePtr(),
		in.TrichYeu, in.DonViXuLyID, in.DonViNhanText, in.KyNhan, in.GhiChu, id,
	)
	return scanVanBanDen(row)
}

func (r *VanBanDenRepository) Delete(ctx context.Context, id int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM van_ban_den WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// vanBanDenSortColumns là whitelist cột được phép sắp xếp (chống SQL injection).
var vanBanDenSortColumns = map[string]string{
	"so_den":       "so_den",
	"ngay_den":     "ngay_den",
	"ngay_van_ban": "ngay_van_ban",
	"nam":          "nam",
	"created_at":   "created_at",
}

func (r *VanBanDenRepository) List(ctx context.Context, f model.VanBanDenFilter) ([]model.VanBanDen, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.Nam != nil {
		where = append(where, fmt.Sprintf("nam = $%d", argN))
		args = append(args, *f.Nam)
		argN++
	}
	if f.SoDen != nil {
		where = append(where, fmt.Sprintf("so_den = $%d", argN))
		args = append(args, *f.SoDen)
		argN++
	}
	if f.SoKyHieu != nil {
		where = append(where, fmt.Sprintf("so_ky_hieu ILIKE $%d", argN))
		args = append(args, "%"+*f.SoKyHieu+"%")
		argN++
	}
	if f.NoiGuiID != nil {
		where = append(where, fmt.Sprintf("noi_gui_id = $%d", argN))
		args = append(args, *f.NoiGuiID)
		argN++
	}
	if f.DonViXuLyID != nil {
		where = append(where, fmt.Sprintf("don_vi_xu_ly_id = $%d", argN))
		args = append(args, *f.DonViXuLyID)
		argN++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("trich_yeu ILIKE $%d", argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	orderBy, err := buildOrderBy(vanBanDenSortColumns, f.SortBy, f.SortDir, "nam DESC, so_den DESC")
	if err != nil {
		return nil, 0, err
	}

	var total int64
	countQuery := `SELECT count(*) FROM van_ban_den WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + vanBanDenColumns + ` FROM van_ban_den WHERE ` + whereClause +
		" ORDER BY " + orderBy +
		fmt.Sprintf(" LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.VanBanDen, 0)
	for rows.Next() {
		m, err := scanVanBanDen(rows)
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
