import { useState } from 'react';
import { IconBack, IconCheck } from '../components/Icons.jsx';

const blankForm = {
  soDen: '', ngayDen: '', noiGui: '', donViNhan: '', nguoiKyNhan: '',
  soKyHieu: '', loaiVB: '', nguoiKy: '', noiNhan: '', soLuong: '',
  ngayVB: '', trichYeu: '', ghiChu: '',
};

function validate(form, type) {
  const e = {};
  if (!form.trichYeu.trim()) e.trichYeu = 'Vui lòng nhập trích yếu.';
  if (!form.ngayVB) e.ngayVB = 'Vui lòng chọn ngày văn bản.';
  if (type === 'den') {
    if (!form.soDen.trim()) e.soDen = 'Bắt buộc.';
    if (!form.ngayDen) e.ngayDen = 'Bắt buộc.';
    if (!form.noiGui.trim()) e.noiGui = 'Bắt buộc.';
  } else {
    if (!form.soKyHieu.trim()) e.soKyHieu = 'Bắt buộc.';
    if (type === 'diLoai' && !form.loaiVB) e.loaiVB = 'Bắt buộc.';
  }
  return e;
}

function Field({ label, required, error, children }) {
  return (
    <div>
      <label className="vl-form-label">{label} {required && <span style={{ color: '#D71920' }}>*</span>}</label>
      {children}
      {error && <div className="vl-form-error">{error}</div>}
    </div>
  );
}

export default function VanBanForm({ type, onTypeChange, onBack, onSaved }) {
  const [form, setForm] = useState(blankForm);
  const [submitted, setSubmitted] = useState(false);

  const errors = submitted ? validate(form, type) : {};
  const isDen = type === 'den';

  const set = (k) => (e) => setForm(f => ({ ...f, [k]: e.target.value }));
  const inputCls = (k) => 'vl-input' + (errors[k] ? ' error' : '');

  const switchType = (t) => { onTypeChange(t); setSubmitted(false); };

  const save = () => {
    setSubmitted(true);
    if (Object.keys(validate(form, type)).length === 0) onSaved();
  };

  const title = isDen ? 'Thêm văn bản đến' : (type === 'diLoai' ? 'Thêm văn bản đi (có tên loại)' : 'Thêm văn bản đi (không tên loại)');

  return (
    <div style={{ maxWidth: 880, margin: '0 auto' }}>
      <button className="vl-back-btn" onClick={onBack}>
        <IconBack />
        Quay lại danh sách
      </button>
      <div className="vl-card" style={{ borderRadius: 12, overflow: 'hidden' }}>
        <div style={{ padding: '18px 24px', borderBottom: '1px solid #EFEFF1' }}>
          <h1 style={{ margin: 0, fontSize: 19, fontWeight: 700 }}>{title}</h1>
          <p style={{ margin: '5px 0 0', fontSize: 13, color: '#70707a' }}>
            Trường có dấu <span style={{ color: '#D71920', fontWeight: 700 }}>*</span> là bắt buộc
          </p>
        </div>

        {/* Segmented */}
        <div style={{ padding: '18px 24px 0' }}>
          <div className="vl-seg-wrap">
            <button className={'vl-seg' + (type === 'den' ? ' active' : '')} onClick={() => switchType('den')}>Văn bản đến</button>
            <button className={'vl-seg' + (type === 'diLoai' ? ' active' : '')} onClick={() => switchType('diLoai')}>Đi có tên loại</button>
            <button className={'vl-seg' + (type === 'diKhong' ? ' active' : '')} onClick={() => switchType('diKhong')}>Đi không tên loại</button>
          </div>
        </div>

        <div style={{ padding: '20px 24px 8px', display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '16px 18px' }}>
          {isDen ? (
            <>
              <Field label="Số đến" required error={errors.soDen}>
                <input className={inputCls('soDen')} value={form.soDen} onChange={set('soDen')} placeholder="VD: 246" />
              </Field>
              <Field label="Ngày đến" required error={errors.ngayDen}>
                <input type="date" className={inputCls('ngayDen')} value={form.ngayDen} onChange={set('ngayDen')} />
              </Field>
              <Field label="Nơi gửi" required error={errors.noiGui}>
                <input className={inputCls('noiGui')} value={form.noiGui} onChange={set('noiGui')} placeholder="VD: Bộ Tài chính" />
              </Field>
              <Field label="Đơn vị nhận">
                <select className="vl-input" value={form.donViNhan} onChange={set('donViNhan')}>
                  <option value="">— Chọn đơn vị —</option>
                  <option>Ban Tổng Giám đốc</option>
                  <option>Văn phòng</option>
                  <option>Phòng Tài chính - Kế toán</option>
                  <option>Phòng Kinh doanh</option>
                  <option>Phòng Công nghệ thông tin</option>
                </select>
              </Field>
              <Field label="Số ký hiệu">
                <input className="vl-input" value={form.soKyHieu} onChange={set('soKyHieu')} placeholder="VD: 1124/BTC-TCT" />
              </Field>
              <Field label="Ngày văn bản" required error={errors.ngayVB}>
                <input type="date" className={inputCls('ngayVB')} value={form.ngayVB} onChange={set('ngayVB')} />
              </Field>
              <Field label="Người ký nhận">
                <input className="vl-input" value={form.nguoiKyNhan} onChange={set('nguoiKyNhan')} placeholder="VD: Nguyễn Thị Thu" />
              </Field>
            </>
          ) : (
            <>
              <Field label="Số ký hiệu" required error={errors.soKyHieu}>
                <input className={inputCls('soKyHieu')} value={form.soKyHieu} onChange={set('soKyHieu')} placeholder="VD: 513/QĐ-VIETLOTT" />
              </Field>
              {type === 'diLoai' && (
                <Field label="Loại văn bản" required error={errors.loaiVB}>
                  <select className={inputCls('loaiVB')} value={form.loaiVB} onChange={set('loaiVB')}>
                    <option value="">— Chọn loại —</option>
                    <option>Quyết định</option><option>Thông báo</option><option>Tờ trình</option><option>Báo cáo</option>
                    <option>Ủy quyền</option><option>Chỉ thị</option><option>Giấy mời</option><option>Công văn</option>
                  </select>
                </Field>
              )}
              <Field label="Người ký">
                <input className="vl-input" value={form.nguoiKy} onChange={set('nguoiKy')} placeholder="VD: Nguyễn Thanh Đạm" />
              </Field>
              <Field label="Nơi nhận">
                <input className="vl-input" value={form.noiNhan} onChange={set('noiNhan')} placeholder="VD: Các phòng ban" />
              </Field>
              <Field label="Số lượng bản">
                <input type="number" className="vl-input" value={form.soLuong} onChange={set('soLuong')} placeholder="VD: 15" />
              </Field>
              <Field label="Ngày văn bản" required error={errors.ngayVB}>
                <input type="date" className={inputCls('ngayVB')} value={form.ngayVB} onChange={set('ngayVB')} />
              </Field>
            </>
          )}
        </div>

        <div style={{ padding: '4px 24px 8px' }}>
          <Field label="Trích yếu nội dung" required error={errors.trichYeu}>
            <textarea
              className={'vl-textarea' + (errors.trichYeu ? ' error' : '')}
              value={form.trichYeu}
              onChange={set('trichYeu')}
              placeholder="Nhập trích yếu nội dung văn bản…"
            />
          </Field>
        </div>
        <div style={{ padding: '12px 24px 8px' }}>
          <Field label="Ghi chú">
            <textarea
              className="vl-textarea"
              style={{ minHeight: 56 }}
              value={form.ghiChu}
              onChange={set('ghiChu')}
              placeholder="Ghi chú thêm (tùy chọn)…"
            />
          </Field>
        </div>

        <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 10, padding: '16px 24px', borderTop: '1px solid #EFEFF1', background: '#FAFAFB' }}>
          <button
            onClick={onBack}
            style={{ height: 40, padding: '0 20px', border: '1px solid #D4D4D8', background: '#fff', color: '#333', borderRadius: 8, fontSize: 14, fontWeight: 600, cursor: 'pointer' }}
          >
            Hủy
          </button>
          <button
            onClick={save}
            style={{ height: 40, padding: '0 24px', border: 'none', background: '#D71920', color: '#fff', borderRadius: 8, fontSize: 14, fontWeight: 700, cursor: 'pointer', boxShadow: '0 1px 2px rgba(215,25,32,.3)', display: 'flex', alignItems: 'center', gap: 7 }}
          >
            <IconCheck />
            Lưu văn bản
          </button>
        </div>
      </div>
    </div>
  );
}
