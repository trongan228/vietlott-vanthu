// Dữ liệu mẫu theo thiết kế "Quản lý văn thư Vietlott"

export const donViData = [
  ['DV01', 'Ban Tổng Giám đốc', 'Nguyễn Thanh Đạm', '5'],
  ['DV02', 'Văn phòng', 'Trần Thị Mai', '12'],
  ['DV03', 'Phòng Tài chính - Kế toán', 'Lê Thị Hoa', '9'],
  ['DV04', 'Phòng Kinh doanh', 'Trần Văn Lộc', '18'],
  ['DV05', 'Phòng Công nghệ thông tin', 'Phạm Quốc Anh', '11'],
  ['DV06', 'Phòng Tổ chức - Hành chính', 'Đỗ Thị Lan', '8'],
  ['DV07', 'Phòng Kế hoạch', 'Vũ Đình Nam', '7'],
  ['DV08', 'Ban Kiểm soát nội bộ', 'Hoàng Văn Sơn', '4'],
];

export const canBoData = [
  ['Nguyễn Thanh Đạm', 'Tổng Giám đốc', 'Ban Tổng Giám đốc', 'dam.nt@vietlott.vn'],
  ['Lê Văn Hùng', 'Phó Tổng Giám đốc', 'Ban Tổng Giám đốc', 'hung.lv@vietlott.vn'],
  ['Phạm Minh Tuấn', 'Phó Tổng Giám đốc', 'Ban Tổng Giám đốc', 'tuan.pm@vietlott.vn'],
  ['Trần Thị Mai', 'Chánh Văn phòng', 'Văn phòng', 'mai.tt@vietlott.vn'],
  ['Lê Thị Hoa', 'Trưởng phòng TC-KT', 'Phòng Tài chính - Kế toán', 'hoa.lt@vietlott.vn'],
  ['Trần Văn Lộc', 'Trưởng phòng KD', 'Phòng Kinh doanh', 'loc.tv@vietlott.vn'],
  ['Phạm Quốc Anh', 'Trưởng phòng CNTT', 'Phòng Công nghệ thông tin', 'anh.pq@vietlott.vn'],
  ['Nguyễn Thị Thu', 'Nhân viên văn thư', 'Văn phòng', 'thu.nt@vietlott.vn'],
];

export const loaiVBData = [
  ['QĐ', 'Quyết định', '.../QĐ-VIETLOTT', '128'],
  ['TB', 'Thông báo', '.../TB-VIETLOTT', '96'],
  ['TTr', 'Tờ trình', '.../TTr-VIETLOTT', '54'],
  ['BC', 'Báo cáo', '.../BC-VIETLOTT', '73'],
  ['UQ', 'Ủy quyền', '.../UQ-VIETLOTT', '21'],
  ['CT', 'Chỉ thị', '.../CT-VIETLOTT', '18'],
  ['GM', 'Giấy mời', '.../GM-VIETLOTT', '45'],
  ['CV', 'Công văn', '.../CV-VIETLOTT', '162'],
];

export const khoStore = {
  '512/QĐ-VIETLOTT': {
    info: { 'Số ký hiệu': '512/QĐ-VIETLOTT', 'Loại văn bản': 'Quyết định', 'Ngày văn bản': '27/06/2026', 'Người ký': 'Nguyễn Thanh Đạm', 'Trích yếu': 'V/v bổ nhiệm Trưởng phòng Công nghệ thông tin' },
    hoSo: 'HS.2026.QĐ.03', hoSoSub: 'Hồ sơ Quyết định nhân sự 2026',
    hop: 'Hộp số 12', hopSub: 'Nhóm văn bản đi 2026',
    thung: 'Thùng T-05', thungSub: 'Lưu trữ dài hạn',
    location: 'Kho tầng 2 · Dãy B · Kệ 04',
  },
  '168/TB-VIETLOTT': {
    info: { 'Số ký hiệu': '168/TB-VIETLOTT', 'Loại văn bản': 'Thông báo', 'Ngày văn bản': '26/06/2026', 'Người ký': 'Lê Văn Hùng', 'Trích yếu': 'Về lịch nghỉ lễ Quốc khánh 02/9/2026' },
    hoSo: 'HS.2026.TB.01', hoSoSub: 'Hồ sơ Thông báo chung 2026',
    hop: 'Hộp số 08', hopSub: 'Nhóm văn bản đi 2026',
    thung: 'Thùng T-03', thungSub: 'Lưu trữ 5 năm',
    location: 'Kho tầng 2 · Dãy A · Kệ 02',
  },
  '245': {
    info: { 'Số đến': '245', 'Nơi gửi': 'Bộ Tài chính', 'Số ký hiệu': '1123/BTC-TCT', 'Ngày văn bản': '25/06/2026', 'Trích yếu': 'V/v hướng dẫn quyết toán thuế TNDN năm 2025' },
    hoSo: 'HS.2026.ĐEN.06', hoSoSub: 'Hồ sơ văn bản đến Q2/2026',
    hop: 'Hộp số 21', hopSub: 'Văn bản đến 2026',
    thung: 'Thùng T-11', thungSub: 'Lưu trữ 5 năm',
    location: 'Kho tầng 1 · Dãy C · Kệ 07',
  },
};

export const khoSuggestions = ['512/QĐ-VIETLOTT', '168/TB-VIETLOTT', '245'];
