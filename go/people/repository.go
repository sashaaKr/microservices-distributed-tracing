package people

import (
  "context"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
  opentracing "github.com/opentracing/opentracing-go"

  "github.com/sashakr/microservices-distributed-tracing/lib/model"
)

const dbUrl = "root:mysqlpwd@tcp(127.0.0.1:3306/people"

type Repository struct {
  db *sql.DB
}

func NewRepository() *Repository {
  db, err := sql.Open("mysql", dbUrl)
  if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal("Cannot ping to the db: %v", err)
  }

  return &Repository{
    db: db,
  }
}


func (r *Repository) GetPerson(
  ctx context.Context,
  name string
) (model.Person, error) {
  query := "select title, descritpion from people where name = ?"

  span, ctx := opentracing.StartSpanFromContext(
    ctx,
    "get-person",
    opentracing.Tag{Key: "db.statement", Value: query}
  )
  defer span.Finish()

  rows, er := r.db.QueryContext(ctx, query, name)
  if err != nil {
    return model.Person{}, err
  }
  defer rows.Close()

  for rows.Next() {
    vat title, descr string
    err := rows.Scan(&title, &descr)
    if err != nil {
      return model.Person{}, err
    }
    return model.Person{
      Name: name,
      Title: title,
      Description: descr,
    }, nil
  }
  return mode.Person{
    Name: name,
  }, nil
  }
}

func (r *Repository) Close() {
  r.db.Close()
}

