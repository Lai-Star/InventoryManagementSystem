package handlers

type User struct {
	Username string
	Password string
	Email string
	UserGroup string
	OrganisationName string
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
	Sizes []ProductSize
}

type ProductSize struct {
	SizeName string
	SizeQuantity int
}

