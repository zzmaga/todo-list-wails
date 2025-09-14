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
  const [error, setError] = useState<string>('');

  const submit = (e: React.FormEvent) => {
    e.preventDefault();
    const trimmed = title.trim();
    
    // Валидация
    if (!trimmed) {
      setError('Название задачи не может быть пустым');
      return;
    }
    
    if (trimmed.length > 200) {
      setError('Название задачи слишком длинное (максимум 200 символов)');
      return;
    }
    
    // Проверка дедлайна
    if (due && new Date(due) < new Date()) {
      setError('Дедлайн не может быть в прошлом');
      return;
    }
    
    setError(''); // Очищаем ошибку при успешной валидации
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
    setError('');
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
          <span className="icon">+</span> Добавить задачу
        </button>
      </div>
    );
  }

  // Если форма развернута, показываем полную форму
  return (
    <div className="add-task-form-container">
      <form className="add-form" onSubmit={submit}>
        
        <div className="form-fields">
          <input 
            className={`input ${error ? 'error' : ''}`}
            placeholder="Введите название задачи..." 
            value={title} 
            onChange={e => {
              setTitle(e.target.value);
              if (error) setError(''); // Очищаем ошибку при вводе
            }}
            autoFocus
          />
          {error && <div className="error-message">{error}</div>}
          
          <div className="form-row">
            <div className="input-group">
              <label className="input-label">Дедлайн:</label>
              <input 
                type="datetime-local" 
                value={due} 
                onChange={e => setDue(e.target.value)} 
                title="Дедлайн (необязательно)"
                className="input"
              />
            </div>
            
            <div className="input-group">
              <label className="input-label">Приоритет:</label>
              <select 
                value={priority} 
                onChange={e => setPriority(e.target.value as Priority)}
                className="input"
              >
                <option value="low">● Низкий</option>
                <option value="medium">● Средний</option>
                <option value="high">● Высокий</option>
              </select>
            </div>
          </div>
        </div>
        
        <div className="form-actions">
          <button type="button" className="btn cancel-btn" onClick={cancel}>
            <span className="icon">↩</span> Отмена
          </button>
          <button className="btn primary" type="submit">
            <span className="icon">✓</span> Добавить задачу
          </button>
        </div>
      </form>
    </div>
  );
}
