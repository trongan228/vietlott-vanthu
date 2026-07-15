import { useState } from 'react';
import { useTraCuuVanBan, useThungByMa } from '../hooks/useVanBan.ts';
import { formatDate } from '../lib/api.ts';
import { IconSearch, IconPin, IconDoc, IconFolder, IconBox, IconChevronRight } from '../components/Icons.jsx';

const SUGGESTIONS = ['01/QĐ-VIETLOTT', '04/TB-VIETLOTT', '245'];

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

function StateCard({ title, sub, children }) {
  return (
    <div className="vl-card vl-fade" style={{ borderRadius: 12, padding: '48px 24px', textAlign: 'center' }}>
      <div style={{ width: 60, height: 60, borderRadius: '50%', background: '#F7F7F8', display: 'flex', alignItems: 'center', justifyContent: 'center', margin: '0 auto 14px' }}>
        <IconSearch size={28} color="#C7C7CE" />
      </div>
      <div style={{ fontSize: 15, fontWeight: 700, color: '#3a3a42' }}>{title}</div>
      <div style={{ fontSize: 13, color: '#9a9aa2', marginTop: 5 }}>{sub}</div>
      {children}
    </div>
  );
}

export default function TraCuuLuuKho({ kho, onChange }) {
  const active = (kho.active || '').trim();
  const { data: rows, isPending, isError, error, refetch, isFetching } = useTraCuuVanBan(active);
  const [selKey, setSelKey] = useState(null);

  const selected = rows?.find(r => r.key === selKey) ?? rows?.[0] ?? null;
  const viTri = selected?.viTri ?? null;
  const { data: thung } = useThungByMa(viTri?.ma_thung);

  const search = () => { setSelKey(null); onChange({ ...kho, active: kho.query.trim() }); };
  const pick = (q) => { setSelKey(null); onChange({ query: q, active: q }); };

  return (
    <div style={{ maxWidth: 1060, margin: '0 auto' }}>
      <h1 className="vl-page-title" style={{ marginBottom: 4 }}>Tra cứu vị trí lưu kho</h1>
      <p style={{ margin: '0 0 18px', fontSize: 13.5, color: '#70707a' }}>
        Nhập số đến, số ký hiệu hoặc trích yếu để xem đường dẫn lưu trữ vật lý của văn bản (hồ sơ → hộp → thùng).
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
              placeholder="VD: 01/QĐ-VIETLOTT hoặc 245"
            />
          </div>
          <button className="vl-kho-btn" onClick={search}>Tra cứu</button>
        </div>
        <div style={{ display: 'flex', gap: 8, alignItems: 'center', marginTop: 14, flexWrap: 'wrap' }}>
          <span style={{ fontSize: 12.5, color: '#9a9aa2' }}>Gợi ý:</span>
          {SUGGESTIONS.map(q => (
            <button key={q} className="vl-kho-chip" onClick={() => pick(q)}>{q}</button>
          ))}
        </div>
      </div>

      {!active && (
        <StateCard title="Chưa có từ khóa tra cứu" sub="Nhập số văn bản hoặc trích yếu rồi bấm Tra cứu." />
      )}

      {active && isPending && (
        <div className="vl-card vl-fade" style={{ borderRadius: 12, padding: '22px 24px' }}>
          <div className="vl-skeleton" style={{ width: '35%', height: 18, marginBottom: 14 }} />
          <div className="vl-skeleton" style={{ width: '80%', height: 14, marginBottom: 10 }} />
          <div className="vl-skeleton" style={{ width: '65%', height: 14 }} />
        </div>
      )}

      {active && isError && (
        <StateCard title="Không tra cứu được" sub={error?.message || 'Đã xảy ra lỗi không xác định.'}>
          <button className="vl-btn" style={{ marginTop: 14 }} onClick={refetch}>Thử lại</button>
        </StateCard>
      )}

      {active && !isPending && !isError && rows.length === 0 && (
        <StateCard
          title="Không tìm thấy văn bản"
          sub={`Không có văn bản nào khớp với “${active}”. Kiểm tra lại số văn bản.`}
        />
      )}

      {active && !isPending && !isError && rows.length > 0 && (
        <div className="vl-fade" style={{ opacity: isFetching ? 0.55 : 1 }}>
          {/* Danh sách kết quả khớp (khi có nhiều hơn 1) */}
          {rows.length > 1 && (
            <div className="vl-card" style={{ borderRadius: 12, padding: '14px 16px', marginBottom: 14 }}>
              <div style={{ fontSize: 12.5, color: '#9a9aa2', marginBottom: 8 }}>
                Tìm thấy {rows.length} văn bản khớp — chọn để xem vị trí:
              </div>
              <div style={{ display: 'flex', flexDirection: 'column', gap: 4 }}>
                {rows.map(r => {
                  const isSel = r.key === selected?.key;
                  return (
                    <button
                      key={r.key}
                      onClick={() => setSelKey(r.key)}
                      style={{
                        display: 'flex', alignItems: 'center', gap: 10, textAlign: 'left',
                        padding: '9px 12px', borderRadius: 9, cursor: 'pointer',
                        background: isSel ? '#FCEBEC' : 'transparent',
                        border: `1px solid ${isSel ? '#F0C9CB' : 'transparent'}`,
                      }}
                    >
                      <span style={{
                        fontSize: 11, fontWeight: 700, whiteSpace: 'nowrap', padding: '3px 8px', borderRadius: 12,
                        color: r.den ? '#2b6fd6' : '#D71920', background: r.den ? '#EAF2FC' : '#FCEBEC',
                      }}>{r.loaiLabel}</span>
                      <span style={{ fontSize: 13.5, fontWeight: 700, color: '#1A1A1A', whiteSpace: 'nowrap' }}>{r.tieuDe}</span>
                      <span style={{ fontSize: 12.5, color: '#70707a', overflow: 'hidden', textOverflow: 'ellipsis', whiteSpace: 'nowrap' }}>{r.trichYeu}</span>
                      <span style={{ fontSize: 12, color: '#9a9aa2', marginLeft: 'auto', whiteSpace: 'nowrap' }}>{r.ngay}</span>
                    </button>
                  );
                })}
              </div>
            </div>
          )}

          <div className="vl-card" style={{ borderRadius: 12, padding: '22px 24px', marginBottom: 14 }}>
            {viTri ? (
              <>
                <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 4 }}>
                  <span style={{ fontSize: 12, fontWeight: 700, color: '#1B7A43', background: '#E7F5EC', padding: '4px 10px', borderRadius: 20 }}>✓ Đã xác định vị trí</span>
                  <span style={{ fontSize: 13, color: '#70707a' }}>Đường dẫn lưu trữ</span>
                </div>

                {/* Sơ đồ 4 tầng: văn bản → hồ sơ → hộp → thùng */}
                <div style={{ display: 'flex', alignItems: 'stretch', gap: 0, marginTop: 18, flexWrap: 'wrap' }}>
                  <PathCard label="Văn bản" value={selected.tieuDe} sub={selected.loaiLabel} icon={<IconDoc />} highlight arrow />
                  <PathCard label="Hồ sơ" value={`Hồ sơ số ${viTri.so_ho_so}`} sub={viTri.ho_so_tieu_de} icon={<IconFolder />} arrow />
                  <PathCard
                    label="Hộp"
                    value={`Hộp số ${viTri.so_hop}`}
                    sub={viTri.loai_hop === 'vinh_vien' ? 'Bảo quản vĩnh viễn' : 'Bảo quản có thời hạn'}
                    icon={<IconBox />}
                    arrow
                  />
                  <PathCard
                    label="Thùng"
                    value={viTri.ma_thung ?? 'Chưa xếp thùng'}
                    sub={viTri.so_serial ? `Serial ${viTri.so_serial}` : (viTri.ma_thung ? '' : 'Hộp chưa đưa vào thùng')}
                    icon={<IconBox />}
                  />
                </div>

                {viTri.ma_thung && (
                  <div style={{ marginTop: 20, padding: '14px 16px', background: '#FCF7F7', border: '1px dashed #F0C9CB', borderRadius: 10, display: 'flex', alignItems: 'center', gap: 12 }}>
                    <IconPin />
                    <div>
                      <span style={{ fontSize: 13, color: '#70707a' }}>Vị trí kho: </span>
                      <span style={{ fontSize: 14, fontWeight: 700, color: '#1A1A1A' }}>
                        {thung?.vi_tri_kho || 'Chưa cập nhật'}
                      </span>
                      {thung?.dot_luu_kho != null && (
                        <span style={{ fontSize: 13, color: '#70707a' }}> · Đợt lưu kho {thung.dot_luu_kho}</span>
                      )}
                      {thung?.ngay_nhap_kho && (
                        <span style={{ fontSize: 13, color: '#70707a' }}> · Nhập kho {formatDate(thung.ngay_nhap_kho)}</span>
                      )}
                    </div>
                  </div>
                )}
              </>
            ) : (
              <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
                <span style={{ fontSize: 12, fontWeight: 700, color: '#B9740A', background: '#FDF3E0', padding: '4px 10px', borderRadius: 20 }}>Chưa xếp kho</span>
                <span style={{ fontSize: 13, color: '#70707a' }}>Văn bản này chưa được gán vào hồ sơ lưu trữ nên chưa có vị trí hộp/thùng.</span>
              </div>
            )}
          </div>

          {/* Chi tiết văn bản */}
          <div className="vl-card" style={{ borderRadius: 12, padding: '20px 24px' }}>
            <h3 style={{ margin: '0 0 14px', fontSize: 14, fontWeight: 700, color: '#3a3a42' }}>Thông tin văn bản</h3>
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '14px 28px' }}>
              {selected.info.map(({ k, v, strong }) => (
                <div key={k} style={{ display: 'flex', flexDirection: 'column', gap: 3 }}>
                  <span style={{ fontSize: 11.5, color: '#9a9aa2', fontWeight: 600 }}>{k}</span>
                  <span style={{ fontSize: 14, color: '#1A1A1A', fontWeight: strong ? 700 : 500 }}>{v}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
