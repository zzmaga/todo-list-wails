export type Priority = 'low' | 'medium' | 'high';

export interface Task {
  id: string;
  title: string;
  done: boolean;
  createdAt: string;     // ISO
  dueAt?: string | null; // ISO | null
  priority: Priority;
}

export type Filter = 'all' | 'active' | 'done';
export type Sort = 'createdAsc' | 'createdDesc' | 'priority' | 'due';
