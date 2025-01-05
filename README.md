# Real-Time Collaborative Task Manager

## Brief Description
A collaborative task management application where multiple users can create, edit, and complete tasks in real-time. This application will leverage WebSocket for real-time communication, RabbitMQ for asynchronous message handling, and Redis for data caching or user session storage.

## Key Features
- **Real-Time Task Updates (WebSocket):**
    1. All users in the team can see changes to the task list instantly.
    2. For example, if user A adds a new task, user B will immediately see the task without needing to refresh the page.
- **Notification Queue (RabbitMQ):**
    1. When a task is created, updated, or completed, notifications are sent to all team members through the RabbitMQ message queue.
    2. This also allows for sending individual emails/messages to users.
- **Caching for Faster Performance (Redis):**
    1. Frequently accessed task lists are temporarily stored in Redis to speed up data retrieval.
    2. Redis is also used for more efficient user session management.

## Architecture
- **Frontend:**
    - Utilizes Vue.js for a real-time interface that connects to the backend via WebSocket.
- **Backend:**
    - Built with Go Language (using frameworks Fiber).
- **Database:**
    - Employs PostgreSQL for primary data storage.
- **Redis:**
    - Used as a cache and for storing user sessions.
- **RabbitMQ:**
    - For managing message queues such as sending notifications or processing background jobs.
