// API client tập trung cho backend Go (quản lý văn thư).
// Dev: gọi đường dẫn tương đối /api/v1/... -> Vite proxy chuyển sang VITE_API_URL
// (backend chưa bật CORS). Production: gọi thẳng VITE_API_URL.

const BASE: string = import.meta.env.DEV ? '' : (import.meta.env.VITE_API_URL ?? '');

export class ApiError extends Error {
  status: number;
  url: string;

  constructor(message: string, status: number, url: string) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
    this.url = url;
  }
}

export interface Pagination {
  page: number;
  page_size: number;
  total: number;
}

export interface Paginated<T> {
  items: T[];
  pagination: Pagination;
}

export interface VanBanDenDto {
  id: number;
  nam: number;
  so_den: number;
  ngay_den: string | null;
  noi_gui_id: number | null;
  noi_gui_text: string | null;
  so_ky_hieu: string | null;
  ngay_van_ban: string | null;
  trich_yeu: string | null;
  don_vi_xu_ly_id: number | null;
  don_vi_nhan_text: string | null;
  ky_nhan: string | null;
  ghi_chu: string | null;
  created_at: string;
  updated_at: string;
}

export interface VanBanDiDto {
  id: number;
  nam: number;
  loai_van_ban_id?: number | null; // chỉ có ở văn bản đi có tên loại
  so_ky_hieu: string | null;
  ngay_van_ban: string | null;
  trich_yeu: string | null;
  nguoi_ky_id: number | null;
  nguoi_ky_text: string | null;
  noi_nhan_text: string | null;
  so_luong_ban: number | null;
  ghi_chu: string | null;
  created_at: string;
  updated_at: string;
}

// Danh sách năm cho bộ lọc: từ năm sớm nhất trong dữ liệu (2017) đến năm hiện
// tại, giảm dần. Tính động để không phải sửa tay mỗi năm.
const FILTER_START_YEAR = 2017;
export function filterYears(): number[] {
  const end = new Date().getFullYear();
  const years: number[] = [];
  for (let y = Math.max(end, FILTER_START_YEAR); y >= FILTER_START_YEAR; y--) years.push(y);
  return years;
}

export interface ListParams {
  page?: number;
  pageSize?: number;
  nam?: number | string;
  q?: string;
}

async function fetchJson<T>(path: string, params: Record<string, unknown> = {}, signal?: AbortSignal): Promise<T> {
  const url = new URL(BASE + path, window.location.origin);
  for (const [k, v] of Object.entries(params)) {
    if (v !== undefined && v !== null && v !== '') url.searchParams.set(k, String(v));
  }
  let res: Response;
  try {
    res = await fetch(url, { signal, headers: { Accept: 'application/json' } });
  } catch (err) {
    if (err instanceof DOMException && err.name === 'AbortError') throw err;
    throw new ApiError('Không kết nối được máy chủ. Kiểm tra backend đang chạy.', 0, url.toString());
  }
  if (!res.ok) {
    throw new ApiError(`Máy chủ trả về lỗi ${res.status}`, res.status, url.toString());
  }
  return res.json() as Promise<T>;
}

function toListQuery(p: ListParams) {
  return { page: p.page, page_size: p.pageSize, nam: p.nam, q: p.q };
}

export function getVanBanDen(p: ListParams = {}, signal?: AbortSignal) {
  return fetchJson<Paginated<VanBanDenDto>>('/api/v1/van-ban-den', toListQuery(p), signal);
}

export function getVanBanDiCoTenLoai(p: ListParams = {}, signal?: AbortSignal) {
  return fetchJson<Paginated<VanBanDiDto>>('/api/v1/van-ban-di-co-ten-loai', toListQuery(p), signal);
}

export function getVanBanDiKhongTenLoai(p: ListParams = {}, signal?: AbortSignal) {
  return fetchJson<Paginated<VanBanDiDto>>('/api/v1/van-ban-di-khong-ten-loai', toListQuery(p), signal);
}

// ---------- Tra cứu lưu kho (văn bản → hồ sơ → hộp → thùng) ----------

export interface ViTriLuuKhoDto {
  ho_so_id: number;
  so_ho_so: number;
  ho_so_tieu_de: string;
  so_hop: number;
  loai_hop: 'thoi_han' | 'vinh_vien';
  ma_thung: string | null; // NULL nếu hộp chưa xếp vào thùng
  so_serial: string | null;
}

// Kết quả GET /van-ban/tim-kiem: văn bản khớp trên cả 3 bảng, kèm vị trí
// lưu kho (vi_tri) nếu văn bản đã được gán hồ sơ lưu trữ.
export interface TimKiemVanBanDto {
  van_ban_den: (VanBanDenDto & { vi_tri: ViTriLuuKhoDto | null })[];
  van_ban_di_co_ten_loai: (VanBanDiDto & { vi_tri: ViTriLuuKhoDto | null })[];
  van_ban_di_khong_ten_loai: (VanBanDiDto & { vi_tri: ViTriLuuKhoDto | null })[];
}

export function timKiemVanBan(q: string, limit = 10, signal?: AbortSignal) {
  return fetchJson<TimKiemVanBanDto>('/api/v1/van-ban/tim-kiem', { q, limit }, signal);
}

export interface ThungDto {
  id: number;
  ma_thung: string;
  so_serial: string | null;
  dot_luu_kho: number | null;
  vi_tri_kho: string | null;
  ngay_nhap_kho: string | null;
  ghi_chu: string | null;
  created_at: string;
}

// Tra chi tiết thùng theo mã (GET /thung?q=...) để lấy vị trí kho, đợt lưu
// kho, ngày nhập kho. Trả null nếu không có thùng khớp đúng mã.
export async function getThungByMa(maThung: string, signal?: AbortSignal): Promise<ThungDto | null> {
  const res = await fetchJson<Paginated<ThungDto>>('/api/v1/thung', { q: maThung, page_size: 5 }, signal);
  return res.items.find(t => t.ma_thung === maThung) ?? null;
}

// ---------- Chuyển đổi DTO -> shape màn hình (theo thiết kế) ----------

const EMPTY = '—';

export function formatDate(iso: string | null): string {
  if (!iso) return EMPTY;
  const m = /^(\d{4})-(\d{2})-(\d{2})/.exec(iso);
  return m ? `${m[3]}/${m[2]}/${m[1]}` : iso;
}

// Không có endpoint tra tên loại: suy ra từ tiền tố số ký hiệu (VD "79/QĐ-VIETLOTT" -> Quyết định).
const LOAI_BY_PREFIX: Record<string, string> = {
  'QĐ': 'Quyết định', 'QD': 'Quyết định',
  'TB': 'Thông báo',
  'TTR': 'Tờ trình',
  'BC': 'Báo cáo',
  'UQ': 'Ủy quyền',
  'CT': 'Chỉ thị',
  'GM': 'Giấy mời',
  'CV': 'Công văn',
};
export function loaiFromSoKyHieu(soKyHieu: string | null): string {
  if (!soKyHieu) return EMPTY;
  const m = /^\s*\d+[\/.]?\s*([A-Za-zĐđ]+)/.exec(soKyHieu);
  if (!m) return EMPTY;
  return LOAI_BY_PREFIX[m[1].toUpperCase()] ?? EMPTY;
}

export interface VanBanDenRow {
  id: number;
  soDen: string;
  ngayDen: string;
  noiGui: string;
  soKyHieu: string;
  ngayVB: string;
  trichYeu: string;
  donViNhan: string;
  nguoiKyNhan: string;
  ghiChu: string;
}

export function mapVanBanDen(d: VanBanDenDto): VanBanDenRow {
  return {
    id: d.id,
    soDen: String(d.so_den),
    ngayDen: formatDate(d.ngay_den),
    noiGui: d.noi_gui_text ?? EMPTY,
    soKyHieu: d.so_ky_hieu ?? EMPTY,
    ngayVB: formatDate(d.ngay_van_ban),
    trichYeu: d.trich_yeu ?? EMPTY,
    donViNhan: d.don_vi_nhan_text ?? EMPTY,
    nguoiKyNhan: d.ky_nhan ?? EMPTY,
    ghiChu: d.ghi_chu ?? '',
  };
}

export interface VanBanDiRow {
  id: number;
  soKyHieu: string;
  loaiVB?: string;
  ngayVB: string;
  trichYeu: string;
  nguoiKy: string;
  noiNhan: string;
  soLuong: string;
  ghiChu: string;
}

export function mapVanBanDi(d: VanBanDiDto, coTenLoai: boolean): VanBanDiRow {
  return {
    id: d.id,
    soKyHieu: d.so_ky_hieu ?? EMPTY,
    ...(coTenLoai ? { loaiVB: loaiFromSoKyHieu(d.so_ky_hieu) } : {}),
    ngayVB: formatDate(d.ngay_van_ban),
    trichYeu: d.trich_yeu ?? EMPTY,
    nguoiKy: d.nguoi_ky_text ?? EMPTY,
    noiNhan: d.noi_nhan_text ?? EMPTY,
    soLuong: d.so_luong_ban != null ? String(d.so_luong_ban) : EMPTY,
    ghiChu: d.ghi_chu ?? '',
  };
}

// Một kết quả tra cứu lưu kho (đã gộp 3 bảng văn bản về cùng shape).
export interface TraCuuRow {
  key: string;        // duy nhất trên cả 3 bảng: `${loai}-${id}`
  loaiLabel: string;  // Văn bản đến / Văn bản đi...
  den: boolean;
  tieuDe: string;     // số ký hiệu, hoặc "Số đến N/năm" nếu không có
  ngay: string;
  trichYeu: string;
  info: { k: string; v: string; strong?: boolean }[];
  viTri: ViTriLuuKhoDto | null;
}

export function mapTraCuu(data: TimKiemVanBanDto): TraCuuRow[] {
  const rows: TraCuuRow[] = [];

  for (const d of data.van_ban_den) {
    rows.push({
      key: `den-${d.id}`,
      loaiLabel: 'Văn bản đến',
      den: true,
      tieuDe: d.so_ky_hieu || `Số đến ${d.so_den}/${d.nam}`,
      ngay: formatDate(d.ngay_den),
      trichYeu: d.trich_yeu ?? EMPTY,
      info: [
        { k: 'Số đến', v: `${d.so_den} (năm ${d.nam})`, strong: true },
        { k: 'Ngày đến', v: formatDate(d.ngay_den) },
        { k: 'Nơi gửi', v: d.noi_gui_text ?? EMPTY },
        { k: 'Số ký hiệu', v: d.so_ky_hieu ?? EMPTY, strong: true },
        { k: 'Ngày văn bản', v: formatDate(d.ngay_van_ban) },
        { k: 'Trích yếu', v: d.trich_yeu ?? EMPTY },
        { k: 'Đơn vị nhận', v: d.don_vi_nhan_text ?? EMPTY },
        { k: 'Người ký nhận', v: d.ky_nhan ?? EMPTY },
      ],
      viTri: d.vi_tri,
    });
  }

  const pushDi = (items: TimKiemVanBanDto['van_ban_di_co_ten_loai'], loaiLabel: string, prefix: string) => {
    for (const d of items) {
      rows.push({
        key: `${prefix}-${d.id}`,
        loaiLabel,
        den: false,
        tieuDe: d.so_ky_hieu ?? EMPTY,
        ngay: formatDate(d.ngay_van_ban),
        trichYeu: d.trich_yeu ?? EMPTY,
        info: [
          { k: 'Số ký hiệu', v: d.so_ky_hieu ?? EMPTY, strong: true },
          { k: 'Năm', v: String(d.nam) },
          { k: 'Ngày văn bản', v: formatDate(d.ngay_van_ban) },
          { k: 'Trích yếu', v: d.trich_yeu ?? EMPTY },
          { k: 'Người ký', v: d.nguoi_ky_text ?? EMPTY },
          { k: 'Nơi nhận', v: d.noi_nhan_text ?? EMPTY },
        ],
        viTri: d.vi_tri,
      });
    }
  };
  pushDi(data.van_ban_di_co_ten_loai, 'Văn bản đi (có tên loại)', 'di-ctl');
  pushDi(data.van_ban_di_khong_ten_loai, 'Văn bản đi (không tên loại)', 'di-ktl');

  return rows;
}
