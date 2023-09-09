<<<<<<< HEAD
package bank

type Customer struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      int    `json:"age"`
	Balance  int    `json:"balance"`
	Password string `json:"password"`
}
=======
package bank

type Customer struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Balance  int    `json:"-"`
	Password string `json:"password" binding:"required"`
}
>>>>>>> f48611f (first commit)
