package constants

//CATEGORIES
var GetCategoriesListQuery = "SELECT * FROM categories"
var CreateCategoryQuery = "INSERT INTO categories (name) VALUES (?)"
var UpdateCategoryQuery = "UPDATE categories SET name = ? WHERE id = ?"
var DeleteCategoryQuery = "DELETE FROM categories WHERE id = ?"
var GetDetailCategoryQuery = "SELECT * FROM categories WHERE id = ?"

//PRODUCTS
var GetListProductQuery = "SELECT id, name, price, thumbnail, description, created_at, updated_at, category_id FROM products"
var CreateProductQuery = "INSERT INTO products (name, price, thumbnail, description, category_id) VALUES (?, ?, ?, ?, ?)"
var GetDetailProductQuery = "SELECT * FROM products WHERE id = ?"
var GetDetailProductUpdateQuery = "SELECT name, price, thumbnail, description FROM products WHERE id = ?"
var UpdateProductQuery = "UPDATE products SET name = ?, price = ?, thumbnail = ?, description = ? WHERE id = ?"
var GetThumbnailQuery = "SELECT thumbnail FROM products WHERE id = ?"
var DeleteProductQuery = "DELETE FROM products WHERE id = ?"
