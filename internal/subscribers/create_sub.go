package subscribers

import (
	"database/sql"
)

const (
	addQuery = `insert into subscribers (email) values ($1)`
)

type Subscriber struct {
	Id    int64  `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}

type StorageSub struct {
	db *sql.DB
}

func (s *StorageSub) CreateSub(email string) (*Subscriber, error) {
	var sub Subscriber
	rows, err := s.db.Query(addQuery, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&sub.Id, &sub.Email); err != nil {
			return nil, err
		}
	}
	return &Subscriber{Id: sub.Id, Email: sub.Email}, nil
}
