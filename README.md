## TaskFlow 

#### Live Demo

- **Frontend (Vercel):** [https://z-frontend-kappa.vercel.app/signup](https://z-frontend-kappa.vercel.app/signup)
- **Backend (Render/Fly.io):** [z-backend-production.up.railway.app](z-backend-production.up.railway.app)


#### Tech Stack

#### Backend
- Golang + Fiber
- JWT Authentication (HTTP-only cookies)
- MySQL + GORM
- WebSockets for real-time task sync
- Docker + Docker Compose

#### Frontend
- Next.js (App Router)
- Tailwind CSS (shadcn-inspired)
- react-hot-toast for UX feedback
- WebSocket client for live updates




#### Deployment

- **Frontend:** Deployed on Vercel
- **Backend:** Deployed on Fly.io or Render
- **MySQL:** Docker container


#### AI Tools Used

- **ChatGPT**: code scaffolding, refactoring, and feature guidance
- **GitHub Copilot**: UI components and TypeScript typing
- **OpenAI API**: Used in `/api/task/ai-chat` to generate task suggestions

