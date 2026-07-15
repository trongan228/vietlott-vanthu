import { useRef, useState } from 'react';
import Header from './components/Header.jsx';
import NavBar from './components/NavBar.jsx';
import Toast from './components/Toast.jsx';
import DetailDrawer from './components/DetailDrawer.jsx';
import Dashboard from './screens/Dashboard.jsx';
import VanBanDen from './screens/VanBanDen.jsx';
import VanBanDi from './screens/VanBanDi.jsx';
import VanBanForm from './screens/VanBanForm.jsx';
import TraCuuLuuKho from './screens/TraCuuLuuKho.jsx';
import DanhMuc from './screens/DanhMuc.jsx';

function detailFields(d, kind) {
  if (kind === 'den') {
    return [
      { k: 'Số đến', v: d.soDen, w: 700 },
      { k: 'Ngày đến', v: d.ngayDen, w: 500 },
      { k: 'Nơi gửi', v: d.noiGui, w: 600 },
      { k: 'Số ký hiệu', v: d.soKyHieu, w: 500 },
      { k: 'Ngày văn bản', v: d.ngayVB, w: 500 },
      { k: 'Trích yếu', v: d.trichYeu, w: 500 },
      { k: 'Đơn vị nhận', v: d.donViNhan, w: 500 },
      { k: 'Người ký nhận', v: d.nguoiKyNhan, w: 500 },
    ];
  }
  const f = [{ k: 'Số ký hiệu', v: d.soKyHieu, w: 700 }];
  if (d.loaiVB) f.push({ k: 'Loại văn bản', v: d.loaiVB, w: 600 });
  f.push(
    { k: 'Ngày văn bản', v: d.ngayVB, w: 500 },
    { k: 'Trích yếu', v: d.trichYeu, w: 500 },
    { k: 'Người ký', v: d.nguoiKy, w: 600 },
    { k: 'Nơi nhận', v: d.noiNhan, w: 500 },
    { k: 'Số lượng bản', v: d.soLuong, w: 500 },
  );
  return f;
}

export default function App() {
  const [screen, setScreen] = useState('dashboard');
  const [formType, setFormType] = useState('den');
  const [detail, setDetail] = useState(null);
  const [toast, setToast] = useState(null);
  const [kho, setKho] = useState({ query: '', active: '' });
  const toastTimer = useRef(null);

  const go = (s) => { setScreen(s); setDetail(null); };

  const showToast = (msg) => {
    setToast(msg);
    clearTimeout(toastTimer.current);
    toastTimer.current = setTimeout(() => setToast(null), 2600);
  };

  const openForm = (type) => { setFormType(type); setScreen('form'); };

  const openDetail = (d, kind) => {
    const hasSoKyHieu = d.soKyHieu && d.soKyHieu !== '—';
    const dir = kind === 'den'
      ? { dirLabel: 'VĂN BẢN ĐẾN', dirBg: '#EAF2FC', dirColor: '#2b6fd6', so: hasSoKyHieu ? d.soKyHieu : 'Số đến ' + d.soDen, khoKey: d.soDen }
      : { dirLabel: 'VĂN BẢN ĐI', dirBg: '#FCEBEC', dirColor: '#D71920', so: d.soKyHieu, khoKey: d.soKyHieu };
    setDetail({ ...dir, fields: detailFields(d, kind) });
  };

  const detailToKho = () => {
    const k = detail ? detail.khoKey : '';
    setDetail(null);
    setKho({ query: k, active: k });
    setScreen('luukho');
  };

  const exportExcel = () => showToast('Đang xuất dữ liệu ra file Excel (.xlsx)…');
  const importExcel = () => showToast('Chọn file Excel để nhập dữ liệu…');

  return (
    <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column', background: '#F7F7F8' }}>
      <Header />
      <NavBar screen={screen === 'form' ? (formType === 'den' ? 'den' : 'di') : screen} onGo={go} />

      <main style={{ flex: 1, overflow: 'auto', padding: '22px 26px 40px' }}>
        {screen === 'dashboard' && (
          <Dashboard onGoDen={() => go('den')} onOpenDetail={openDetail} onExport={exportExcel} />
        )}
        {screen === 'den' && (
          <VanBanDen onOpenDetail={openDetail} onOpenForm={() => openForm('den')} onExport={exportExcel} onImport={importExcel} />
        )}
        {screen === 'di' && (
          <VanBanDi onOpenDetail={openDetail} onOpenForm={() => openForm('diLoai')} onExport={exportExcel} onImport={importExcel} />
        )}
        {screen === 'form' && (
          <VanBanForm
            type={formType}
            onTypeChange={setFormType}
            onBack={() => go(formType === 'den' ? 'den' : 'di')}
            onSaved={() => { go(formType === 'den' ? 'den' : 'di'); showToast('Đã lưu văn bản thành công.'); }}
          />
        )}
        {screen === 'luukho' && (
          <TraCuuLuuKho kho={kho} onChange={setKho} />
        )}
        {screen === 'danhmuc' && (
          <DanhMuc onExport={exportExcel} onAdd={() => showToast('Mở form thêm mới danh mục.')} />
        )}
      </main>

      <DetailDrawer detail={detail} onClose={() => setDetail(null)} onToKho={detailToKho} />
      <Toast message={toast} />
    </div>
  );
}
