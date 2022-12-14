//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=structs_mock.go

package handlers

import "database/sql"

type MyTesting struct {
	db *sql.DB
}
