## About this App

This application was created for my university thesis.
The title of my thesis is The Potential and Reality of Suica as a Means of Electronic Payment in East Africa.
The content of the thesis is to examine the integration of Suica, a widely used application in Japan, into the current state of public transport improvement in East Africa.

## Aims

Creating a contactless IC infrastructure for public transportation in East Africa and releasing it along with native apps would provide a security risk hedge against cash while at the same time capturing train and bus diagrams.

## Target User

The number of people using public transportation in East Africa is very high; for example, millions of people use public transportation every day in Nairobi alone(Friedrich-Ebert-Stiftung. (2014)), the capital of Kenya. Many people use buses and trains to travel within the city or to other parts of the country. In addition, people of all backgrounds use public transportation, including students, commuters, and tourists.
Suica targets these users who are looking for a more efficient and convenient way to pay for public transportation; Suica allows them to pay quickly and smoothly without having to carry cash or purchase tickets in advance. Businesses and service providers can also improve convenience for their customers by accepting Suica to help public transportation users pay for their services.

## Main Functions

- Register
- Login
- Search train route with google map
- Top up
- Add credit card
- Check balance
- History

## Technology stack

### Backend

- Golang(echo)
- Gorm
- Sql-migrate
- JWT
- Uuid
- PostgreSQL

### Frontend

- Next.js(App router)
- Material UI
- TailwindCSS
- Google map API

## How to run this App

### you can init first

```bash
cd backend/
docker compose up -d
touch dbconfig.yml (sample are below)
sql-migrate up
```

```yml
development:
  dialect: postgres
  datasource: postgresql://yourusername:example@localhost:5432/postgres?sslmode=disable
  dir: ./internal/database/migrations/
staging:
  dialect: postgres
  datasource: postgresql://yourusername:example@localhost:5432/postgres?sslmode=disable
  dir: ./internal/database/migrations/
```

### Backend

```bash
cd backend
go run main.go
// or
make start-local
```

### Frontend

```bash
cd frontend/
npm i
npm run dev
```
