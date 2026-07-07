// Trạng thái bảng: đang tải (skeleton), lỗi (kèm nút thử lại), rỗng.
// Luôn render bên trong <tbody> để giữ khung bảng, không để trắng màn hình.

export function TableLoading({ colSpan, rows = 8 }) {
  return (
    <>
      {Array.from({ length: rows }).map((_, i) => (
        <tr key={i}>
          <td colSpan={colSpan} style={{ padding: '0 16px', height: 52 }}>
            <div className="vl-skeleton" style={{ width: `${88 - (i % 4) * 9}%` }} />
          </td>
        </tr>
      ))}
    </>
  );
}

export function TableError({ colSpan, error, onRetry }) {
  const msg = error?.message || 'Đã xảy ra lỗi không xác định.';
  return (
    <tr>
      <td colSpan={colSpan} style={{ padding: '40px 16px', textAlign: 'center' }}>
        <div style={{ width: 52, height: 52, borderRadius: '50%', background: '#FDF3E0', display: 'flex', alignItems: 'center', justifyContent: 'center', margin: '0 auto 12px' }}>
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
            <path d="M12 9v4m0 4h.01M10.3 3.9 1.8 18a2 2 0 0 0 1.7 3h17a2 2 0 0 0 1.7-3L13.7 3.9a2 2 0 0 0-3.4 0Z" stroke="#B9740A" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
          </svg>
        </div>
        <div style={{ fontSize: 14.5, fontWeight: 700, color: '#3a3a42' }}>Không tải được dữ liệu</div>
        <div style={{ fontSize: 13, color: '#9a9aa2', marginTop: 5 }}>{msg}</div>
        {onRetry && (
          <button className="vl-btn" style={{ marginTop: 14 }} onClick={onRetry}>Thử lại</button>
        )}
      </td>
    </tr>
  );
}

export function TableEmpty({ colSpan, message = 'Không có văn bản nào khớp bộ lọc.' }) {
  return (
    <tr>
      <td colSpan={colSpan} style={{ padding: '40px 16px', textAlign: 'center', color: '#9a9aa2', fontSize: 13.5 }}>
        {message}
      </td>
    </tr>
  );
}
