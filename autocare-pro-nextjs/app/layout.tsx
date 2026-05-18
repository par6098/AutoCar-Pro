import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "AutoCare Pro",
  description: "AutoCare Pro service booking and operations web app"
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
