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

type LoaiVanBanRepository struct {
	db *pgxpool.Pool
}

func NewLoaiVanBanRepository(db *pgxpool.Pool) *LoaiVanBanRepository {
	return &LoaiVanBanRepository{db: db}
}

const loaiVanBanColumns = `id, ma_loai, ten_loai, thoi_han_bao_quan_nam, ghi_chu`

func scanLoaiVanBan(row pgx.Row) (*model.LoaiVanBan, error) {
	var m model.LoaiVanBan
	err := row.Scan(&m.ID, &m.MaLoai, &m.TenLoai, &m.ThoiHanBaoQuanNam, &m.GhiChu)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}

func (r *LoaiVanBanRepository) Create(ctx context.Context, in *model.LoaiVanBanInput) (*model.LoaiVanBan, error) {
	query := `INSERT INTO loai_van_ban (ma_loai, ten_loai, thoi_han_bao_quan_nam, ghi_chu)
		VALUES ($1,$2,$3,$4)
		RETURNING ` + loaiVanBanColumns

	row := r.db.QueryRow(ctx, query, in.MaLoai, in.TenLoai, in.ThoiHanBaoQuanNam, in.GhiChu)
	return scanLoaiVanBan(row)
}

func (r *LoaiVanBanRepository) GetByID(ctx context.Context, id int32) (*model.LoaiVanBan, error) {
	query := `SELECT ` + loaiVanBanColumns + ` FROM loai_van_ban WHERE id = $1`
	return scanLoaiVanBan(r.db.QueryRow(ctx, query, id))
}

func (r *LoaiVanBanRepository) Update(ctx context.Context, id int32, in *model.LoaiVanBanInput) (*model.LoaiVanBan, error) {
	query := `UPDATE loai_van_ban SET
		ma_loai = $1, ten_loai = $2, thoi_han_bao_quan_nam = $3, ghi_chu = $4
		WHERE id = $5
		RETURNING ` + loaiVanBanColumns

	row := r.db.QueryRow(ctx, query, in.MaLoai, in.TenLoai, in.ThoiHanBaoQuanNam, in.GhiChu, id)
	return scanLoaiVanBan(row)
}

func (r *LoaiVanBanRepository) Delete(ctx context.Context, id int32) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM loai_van_ban WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *LoaiVanBanRepository) List(ctx context.Context, f model.LoaiVanBanFilter) ([]model.LoaiVanBan, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.Search != "" {
		where = append(where, fmt.Sprintf("(ma_loai ILIKE $%d OR ten_loai ILIKE $%d)", argN, argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	var total int64
	countQuery := `SELECT count(*) FROM loai_van_ban WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + loaiVanBanColumns + ` FROM loai_van_ban WHERE ` + whereClause +
		fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.LoaiVanBan, 0)
	for rows.Next() {
		m, err := scanLoaiVanBan(rows)
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
