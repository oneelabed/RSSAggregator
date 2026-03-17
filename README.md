# Israel Conflict Monitor (ICM)
A high-performance, multi-lingual intelligence engine built in Go. ICM provides a 
unified, real-time dashboard for verified news sources, specifically designed to 
filter and aggregate critical updates during the ongoing conflict in Israel.

## 🚀 Features
- **Concurrent Scraper**: A background worker pool that uses goroutines and WaitGroups to fetch multiple RSS feeds simultaneously without blocking the main API.
- **Type-Safe SQL**: Utilizes SQLC to generate compile-time safe Go code from raw PostgreSQL queries.
- **Database Migrations**: Managed schema evolution using Goose to handle versioned SQL migrations.
- **Custom Authentication**: Secure API access via API Key-based authentication implemented through custom middleware.
- **Automated Scheduling**: Intelligent fetching logic that prioritizes feeds based on their last fetched timestamp.
- **Full-Text Search**: Integrated PostgreSQL's full-text search engine to allow users to filter and find specific posts by title or description across all followed feeds.

- ## 🛠️ Tech Stack
- **Backend**: Go (Golang) with Chi Router
- **Frontend**: Next.js 14, TypeScript, Tailwind CSS
- **UI Components**: Shadcn UI & Lucide Icons
- **Database**: PostgreSQL
- **SQL Toolkit**: SQLC (Type-safe query generation)
- **Migrations**: Goose

## 🔒 Authentication
Unique API Keys are generated upon user registration and required in the Authorization header to verify identity for all protected routes.

## 🌐 Deployment
Frontend → Vercel \
Backend → Render \
Database → PostgreSQL (Neon) \
Configured with a custom domain + SSL.

## Link
[Israel Conflict Monitor](https://israelconflictmonitor.org)
