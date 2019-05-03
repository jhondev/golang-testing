package testing101

import (
  "flag"
  "fmt"
  "os"
  "testing"
)

var db struct {
  Url string
}

func TestMain(m *testing.M) {
  // Pretend to open our DB connection
  db.Url = os.Getenv("DATABASE_URL")
  if db.Url == "" {
    db.Url = "localhost:5432"
  }

  flag.Parse()
  exitCode := m.Run()

  // Pretend to close our DB connection
  db.Url = ""

  // Exit
  os.Exit(exitCode)
}

func TestDatabase(t *testing.T) {
  // Pretend to use the db
  fmt.Println(db.Url)
}