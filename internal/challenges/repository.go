package challenges

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	GetTotalsByCustomerCondition() (domain.TotalsByCustomerCondition, error)
	GetTopSoldProducts() ([]domain.TopSoldProduct, error)
	GetTopActiveCustomersSpent() ([]domain.TopActiveCustomerSpent, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetTotalsByCustomerCondition() (domain.TotalsByCustomerCondition, error) {
	query := `select c.condition, round(sum(i.total),2) from customers c inner join invoices i on c.id = i.customer_id group by c.condition;`
	rows, err := r.db.Query(query)
	if err != nil {
		return domain.TotalsByCustomerCondition{}, err
	}
	defer rows.Close()
	var tbcc domain.TotalsByCustomerCondition
	var cond int
	if rows.Next() {
		rows.Scan(&cond, &tbcc.Inactive)
		if err != nil {
			return domain.TotalsByCustomerCondition{}, err
		}
	}
	if rows.Next() {
		rows.Scan(&cond, &tbcc.Active)
		if err != nil {
			return domain.TotalsByCustomerCondition{}, err
		}
	}
	return tbcc, nil
}

func (r *repository) GetTopSoldProducts() ([]domain.TopSoldProduct, error) {
	query := `select description,sum(quantity) as total from sales s inner join products p on s.product_id = p.id group by p.id order by total desc limit 5;`
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.TopSoldProduct{}, err
	}
	defer rows.Close()
	topSoldProducts := make([]domain.TopSoldProduct, 0)
	for rows.Next() {
		var tsp domain.TopSoldProduct
		err := rows.Scan(&tsp.Description, &tsp.Total)
		if err != nil {
			return []domain.TopSoldProduct{}, err
		}
		topSoldProducts = append(topSoldProducts, tsp)
	}
	return topSoldProducts, nil
}

func (r *repository) GetTopActiveCustomersSpent() ([]domain.TopActiveCustomerSpent, error) {
	query := `select last_name,first_name,round(sum(i.total),2) as amount from fantasy_products.customers c inner join fantasy_products.invoices i on c.id = i.customer_id where c.condition=1 group by c.id order by amount desc limit 4;`
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.TopActiveCustomerSpent{}, err
	}
	defer rows.Close()
	topActiveCustomersSpent := make([]domain.TopActiveCustomerSpent, 0)
	for rows.Next() {
		var tacs domain.TopActiveCustomerSpent
		err := rows.Scan(&tacs.LastName, &tacs.FirstName, &tacs.Amount)
		if err != nil {
			return []domain.TopActiveCustomerSpent{}, err
		}
		topActiveCustomersSpent = append(topActiveCustomersSpent, tacs)
	}
	return topActiveCustomersSpent, nil
}
