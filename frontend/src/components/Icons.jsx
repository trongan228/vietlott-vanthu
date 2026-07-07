// Bộ icon SVG dùng chung, stroke theo currentColor trừ khi truyền màu riêng
export const IconLogo = () => (
  <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
    <path d="M6 20V9l6-4 6 4v11" stroke="#D71920" strokeWidth="2" strokeLinejoin="round" />
    <path d="M9 20v-5h6v5" stroke="#D71920" strokeWidth="2" strokeLinejoin="round" />
    <circle cx="12" cy="4" r="1.6" fill="#D71920" />
  </svg>
);

export const IconSearch = ({ size = 16, color = 'currentColor', style }) => (
  <svg style={style} width={size} height={size} viewBox="0 0 24 24" fill="none">
    <circle cx="11" cy="11" r="7" stroke={color} strokeWidth="2" />
    <path d="m20 20-3.5-3.5" stroke={color} strokeWidth="2" strokeLinecap="round" />
  </svg>
);

export const IconBell = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M18 8a6 6 0 1 0-12 0c0 7-3 9-3 9h18s-3-2-3-9" stroke="#fff" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
    <path d="M13.7 21a2 2 0 0 1-3.4 0" stroke="#fff" strokeWidth="2" strokeLinecap="round" />
  </svg>
);

export const IconDashboard = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <rect x="3" y="3" width="8" height="8" rx="1.5" stroke="currentColor" strokeWidth="2" />
    <rect x="13" y="3" width="8" height="5" rx="1.5" stroke="currentColor" strokeWidth="2" />
    <rect x="13" y="10" width="8" height="11" rx="1.5" stroke="currentColor" strokeWidth="2" />
    <rect x="3" y="13" width="8" height="8" rx="1.5" stroke="currentColor" strokeWidth="2" />
  </svg>
);

export const IconInbox = ({ arrow = 'in' }) => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M4 13h4l2 3h4l2-3h4" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
    <path d="M4 13V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v7" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
    <path d="M4 13v3a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-3" stroke="currentColor" strokeWidth="2" />
    {arrow === 'in'
      ? <path d="M12 4v6m0 0-2-2m2 2 2-2" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
      : <path d="M12 10V4m0 0-2 2m2-2 2 2" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />}
  </svg>
);

export const IconArchive = ({ size = 18 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none">
    <path d="M3 7l2-3h14l2 3" stroke="currentColor" strokeWidth="2" strokeLinejoin="round" />
    <rect x="3" y="7" width="18" height="13" rx="1.5" stroke="currentColor" strokeWidth="2" />
    <path d="M9 11h6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
  </svg>
);

export const IconList = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M4 6h16M4 12h16M4 18h16" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
    <circle cx="8" cy="6" r="1.6" fill="currentColor" />
    <circle cx="14" cy="12" r="1.6" fill="currentColor" />
    <circle cx="10" cy="18" r="1.6" fill="currentColor" />
  </svg>
);

export const IconExport = ({ size = 16 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none">
    <path d="M12 3v12m0 0 4-4m-4 4-4-4" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
    <path d="M4 17v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
  </svg>
);

export const IconImport = ({ size = 16 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none">
    <path d="M12 15V3m0 12-4-4m4 4 4-4" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" transform="rotate(180 12 9)" />
    <path d="M4 17v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
  </svg>
);

export const IconPlus = () => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
    <path d="M12 5v14M5 12h14" stroke="#fff" strokeWidth="2.2" strokeLinecap="round" />
  </svg>
);

export const IconCheck = ({ size = 16, width = 2.2 }) => (
  <svg width={size} height={size} viewBox="0 0 24 24" fill="none">
    <path d="M5 12l4 4L19 6" stroke="#fff" strokeWidth={width} strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconBack = () => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
    <path d="M15 18l-6-6 6-6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconClose = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M6 6l12 12M18 6 6 18" stroke="#5a5a62" strokeWidth="2" strokeLinecap="round" />
  </svg>
);

export const IconChevronRight = () => (
  <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
    <path d="M9 6l6 6-6 6" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconEdit = () => (
  <svg width="15" height="15" viewBox="0 0 24 24" fill="none">
    <path d="M4 20h4L18 10l-4-4L4 16v4Z" stroke="#5a5a62" strokeWidth="2" strokeLinejoin="round" />
    <path d="m13 7 4 4" stroke="#5a5a62" strokeWidth="2" />
  </svg>
);

export const IconTrash = () => (
  <svg width="15" height="15" viewBox="0 0 24 24" fill="none">
    <path d="M5 7h14M9 7V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2m-8 0v12a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V7" stroke="#c0392b" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconPin = () => (
  <svg width="20" height="20" viewBox="0 0 24 24" fill="none">
    <path d="M12 21s7-5.5 7-11a7 7 0 1 0-14 0c0 5.5 7 11 7 11Z" stroke="#D71920" strokeWidth="2" strokeLinejoin="round" />
    <circle cx="12" cy="10" r="2.5" stroke="#D71920" strokeWidth="2" />
  </svg>
);

export const IconDoc = ({ color = '#D71920' }) => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
    <path d="M14 3v5h5M14 3H6a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V8l-5-5Z" stroke={color} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconFolder = ({ color = '#63636b' }) => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
    <path d="M3 7a1 1 0 0 1 1-1h5l2 2h9a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7Z" stroke={color} strokeWidth="2" strokeLinejoin="round" />
  </svg>
);

export const IconBox = ({ color = '#63636b' }) => (
  <svg width="16" height="16" viewBox="0 0 24 24" fill="none">
    <path d="M3 7l2-3h14l2 3M3 7h18v12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7Zm6 4h6" stroke={color} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconStatDen = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M4 13h4l2 3h4l2-3h4M4 13V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-3" stroke="#D71920" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconStatDi = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M4 13h4l2 3h4l2-3h4M4 13V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-3" stroke="#C77A11" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconClock = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M12 7v5l3 2M12 21a9 9 0 1 1 0-18 9 9 0 0 1 0 18Z" stroke="#C77A11" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);

export const IconDone = () => (
  <svg width="18" height="18" viewBox="0 0 24 24" fill="none">
    <path d="M5 12l4 4L19 6" stroke="#1B8A4B" strokeWidth="2.4" strokeLinecap="round" strokeLinejoin="round" />
  </svg>
);
