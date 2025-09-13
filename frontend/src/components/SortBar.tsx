import { Sort } from '../lib/types';

export default function SortBar({ value, onChange }: { value: Sort; onChange: (s: Sort)=>void }) {
  return (
    <div className="sortbar">
      <label>Сортировка:</label>
      <select value={value} onChange={e=>onChange(e.target.value as Sort)}>
        <option value="createdAsc">По дате добавления ↑</option>
        <option value="createdDesc">По дате добавления ↓</option>
        <option value="priority">По приоритету</option>
        <option value="due">По дедлайну</option>
      </select>
    </div>
  );
}
