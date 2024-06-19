# Recything - Backend

Recything is a mobile platform designed to facilitate recycling education, littering reporting, and gamified achievements for users. The app allows end users to find recycling information, report improper trash disposal, and earn rewards for completing tasks, while a web dashboard enables super admins to manage accounts and access all data.

![dashboard admin](docs/Dashboard%20-%20Website.png)
[( Recything - Web )](https://recything.netlify.app/)

## Features

### User

- Register / Login account
- Edit user detail
- Homepage mobile
- Reporting Littering / Rubbish
- Customer service with an AI
- Article content for an education
- Video content for an education
- Doing task challenge to earn rewards
- About our team
- Achievements detail (leaderboard)

### Superadmin / Admin

- Login Superadmin / Admin
- Dashboard admin
- Manage Admins data (only superadmin)
- Manage Users data
- Manage Reports (approve/reject report from user)
- Manage Articles (add/update/delete)
- Manage Videos (add/update/delete)
- Manage Achievement (update target point for an each badge)
- Manage Custom Data for dataset AI
- Manage Tasks (approving/rejecting task user)

## TechStacks

- [Echo](https://github.com/labstack/echo) (Web Framework Go)
- [Cloudinary](https://github.com/cloudinary/cloudinary-go/) (Cloud storage free)
- [Viper](https://github.com/spf13/viper) (Configuration)
- [Validator](https://github.com/go-playground/validator) (Type validation)
- [JWT](https://github.com/golang-jwt/jwt) (Middleware)
- [OpenAI](https://github.com/sashabaranov/go-openai) (Chat Bot)
- MySQL (SQL)
- [GORM](https://gorm.io/docs/) (ORM)
- AWS EC2 (Deployment)

## API Documentation

[( Swagger API )](https://recything.site/)
![Swagger](docs/Swagger.png)

[( Postman API )](https://www.postman.com/sawalrever23/workspace/capstone/collection/34865902-43aa5087-a7e3-4c4b-89b6-749fafe0a359?action=share&creator=34865902)
![Postman](docs/Postman.png)

## ERD

[( ERD - draw.io )](https://drive.google.com/file/d/1fbE-hpS4z3lMEEUfUAiL2XWvwETfgipp/view)
![ERD](docs/ERD%20Recything.drawio.png)
