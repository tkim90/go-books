package model

type Book struct {
  isbn   string `db:isbn`
  title  string `db:title`
  author string `db:author`
  price  float32 `db:price`
}

type User struct {
	ID             int    `db:"id"`
	Username       string `db:"username"`
	HashedPassword string `db:"password"`
	UUID           string `db:"uuid"`
}

// Identity holds info for a GetUser request
type Identity struct {
	Username   string `db:"username"`
	UUID       string `db:"uuid"`
	IsInternal bool   `db:"is_internal"`
}

// IdentityTable is the table name for a user identity
const IdentityTable = "identity"
