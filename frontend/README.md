# Vietlott Văn Thư — Frontend

Ứng dụng quản lý văn thư (công văn đến/đi, tra cứu lưu kho, danh mục) cho Vietlott.
Xây dựng bằng **React 18 + Vite**, dữ liệu lấy từ **backend Go** (thư mục `../backend`) qua REST API.

## Công nghệ chính

- **React 18** — UI, dùng function component + hooks, không dùng router (điều hướng bằng state trong `App.jsx`).
- **Vite 5** — dev server & build tool.
- **TanStack Query (`@tanstack/react-query`)** — quản lý gọi API, cache, loading/error state cho các danh sách văn bản.
- **TypeScript** — chỉ dùng cho lớp gọi API/hook (`src/lib`, `src/hooks`), phần UI còn lại là `.jsx` thuần.

## Cách chạy

```bash
npm install
npm run dev       # dev server tại http://localhost:5173
npm run build      # build production vào dist/
npm run preview    # xem thử bản build
```

Cấu hình URL backend trong [.env](.env) qua biến `VITE_API_URL`. Khi dev, Vite proxy mọi request `/api/**`
sang `VITE_API_URL` (xem [vite.config.js](vite.config.js)) để tránh lỗi CORS vì backend Go chưa bật CORS.

## Cấu trúc thư mục

```
frontend/
├── index.html              # HTML gốc, mount điểm #root cho React
├── vite.config.js          # cấu hình Vite + proxy /api -> backend Go
├── .env                     # VITE_API_URL (URL backend)
├── design/                  # file thiết kế tĩnh (HTML) tham khảo giao diện gốc
└── src/
    ├── main.jsx             # entry point: tạo QueryClient, render <App/>
    ├── App.jsx              # component gốc: điều hướng màn hình (state, không router),
    │                        #   quản lý toast, drawer chi tiết, trạng thái tra cứu kho
    ├── data.js               # dữ liệu mẫu tĩnh cho Danh mục & Tra cứu lưu kho (chưa có API)
    ├── styles.css             # toàn bộ CSS thuần (class tiền tố "vl-")
    ├── lib/
    │   └── api.ts             # client gọi REST API backend Go: fetch, kiểu DTO,
    │                          #   hàm map DTO -> shape hiển thị (mapVanBanDen, mapVanBanDi...)
    ├── hooks/
    │   └── useVanBan.ts        # custom hooks TanStack Query (useVanBanDen,
    │                          #   useVanBanDiCoTenLoai, useVanBanDiKhongTenLoai)
    ├── components/            # thành phần UI dùng chung, không gắn với 1 màn hình cụ thể
    │   ├── Header.jsx          # thanh header trên cùng (logo, tìm kiếm, avatar người dùng)
    │   ├── NavBar.jsx          # thanh điều hướng giữa các màn hình chính
    │   ├── DetailDrawer.jsx    # drawer trượt hiển thị chi tiết 1 văn bản
    │   ├── TableState.jsx      # 3 trạng thái bảng dùng chung: Loading (skeleton), Error, Empty
    │   ├── Toast.jsx           # thông báo nổi góc màn hình
    │   └── Icons.jsx           # bộ icon SVG dùng chung toàn app
    └── screens/                # mỗi file là 1 màn hình chính, gắn với 1 mục trong NavBar
        ├── Dashboard.jsx        # Tổng quan: số liệu thống kê, biểu đồ, văn bản mới nhất
        ├── VanBanDen.jsx        # Danh sách văn bản đến (lọc, phân trang)
        ├── VanBanDi.jsx         # Danh sách văn bản đi (2 tab: có/không tên loại)
        ├── VanBanForm.jsx       # Form thêm văn bản đến/đi (chuyển đổi 3 loại bằng segmented control)
        ├── TraCuuLuuKho.jsx     # Tra cứu vị trí lưu kho vật lý theo số văn bản
        └── DanhMuc.jsx          # Danh mục hệ thống: Đơn vị / Cán bộ / Loại văn bản (3 tab)
```

## Luồng dữ liệu

1. `screens/*` gọi hook trong `hooks/useVanBan.ts`.
2. Hook gọi hàm fetch trong `lib/api.ts` (endpoint `/api/v1/van-ban-den`, `/api/v1/van-ban-di-co-ten-loai`,
   `/api/v1/van-ban-di-khong-ten-loai`), nhận DTO thô từ backend rồi map sang shape hiển thị (`VanBanDenRow`, `VanBanDiRow`).
3. TanStack Query lo cache, giữ dữ liệu trang trước khi chuyển trang (`keepPreviousData`) để bảng không nhấp nháy,
   và tự cung cấp trạng thái `isPending` / `isError` / `isFetching` cho UI (`components/TableState.jsx`).

**Lưu ý:** phần *Danh mục* (`DanhMuc.jsx`) và *Tra cứu lưu kho* (`TraCuuLuuKho.jsx`) hiện dùng dữ liệu mẫu tĩnh
trong `data.js`, backend chưa có endpoint tương ứng. Chức năng Xuất/Nhập Excel trong `App.jsx`
(`exportExcel`, `importExcel`) hiện chỉ hiện toast, chưa nối API thật.

## Điều hướng màn hình

Không dùng thư viện router — `App.jsx` giữ `screen` trong state (`dashboard | den | di | form | luukho | danhmuc`)
và render component tương ứng. `NavBar` gọi `onGo(key)` để đổi màn hình.
