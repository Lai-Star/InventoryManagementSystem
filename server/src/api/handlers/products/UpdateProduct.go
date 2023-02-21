package handlers_products

// func UpdateProduct(w http.ResponseWriter, req *http.Request) {
// 	// Set Headers
// 	w.Header().Set("Content-Type", "application/json")
// 	var updateProduct ProductJson

// 	// Reading the request body and Unmarshal the body to the ProductJson struct
// 	bs, _ := io.ReadAll(req.Body)
// 	if err := json.Unmarshal(bs, &updateProduct); err != nil {
// 		utils.InternalServerError(w, "Internal Server Error in UnMarshal JSON body in CreateProduct route: ", err)
// 		return;
// 	}

// 	// Check user group for products
// 	if !CheckProductsUserGroup(w, req) {return}

// 	// Get productid from url params
// 	productIdStr := chi.URLParam(req, "product_id")
// 	updateProduct.ProductId, _ = strconv.Atoi(productIdStr)

// 	// Check if product exists in database
// 	if !database.ProductIdExists(updateProduct.ProductId) {
// 		utils.ResponseJson(w, http.StatusNotFound, "Product does not exist in database. Please try again.")
// 		return
// 	}

// 	// Update product with current product data (if none provided)
// 	updateProduct, result := UpdateCurrentData(w, updateProduct)
// 	if !result {return}

// 	// Product Form Validation
// 	if !ProductFormValidation(w, updateProduct) {return}

// 	// Update products table
// 	err := database.UpdateProduct(updateProduct.ProductId, updateProduct.ProductName, updateProduct.ProductDescription, updateProduct.ProductSku, updateProduct.ProductColour, updateProduct.ProductCategory, updateProduct.ProductBrand, updateProduct.ProductCost)
// 	if err != nil {
// 		utils.InternalServerError(w, "Internal Server Error in UpdateProduct: ", err)
// 		return
// 	}

// 	utils.ResponseJson(w, http.StatusOK, "Successfully updated product!")

// }

// func UpdateCurrentData(w http.ResponseWriter, updateProduct ProductJson) (ProductJson, bool) {
// 	currentProductData, err := GetCurrentProductData(w, updateProduct.ProductId)
// 	if err != nil {
// 		utils.InternalServerError(w, "Internal Server Error when getting current product data: ", err)
// 		return ProductJson{}, false
// 	}

// 	// Fill empty product name
// 	if updateProduct.ProductName == "" {
// 		updateProduct.ProductName = currentProductData.ProductName
// 	}

// 	// Fill empty product description
// 	if updateProduct.ProductDescription == "" {
// 		updateProduct.ProductDescription = currentProductData.ProductDescription
// 	}

// 	// Fill empty product sku
// 	if updateProduct.ProductSku == "" {
// 		updateProduct.ProductSku = currentProductData.ProductSku
// 	}

// 	// Fill empty product colour
// 	if updateProduct.ProductColour == "" {
// 		updateProduct.ProductColour = currentProductData.ProductColour
// 	}

// 	// Fill empty product category
// 	if updateProduct.ProductCategory == "" {
// 		updateProduct.ProductCategory = currentProductData.ProductCategory
// 	}

// 	// Fill empty product brand
// 	if updateProduct.ProductBrand == "" {
// 		updateProduct.ProductBrand = currentProductData.ProductBrand
// 	}

// 	// Fill empty product cost
// 	if updateProduct.ProductCost == 0.0 {
// 		updateProduct.ProductCost = currentProductData.ProductCost
// 	}

// 	return updateProduct, true
// }

// func GetCurrentProductData(w http.ResponseWriter, product_id int) (handlers.Product, error) {
// 	var productName, productDescription, productSku, productColour, productCategory, productBrand sql.NullString
// 	var productCost sql.NullFloat64
// 	result := database.GetProductByProductId(product_id)

// 	var currentProductData handlers.Product

// 	err := result.Scan(&productName, &productDescription, &productSku, &productColour, &productCategory, &productBrand, &productCost)
// 	if err != sql.ErrNoRows {
// 		currentProductData.ProductName = productName.String
// 		currentProductData.ProductDescription = productDescription.String
// 		currentProductData.ProductSku = productSku.String
// 		currentProductData.ProductColour = productColour.String
// 		currentProductData.ProductCategory = productCategory.String
// 		currentProductData.ProductBrand = productBrand.String
// 		currentProductData.ProductCost = float32(productCost.Float64)
// 	} else {
// 		utils.InternalServerError(w, "Internal Server Error in GetCurrentProductData: ", err)
// 		return handlers.Product{}, err
// 	}

// 	return currentProductData, nil
// }