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
