package repository

import (
	"context"
	"grpc/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgsqlCon struct {
	con *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *PgsqlCon {
	return &PgsqlCon{con: db}
}

func (db *PgsqlCon) GetData() ([]domain.Customer, error) {
	rows, err := db.con.Query(context.Background(), "SELECT id,name,address,email,phone FROM m_customer")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := make([]domain.Customer, 0)
	for rows.Next() {
		t := new(domain.Customer)
		err = rows.Scan(&t.Id,
			&t.Name,
			&t.Address,
			&t.Email,
			&t.Phone,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, *t)
	}
	return result, nil

}

func (db *PgsqlCon) Insert(customer domain.Customer) (int64, error) {
	var id int64
	err := db.con.QueryRow(context.Background(), "INSERT INTO m_customer (name,address,email,phone) VALUES($1,$2,$3,$4) RETURNING id", customer.Name, customer.Address, customer.Email, customer.Phone).Scan(&id)
	return id, err
}
