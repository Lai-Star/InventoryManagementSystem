package handlers

type User struct {
	Username string
	Password string
	Email string
	UserGroup string
	CompanyName string
	IsActive int
	AddedDate string
	UpdatedDate string
}

type Product struct {
	ProductName string
	ProductDescription string
	ProductSku string
	ProductColour string
	ProductCategory string
	ProductBrand string
	ProductCost float32
}
