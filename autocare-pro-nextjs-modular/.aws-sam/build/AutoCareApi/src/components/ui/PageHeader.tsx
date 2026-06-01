export function PageHeader({ title, subtitle, actions }: { title: string; subtitle?: string; actions?: React.ReactNode }) {
  return <div className="ph"><div><div className="ph-title">{title}</div>{subtitle && <div className="ph-sub">{subtitle}</div>}</div><div>{actions}</div></div>;
}
