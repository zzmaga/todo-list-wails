export default function ConfirmDialog({ open, text, onOk, onCancel }:{
    open:boolean; text:string; onOk:()=>void; onCancel:()=>void
  }){
    if (!open) return null;
    return (
      <div className="modal">
        <div className="modal-content">
          <p>{text}</p>
          <div className="row">
            <button className="btn" onClick={onCancel}>Отмена</button>
            <button className="btn danger" onClick={onOk}>Удалить</button>
          </div>
        </div>
      </div>
    );
  }