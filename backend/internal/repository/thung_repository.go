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

type ThungRepository struct {
	db *pgxpool.Pool
}

func NewThungRepository(db *pgxpool.Pool) *ThungRepository {
	return &ThungRepository{db: db}
}

const thungColumns = `id, ma_thung, so_serial, dot_luu_kho, vi_tri_kho, ngay_nhap_kho, ghi_chu, created_at`

func scanThung(row pgx.Row) (*model.Thung, error) {
	var (
		m           model.Thung
		ngayNhapKho *time.Time
	)
	err := row.Scan(&m.ID, &m.MaThung, &m.SoSerial, &m.DotLuuKho, &m.ViTriKho,
		&ngayNhapKho, &m.GhiChu, &m.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	m.NgayNhapKho = model.DatePtrFromTimePtr(ngayNhapKho)
	return &m, nil
}

func (r *ThungRepository) Create(ctx context.Context, in *model.ThungInput) (*model.Thung, error) {
	query := `INSERT INTO thung (ma_thung, so_serial, dot_luu_kho, vi_tri_kho, ngay_nhap_kho, ghi_chu)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING ` + thungColumns

	row := r.db.QueryRow(ctx, query, in.MaThung, in.SoSerial, in.DotLuuKho, in.ViTriKho,
		in.NgayNhapKho.TimePtr(), in.GhiChu)
	return scanThung(row)
}

func (r *ThungRepository) GetByID(ctx context.Context, id int32) (*model.Thung, error) {
	query := `SELECT ` + thungColumns + ` FROM thung WHERE id = $1`
	return scanThung(r.db.QueryRow(ctx, query, id))
}

func (r *ThungRepository) Update(ctx context.Context, id int32, in *model.ThungInput) (*model.Thung, error) {
	query := `UPDATE thung SET
		ma_thung = $1, so_serial = $2, dot_luu_kho = $3, vi_tri_kho = $4, ngay_nhap_kho = $5, ghi_chu = $6
		WHERE id = $7
		RETURNING ` + thungColumns

	row := r.db.QueryRow(ctx, query, in.MaThung, in.SoSerial, in.DotLuuKho, in.ViTriKho,
		in.NgayNhapKho.TimePtr(), in.GhiChu, id)
	return scanThung(row)
}

func (r *ThungRepository) Delete(ctx context.Context, id int32) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM thung WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *ThungRepository) List(ctx context.Context, f model.ThungFilter) ([]model.Thung, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.DotLuuKho != nil {
		where = append(where, fmt.Sprintf("dot_luu_kho = $%d", argN))
		args = append(args, *f.DotLuuKho)
		argN++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("(ma_thung ILIKE $%d OR so_serial ILIKE $%d)", argN, argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	var total int64
	countQuery := `SELECT count(*) FROM thung WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + thungColumns + ` FROM thung WHERE ` + whereClause +
		fmt.Sprintf(" ORDER BY ma_thung LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.Thung, 0)
	for rows.Next() {
		m, err := scanThung(rows)
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
