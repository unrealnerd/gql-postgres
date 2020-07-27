package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" //postgres drivers for initialization
	"github.com/unrealnerd/gql-postgres/graph/model"
)

//InventoryRepo ...
type InventoryRepo struct {
}

var connStr = os.Getenv("inventoryDBConnectionString")
const maxLimit int = 100

//GetSingleProduct ... runs a query with single row output
func (r *InventoryRepo) GetSingleProduct() {

	db, _ := sql.Open("postgres", connStr)
	defer db.Close()

	productID := "30815858"
	row := db.QueryRow("SELECT product_id, name, description FROM public.product where product_id=$1", productID)

	p := model.Product{}
	err := row.Scan(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p)
	}

}

//GetProducts ... pulls &limit no.of products whose productid is greater then &productid 
func (r *InventoryRepo) GetProducts(limit int, productID int) ([]*model.Product, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	if limit > maxLimit { 
		return nil, errors.New("You are exceeding the allowed number of records")
	} else if limit == 0 { //if no limit passed the pull max records
		limit = maxLimit
	}


	query := "SELECT product_id, name, description FROM public.product where product_id >= $1 LIMIT $2"

	rows, err := db.Query(query, productID, limit)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var products []*model.Product

	for rows.Next() {
		q := &model.Product{}
		err = rows.Scan(&q.ProductID, &q.Name, &q.Description)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, q)
	}
	log.Printf("records: %v", products)
	return products, nil
}

//Ping ...  Ping the DB to verify if the able to connect to the db
func Ping() {

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if db.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("successfully able to connect!")
	}

}
