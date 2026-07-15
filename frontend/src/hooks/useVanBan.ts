// Custom hook TanStack Query cho từng loại văn bản.
// Giữ dữ liệu trang trước khi chuyển trang (placeholderData) để bảng không nhấp nháy.
import { keepPreviousData, useQuery } from '@tanstack/react-query';
import {
  getThungByMa,
  getVanBanDen,
  getVanBanDiCoTenLoai,
  getVanBanDiKhongTenLoai,
  mapTraCuu,
  mapVanBanDen,
  mapVanBanDi,
  timKiemVanBan,
  type ListParams,
} from '../lib/api';

const common = {
  placeholderData: keepPreviousData,
  staleTime: 30_000,
  retry: 1,
} as const;

export function useVanBanDen(params: ListParams = {}) {
  return useQuery({
    queryKey: ['van-ban-den', params],
    queryFn: ({ signal }) => getVanBanDen(params, signal),
    select: (data) => ({
      rows: data.items.map(mapVanBanDen),
      pagination: data.pagination,
    }),
    ...common,
  });
}

export function useVanBanDiCoTenLoai(params: ListParams = {}) {
  return useQuery({
    queryKey: ['van-ban-di-co-ten-loai', params],
    queryFn: ({ signal }) => getVanBanDiCoTenLoai(params, signal),
    select: (data) => ({
      rows: data.items.map((d) => mapVanBanDi(d, true)),
      pagination: data.pagination,
    }),
    ...common,
  });
}

export function useVanBanDiKhongTenLoai(params: ListParams = {}) {
  return useQuery({
    queryKey: ['van-ban-di-khong-ten-loai', params],
    queryFn: ({ signal }) => getVanBanDiKhongTenLoai(params, signal),
    select: (data) => ({
      rows: data.items.map((d) => mapVanBanDi(d, false)),
      pagination: data.pagination,
    }),
    ...common,
  });
}

// Tra cứu vị trí lưu kho: tìm văn bản trên cả 3 bảng kèm vị trí
// hồ sơ → hộp → thùng. Chỉ chạy khi có từ khóa.
export function useTraCuuVanBan(q: string) {
  return useQuery({
    queryKey: ['tra-cuu-van-ban', q],
    queryFn: ({ signal }) => timKiemVanBan(q, 10, signal),
    select: mapTraCuu,
    enabled: q !== '',
    ...common,
  });
}

// Chi tiết thùng theo mã (vị trí kho, đợt lưu kho, ngày nhập kho).
export function useThungByMa(maThung: string | null | undefined) {
  return useQuery({
    queryKey: ['thung-theo-ma', maThung],
    queryFn: ({ signal }) => getThungByMa(maThung!, signal),
    enabled: !!maThung,
    staleTime: 5 * 60_000,
    retry: 1,
  });
}
