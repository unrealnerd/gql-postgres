package repo

import (
	"database/sql"
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

//Find ... runs a query for the given where condition with multiple row output
func (r *InventoryRepo) Find(whereCondition string, args ...interface{}) []*model.Product {

	db, _ := sql.Open("postgres", connStr)
	defer db.Close()

	query := fmt.Sprintf("SELECT product_id, name, description FROM public.product where %s", whereCondition)
	rows, err := db.Query(query, args...)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	var quotes []*model.Product

	for rows.Next() {
		q := &model.Product{}
		err = rows.Scan(&q.ProductID, &q.Name, &q.Description)	

		if err != nil {
			log.Println(err)
			return nil
		}
		quotes = append(quotes, q)
	}
	log.Printf("records: %v", quotes)
	return quotes
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
