import { useEffect, useReducer } from 'react';
import type { Task, Filter, Sort } from '../lib/types';
import * as api from '../lib/api';

interface State {
  tasks: Task[];
  filter: Filter;
  sort: Sort;
  loading: boolean;
}

type Action =
  | { type: 'set'; tasks: Task[] }
  | { type: 'loading'; value: boolean }
  | { type: 'filter'; value: Filter }
  | { type: 'sort'; value: Sort };

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case 'set': return { ...state, tasks: action.tasks };
    case 'loading': return { ...state, loading: action.value };
    case 'filter': return { ...state, filter: action.value };
    case 'sort': return { ...state, sort: action.value };
    default: return state;
  }
}

export function useTasks() {
  const [state, dispatch] = useReducer(reducer, {
    tasks: [], filter: 'all', sort: 'createdAsc', loading: true
  });

  useEffect(() => { (async () => {
    dispatch({ type: 'loading', value: true });
    const t = await api.list();
    dispatch({ type: 'set', tasks: t });
    dispatch({ type: 'loading', value: false });
  })(); }, []);

  const create = async (title: string, dueISO?: string, priority: 'low'|'medium'|'high' = 'medium') => {
    const t = await api.add(title, dueISO, priority);
    dispatch({ type: 'set', tasks: [...state.tasks, t] });
  };

  const toggle = async (id: string) => {
    await api.toggle(id);
    const tasks = state.tasks.map(t => t.id === id ? { ...t, done: !t.done } : t);
    dispatch({ type: 'set', tasks });
  };

  const remove = async (id: string) => {
    await api.remove(id);
    dispatch({ type: 'set', tasks: state.tasks.filter(t => t.id !== id) });
  };

  const setFilter = (value: Filter) => dispatch({ type: 'filter', value });
  const setSort = (value: Sort) => dispatch({ type: 'sort', value });

  const view = (() => {
    let list = [...state.tasks];
    if (state.filter === 'active') list = list.filter(t => !t.done);
    if (state.filter === 'done') list = list.filter(t => t.done);
    switch (state.sort) {
      case 'createdAsc': list.sort((a,b)=> a.createdAt.localeCompare(b.createdAt)); break;
      case 'createdDesc': list.sort((a,b)=> b.createdAt.localeCompare(a.createdAt)); break;
      case 'priority': {
        const rank = { high: 0, medium: 1, low: 2 } as const;
        list.sort((a,b)=> rank[a.priority] - rank[b.priority]);
        break;
      }
      case 'due': list.sort((a,b)=> (a.dueAt||'').localeCompare(b.dueAt||'')); break;
    }
    return list;
  })();

  const active = state.tasks.filter(t=>!t.done);
  const done = state.tasks.filter(t=>t.done);

  return { ...state, tasks: view, activeCount: active.length, doneCount: done.length, create, toggle, remove, setFilter, setSort };
}
