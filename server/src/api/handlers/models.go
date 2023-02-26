package handlers

type User struct {
	UserId int
	Username string
	Email string
	IsActive int
	OrganisationName string
	UserGroup []string
	AddedDate string
	UpdatedDate string
}

type Product struct {
	ProductId int
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

