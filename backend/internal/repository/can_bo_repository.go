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

type CanBoRepository struct {
	db *pgxpool.Pool
}

func NewCanBoRepository(db *pgxpool.Pool) *CanBoRepository {
	return &CanBoRepository{db: db}
}

const canBoColumns = `id, ho_ten, chuc_danh, don_vi_id, so_dien_thoai, email, is_van_thu, is_active, ghi_chu, created_at, updated_at`

func scanCanBo(row pgx.Row) (*model.CanBo, error) {
	var m model.CanBo
	err := row.Scan(&m.ID, &m.HoTen, &m.ChucDanh, &m.DonViID, &m.SoDienThoai, &m.Email,
		&m.IsVanThu, &m.IsActive, &m.GhiChu, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}

func (r *CanBoRepository) Create(ctx context.Context, in *model.CanBoInput) (*model.CanBo, error) {
	query := `INSERT INTO can_bo (ho_ten, chuc_danh, don_vi_id, so_dien_thoai, email, is_van_thu, is_active, ghi_chu)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING ` + canBoColumns

	row := r.db.QueryRow(ctx, query, in.HoTen, in.ChucDanh, in.DonViID, in.SoDienThoai, in.Email,
		*in.IsVanThu, *in.IsActive, in.GhiChu)
	return scanCanBo(row)
}

func (r *CanBoRepository) GetByID(ctx context.Context, id int32) (*model.CanBo, error) {
	query := `SELECT ` + canBoColumns + ` FROM can_bo WHERE id = $1`
	return scanCanBo(r.db.QueryRow(ctx, query, id))
}

// Update tự set updated_at = now() vì bảng can_bo không có trigger cập nhật.
func (r *CanBoRepository) Update(ctx context.Context, id int32, in *model.CanBoInput) (*model.CanBo, error) {
	query := `UPDATE can_bo SET
		ho_ten = $1, chuc_danh = $2, don_vi_id = $3, so_dien_thoai = $4, email = $5,
		is_van_thu = $6, is_active = $7, ghi_chu = $8, updated_at = now()
		WHERE id = $9
		RETURNING ` + canBoColumns

	row := r.db.QueryRow(ctx, query, in.HoTen, in.ChucDanh, in.DonViID, in.SoDienThoai, in.Email,
		*in.IsVanThu, *in.IsActive, in.GhiChu, id)
	return scanCanBo(row)
}

func (r *CanBoRepository) Delete(ctx context.Context, id int32) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM can_bo WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *CanBoRepository) List(ctx context.Context, f model.CanBoFilter) ([]model.CanBo, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.DonViID != nil {
		where = append(where, fmt.Sprintf("don_vi_id = $%d", argN))
		args = append(args, *f.DonViID)
		argN++
	}
	if f.IsVanThu != nil {
		where = append(where, fmt.Sprintf("is_van_thu = $%d", argN))
		args = append(args, *f.IsVanThu)
		argN++
	}
	if f.IsActive != nil {
		where = append(where, fmt.Sprintf("is_active = $%d", argN))
		args = append(args, *f.IsActive)
		argN++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("ho_ten ILIKE $%d", argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	var total int64
	countQuery := `SELECT count(*) FROM can_bo WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + canBoColumns + ` FROM can_bo WHERE ` + whereClause +
		fmt.Sprintf(" ORDER BY ho_ten LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.CanBo, 0)
	for rows.Next() {
		m, err := scanCanBo(rows)
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
