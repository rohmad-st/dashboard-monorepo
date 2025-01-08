# Monorepo for Frontend & Backend Applications

This repository is structured as a monorepo using **Turborepo** to manage both the frontend (Next.js) and backend (Go) applications. By leveraging shared models, utilities, and configurations, the architecture promotes reusability and consistency across the codebase.

## **Architecture Overview**

The repository contains the following apps and packages:

- **apps/frontend**: A Next.js application for the client-facing frontend, utilizing React and Material-UI.
- **apps/backend**: An Golang server application that provides RESTful APIs and interacts with Firebase for authentication.
- **packages/ui**: A library of reusable UI components built with Material-UI and styled with custom themes, used in the frontend apps.
- **packages/theme**: A library of reusable custom themes for Material UI, used in the frontend apps.
- **packages/models**: Shared TypeScript interfaces and types used across the apps for consistency.
- **packages/eslint-config**: Shared configurations of ESLint settings.
- **packages/typescript-config**: Shared configurations of `tsconfig` settings.

Turborepo handles the build and development processes for the monorepo, ensuring efficient dependency management and build caching.

---

## **Getting Started**

### **Prerequisites**

Ensure you have the following installed on your machine:

- [Node.js](https://nodejs.org/) (v16 or higher)
- [npm](https://www.npmjs.com/) (v7 or higher)
- [Go](https://go.dev/) (v1.23.4 or higher)
- Firebase CLI (`npm install -g firebase-tools`)

### **Installation**

#### 1. Clone the repository:

```bash
git clone https://github.com/rohmad-st/dashboard-monorepo
cd dashboard-monorepo
```

#### 2. Install dependencies using npm:

```bash
npm install
```

#### 3. Running Locally

```bash
# run frontend
npm run dev:frontend
# run backend
npm run dev:backend
# run both frontend & backend
npm run dev
```

_By default._ The frontend will be available at [http://localhost:3001](http://localhost:3001), and the backend will run at [http://localhost:8000](http://localhost:8000).
