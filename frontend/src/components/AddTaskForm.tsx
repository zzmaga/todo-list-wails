import { useState } from 'react';
import type { Priority } from '../lib/types';

export default function AddTaskForm({ onAdd }: { onAdd: (title: string, dueISO?: string, priority?: Priority)=>void }) {
  const [title, setTitle] = useState('');
  const [due, setDue] = useState<string>('');
  const [priority, setPriority] = useState<Priority>('medium');

  const submit = (e: React.FormEvent) => {
    e.preventDefault();
    const trimmed = title.trim();
    if (!trimmed) return; // валидация пустого ввода
    onAdd(trimmed, due ? new Date(due).toISOString() : undefined, priority);
    setTitle(''); setDue(''); setPriority('medium');
  };

  return (
    <form className="add-form" onSubmit={submit}>
      <input className="input" placeholder="Новая задача…" value={title} onChange={e=>setTitle(e.target.value)} />
      <input type="datetime-local" value={due} onChange={e=>setDue(e.target.value)} title="Дедлайн" />
      <select value={priority} onChange={e=>setPriority(e.target.value as Priority)}>
        <option value="low">Низкий</option>
        <option value="medium">Средний</option>
        <option value="high">Высокий</option>
      </select>
      <button className="btn primary" type="submit">Добавить</button>
    </form>
  );
}
