package dto

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignUpReq struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CompanyId    uint   `json:"companyId"`
	DepartmentId uint   `json:"departmentId"`
}

type UserSearchReq struct {
	PageParam
	Status       *int8  `json:"status"`
	Admin        *int8  `json:"admin"`
	Username     string `json:"username"`
	CompanyId    *uint  `json:"companyId"`
	DepartmentId *uint  `json:"departmentId"`
}
type UserStatusReq struct {
	Id     uint `json:"id"`
	Status int8 `json:"status"`
}

type UserAdminReq struct {
	Id    uint `json:"id"`
	Admin int8 `json:"admin"`
}

type UserResetReq struct {
	Id       uint   `json:"id"`
	Password string `json:"password"`
}
