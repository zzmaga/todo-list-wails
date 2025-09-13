import { Task } from '../lib/types';

export default function TaskItem({
  t,
  onToggle,
  onDelete
}: {
  t: Task;
  onToggle: (id: string) => void;
  onDelete: (id: string) => void;
}) {
  return (
    <div className={`task ${t.done ? 'done' : ''}`}>
      <input type="checkbox" checked={t.done} onChange={() => onToggle(t.id)} />
      <div className="meta">
        <div className="title">
          <span className={`prio ${t.priority}`}></span>
          {t.title}
        </div>
        <div className="sub">
          Создано: {new Date(t.createdAt).toLocaleString()}{' '}
          {t.dueAt && <> • Дедлайн: {new Date(t.dueAt).toLocaleString()}</>}
        </div>
      </div>
      <button className="icon danger" title="Удалить" onClick={() => onDelete(t.id)}>
        🗑️
      </button>
    </div>
  );
}
