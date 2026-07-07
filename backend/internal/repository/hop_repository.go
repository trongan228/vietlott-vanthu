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

type HopRepository struct {
	db *pgxpool.Pool
}

func NewHopRepository(db *pgxpool.Pool) *HopRepository {
	return &HopRepository{db: db}
}

const hopColumns = `id, so_hop, loai_hop, thung_id, ghi_chu, created_at`

func scanHop(row pgx.Row) (*model.Hop, error) {
	var m model.Hop
	err := row.Scan(&m.ID, &m.SoHop, &m.LoaiHop, &m.ThungID, &m.GhiChu, &m.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}

func (r *HopRepository) Create(ctx context.Context, in *model.HopInput) (*model.Hop, error) {
	query := `INSERT INTO hop (so_hop, loai_hop, thung_id, ghi_chu)
		VALUES ($1,$2,$3,$4)
		RETURNING ` + hopColumns

	row := r.db.QueryRow(ctx, query, in.SoHop, in.LoaiHop, in.ThungID, in.GhiChu)
	return scanHop(row)
}

func (r *HopRepository) GetByID(ctx context.Context, id int32) (*model.Hop, error) {
	query := `SELECT ` + hopColumns + ` FROM hop WHERE id = $1`
	return scanHop(r.db.QueryRow(ctx, query, id))
}

func (r *HopRepository) Update(ctx context.Context, id int32, in *model.HopInput) (*model.Hop, error) {
	query := `UPDATE hop SET
		so_hop = $1, loai_hop = $2, thung_id = $3, ghi_chu = $4
		WHERE id = $5
		RETURNING ` + hopColumns

	row := r.db.QueryRow(ctx, query, in.SoHop, in.LoaiHop, in.ThungID, in.GhiChu, id)
	return scanHop(row)
}

func (r *HopRepository) Delete(ctx context.Context, id int32) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM hop WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *HopRepository) List(ctx context.Context, f model.HopFilter) ([]model.Hop, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.SoHop != nil {
		where = append(where, fmt.Sprintf("so_hop = $%d", argN))
		args = append(args, *f.SoHop)
		argN++
	}
	if f.LoaiHop != nil {
		where = append(where, fmt.Sprintf("loai_hop = $%d", argN))
		args = append(args, *f.LoaiHop)
		argN++
	}
	if f.ThungID != nil {
		where = append(where, fmt.Sprintf("thung_id = $%d", argN))
		args = append(args, *f.ThungID)
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	var total int64
	countQuery := `SELECT count(*) FROM hop WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + hopColumns + ` FROM hop WHERE ` + whereClause +
		fmt.Sprintf(" ORDER BY loai_hop, so_hop LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.Hop, 0)
	for rows.Next() {
		m, err := scanHop(rows)
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
