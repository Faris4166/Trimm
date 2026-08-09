# Trimm - Project Documentation

This project is a Full-Stack Web Application separated into Frontend, Backend, and Database components. It is containerized and can be easily managed and run using Docker Compose.

## 🛠️ Tech Stack & Frameworks

### 1. Frontend
Developed with modern web technologies, focusing on performance, beautiful UI, and responsiveness.
- **Next.js (v16.3)**: The core React framework used for Server-Side Rendering (SSR) and internal routing.
- **React (v19.2)**: The core library for building UI components.
- **TypeScript**: Adds static typing to JavaScript, ensuring type safety and reducing potential bugs.
- **Tailwind CSS (v4)**: A utility-first CSS framework for rapid and highly customizable styling.
- **Shadcn UI & Base UI**: Beautiful, accessible, and flexible pre-built UI components.
- **Axios**: A promise-based HTTP client for making API requests to the Backend.
- **Lucide React**: A collection of clean and minimalist icons.

### 2. Backend
Developed in Go (Golang) for high-performance and lightweight resource usage.
- **Go (v1.26)**: The primary programming language used for the backend logic.
- **Gin Framework (`gin-gonic/gin`)**: A fast and lightweight Web framework for Go, used for building RESTful APIs and managing routes.
- **GORM (`gorm.io/gorm`)**: An ORM (Object Relational Mapping) library for Go, used to interact with the database using Go structs.
- **PostgreSQL Driver (`gorm.io/driver/postgres`)**: Used alongside GORM to connect to the PostgreSQL database.
- **Godotenv**: Used to load configuration and environment variables from a `.env` file.
- **Gin CORS**: A middleware for Gin to handle CORS (Cross-Origin Resource Sharing), allowing the frontend to communicate with the backend APIs securely.

### 3. Database
- **PostgreSQL (v16)**: A powerful, open-source object-relational database system used as the primary data store for the application.

### 4. Infrastructure & Deployment
- **Docker**: Used for containerizing the application, ensuring that it runs consistently across different environments.
- **Docker Compose**: Used to define and run the multi-container application (Frontend, Backend, and Database) with a single command.

---

## 🚀 How to Run the Project

### Prerequisites
1. Ensure you have **Docker** and **Docker Compose** installed on your machine.
2. Ensure that ports `8080` (Backend), `3000` (Frontend), and `5432` (Database) are available and not being used by other services.

### Environment Variables Setup (.env)
Before running the project, you need to set up the environment variables for both the backend and the frontend.

**1. Backend Environment Variables:**
Create a `.env` file inside the `backend` directory (`backend/.env`) and add the following database configuration:
```env
DB_USERNAME=appuser
DB_PASSWORD=apppass
DB_HOST=postgres
DB_PORT=5432
DB_DATABASE=appdb
```

**2. Frontend Environment Variables:**
Create a `.env.local` file inside the `frontend` directory (`frontend/.env.local`) and add the backend API URL:
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

### Running the Application (Docker Compose)

1. Open your Terminal or Command Prompt.
2. Navigate to the root directory of the project (where the `docker-compose.yml` file is located).
3. Run the following command to build and start the containers in detached mode:
   ```bash
   docker-compose up --build -d
   ```
4. Wait a moment for the containers to start. You can check their status using:
   ```bash
   docker-compose ps
   ```

**Accessing the Application:**
- **Frontend (Web UI):** Open your browser and go to `http://localhost:3000`
- **Backend API:** The API will be accessible at `http://localhost:8080`

**Stopping the Application:**
To stop all running containers, execute:
```bash
docker-compose down
```
*(If you also want to remove the database volumes and clear the data, use `docker-compose down -v`)*

---

### Project Structure
- `/frontend`: Contains the Next.js frontend application code.
- `/backend`: Contains the Go + Gin backend application code.
- `docker-compose.yml`: The configuration file used by Docker Compose to orchestrate the services.
