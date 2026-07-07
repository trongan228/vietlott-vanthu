// Custom hook TanStack Query cho từng loại văn bản.
// Giữ dữ liệu trang trước khi chuyển trang (placeholderData) để bảng không nhấp nháy.
import { keepPreviousData, useQuery } from '@tanstack/react-query';
import {
  getVanBanDen,
  getVanBanDiCoTenLoai,
  getVanBanDiKhongTenLoai,
  mapVanBanDen,
  mapVanBanDi,
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
