# My Blog

**My Blog** is a simple Go web application that allows users to create, edit, delete, and view articles. The project uses JWT and cookies for authentication.

## Features

- User registration and login (with JWT and cookies).
- Create new articles with a title, publishing date, and content.
- View all articles on the homepage and in the Dashboard.
- Edit existing articles (only by the author).
- Delete articles (only by the author).
- Articles are sorted by publishing date (newest first).
- Responsive design for mobile and tablet devices.
- Stylish article cards with CSS hover and delete animations.
- Security checks for authorization and user access rights.

## Installation and Running

1. Clone the repository:

```bash
git clone https://github.com/LoLoL200/my_blog.git
cd my_blog
```
2.Install Go dependencies:
```bash
go mod tidy
```
3. Create a .env file in the project root and add your JWT secret:
```env
JWT_SECRET=your_secret_key_here
```
4. Run the server:
```bash
go run main.go
```
5.Open your browser at:
```
http://localhost:8080
```
## Project Structure
```
my_blog/
├─ articles/        # JSON files for articles
├─ handlers/        # HTTP handlers
├─ models/          # Data structures and templates
├─ templates/       # HTML templates
├─ images/          # Images (if any)
├─ main.go          # Entry point
├─ go.mod
└─ .gitignore       # Ignored files (secrets, .env)
```
