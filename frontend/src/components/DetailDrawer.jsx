import { IconClose, IconArchive } from './Icons.jsx';

export default function DetailDrawer({ detail, onClose, onToKho }) {
  if (!detail) return null;
  return (
    <div style={{ position: 'fixed', inset: 0, zIndex: 50 }}>
      <div className="vl-drawer-overlay" onClick={onClose} />
      <div className="vl-drawer">
        <div style={{ padding: '18px 22px', borderBottom: '1px solid #EFEFF1', display: 'flex', alignItems: 'flex-start', justifyContent: 'space-between', gap: 12 }}>
          <div>
            <span className="vl-dir-badge" style={{ background: detail.dirBg, color: detail.dirColor, marginBottom: 8 }}>{detail.dirLabel}</span>
            <div style={{ fontSize: 18, fontWeight: 700, color: '#D71920' }}>{detail.so}</div>
          </div>
          <button className="vl-drawer-close" onClick={onClose}><IconClose /></button>
        </div>
        <div style={{ flex: 1, overflow: 'auto', padding: '20px 22px' }}>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 16 }}>
            {detail.fields.map(f => (
              <div key={f.k} style={{ display: 'flex', flexDirection: 'column', gap: 4, paddingBottom: 14, borderBottom: '1px solid #F2F2F4' }}>
                <span style={{ fontSize: 11.5, color: '#9a9aa2', fontWeight: 600, textTransform: 'uppercase', letterSpacing: '.02em' }}>{f.k}</span>
                <span style={{ fontSize: 14.5, color: '#1A1A1A', lineHeight: 1.45, fontWeight: f.w }}>{f.v}</span>
              </div>
            ))}
          </div>
        </div>
        <div style={{ padding: '16px 22px', borderTop: '1px solid #EFEFF1', display: 'flex', gap: 10, background: '#FAFAFB' }}>
          <button className="vl-drawer-action" onClick={onToKho}>
            <IconArchive size={16} />
            Tra cứu lưu kho
          </button>
          <button className="vl-drawer-edit">Chỉnh sửa</button>
        </div>
      </div>
    </div>
  );
}
