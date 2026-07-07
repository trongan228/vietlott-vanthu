import { IconDashboard, IconInbox, IconArchive, IconList } from './Icons.jsx';

const items = [
  { key: 'dashboard', label: 'Tổng quan', icon: <IconDashboard /> },
  { key: 'den', label: 'Văn bản đến', icon: <IconInbox arrow="in" /> },
  { key: 'di', label: 'Văn bản đi', icon: <IconInbox arrow="out" /> },
  { key: 'luukho', label: 'Tra cứu lưu kho', icon: <IconArchive /> },
  { key: 'danhmuc', label: 'Danh mục', icon: <IconList /> },
];

export default function NavBar({ screen, onGo }) {
  return (
    <nav className="vl-nav">
      {items.map(it => (
        <button
          key={it.key}
          className={'vl-nav-item' + (screen === it.key ? ' active' : '')}
          onClick={() => onGo(it.key)}
        >
          {it.icon}
          {it.label}
        </button>
      ))}
    </nav>
  );
}
