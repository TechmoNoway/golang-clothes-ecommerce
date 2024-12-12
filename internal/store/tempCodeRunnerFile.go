type Product struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Size        string `json:"size"`
	Color       string `json:"color"`
	CategoryID  int64  `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}