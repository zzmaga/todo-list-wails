import { useState } from 'react';
import type { Priority } from '../lib/types';

interface AddTaskFormProps {
  onAdd: (title: string, dueISO?: string, priority?: Priority) => void;
}

export default function AddTaskForm({ onAdd }: AddTaskFormProps) {
  const [isExpanded, setIsExpanded] = useState(false);
  const [title, setTitle] = useState('');
  const [due, setDue] = useState<string>('');
  const [priority, setPriority] = useState<Priority>('medium');

  const submit = (e: React.FormEvent) => {
    e.preventDefault();
    const trimmed = title.trim();
    if (!trimmed) return; // валидация пустого ввода
    onAdd(trimmed, due ? new Date(due).toISOString() : undefined, priority);
    setTitle(''); 
    setDue(''); 
    setPriority('medium');
    setIsExpanded(false);
  };

  const cancel = () => {
    setTitle('');
    setDue('');
    setPriority('medium');
    setIsExpanded(false);
  };

  // Если форма свернута, показываем только кнопку
  if (!isExpanded) {
    return (
      <div className="add-task-button">
        <button 
          className="btn primary add-btn" 
          onClick={() => setIsExpanded(true)}
        >
          ➕ Добавить задачу
        </button>
      </div>
    );
  }

  // Если форма развернута, показываем полную форму
  return (
    <div className="add-task-form-container">
      <form className="add-form" onSubmit={submit}>
        <div className="form-header">
          <h3>Новая задача</h3>
          <button 
            type="button" 
            className="btn icon" 
            onClick={cancel}
            title="Отмена"
          >
            ✕
          </button>
        </div>
        
        <div className="form-fields">
          <input 
            className="input" 
            placeholder="Введите название задачи..." 
            value={title} 
            onChange={e => setTitle(e.target.value)}
            autoFocus
          />
          
          <div className="form-row">
            <input 
              type="datetime-local" 
              value={due} 
              onChange={e => setDue(e.target.value)} 
              title="Дедлайн (необязательно)"
              className="input"
            />
            
            <select 
              value={priority} 
              onChange={e => setPriority(e.target.value as Priority)}
              className="input"
            >
              <option value="low">🟢 Низкий</option>
              <option value="medium">🟡 Средний</option>
              <option value="high">🔴 Высокий</option>
            </select>
          </div>
        </div>
        
        <div className="form-actions">
          <button type="button" className="btn" onClick={cancel}>
            Отмена
          </button>
          <button className="btn primary" type="submit">
            ✅ Добавить задачу
          </button>
        </div>
      </form>
    </div>
  );
}
