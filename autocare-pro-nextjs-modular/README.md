# AutoCare Pro - Modular Next.js App

This version breaks the original single-page HTML into a modular Next.js application with feature modules, typed domain models, and service classes ready to consume backend APIs.

## Run

```bash
npm install
npm run dev
```

Open `http://localhost:3000`.

## Backend integration

Set `.env.local`:

```env
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080/api
NEXT_PUBLIC_USE_MOCKS=false
```

Expected API endpoints:

- `GET /dashboard/summary`
- `GET /bookings`
- `POST /bookings`
- `PATCH /bookings/{id}`
- `GET /packages`
- `GET /employees`
- `GET /messages/threads`
- `GET /billing/transactions`

Services are in `src/services/*Service.ts`. Each service calls `httpClient` and falls back to mock data when `NEXT_PUBLIC_USE_MOCKS=true`.

## Structure

```text
src/app                 Next.js app entry
src/components/layout   Shell, sidebar, topbar
src/components/ui       Reusable UI components
src/features            Feature modules/pages
src/services            Backend-facing service layer
src/types               Shared domain types
src/config              API configuration
```
