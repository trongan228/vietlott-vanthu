package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"vanthu-backend/internal/model"
)

// ViTriVanBanRepository đọc 3 view v_vi_tri_van_ban_* (chỉ đọc,
// dùng để tra cứu văn bản đang nằm ở hồ sơ/hộp/thùng nào).
type ViTriVanBanRepository struct {
	db *pgxpool.Pool
}

func NewViTriVanBanRepository(db *pgxpool.Pool) *ViTriVanBanRepository {
	return &ViTriVanBanRepository{db: db}
}

func (r *ViTriVanBanRepository) TraCuuVanBanDen(ctx context.Context, nam *int16, soDen *int32) ([]model.ViTriVanBanDen, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if nam != nil {
		where = append(where, fmt.Sprintf("nam = $%d", argN))
		args = append(args, *nam)
		argN++
	}
	if soDen != nil {
		where = append(where, fmt.Sprintf("so_den = $%d", argN))
		args = append(args, *soDen)
		argN++
	}

	query := `SELECT van_ban_id, nam, so_den, ngay_den, trich_yeu, so_ho_so, ho_so_tieu_de,
		so_hop, ma_thung, so_serial
		FROM v_vi_tri_van_ban_den WHERE ` + strings.Join(where, " AND ") +
		fmt.Sprintf(" ORDER BY nam DESC, so_den LIMIT $%d", argN)
	args = append(args, 200)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]model.ViTriVanBanDen, 0)
	for rows.Next() {
		var (
			m       model.ViTriVanBanDen
			ngayDen time.Time
		)
		if err := rows.Scan(&m.VanBanID, &m.Nam, &m.SoDen, &ngayDen, &m.TrichYeu,
			&m.SoHoSo, &m.HoSoTieuDe, &m.SoHop, &m.MaThung, &m.SoSerial); err != nil {
			return nil, err
		}
		m.NgayDen = model.DateFromTime(ngayDen)
		items = append(items, m)
	}
	return items, rows.Err()
}

// traCuuVanBanDi dùng chung cho 2 view văn bản đi (cùng cấu trúc cột).
// viewName chỉ nhận giá trị cố định từ code, không nhận từ người dùng.
func (r *ViTriVanBanRepository) traCuuVanBanDi(ctx context.Context, viewName string, nam *int16, soKyHieu *string) ([]model.ViTriVanBanDi, error) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if nam != nil {
		where = append(where, fmt.Sprintf("nam = $%d", argN))
		args = append(args, *nam)
		argN++
	}
	if soKyHieu != nil {
		where = append(where, fmt.Sprintf("so_ky_hieu = $%d", argN))
		args = append(args, *soKyHieu)
		argN++
	}

	query := `SELECT van_ban_id, nam, so_ky_hieu, ngay_van_ban, trich_yeu, so_ho_so, ho_so_tieu_de,
		so_hop, ma_thung, so_serial
		FROM ` + viewName + ` WHERE ` + strings.Join(where, " AND ") +
		fmt.Sprintf(" ORDER BY nam DESC, ngay_van_ban LIMIT $%d", argN)
	args = append(args, 200)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]model.ViTriVanBanDi, 0)
	for rows.Next() {
		var (
			m          model.ViTriVanBanDi
			ngayVanBan time.Time
		)
		if err := rows.Scan(&m.VanBanID, &m.Nam, &m.SoKyHieu, &ngayVanBan, &m.TrichYeu,
			&m.SoHoSo, &m.HoSoTieuDe, &m.SoHop, &m.MaThung, &m.SoSerial); err != nil {
			return nil, err
		}
		m.NgayVanBan = model.DateFromTime(ngayVanBan)
		items = append(items, m)
	}
	return items, rows.Err()
}

func (r *ViTriVanBanRepository) TraCuuVanBanDiCoTenLoai(ctx context.Context, nam *int16, soKyHieu *string) ([]model.ViTriVanBanDi, error) {
	return r.traCuuVanBanDi(ctx, "v_vi_tri_van_ban_di_co_ten_loai", nam, soKyHieu)
}

func (r *ViTriVanBanRepository) TraCuuVanBanDiKhongTenLoai(ctx context.Context, nam *int16, soKyHieu *string) ([]model.ViTriVanBanDi, error) {
	return r.traCuuVanBanDi(ctx, "v_vi_tri_van_ban_di_khong_ten_loai", nam, soKyHieu)
}
