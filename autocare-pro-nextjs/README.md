# AutoCare Pro Next.js Web App

This project wraps the provided AutoCare Pro HTML application inside a Next.js app shell.

## Run locally

```bash
npm install
npm run dev
```

Open http://localhost:3000

## Structure

- `app/page.tsx` - Next.js home page
- `app/AutoCareAppFrame.tsx` - client component that loads the HTML app
- `public/autocare_pro_nextjs_webapp.html` - original AutoCare Pro HTML web app

## Notes

This keeps the original UI and JavaScript behavior intact. For a production-grade implementation, the next step would be to refactor the original HTML pages into reusable React components and connect real APIs.
