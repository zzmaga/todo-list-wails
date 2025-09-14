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
        <div className="header-left">
          <h1>📝 Мои Задачи</h1>
          <div className="app-subtitle">Управляйте своими задачами эффективно</div>
        </div>
        <div className="spacer" />
        <div className="header-actions">
          <div className="task-stats">
            <span className="stat-item">
              <span className="stat-icon">📋</span>
              <span className="stat-text">Всего: {tasks.length}</span>
            </span>
            <span className="stat-item">
              <span className="stat-icon">⏳</span>
              <span className="stat-text">Активных: {activeCount}</span>
            </span>
            <span className="stat-item">
              <span className="stat-icon">✅</span>
              <span className="stat-text">Выполнено: {doneCount}</span>
            </span>
          </div>
          <button 
            className="btn theme-toggle" 
            onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}
            title={`Переключить на ${theme === 'light' ? 'тёмную' : 'светлую'} тему`}
          >
            {theme === 'light' ? '🌙' : '☀️'}
          </button>
        </div>
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
          <div className="task-column">
            <div className="column-header">
              <h2>⏳ Активные задачи</h2>
              <span className="task-count">{active.length}</span>
            </div>
            {active.length === 0 ? (
              <div className="empty-state">
                <div className="empty-icon">📝</div>
                <div className="empty-text">Нет активных задач</div>
                <div className="empty-subtext">Нажмите "Добавить задачу" чтобы создать первую</div>
              </div>
            ) : (
              <div className="task-list">
                {active.map((t) => (
                  <TaskItem key={t.id} t={t} onToggle={toggle} onDelete={onDelete} />
                ))}
              </div>
            )}
          </div>
          
          <div className="task-column">
            <div className="column-header">
              <h2>✅ Выполненные задачи</h2>
              <span className="task-count">{done.length}</span>
            </div>
            {done.length === 0 ? (
              <div className="empty-state">
                <div className="empty-icon">🎯</div>
                <div className="empty-text">Нет выполненных задач</div>
                <div className="empty-subtext">Отмечайте задачи как выполненные</div>
              </div>
            ) : (
              <div className="task-list">
                {done.map((t) => (
                  <TaskItem key={t.id} t={t} onToggle={toggle} onDelete={onDelete} />
                ))}
              </div>
            )}
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
