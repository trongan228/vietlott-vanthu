import { IconLogo, IconSearch, IconBell } from './Icons.jsx';

export default function Header() {
  return (
    <header className="vl-header">
      <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
        <div style={{ width: 38, height: 38, background: '#fff', borderRadius: 8, display: 'flex', alignItems: 'center', justifyContent: 'center', boxShadow: '0 1px 2px rgba(0,0,0,.15)' }}>
          <IconLogo />
        </div>
        <div style={{ lineHeight: 1.1 }}>
          <div style={{ color: '#fff', fontWeight: 700, fontSize: 16, letterSpacing: '.02em' }}>VIETLOTT</div>
          <div style={{ color: 'rgba(255,255,255,.82)', fontSize: 11.5, fontWeight: 500 }}>Quản lý văn thư</div>
        </div>
      </div>

      <div style={{ flex: 1, display: 'flex', justifyContent: 'center' }}>
        <div style={{ position: 'relative', width: '100%', maxWidth: 420 }}>
          <IconSearch color="rgba(255,255,255,.8)" style={{ position: 'absolute', left: 12, top: '50%', transform: 'translateY(-50%)' }} />
          <input className="vl-header-search" placeholder="Tìm nhanh số văn bản, trích yếu…" />
        </div>
      </div>

      <div style={{ display: 'flex', alignItems: 'center', gap: 14 }}>
        <button className="vl-header-iconbtn">
          <IconBell />
          <span style={{ position: 'absolute', top: 6, right: 7, width: 8, height: 8, background: '#FFD34E', borderRadius: '50%', border: '1.5px solid #B3141A' }} />
        </button>
        <div style={{ display: 'flex', alignItems: 'center', gap: 10, paddingLeft: 6 }}>
          <div style={{ width: 36, height: 36, borderRadius: '50%', background: '#fff', color: '#D71920', display: 'flex', alignItems: 'center', justifyContent: 'center', fontWeight: 700, fontSize: 14 }}>NT</div>
          <div style={{ lineHeight: 1.15 }}>
            <div style={{ color: '#fff', fontSize: 13, fontWeight: 600 }}>Nguyễn Thị Thu</div>
            <div style={{ color: 'rgba(255,255,255,.8)', fontSize: 11 }}>Nhân viên văn thư</div>
          </div>
        </div>
      </div>
    </header>
  );
}
