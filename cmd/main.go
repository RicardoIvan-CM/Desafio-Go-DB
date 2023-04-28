package main

import (
	"database/sql"
	"encoding/json"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bootcamp-go/desafio-cierre-db.git/cmd/router"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("mysql", "user1:secret_password@/fantasy_products?parseTime=true")
	if err != nil {
		panic(err)
	}

	_, err = db.Query("DELETE FROM sales")
	_, err = db.Query("ALTER TABLE sales AUTO_INCREMENT = 1")
	_, err = db.Query("DELETE FROM products")
	_, err = db.Query("ALTER TABLE products AUTO_INCREMENT = 1")
	_, err = db.Query("DELETE FROM invoices")
	_, err = db.Query("ALTER TABLE invoices AUTO_INCREMENT = 1")
	_, err = db.Query("DELETE FROM customers")
	_, err = db.Query("ALTER TABLE customers AUTO_INCREMENT = 1")

	if err != nil {
		panic(err)
	}

	var customers []domain.Customers
	jsonFile, err := os.ReadFile("datos/customers.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonFile, &customers)
	if err != nil {
		panic(err)
	}
	for _, customer := range customers {
		stmt, err := db.Prepare("INSERT INTO customers (last_name,first_name,`condition`) VALUES (?,?,?)")
		defer stmt.Close()
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(customer.LastName, customer.FirstName, customer.Condition)
		if err != nil {
			panic(err)
		}
	}

	var invoices []domain.Invoices
	jsonFile, err = os.ReadFile("datos/invoices.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonFile, &invoices)
	if err != nil {
		panic(err)
	}
	for _, invoice := range invoices {
		stmt, err := db.Prepare("INSERT INTO invoices (`datetime`,customer_id,total) VALUES (?,?,?)")
		defer stmt.Close()
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(invoice.Datetime, invoice.CustomerId, invoice.Total)
		if err != nil {
			panic(err)
		}
	}

	var products []domain.Product
	jsonFile, err = os.ReadFile("datos/products.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonFile, &products)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		stmt, err := db.Prepare("INSERT INTO products (description,price) VALUES (?,?)")
		defer stmt.Close()
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(product.Description, product.Price)
		if err != nil {
			panic(err)
		}
	}

	var sales []domain.Sales
	jsonFile, err = os.ReadFile("datos/sales.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonFile, &sales)
	if err != nil {
		panic(err)
	}

	for _, sale := range sales {
		stmt, err := db.Prepare("INSERT INTO sales (product_id,invoice_id,quantity) VALUES (?,?,?)")
		defer stmt.Close()
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(sale.ProductId, sale.InvoicesId, sale.Quantity)
		if err != nil {
			panic(err)
		}
	}

	//Actualizar invoices
	rows, err := db.Query("select id from invoices")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		stmt, err := db.Prepare("UPDATE invoices SET total = (SELECT round(sum(quantity*price),2) FROM fantasy_products.sales s inner join fantasy_products.products p on s.product_id = p.id group by s.invoice_id having s.invoice_id = ?) WHERE id = ?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(id, id)
		if err != nil {
			panic(err)
		}
	}

	router.NewRouter(r, db).MapRoutes()

	r.Run()
}
