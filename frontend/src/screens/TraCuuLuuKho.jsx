import { khoStore, khoSuggestions } from '../data.js';
import { IconSearch, IconPin, IconDoc, IconFolder, IconBox, IconChevronRight } from '../components/Icons.jsx';

function PathCard({ label, value, sub, icon, highlight, arrow }) {
  return (
    <div style={{ display: 'flex', alignItems: 'center' }}>
      <div style={{
        minWidth: 180,
        background: highlight ? '#FCEBEC' : '#FAFAFB',
        border: `1.5px solid ${highlight ? '#F0C9CB' : '#E5E5E8'}`,
        borderRadius: 12, padding: '16px 18px',
      }}>
        <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 10 }}>
          <span style={{ width: 32, height: 32, borderRadius: 9, background: '#fff', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>{icon}</span>
          <span style={{ fontSize: 11.5, fontWeight: 700, color: highlight ? '#D71920' : '#63636b', textTransform: 'uppercase', letterSpacing: '.03em' }}>{label}</span>
        </div>
        <div style={{ fontSize: 14.5, fontWeight: 700, color: '#1A1A1A', lineHeight: 1.3 }}>{value}</div>
        <div style={{ fontSize: 12, color: '#70707a', marginTop: 3 }}>{sub}</div>
      </div>
      {arrow && (
        <div style={{ padding: '0 6px', color: '#C7C7CE', flexShrink: 0 }}><IconChevronRight /></div>
      )}
    </div>
  );
}

export default function TraCuuLuuKho({ kho, onChange }) {
  const rec = khoStore[kho.active];

  const search = () => onChange({ ...kho, active: kho.query.trim() });
  const pick = (q) => onChange({ query: q, active: q });

  return (
    <div style={{ maxWidth: 1060, margin: '0 auto' }}>
      <h1 className="vl-page-title" style={{ marginBottom: 4 }}>Tra cứu vị trí lưu kho</h1>
      <p style={{ margin: '0 0 18px', fontSize: 13.5, color: '#70707a' }}>
        Nhập số đến hoặc số ký hiệu để xem đường dẫn lưu trữ vật lý của văn bản.
      </p>

      <div className="vl-card" style={{ borderRadius: 12, padding: '22px 24px', marginBottom: 18 }}>
        <div style={{ display: 'flex', gap: 12, alignItems: 'center' }}>
          <div style={{ position: 'relative', flex: 1 }}>
            <IconSearch size={18} color="#9a9aa2" style={{ position: 'absolute', left: 14, top: '50%', transform: 'translateY(-50%)' }} />
            <input
              className="vl-kho-input"
              value={kho.query}
              onChange={e => onChange({ ...kho, query: e.target.value })}
              onKeyDown={e => { if (e.key === 'Enter') search(); }}
              placeholder="VD: 512/QĐ-VIETLOTT hoặc 245"
            />
          </div>
          <button className="vl-kho-btn" onClick={search}>Tra cứu</button>
        </div>
        <div style={{ display: 'flex', gap: 8, alignItems: 'center', marginTop: 14, flexWrap: 'wrap' }}>
          <span style={{ fontSize: 12.5, color: '#9a9aa2' }}>Gợi ý:</span>
          {khoSuggestions.map(q => (
            <button key={q} className="vl-kho-chip" onClick={() => pick(q)}>{q}</button>
          ))}
        </div>
      </div>

      {rec ? (
        <div className="vl-fade">
          <div className="vl-card" style={{ borderRadius: 12, padding: '22px 24px', marginBottom: 14 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 4 }}>
              <span style={{ fontSize: 12, fontWeight: 700, color: '#1B7A43', background: '#E7F5EC', padding: '4px 10px', borderRadius: 20 }}>✓ Tìm thấy</span>
              <span style={{ fontSize: 13, color: '#70707a' }}>Đường dẫn lưu trữ</span>
            </div>

            {/* Sơ đồ 4 tầng */}
            <div style={{ display: 'flex', alignItems: 'stretch', gap: 0, marginTop: 18, flexWrap: 'wrap' }}>
              <PathCard label="Văn bản" value={kho.active} sub={rec.info['Loại văn bản'] || 'Văn bản đến'} icon={<IconDoc />} highlight arrow />
              <PathCard label="Hồ sơ" value={rec.hoSo} sub={rec.hoSoSub} icon={<IconFolder />} arrow />
              <PathCard label="Hộp" value={rec.hop} sub={rec.hopSub} icon={<IconBox />} arrow />
              <PathCard label="Thùng" value={rec.thung} sub={rec.thungSub} icon={<IconBox />} />
            </div>

            <div style={{ marginTop: 20, padding: '14px 16px', background: '#FCF7F7', border: '1px dashed #F0C9CB', borderRadius: 10, display: 'flex', alignItems: 'center', gap: 12 }}>
              <IconPin />
              <div>
                <span style={{ fontSize: 13, color: '#70707a' }}>Vị trí vật lý: </span>
                <span style={{ fontSize: 14, fontWeight: 700, color: '#1A1A1A' }}>{rec.location}</span>
              </div>
            </div>
          </div>

          {/* Chi tiết văn bản */}
          <div className="vl-card" style={{ borderRadius: 12, padding: '20px 24px' }}>
            <h3 style={{ margin: '0 0 14px', fontSize: 14, fontWeight: 700, color: '#3a3a42' }}>Thông tin văn bản</h3>
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '14px 28px' }}>
              {Object.entries(rec.info).map(([k, v]) => (
                <div key={k} style={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
                  <span style={{ fontSize: 11.5, color: '#9a9aa2', fontWeight: 600 }}>{k}</span>
                  <span style={{ fontSize: 14, color: '#1A1A1A', fontWeight: k === 'Số ký hiệu' || k === 'Số đến' ? 700 : 500 }}>{v}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      ) : (
        <div className="vl-card vl-fade" style={{ borderRadius: 12, padding: '48px 24px', textAlign: 'center' }}>
          <div style={{ width: 60, height: 60, borderRadius: '50%', background: '#F7F7F8', display: 'flex', alignItems: 'center', justifyContent: 'center', margin: '0 auto 14px' }}>
            <IconSearch size={28} color="#C7C7CE" />
          </div>
          <div style={{ fontSize: 15, fontWeight: 700, color: '#3a3a42' }}>Không tìm thấy văn bản</div>
          <div style={{ fontSize: 13, color: '#9a9aa2', marginTop: 5 }}>
            Không có văn bản nào khớp với “{kho.query}”. Kiểm tra lại số văn bản.
          </div>
        </div>
      )}
    </div>
  );
}
