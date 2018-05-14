package sql

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
	cart "github.com/marcusolsson/coverage-cravings"
)

type OrderRepository struct {
	dbname   string
	host     string
	port     string
	user     string
	password string
	sslMode  string

	db *sql.DB
}

func SetDBName(dbname string) func(*OrderRepository) {
	return func(r *OrderRepository) { r.dbname = dbname }
}

func SetHost(host string) func(*OrderRepository) {
	return func(r *OrderRepository) { r.host = host }
}

func SetPort(port string) func(*OrderRepository) {
	return func(r *OrderRepository) { r.port = port }
}

func SetUser(user string) func(*OrderRepository) {
	return func(r *OrderRepository) { r.user = user }
}

func SetPassword(pass string) func(*OrderRepository) {
	return func(r *OrderRepository) { r.password = pass }
}

func SetSSLMode(mode string) func(*OrderRepository) {
	return func(r *OrderRepository) { r.sslMode = mode }
}

func NewOrderRepository(opts ...func(*OrderRepository)) (*OrderRepository, error) {
	repo := new(OrderRepository)

	for _, opt := range opts {
		opt(repo)
	}

	db, err := sql.Open("postgres", repo.connString())
	if err != nil {
		return nil, err
	}
	repo.db = db

	return repo, nil
}

func (r *OrderRepository) connString() string {
	vals := map[string]string{
		"host":     r.host,
		"port":     r.port,
		"dbname":   r.dbname,
		"user":     r.user,
		"password": r.password,
		"sslmode":  r.sslMode,
	}

	var kvs []string
	for k, v := range vals {
		if v != "" {
			kvs = append(kvs, k+"="+v)
		}
	}

	return strings.Join(kvs, " ")
}

func (r *OrderRepository) Save(o *cart.Order) error {
	return nil
}

func (r *OrderRepository) FindByID(id string) (*cart.Order, error) {
	return nil, nil
}
