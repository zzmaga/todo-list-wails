import { useEffect, useState } from 'react';
import { useTasks } from './state/useTasks';
import AddTaskForm from './components/AddTaskForm';
import TaskItem from './components/TaskItem';
import Filters from './components/Filters';
import SortBar from './components/SortBar';
import ConfirmDialog from './components/ConfirmDialog';
import './style.css';

export default function App() {
  const {
    tasks,
    loading,
    create,
    toggle,
    remove,
    filter,
    sort,
    setFilter,
    setSort,
    activeCount,
    doneCount
  } = useTasks();

  const [theme, setTheme] = useState<'light' | 'dark'>(
    () => (localStorage.getItem('theme') as 'light' | 'dark') || 'light'
  );
  const [confirm, setConfirm] = useState<{ open: boolean; id?: string }>({ open: false });

  const onDelete = (id: string) => setConfirm({ open: true, id });
  const doDelete = () => {
    if (confirm.id) remove(confirm.id);
    setConfirm({ open: false });
  };

  // корректно: сайд-эффект через useEffect
  useEffect(() => {
    document.documentElement.dataset.theme = theme;
    localStorage.setItem('theme', theme);
  }, [theme]);

  const active = tasks.filter((t) => !t.done);
  const done = tasks.filter((t) => t.done);

  return (
    <div className="container">
      <header>
        <h1>To-Do (Wails)</h1>
        <div className="spacer" />
        <button className="btn" onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}>
          {theme === 'light' ? '🌙 Тёмная' : '☀️ Светлая'}
        </button>
      </header>

      <AddTaskForm onAdd={create} />

      <section className="toolbar">
        <Filters value={filter} onChange={setFilter} />
        <SortBar value={sort} onChange={setSort} />
        <div className="counts">
          Активные: {activeCount} • Выполненные: {doneCount}
        </div>
      </section>

      {loading ? (
        <div>Загрузка…</div>
      ) : (
        <div className="columns">
          <div>
            <h2>Активные</h2>
            {active.map((t) => (
              <TaskItem key={t.id} t={t} onToggle={toggle} onDelete={onDelete} />
            ))}
          </div>
          <div>
            <h2>Выполненные</h2>
            {done.map((t) => (
              <TaskItem key={t.id} t={t} onToggle={toggle} onDelete={onDelete} />
            ))}
          </div>
        </div>
      )}

      <ConfirmDialog
        open={confirm.open}
        text="Удалить задачу?"
        onOk={doDelete}
        onCancel={() => setConfirm({ open: false })}
      />
    </div>
  );
}
