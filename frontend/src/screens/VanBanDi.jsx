import { useState } from 'react';
import { useVanBanDiCoTenLoai, useVanBanDiKhongTenLoai } from '../hooks/useVanBan.ts';
import { TableLoading, TableError, TableEmpty } from '../components/TableState.jsx';
import { IconExport, IconImport, IconPlus } from '../components/Icons.jsx';

const PAGE_SIZE = 10;

export default function VanBanDi({ onOpenDetail, onOpenForm, onExport, onImport }) {
  const [tab, setTab] = useState('coloai');
  const [keyword, setKeyword] = useState('');
  const [year, setYear] = useState('');
  const [applied, setApplied] = useState({ q: '', nam: '' });
  const [page, setPage] = useState(1);

  const coLoai = tab === 'coloai';
  const params = {
    page,
    pageSize: PAGE_SIZE,
    nam: applied.nam || undefined,
    q: applied.q || undefined,
  };
  // Hooks phải gọi vô điều kiện; query của tab còn lại đóng vai trò prefetch nhẹ.
  const qCoLoai = useVanBanDiCoTenLoai(params);
  const qKhongLoai = useVanBanDiKhongTenLoai(params);
  const { data, isPending, isError, error, refetch, isFetching } = coLoai ? qCoLoai : qKhongLoai;

  const rows = data?.rows ?? [];
  const total = data?.pagination.total ?? 0;
  const totalPages = Math.max(1, Math.ceil(total / PAGE_SIZE));
  const cols = coLoai ? 7 : 6;

  const switchTab = (t) => { setTab(t); setPage(1); };
  const applyFilter = () => {
    setApplied({ q: keyword.trim(), nam: year });
    setPage(1);
  };

  return (
    <div style={{ maxWidth: 1320, margin: '0 auto' }}>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 16, flexWrap: 'wrap', gap: 12 }}>
        <div>
          <h1 className="vl-page-title">Văn bản đi</h1>
          <p className="vl-page-sub">{isPending ? 'Đang tải…' : `${total.toLocaleString('vi-VN')} văn bản`}{applied.nam ? ` · Năm ${applied.nam}` : ''}</p>
        </div>
        <div style={{ display: 'flex', gap: 10 }}>
          <button className="vl-btn import" onClick={onImport}>
            <IconImport />
            Nhập Excel
          </button>
          <button className="vl-btn export" onClick={onExport}>
            <IconExport />
            Xuất Excel
          </button>
          <button className="vl-btn-primary" onClick={onOpenForm}>
            <IconPlus />
            Thêm văn bản đi
          </button>
        </div>
      </div>

      {/* Tabs */}
      <div className="vl-tabs">
        <button className={'vl-tab' + (coLoai ? ' active' : '')} onClick={() => switchTab('coloai')}>Có tên loại</button>
        <button className={'vl-tab' + (!coLoai ? ' active' : '')} onClick={() => switchTab('khongloai')}>Không tên loại</button>
      </div>

      {/* Filter bar */}
      <div className="vl-filterbar">
        <div className="vl-field">
          <label>Năm</label>
          <select style={{ minWidth: 90 }} value={year} onChange={e => setYear(e.target.value)}>
            <option value="">Tất cả</option>
            <option>2026</option><option>2025</option><option>2024</option>
          </select>
        </div>
        {coLoai && (
          <div className="vl-field">
            <label>Loại văn bản</label>
            <select style={{ minWidth: 150 }}>
              <option>Tất cả</option><option>Quyết định</option><option>Thông báo</option><option>Tờ trình</option>
              <option>Báo cáo</option><option>Chỉ thị</option><option>Giấy mời</option><option>Công văn</option><option>Ủy quyền</option>
            </select>
          </div>
        )}
        <div className="vl-field">
          <label>Người ký</label>
          <select style={{ minWidth: 170 }}>
            <option>Tất cả</option>
          </select>
        </div>
        <div className="vl-field" style={{ flex: 1, minWidth: 180 }}>
          <label>Từ khóa trích yếu</label>
          <input
            value={keyword}
            onChange={e => setKeyword(e.target.value)}
            onKeyDown={e => { if (e.key === 'Enter') applyFilter(); }}
            placeholder="Nhập từ khóa…"
            style={{ width: '100%' }}
          />
        </div>
        <button className="vl-btn-filter" onClick={applyFilter}>Lọc</button>
      </div>

      {/* Table */}
      <div className="vl-card" style={{ overflow: 'hidden' }}>
        <div style={{ overflowX: 'auto' }}>
          <table className="vl-table" style={{ minWidth: 1000 }}>
            <thead>
              <tr>
                <th>Số ký hiệu</th>
                {coLoai && <th>Loại VB</th>}
                <th>Ngày VB</th>
                <th>Trích yếu</th>
                <th>Người ký</th>
                <th>Nơi nhận</th>
                <th style={{ textAlign: 'center' }}>SL bản</th>
              </tr>
            </thead>
            <tbody style={{ opacity: isFetching && !isPending ? 0.55 : 1 }}>
              {isPending && <TableLoading colSpan={cols} rows={PAGE_SIZE} />}
              {isError && <TableError colSpan={cols} error={error} onRetry={refetch} />}
              {!isPending && !isError && rows.length === 0 && <TableEmpty colSpan={cols} />}
              {!isPending && !isError && rows.map(d => (
                <tr key={d.id} className="clickable" onClick={() => onOpenDetail(d, 'di')}>
                  <td>
                    <span style={{ display: 'inline-block', fontWeight: 700, color: '#D71920', fontSize: 13, background: '#FCEBEC', padding: '4px 9px', borderRadius: 6, whiteSpace: 'nowrap' }}>{d.soKyHieu}</span>
                  </td>
                  {coLoai && <td style={{ fontSize: 12.5, color: '#3a3a42', whiteSpace: 'nowrap' }}>{d.loaiVB}</td>}
                  <td style={{ color: '#4a4a52', whiteSpace: 'nowrap' }}>{d.ngayVB}</td>
                  <td style={{ maxWidth: 280 }}>{d.trichYeu}</td>
                  <td style={{ fontSize: 12.5, color: '#3a3a42', whiteSpace: 'nowrap' }}>{d.nguoiKy}</td>
                  <td style={{ fontSize: 12.5, color: '#5a5a62', maxWidth: 140 }}>{d.noiNhan}</td>
                  <td style={{ color: '#4a4a52', textAlign: 'center' }}>{d.soLuong}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <div className="vl-pagination">
          <span>
            {isError ? '—' : `Hiển thị ${rows.length} / ${total.toLocaleString('vi-VN')} văn bản · Trang ${page}/${totalPages.toLocaleString('vi-VN')}`}
          </span>
          <div style={{ display: 'flex', gap: 6 }}>
            <button className={'vl-page-btn' + (page <= 1 ? ' muted' : '')} disabled={page <= 1} onClick={() => setPage(p => p - 1)}>‹</button>
            <button className="vl-page-btn current">{page}</button>
            <button className={'vl-page-btn' + (page >= totalPages ? ' muted' : '')} disabled={page >= totalPages} onClick={() => setPage(p => p + 1)}>›</button>
          </div>
        </div>
      </div>
    </div>
  );
}
