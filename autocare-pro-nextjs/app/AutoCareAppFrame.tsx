"use client";

import { useEffect, useState } from "react";

export default function AutoCareAppFrame() {
  const [srcDoc, setSrcDoc] = useState<string>("");

  useEffect(() => {
    fetch("/autocare_pro_nextjs_webapp.html")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Unable to load AutoCare Pro HTML asset");
        }
        return response.text();
      })
      .then(setSrcDoc)
      .catch((error) => {
        setSrcDoc(`<!doctype html><html><body style="font-family:sans-serif;padding:24px"><h1>AutoCare Pro</h1><p>${error.message}</p></body></html>`);
      });
  }, []);

  return (
    <main style={{ width: "100vw", height: "100vh", overflow: "hidden" }}>
      <iframe
        title="AutoCare Pro Web App"
        srcDoc={srcDoc}
        style={{ width: "100%", height: "100%", border: 0, display: "block" }}
      />
    </main>
  );
}
