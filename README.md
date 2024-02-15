# Fiber Go Boilerplate

Is a production-ready golang based boilerplate application with Fiber Go web framework.
For any fiber go application, just clone the repo & rename the application name.

[Fiber](https://gofiber.io/) is an Express.js inspired web framework build on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind.

## Tools & Libraries used
- [Golang 1.16+](https://golang.org/doc/go1.16)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/get-started)
- [Fiber framework](https://github.com/gofiber/fiber)
- [JWT](https://github.com/golang-jwt/jwt)
- [Swagger docs](https://github.com/swaggo/swag)
- [Air](https://github.com/cosmtrek/air)
- [Swag](https://github.com/swaggo/swag)

## ⚡️ Quick start
- Install **`golang >= 1.21`**, **`air`**, & **`swag`**
- Copy `.env.example` to `.env` (optional)
- Run project by this command:
  ```bash
  air
  ```
- Visit **`http://localhost:5000`** or **`http://localhost:5000/swagger/`**

## 🗄 Project structure

### /app

**Folder with business logic only**. This directory doesn't care about what database driver you're using.

- `/app/controller` folder for functional controller (used in routes)
- `/app/schema` folder for object for API clients
- `/app/model` folder for describe business models and methods of your project

### /cmd
**Main applications for this project.**

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/server` `/cmd/cron`).
Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects,
then it should live in the `/pkg` directory.

### /docs

**Folder with API Documentation.**

This directory contains config files for auto-generated API Docs by Swagger, screenshots
and any other documents related to this project.

### /pkg

**Folder with project-specific functionality.** This directory contains all the project-specific code tailored only for your business use case.

- `/pkg/config` folder for configuration functions
- `/pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `/pkg/route` folder for describe routes of your project
- `/pkg/validator` folder with validation functions

### /platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project,
like setting up the database, logger instance.

- `/platform/database` folder with database setup functions (by default, Sqlite)
- `/platform/logger` folder with better logger setup functions (by default, Logrus)

## ⚙️ Configuration

```bash
# .env
# APP settings:
HOST="0.0.0.0"
PORT=5000
DEBUG=true

# JWT settings:
JWT_SECRET="super_secret_here"
JWT_EXPIRE_SECONDS=1209600

# Database settings:
DB_HOST=fibergb-postgres
DB_PORT=5432
DB_USERNAME=dev
DB_PASSWORD=dev
DB_NAME=fiber_go_api
DB_ECHO=true
```

## ⚠️ License

[MIT](https://opensource.org/licenses/MIT)
