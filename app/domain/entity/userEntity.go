package entity

type InsertUserRequest struct {
	Name     string
	Password string
	Email    string
	Phone    string
	IsActive bool
}

type InsertUserResponse struct {
	Id string
}

type GetUserRequest struct {
	Id    string
	Value string //it can be username or email or phone because there are unique
}

type GetUserResponse struct {
	Id        string
	Name      string
	Password  string
	Email     string
	Phone     string
	IsActive  bool
	CreatedAt string
	UpdateAt  string
	DeletedAt string
}
