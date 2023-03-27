package products

import (
	"net/http"
)

type CreateSizeJson struct {
	SizeName string `json:"size_name"`
}

func CreateSize(w http.ResponseWriter, req *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var newSize CreateSizeJson

	// // Reading the request body and UnMarshal the body to the CreateSizeJson struct
	// bs, _ := io.ReadAll(req.Body)
	// if err := json.Unmarshal(bs, &newSize); err != nil {
	// 	utils.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
	// 	log.Println("Internal Server Error in UnMarshal JSON body in CreateSize route:", err)
	// 	return
	// }

	// // CheckUserGroup: IMS User and Operations
	// if !auth_management.RetrieveIssuer(w, req) {
	// 	return
	// }
	// if !utils.CheckUserGroup(w, w.Header().Get("username"), "InvenNexus User", "Operations") {
	// 	return
	// }

	// // Trim White Spaces in size name
	// newSize.SizeName = strings.TrimSpace(newSize.SizeName)

	// // Size Name Form Validation
	// if !SizeNameFormValidation(w, newSize.SizeName) {
	// 	return
	// }

	// // Check User Organisation
	// username := w.Header().Get("username")
	// organisationName, userId, err := database.GetOrganisationNameAndUserIdByUsername(username)
	// if err != nil {
	// 	utils.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
	// 	log.Println("Internal server error in getting company name from database:", err)
	// 	return
	// }

	// // Check size name to see if it already exists in database (cannot have duplicates within the same organisation or username)
	// var count int
	// if organisationName == "InvenNexus" {
	// 	// Check the Size name for duplicates based on username
	// 	count, err = database.GetSizeNameCountByUsername(userId, newSize.SizeName)
	// } else {
	// 	// Check the Size name for duplicates based on organisation name
	// 	count, err = database.GetSizeNameCountByOrganisation(organisationName, newSize.SizeName)
	// }

	// if err != nil {
	// 	utils.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
	// 	log.Println("Internal server error in getting size name count:", err)
	// 	return
	// }
	// if count >= 1 {
	// 	utils.WriteJSON(w, http.StatusBadRequest, "Size Name already exists. Please try again.")
	// 	return
	// }

	// // Insert size name to `organisation_sizes` or `user_sizes` tables
	// if organisationName == "InvenNexus" {
	// 	// insert into `user_sizes` table
	// 	err = database.InsertIntoUserSizes(userId, newSize.SizeName)
	// } else {
	// 	err = database.InsertIntoOrganisationSizes(organisationName, newSize.SizeName)
	// }

	// if err != nil {
	// 	utils.WriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
	// 	log.Println("Internal server error in inserting size name into database:", err)
	// 	return
	// }

	// utils.WriteJSON(w, http.StatusOK, "Successfully created a size!")
}
