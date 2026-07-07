import { useVanBanDen, useVanBanDiCoTenLoai } from '../hooks/useVanBan.ts';
import { TableLoading, TableError, TableEmpty } from '../components/TableState.jsx';
import { IconExport, IconStatDen, IconStatDi, IconClock, IconDone } from '../components/Icons.jsx';

const stats = [
  { label: 'Văn bản đến (T6)', value: '245', accent: '#D71920', deltaColor: '#1B7A43', delta: '▲ 12% so với tháng 5', iconBg: '#FCEBEC', icon: <IconStatDen /> },
  { label: 'Văn bản đi (T6)', value: '187', accent: '#F0A63C', deltaColor: '#1B7A43', delta: '▲ 8% so với tháng 5', iconBg: '#FDF3E0', icon: <IconStatDi /> },
  { label: 'Chờ xử lý', value: '23', accent: '#F0A63C', deltaColor: '#B9740A', delta: 'Cần xử lý trong tuần', iconBg: '#FDF3E0', icon: <IconClock /> },
  { label: 'Đã xử lý (T6)', value: '209', accent: '#1B8A4B', deltaColor: '#1B7A43', delta: 'Đạt 90% khối lượng', iconBg: '#E7F5EC', icon: <IconDone /> },
];

const chartRaw = [
  ['Công văn', 162], ['Quyết định', 128], ['Thông báo', 96], ['Báo cáo', 73],
  ['Tờ trình', 54], ['Giấy mời', 45], ['Chỉ thị', 18], ['Ủy quyền', 21],
].sort((a, b) => b[1] - a[1]);
const chartMax = Math.max(...chartRaw.map(c => c[1]));
const chartColors = ['#D71920', '#E4553B', '#EE7A4F', '#F0A63C', '#F0C24B', '#8FBF6A', '#5BAE7E', '#4E9CD6'];
const chart = chartRaw.map((c, i) => ({
  label: c[0], count: String(c[1]),
  pct: (c[1] / chartMax * 100).toFixed(0) + '%',
  color: chartColors[i % chartColors.length],
}));

const monthRaw = [['T1', 198, 152], ['T2', 176, 138], ['T3', 221, 169], ['T4', 205, 161], ['T5', 219, 173], ['T6', 245, 187]];
const mMax = 260;
const months = monthRaw.map(m => ({
  label: m[0],
  denH: (m[1] / mMax * 100).toFixed(0) + '%',
  diH: (m[2] / mMax * 100).toFixed(0) + '%',
}));

const dirDen = { dir: 'ĐẾN', dirBg: '#EAF2FC', dirColor: '#2b6fd6' };
const dirDi = { dir: 'ĐI', dirBg: '#FCEBEC', dirColor: '#D71920' };

export default function Dashboard({ onGoDen, onOpenDetail, onExport }) {
  const denQ = useVanBanDen({ page: 1, pageSize: 3 });
  const diQ = useVanBanDiCoTenLoai({ page: 1, pageSize: 3 });

  const isPending = denQ.isPending || diQ.isPending;
  const isError = denQ.isError && diQ.isError; // chỉ coi là lỗi khi cả hai nguồn đều hỏng
  const denRows = (denQ.data?.rows ?? []).map(d => ({ ...dirDen, so: 'Số ' + d.soDen, trichYeu: d.trichYeu, ngay: d.ngayDen, kind: 'den', doc: d }));
  const diRows = (diQ.data?.rows ?? []).map(d => ({ ...dirDi, so: d.soKyHieu, trichYeu: d.trichYeu, ngay: d.ngayVB, kind: 'di', doc: d }));
  // Đan xen đi/đến như bố cục thiết kế
  const latest = [];
  for (let i = 0; i < 3; i++) {
    if (diRows[i]) latest.push(diRows[i]);
    if (denRows[i]) latest.push(denRows[i]);
  }

  return (
    <div style={{ maxWidth: 1320, margin: '0 auto' }}>
      <div style={{ display: 'flex', alignItems: 'flex-end', justifyContent: 'space-between', marginBottom: 18, flexWrap: 'wrap', gap: 12 }}>
        <div>
          <h1 className="vl-page-title">Tổng quan văn thư</h1>
          <p className="vl-page-sub">Số liệu kỳ tháng 6/2026 · Cập nhật 02/07/2026 09:15</p>
        </div>
        <div style={{ display: 'flex', gap: 10 }}>
          <button className="vl-btn export" onClick={onExport}>
            <IconExport />
            Xuất Excel
          </button>
        </div>
      </div>

      {/* Stat cards */}
      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(4,1fr)', gap: 16, marginBottom: 20 }}>
        {stats.map(s => (
          <div key={s.label} className="vl-card" style={{ borderLeft: `4px solid ${s.accent}`, padding: '16px 18px', boxShadow: '0 1px 2px rgba(20,20,30,.04)' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
              <span style={{ fontSize: 12.5, color: '#70707a', fontWeight: 600, letterSpacing: '.01em' }}>{s.label}</span>
              <span style={{ width: 34, height: 34, borderRadius: 9, background: s.iconBg, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>{s.icon}</span>
            </div>
            <div style={{ fontSize: 30, fontWeight: 700, color: '#1A1A1A', marginTop: 8, lineHeight: 1 }}>{s.value}</div>
            <div style={{ fontSize: 12, color: s.deltaColor, marginTop: 8, fontWeight: 600 }}>{s.delta}</div>
          </div>
        ))}
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: '1.15fr 1fr', gap: 16, marginBottom: 16 }}>
        {/* Biểu đồ theo loại */}
        <div className="vl-card" style={{ padding: '18px 20px' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 16 }}>
            <h3 style={{ margin: 0, fontSize: 15, fontWeight: 700 }}>Văn bản đi theo loại</h3>
            <span style={{ fontSize: 12, color: '#9a9aa2' }}>Tháng 6/2026</span>
          </div>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 12 }}>
            {chart.map(c => (
              <div key={c.label} style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
                <span style={{ width: 78, fontSize: 12.5, color: '#4a4a52', textAlign: 'right', flexShrink: 0 }}>{c.label}</span>
                <div style={{ flex: 1, height: 20, background: '#F2F2F4', borderRadius: 5, overflow: 'hidden' }}>
                  <div style={{ height: '100%', width: c.pct, background: c.color, borderRadius: 5 }} />
                </div>
                <span style={{ width: 28, fontSize: 13, fontWeight: 700, color: '#1A1A1A', textAlign: 'right' }}>{c.count}</span>
              </div>
            ))}
          </div>
        </div>

        {/* Đến / đi 6 tháng */}
        <div className="vl-card" style={{ padding: '18px 20px' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 8 }}>
            <h3 style={{ margin: 0, fontSize: 15, fontWeight: 700 }}>Đến / Đi 6 tháng gần đây</h3>
          </div>
          <div style={{ display: 'flex', gap: 16, marginBottom: 14, fontSize: 12, color: '#70707a' }}>
            <span style={{ display: 'flex', alignItems: 'center', gap: 6 }}><span style={{ width: 10, height: 10, borderRadius: 2, background: '#D71920' }} />Đến</span>
            <span style={{ display: 'flex', alignItems: 'center', gap: 6 }}><span style={{ width: 10, height: 10, borderRadius: 2, background: '#F0A63C' }} />Đi</span>
          </div>
          <div style={{ display: 'flex', alignItems: 'flex-end', gap: 14, height: 180, paddingTop: 8 }}>
            {months.map(m => (
              <div key={m.label} style={{ flex: 1, display: 'flex', flexDirection: 'column', alignItems: 'center', gap: 6, height: '100%', justifyContent: 'flex-end' }}>
                <div style={{ display: 'flex', alignItems: 'flex-end', gap: 4, width: '100%', justifyContent: 'center', flex: 1 }}>
                  <div style={{ width: 14, background: '#D71920', borderRadius: '3px 3px 0 0', height: m.denH }} />
                  <div style={{ width: 14, background: '#F0A63C', borderRadius: '3px 3px 0 0', height: m.diH }} />
                </div>
                <span style={{ fontSize: 11.5, color: '#70707a' }}>{m.label}</span>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Văn bản mới nhất */}
      <div className="vl-card" style={{ overflow: 'hidden' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '16px 20px', borderBottom: '1px solid #EFEFF1' }}>
          <h3 style={{ margin: 0, fontSize: 15, fontWeight: 700 }}>Văn bản mới nhất</h3>
          <button className="vl-link-btn" onClick={onGoDen}>Xem tất cả →</button>
        </div>
        <table className="vl-table">
          <thead>
            <tr>
              <th style={{ paddingLeft: 20 }}>Loại</th>
              <th>Số ký hiệu</th>
              <th>Trích yếu</th>
              <th style={{ paddingRight: 20 }}>Ngày</th>
            </tr>
          </thead>
          <tbody>
            {isPending && <TableLoading colSpan={4} rows={6} />}
            {!isPending && isError && (
              <TableError colSpan={4} error={denQ.error || diQ.error} onRetry={() => { denQ.refetch(); diQ.refetch(); }} />
            )}
            {!isPending && !isError && latest.length === 0 && <TableEmpty colSpan={4} message="Chưa có văn bản nào." />}
            {!isPending && !isError && latest.map((d, i) => (
              <tr key={i} className="clickable" onClick={() => onOpenDetail(d.doc, d.kind)}>
                <td style={{ paddingLeft: 20, height: 50 }}>
                  <span className="vl-dir-badge" style={{ background: d.dirBg, color: d.dirColor }}>{d.dir}</span>
                </td>
                <td><span style={{ fontWeight: 700, color: '#D71920', fontSize: 13 }}>{d.so}</span></td>
                <td style={{ maxWidth: 420 }}>{d.trichYeu}</td>
                <td style={{ color: '#5a5a62', whiteSpace: 'nowrap', paddingRight: 20 }}>{d.ngay}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
