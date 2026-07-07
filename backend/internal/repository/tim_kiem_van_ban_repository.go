package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"vanthu-backend/internal/model"
)

// TimKiemVanBanRepository tìm kiếm hợp nhất theo tên (trích yếu) và
// số văn bản trên cả 3 bảng văn bản.
type TimKiemVanBanRepository struct {
	db *pgxpool.Pool
}

func NewTimKiemVanBanRepository(db *pgxpool.Pool) *TimKiemVanBanRepository {
	return &TimKiemVanBanRepository{db: db}
}

// dieuKienChung dựng phần WHERE dùng chung (q → trich_yeu, nam) rồi trả về
// mảng điều kiện + args + chỉ số tham số kế tiếp.
func dieuKienChung(p model.TimKiemVanBanParams) ([]string, []any, int) {
	where := []string{"1=1"}
	args := []any{}
	argN := 1

	if p.Q != "" {
		where = append(where, fmt.Sprintf("trich_yeu ILIKE $%d", argN))
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

func (r *TimKiemVanBanRepository) TimVanBanDen(ctx context.Context, p model.TimKiemVanBanParams) ([]model.VanBanDen, error) {
	where, args, argN := dieuKienChung(p)

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
	where, args, argN := dieuKienChung(p)

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
	where, args, argN := dieuKienChung(p)

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
