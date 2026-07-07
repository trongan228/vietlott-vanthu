package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"vanthu-backend/internal/handler"
	"vanthu-backend/internal/repository"
	"vanthu-backend/internal/service"
)

// New khởi tạo router Gin và wire đủ 3 layer (repository → service → handler)
// cho toàn bộ bảng nghiệp vụ (văn bản đến/đi), danh mục (loại văn bản, đơn vị,
// cán bộ) và lưu kho (thùng, hộp, hồ sơ lưu trữ) + tra cứu vị trí văn bản.
func New(pool *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	vbdHandler := handler.NewVanBanDenHandler(
		service.NewVanBanDenService(
			repository.NewVanBanDenRepository(pool),
		),
	)
	vbdktlHandler := handler.NewVanBanDiKhongTenLoaiHandler(
		service.NewVanBanDiKhongTenLoaiService(
			repository.NewVanBanDiKhongTenLoaiRepository(pool),
		),
	)
	vbdctlHandler := handler.NewVanBanDiCoTenLoaiHandler(
		service.NewVanBanDiCoTenLoaiService(
			repository.NewVanBanDiCoTenLoaiRepository(pool),
		),
	)
	loaiVanBanHandler := handler.NewLoaiVanBanHandler(
		service.NewLoaiVanBanService(
			repository.NewLoaiVanBanRepository(pool),
		),
	)
	donViHandler := handler.NewDonViHandler(
		service.NewDonViService(
			repository.NewDonViRepository(pool),
		),
	)
	canBoHandler := handler.NewCanBoHandler(
		service.NewCanBoService(
			repository.NewCanBoRepository(pool),
		),
	)
	thungHandler := handler.NewThungHandler(
		service.NewThungService(
			repository.NewThungRepository(pool),
		),
	)
	hopHandler := handler.NewHopHandler(
		service.NewHopService(
			repository.NewHopRepository(pool),
		),
	)
	hoSoLuuTruHandler := handler.NewHoSoLuuTruHandler(
		service.NewHoSoLuuTruService(
			repository.NewHoSoLuuTruRepository(pool),
		),
	)
	viTriVanBanHandler := handler.NewViTriVanBanHandler(
		service.NewViTriVanBanService(
			repository.NewViTriVanBanRepository(pool),
		),
	)
	timKiemVanBanHandler := handler.NewTimKiemVanBanHandler(
		service.NewTimKiemVanBanService(
			repository.NewTimKiemVanBanRepository(pool),
		),
	)

	api := r.Group("/api/v1")
	vbdHandler.Register(api)
	vbdktlHandler.Register(api)
	vbdctlHandler.Register(api)
	loaiVanBanHandler.Register(api)
	donViHandler.Register(api)
	canBoHandler.Register(api)
	thungHandler.Register(api)
	hopHandler.Register(api)
	hoSoLuuTruHandler.Register(api)
	viTriVanBanHandler.Register(api)
	timKiemVanBanHandler.Register(api)

	return r
}
