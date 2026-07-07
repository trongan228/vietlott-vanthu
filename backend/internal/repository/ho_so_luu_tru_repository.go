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

type HoSoLuuTruRepository struct {
	db *pgxpool.Pool
}

func NewHoSoLuuTruRepository(db *pgxpool.Pool) *HoSoLuuTruRepository {
	return &HoSoLuuTruRepository{db: db}
}

const hoSoLuuTruColumns = `id, so_ho_so, hop_id, tieu_de, loai_tap, nam, tu_ngay, den_ngay,
	tu_so, den_so, so_to, thoi_han_bao_quan_nam, vinh_vien, ghi_chu, created_at`

func scanHoSoLuuTru(row pgx.Row) (*model.HoSoLuuTru, error) {
	var (
		m       model.HoSoLuuTru
		tuNgay  *time.Time
		denNgay *time.Time
	)
	err := row.Scan(&m.ID, &m.SoHoSo, &m.HopID, &m.TieuDe, &m.LoaiTap, &m.Nam, &tuNgay, &denNgay,
		&m.TuSo, &m.DenSo, &m.SoTo, &m.ThoiHanBaoQuanNam, &m.VinhVien, &m.GhiChu, &m.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	m.TuNgay = model.DatePtrFromTimePtr(tuNgay)
	m.DenNgay = model.DatePtrFromTimePtr(denNgay)
	return &m, nil
}

func (r *HoSoLuuTruRepository) Create(ctx context.Context, in *model.HoSoLuuTruInput) (*model.HoSoLuuTru, error) {
	query := `INSERT INTO ho_so_luu_tru
		(so_ho_so, hop_id, tieu_de, loai_tap, nam, tu_ngay, den_ngay, tu_so, den_so, so_to,
		 thoi_han_bao_quan_nam, vinh_vien, ghi_chu)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
		RETURNING ` + hoSoLuuTruColumns

	row := r.db.QueryRow(ctx, query,
		in.SoHoSo, in.HopID, in.TieuDe, in.LoaiTap, in.Nam, in.TuNgay.TimePtr(), in.DenNgay.TimePtr(),
		in.TuSo, in.DenSo, in.SoTo, in.ThoiHanBaoQuanNam, *in.VinhVien, in.GhiChu,
	)
	return scanHoSoLuuTru(row)
}

func (r *HoSoLuuTruRepository) GetByID(ctx context.Context, id int32) (*model.HoSoLuuTru, error) {
	query := `SELECT ` + hoSoLuuTruColumns + ` FROM ho_so_luu_tru WHERE id = $1`
	return scanHoSoLuuTru(r.db.QueryRow(ctx, query, id))
}

func (r *HoSoLuuTruRepository) Update(ctx context.Context, id int32, in *model.HoSoLuuTruInput) (*model.HoSoLuuTru, error) {
	query := `UPDATE ho_so_luu_tru SET
		so_ho_so = $1, hop_id = $2, tieu_de = $3, loai_tap = $4, nam = $5, tu_ngay = $6,
		den_ngay = $7, tu_so = $8, den_so = $9, so_to = $10, thoi_han_bao_quan_nam = $11,
		vinh_vien = $12, ghi_chu = $13
		WHERE id = $14
		RETURNING ` + hoSoLuuTruColumns

	row := r.db.QueryRow(ctx, query,
		in.SoHoSo, in.HopID, in.TieuDe, in.LoaiTap, in.Nam, in.TuNgay.TimePtr(), in.DenNgay.TimePtr(),
		in.TuSo, in.DenSo, in.SoTo, in.ThoiHanBaoQuanNam, *in.VinhVien, in.GhiChu, id,
	)
	return scanHoSoLuuTru(row)
}

func (r *HoSoLuuTruRepository) Delete(ctx context.Context, id int32) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM ho_so_luu_tru WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *HoSoLuuTruRepository) List(ctx context.Context, f model.HoSoLuuTruFilter) ([]model.HoSoLuuTru, int64, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if f.HopID != nil {
		where = append(where, fmt.Sprintf("hop_id = $%d", argN))
		args = append(args, *f.HopID)
		argN++
	}
	if f.LoaiTap != nil {
		where = append(where, fmt.Sprintf("loai_tap = $%d", argN))
		args = append(args, *f.LoaiTap)
		argN++
	}
	if f.Nam != nil {
		where = append(where, fmt.Sprintf("nam = $%d", argN))
		args = append(args, *f.Nam)
		argN++
	}
	if f.VinhVien != nil {
		where = append(where, fmt.Sprintf("vinh_vien = $%d", argN))
		args = append(args, *f.VinhVien)
		argN++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("tieu_de ILIKE $%d", argN))
		args = append(args, "%"+f.Search+"%")
		argN++
	}
	whereClause := strings.Join(where, " AND ")

	var total int64
	countQuery := `SELECT count(*) FROM ho_so_luu_tru WHERE ` + whereClause
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	listArgs := append(append([]any{}, args...), f.PageSize, (f.Page-1)*f.PageSize)
	query := `SELECT ` + hoSoLuuTruColumns + ` FROM ho_so_luu_tru WHERE ` + whereClause +
		fmt.Sprintf(" ORDER BY so_ho_so LIMIT $%d OFFSET $%d", argN, argN+1)

	rows, err := r.db.Query(ctx, query, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := make([]model.HoSoLuuTru, 0)
	for rows.Next() {
		m, err := scanHoSoLuuTru(rows)
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
