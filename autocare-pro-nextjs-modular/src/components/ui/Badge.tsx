export function Badge({ children, tone = 'gray' }: { children: React.ReactNode; tone?: 'green'|'blue'|'amber'|'gray'|'red' }) {
  return <span className={`badge ${tone}`}>{children}</span>;
}
