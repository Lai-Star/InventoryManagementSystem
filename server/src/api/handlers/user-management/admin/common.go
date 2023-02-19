package handlers_admin


type AdminDeleteUserMgmt struct {
	Username string `json:"username"`
}

// func CheckUserGroupAdmin(w http.ResponseWriter, req *http.Request) bool {
// 	// CheckUserGroup: IMS User and Operations
// 	if !handlers_user_mgmt.RetrieveIssuer(w, req) {return false}
	
// 	checkUserGroupIMSUser := utils.CheckUserGroup(w, w.Header().Get("username"), "Admin")
// 	if !checkUserGroupIMSUser {
// 		utils.ResponseJson(w, http.StatusForbidden, "Access Denied: You do not have permission to access this resource.")
// 		return false
// 	}
// 	return true
// }

