'use client';
import { useEffect, useState } from 'react';
import { PageHeader } from '@/components/ui/PageHeader';
import { messageService } from '@/services/messageService';
import { MessageThread } from '@/types/domain';

export function MessagesPage() {
  const [threads, setThreads] = useState<MessageThread[]>([]);
  const [selected, setSelected] = useState<MessageThread>();
  useEffect(() => { messageService.listThreads().then(data => { setThreads(data); setSelected(data[0]); }); }, []);
  return <><PageHeader title="Messages" subtitle="Customer conversations and campaign communication" actions={<button className="btn btn-accent">Broadcast campaign</button>} />
  <div className="chat"><div>{threads.map(t => <div key={t.id} className={`thread ${selected?.id === t.id ? 'active' : ''}`} onClick={() => setSelected(t)}><b>{t.customerName}</b>{t.unread && <span className="s-badge">new</span>}<div className="muted">{t.lastMessage}</div></div>)}</div><div><div className="card-title" style={{padding:14, margin:0}}>{selected?.customerName || 'Conversation'}<span className="muted">{selected?.vehicle}</span></div><div className="chat-body">{selected?.messages.length ? selected.messages.map((m,i) => <div key={i} className={`bubble ${m.direction === 'out' ? 'out' : ''}`}>{m.text}<div style={{fontSize:10, opacity:.7}}>{m.time}</div></div>) : <div className="muted">No messages yet.</div>}</div><div style={{padding:12, display:'flex', gap:8}}><input className="form-input" style={{flex:1}} placeholder="Type a message..."/><button className="btn btn-primary">Send</button></div></div></div></>;
}
