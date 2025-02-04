## Database Migration Commands

To run the database migrations, use the following `goose` commands:

### cd into the sql/schema directory
```bash
cd sql/schema
```

### Migrate Up

```bash
goose postgres "postgres://marcus:@localhost:5432/chirpy" up
```

### Migrate Down

```bash
goose postgres "postgres://marcus:@localhost:5432/chirpy" down
```
