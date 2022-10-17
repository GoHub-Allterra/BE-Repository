package domain

// import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Nama     string
	HP       string
	Password string
}

type Repository interface { // Data /Repository (berhubungan dg DB)
	Insert(newUser Core) (Core, error)
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
	Delete(ID uint) (Core, error)
	Edit(ID uint, input Core) (Core, error)
}

type Service interface { // Bisnis logic
	AddUser(newUser Core) (Core, error)
	Get(ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
	DeleteUser(ID uint) (Core, error)
	UpdateUser(ID uint, input Core) (Core, error)
}

// type Handler interface {
// 	AddUser() echo.HandlerFunc
// 	ShowAllUser() echo.HandlerFunc
// 	GetUser() echo.HandlerFunc
// 	DeleteUser() echo.HandlerFunc
// }
