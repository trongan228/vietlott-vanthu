package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"vanthu-backend/internal/model"
)

// TimKiemVanBanRepository tìm kiếm hợp nhất theo tên (trích yếu), số văn bản
// và số đến trên cả 3 bảng văn bản.
type TimKiemVanBanRepository struct {
	db *pgxpool.Pool
}

func NewTimKiemVanBanRepository(db *pgxpool.Pool) *TimKiemVanBanRepository {
	return &TimKiemVanBanRepository{db: db}
}

// dieuKienChung dựng phần WHERE dùng chung rồi trả về mảng điều kiện + args
// + chỉ số tham số kế tiếp. q khớp gần đúng (OR) trên các cột trong qCols
// (tên/trích yếu, số văn bản, số đến... tùy bảng), nam khớp chính xác.
func dieuKienChung(p model.TimKiemVanBanParams, qCols ...string) ([]string, []any, int) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if p.Q != "" {
		conds := make([]string, len(qCols))
		for i, col := range qCols {
			conds[i] = fmt.Sprintf("%s ILIKE $%d", col, argN)
		}
		where = append(where, "("+strings.Join(conds, " OR ")+")")
		args = append(args, "%"+p.Q+"%")
		argN++
	}
	if p.Nam != nil {
		where = append(where, fmt.Sprintf("nam = $%d", argN))
		args = append(args, *p.Nam)
		argN++
	}
	return where, args, argN
}

// LayViTriLuuKho trả map ho_so_id → vị trí lưu kho (hồ sơ → hộp → thùng)
// cho các hồ sơ được truyền vào, dùng để gắn vi_tri vào kết quả tìm kiếm.
func (r *TimKiemVanBanRepository) LayViTriLuuKho(ctx context.Context, hoSoIDs []int32) (map[int32]model.ViTriLuuKho, error) {
	result := make(map[int32]model.ViTriLuuKho)
	if len(hoSoIDs) == 0 {
		return result, nil
	}

	query := `SELECT hs.id, hs.so_ho_so, hs.tieu_de, h.so_hop, h.loai_hop, t.ma_thung, t.so_serial
		FROM ho_so_luu_tru hs
		JOIN hop h ON h.id = hs.hop_id
		LEFT JOIN thung t ON t.id = h.thung_id
		WHERE hs.id = ANY($1)`

	rows, err := r.db.Query(ctx, query, hoSoIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v model.ViTriLuuKho
		if err := rows.Scan(&v.HoSoID, &v.SoHoSo, &v.HoSoTieuDe, &v.SoHop, &v.LoaiHop, &v.MaThung, &v.SoSerial); err != nil {
			return nil, err
		}
		result[v.HoSoID] = v
	}
	return result, rows.Err()
}

func (r *TimKiemVanBanRepository) TimVanBanDen(ctx context.Context, p model.TimKiemVanBanParams) ([]model.VanBanDen, error) {
	where, args, argN := dieuKienChung(p, "trich_yeu", "so_ky_hieu", "so_den::text")

	if p.So != "" {
		// Văn bản đến khớp theo so_ky_hieu hoặc số đến (so_den).
		where = append(where, fmt.Sprintf("(so_ky_hieu ILIKE $%d OR so_den::text ILIKE $%d)", argN, argN))
		args = append(args, "%"+p.So+"%")
		argN++
	}

	query := `SELECT ` + vanBanDenColumns + ` FROM van_ban_den WHERE ` + strings.Join(where, " AND ") +
		fmt.Sprintf(" ORDER BY nam DESC, so_den DESC LIMIT $%d", argN)
	args = append(args, p.Limit)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]model.VanBanDen, 0)
	for rows.Next() {
		m, err := scanVanBanDen(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *m)
	}
	return items, rows.Err()
}

func (r *TimKiemVanBanRepository) TimVanBanDiKhongTenLoai(ctx context.Context, p model.TimKiemVanBanParams) ([]model.VanBanDiKhongTenLoai, error) {
	where, args, argN := dieuKienChung(p, "trich_yeu", "so_ky_hieu")

	if p.So != "" {
		where = append(where, fmt.Sprintf("so_ky_hieu ILIKE $%d", argN))
		args = append(args, "%"+p.So+"%")
		argN++
	}

	query := `SELECT ` + vanBanDiKhongTenLoaiColumns + ` FROM van_ban_di_khong_ten_loai WHERE ` +
		strings.Join(where, " AND ") +
		fmt.Sprintf(" ORDER BY nam DESC, ngay_van_ban DESC LIMIT $%d", argN)
	args = append(args, p.Limit)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]model.VanBanDiKhongTenLoai, 0)
	for rows.Next() {
		m, err := scanVanBanDiKhongTenLoai(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *m)
	}
	return items, rows.Err()
}

func (r *TimKiemVanBanRepository) TimVanBanDiCoTenLoai(ctx context.Context, p model.TimKiemVanBanParams) ([]model.VanBanDiCoTenLoai, error) {
	where, args, argN := dieuKienChung(p, "trich_yeu", "so_ky_hieu")

	if p.So != "" {
		where = append(where, fmt.Sprintf("so_ky_hieu ILIKE $%d", argN))
		args = append(args, "%"+p.So+"%")
		argN++
	}

	query := `SELECT ` + vanBanDiCoTenLoaiColumns + ` FROM van_ban_di_co_ten_loai WHERE ` +
		strings.Join(where, " AND ") +
		fmt.Sprintf(" ORDER BY nam DESC, ngay_van_ban DESC LIMIT $%d", argN)
	args = append(args, p.Limit)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]model.VanBanDiCoTenLoai, 0)
	for rows.Next() {
		m, err := scanVanBanDiCoTenLoai(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *m)
	}
	return items, rows.Err()
}
