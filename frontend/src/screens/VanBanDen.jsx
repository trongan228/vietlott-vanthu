import { useState } from 'react';
import { useVanBanDen } from '../hooks/useVanBan.ts';
import { TableLoading, TableError, TableEmpty } from '../components/TableState.jsx';
import { IconExport, IconImport, IconPlus } from '../components/Icons.jsx';

const PAGE_SIZE = 10;
const COLS = 6;

export default function VanBanDen({ onOpenDetail, onOpenForm, onExport, onImport }) {
  const [keyword, setKeyword] = useState('');
  const [year, setYear] = useState('');
  const [applied, setApplied] = useState({ q: '', nam: '' });
  const [page, setPage] = useState(1);

  const { data, isPending, isError, error, refetch, isFetching } = useVanBanDen({
    page,
    pageSize: PAGE_SIZE,
    nam: applied.nam || undefined,
    q: applied.q || undefined,
  });

  const rows = data?.rows ?? [];
  const total = data?.pagination.total ?? 0;
  const totalPages = Math.max(1, Math.ceil(total / PAGE_SIZE));

  const applyFilter = () => {
    setApplied({ q: keyword.trim(), nam: year });
    setPage(1);
  };

  return (
    <div style={{ maxWidth: 1320, margin: '0 auto' }}>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 16, flexWrap: 'wrap', gap: 12 }}>
        <div>
          <h1 className="vl-page-title">Văn bản đến</h1>
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
            Thêm văn bản đến
          </button>
        </div>
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
        <div className="vl-field">
          <label>Từ ngày</label>
          <input type="date" />
        </div>
        <div className="vl-field">
          <label>Đến ngày</label>
          <input type="date" />
        </div>
        <div className="vl-field">
          <label>Nơi gửi</label>
          <input placeholder="Tất cả" style={{ width: 180 }} />
        </div>
        <div className="vl-field" style={{ flex: 1, minWidth: 200 }}>
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
                <th style={{ width: 70 }}>Số đến</th>
                <th>Ngày đến</th>
                <th>Nơi gửi</th>
                <th>Số ký hiệu</th>
                <th>Trích yếu</th>
                <th>Đơn vị nhận</th>
              </tr>
            </thead>
            <tbody style={{ opacity: isFetching && !isPending ? 0.55 : 1 }}>
              {isPending && <TableLoading colSpan={COLS} rows={PAGE_SIZE} />}
              {isError && <TableError colSpan={COLS} error={error} onRetry={refetch} />}
              {!isPending && !isError && rows.length === 0 && <TableEmpty colSpan={COLS} />}
              {!isPending && !isError && rows.map(d => (
                <tr key={d.id} className="clickable" onClick={() => onOpenDetail(d, 'den')}>
                  <td><span style={{ fontWeight: 700, color: '#D71920', fontSize: 14 }}>{d.soDen}</span></td>
                  <td style={{ color: '#4a4a52', whiteSpace: 'nowrap' }}>{d.ngayDen}</td>
                  <td style={{ maxWidth: 170 }}>{d.noiGui}</td>
                  <td style={{ fontSize: 12.5, color: '#5a5a62', whiteSpace: 'nowrap' }}>{d.soKyHieu}</td>
                  <td style={{ maxWidth: 300 }}>{d.trichYeu}</td>
                  <td style={{ fontSize: 12.5, color: '#5a5a62', maxWidth: 150 }}>{d.donViNhan}</td>
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
