package domain

// import "github.com/labstack/echo/v4"

type Core struct {
	ID              uint
	Name            string
	Username        string
	Email           string
	Password        string
	Profile_picture string
	HP              string
	Bio             string
}

type Repository interface { // Data /Repository (berhubungan dg DB)
	Insert(newUser Core) (Core, error)
	Get(ID uint) (Core, error)
	// GetAll() ([]Core, error)
	Delete(ID uint) (Core, error)
	Edit(input Core) (Core, error)
	Login(input Core) (Core, error)
}

type Service interface { // Bisnis logic
	AddUser(newUser Core) (Core, error)
	Get(ID uint) (Core, error)
	// ShowAllUser() ([]Core, error)
	DeleteUser(ID uint) (Core, error)
	UpdateUser(input Core) (Core, error)
	Login(input Core) (interface{}, error)
}
