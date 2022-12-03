package entity

type GetRoleResponse struct {
	Id          string
	RoleName    string
	Scopes      string
	Description string
	IsSuperRole bool
}
