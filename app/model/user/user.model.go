package usermodel

type User struct {
	ID       int64  `json:"id"        db:"id"`
	Name     string `json:"name"      db:"name"`
	Email    string `json:"email"     db:"email"`
	Password string `json:"password"  db:"password"`
	Role     string `json:"role"      db:"role"`
}
