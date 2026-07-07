import { useState } from 'react';
import { donViData, canBoData, loaiVBData } from '../data.js';
import { IconPlus, IconEdit, IconTrash } from '../components/Icons.jsx';

const tabs = [
  { key: 'donvi', label: 'Đơn vị', cols: ['Mã', 'Tên đơn vị', 'Trưởng đơn vị', 'Số cán bộ'], rows: donViData },
  { key: 'canbo', label: 'Cán bộ', cols: ['Họ và tên', 'Chức vụ', 'Đơn vị', 'Email'], rows: canBoData },
  { key: 'loai', label: 'Loại văn bản', cols: ['Mã', 'Tên loại', 'Ký hiệu', 'Số VB đã dùng'], rows: loaiVBData },
];

export default function DanhMuc({ onExport, onAdd }) {
  const [tab, setTab] = useState('donvi');
  const current = tabs.find(t => t.key === tab);

  return (
    <div style={{ maxWidth: 1200, margin: '0 auto' }}>
      <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 16, flexWrap: 'wrap', gap: 12 }}>
        <div>
          <h1 className="vl-page-title">Danh mục hệ thống</h1>
          <p className="vl-page-sub">Quản lý dữ liệu nền dùng chung</p>
        </div>
        <div style={{ display: 'flex', gap: 10 }}>
          <button className="vl-btn export" onClick={onExport}>Xuất Excel</button>
          <button className="vl-btn-primary" onClick={onAdd}>
            <IconPlus />
            Thêm mới
          </button>
        </div>
      </div>

      <div className="vl-tabs">
        {tabs.map(t => (
          <button key={t.key} className={'vl-tab' + (tab === t.key ? ' active' : '')} onClick={() => setTab(t.key)}>
            {t.label}
          </button>
        ))}
      </div>

      <div className="vl-card" style={{ overflow: 'hidden' }}>
        <div style={{ overflowX: 'auto' }}>
          <table className="vl-table" style={{ minWidth: 700 }}>
            <thead>
              <tr>
                {current.cols.map(c => <th key={c} style={{ padding: '12px 16px' }}>{c}</th>)}
                <th style={{ textAlign: 'right', padding: '12px 16px', width: 90 }}>Thao tác</th>
              </tr>
            </thead>
            <tbody>
              {current.rows.map((r, i) => (
                <tr key={i} className="hoverable">
                  {r.map((cell, j) => <td key={j} style={{ padding: '0 16px' }}>{cell}</td>)}
                  <td style={{ padding: '0 16px', textAlign: 'right' }}>
                    <div style={{ display: 'inline-flex', gap: 6 }}>
                      <button className="vl-icon-btn"><IconEdit /></button>
                      <button className="vl-icon-btn"><IconTrash /></button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
