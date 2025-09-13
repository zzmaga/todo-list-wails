import type { Task, Priority } from './types';

export async function list(): Promise<Task[]> {
  // @ts-ignore
  return await window.go.backend.App.List();
}

export async function add(title: string, dueISO?: string, priority: Priority = 'medium'): Promise<Task> {
  // @ts-ignore
  return await window.go.backend.App.Add(title, dueISO ?? null, priority);
}

export async function toggle(id: string) {
  // @ts-ignore
  return await window.go.backend.App.Toggle(id);
}

export async function remove(id: string) {
  // @ts-ignore
  return await window.go.backend.App.Delete(id);
}

export async function update(task: Task) {
  // @ts-ignore
  return await window.go.backend.App.Update(task);
}
