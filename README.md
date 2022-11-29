# Kanban Board

Repository for Hacktiv8 Final Project 3. Already deployed on heroku `https://hacktiv8-fp3.herokuapp.com/api/v1/ping`
## Admin credentials
- Email: devopunda@kb.id
- Password: adminkanbanboard

## Developer's Manual

### Migrations

- First you need to install the golang-migrate to do database migrations.

MacOS

```bash
brew install golang-migrate
```

Windows (use scoop)

```bash
scoop install migrate
```

To run a migrations

```bash
migrate -source file://./db/migrations -database "mysql://root:@tcp(localhost:3306)/kanban_board" up
```

To rollback a migrations

```bash
migrate -source file://./db/migrations -database "mysql://root:@tcp(localhost:3306)/kanban_board" down
```

### How To Run

```bash
go run main.go
```

## Our Team

### Group 3

- Pande Putu Devo Punda Maheswara
- Hanif Fadillah Amrynudin
- I Putu Agus Arya Wiguna
