package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
	//binding required를 통해 값이 안들어오면 오류처리
}

func (c *CreateRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type CreateUserResponse struct {
	*ApiResponse
	*User
}

type UpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	UpdatedAge int64  `json:"updated_age" binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
	*User
}

type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *DeleteRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}

type DeleteUserResponse struct {
	*ApiResponse
	*User
}
