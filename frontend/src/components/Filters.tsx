import { Filter } from '../lib/types';

export default function Filters({ value, onChange }: { value: Filter; onChange: (f: Filter)=>void }) {
  return (
    <div className="filters">
      <button className={value==='all'? 'active':''} onClick={()=>onChange('all')}>Все</button>
      <button className={value==='active'? 'active':''} onClick={()=>onChange('active')}>Активные</button>
      <button className={value==='done'? 'active':''} onClick={()=>onChange('done')}>Выполненные</button>
    </div>
  );
}
