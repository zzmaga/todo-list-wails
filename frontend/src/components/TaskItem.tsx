import { Task } from '../lib/types';

interface TaskItemProps {
  t: Task;
  onToggle: (id: string) => void;
  onDelete: (id: string) => void;
}

export default function TaskItem({ t, onToggle, onDelete }: TaskItemProps) {
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    const now = new Date();
    const diffTime = date.getTime() - now.getTime();
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    
    // Для даты создания показываем только дату
    if (dateString === t.createdAt) {
      return date.toLocaleDateString('ru-RU', { 
        day: 'numeric', 
        month: 'short' 
      });
    }
    
    // Для дедлайна показываем относительное время
    if (diffDays === 0) return 'Сегодня';
    if (diffDays === 1) return 'Завтра';
    if (diffDays === -1) return 'Вчера';
    if (diffDays > 0) return `Через ${diffDays} дн.`;
    return `${Math.abs(diffDays)} дн. назад`;
  };

  const isOverdue = t.dueAt && new Date(t.dueAt) < new Date() && !t.done;

  return (
    <div className={`task ${t.done ? 'done' : ''} ${isOverdue ? 'overdue' : ''}`}>
      <button 
        className={`checkbox ${t.done ? 'checked' : ''}`}
        onClick={() => onToggle(t.id)}
        title={t.done ? 'Отметить как невыполненную' : 'Отметить как выполненную'}
      >
        <span className="icon">{t.done ? '✓' : '○'}</span>
      </button>
      
      <div className="meta">
        <div className="title">
          <span className={`priority-badge ${t.priority}`}>
            {t.priority === 'high' ? '●' : t.priority === 'medium' ? '●' : '●'}
          </span>
          <span className="task-title">{t.title}</span>
          {isOverdue && <span className="overdue-badge">! Просрочено</span>}
        </div>
        
        <div className="sub">
          <span className="created-date">
            <span className="icon">📅</span> {formatDate(t.createdAt)}
          </span>
          {t.dueAt && (
            <span className={`due-date ${isOverdue ? 'overdue' : ''}`}>
              <span className="icon">⏰</span> {formatDate(t.dueAt)}
            </span>
          )}
        </div>
      </div>
      
      <button 
        className="btn icon danger" 
        title="Удалить задачу" 
        onClick={() => onDelete(t.id)}
      >
        <span className="icon">🗑</span>
      </button>
    </div>
  );
}
