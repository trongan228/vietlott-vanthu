import { IconCheck } from './Icons.jsx';

export default function Toast({ message }) {
  if (!message) return null;
  return (
    <div className="vl-toast">
      <span style={{ width: 22, height: 22, borderRadius: '50%', background: '#1B7A43', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <IconCheck size={13} width={3} />
      </span>
      {message}
    </div>
  );
}
