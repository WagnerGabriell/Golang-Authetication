package repository

import (
	"GolangAuthetication/internal/entity"
	"database/sql"
)

type ProductRepositoryMysql struct {
	Db *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		Db: db,
	}
}

func (r *ProductRepositoryMysql) GetProduct() ([]*entity.Product, error) {

	var productList []*entity.Product

	row, err := r.Db.Query("Select id,name,price,idUser from Product")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var product entity.Product
		row.Scan(&product.Id, &product.Name, &product.Price, &product.UserId)
		productList = append(productList, &product)
	}
	return productList, nil
}
func (r *ProductRepositoryMysql) CreateProduct(Product *entity.Product) error {
	_, err := r.Db.Exec("Insert into Product (id,name,price, idUser) values (?,?,?,?)", Product.Id, Product.Name, Product.Price, Product.UserId)
	if err != nil {
		return err
	}
	return nil
}
