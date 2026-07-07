package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"vanthu-backend/internal/model"
)

type DonViRepository struct {
	db *pgxpool.Pool
}

func NewDonViRepository(db *pgxpool.Pool) *DonViRepository {
	return &DonViRepository{db: db}
}

const donViColumns = `id, ma_don_vi, ten_don_vi, loai_don_vi, dia_chi, ghi_chu, is_active, created_at, updated_at`

func scanDonVi(row pgx.Row) (*model.DonVi, error) {
	var m model.DonVi
	err := row.Scan(&m.ID, &m.MaDonVi, &m.TenDonVi, &m.LoaiDonVi, &m.DiaChi, &m.GhiChu,
		&m.IsActive, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}

func (r *DonViRepository) Create(ctx context.Context, in *model.DonViInput) (*model.DonVi, error) {
	query := `INSERT INTO don_vi (ma_don_vi, ten_don_vi, loai_don_vi, dia_chi, ghi_chu, is_active)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING ` + donViColumns

	row := r.db.QueryRow(ctx, query, in.MaDonVi, in.TenDonVi, in.LoaiDonVi, in.DiaChi, in.GhiChu, *in.IsActive)
	return scanDonVi(row)
}

func (r *DonViRepository) GetByID(ctx context.Context, id int32) (*model.DonVi, error) {
	query := `SELECT ` + donViColumns + ` FROM don_vi WHERE id = $1`
	return scanDonVi(r.db.QueryRow(ctx, query, id))
}

// Update tự set updated_at = now() vì bảng don_vi không có trigger cập nhật.
func (r *DonViRepository) Update(ctx context.Context, id int32, in *model.DonViInput) (*model.DonVi, error) {
	query := `UPDATE don_vi SET
		ma_don_vi = $1, ten_don_vi = $2, loai_don_vi = $3, dia_chi = $4, ghi_chu = $5,
		is_active = $6, updated_at = now()
		WHERE id = $7
		RETURNING ` + donViColumns

	row := r.db.QueryRow(ctx, query, in.MaDonVi, in.TenDonVi, in.LoaiDonVi, in.DiaChi, in.GhiChu, *in.IsActive, id)
	return scanDonVi(row)
}

func (r *DonViRepository) Delete(ctx context.Context, id int32) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM don_vi WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *DonViRepository) List(ctx context.Context, f model.DonViFilter) ([]model.DonVi, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.LoaiDonVi != nil {
		where = append(where, fmt.Sprintf("loai_don_vi = $%d", argN))
		args = append(args, *f.LoaiDonVi)
		argN++
	}
	if f.IsActive != nil {
		where = append(where, fmt.Sprintf("is_active = $%d", argN))
		args = append(args, *f.IsActive)
		argN++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("(ten_don_vi ILIKE $%d OR ma_don_vi ILIKE $%d)", argN, argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	var total int64
	countQuery := `SELECT count(*) FROM don_vi WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + donViColumns + ` FROM don_vi WHERE ` + whereClause +
		fmt.Sprintf(" ORDER BY ten_don_vi LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.DonVi, 0)
	for rows.Next() {
		m, err := scanDonVi(rows)
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
