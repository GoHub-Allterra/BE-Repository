package domain

// import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Name     string
	Username string
	Email    string
	Password string
	Images   string
	HP       string
	Bio      string
	Token    string
}

type PostCore struct {
	ID      uint
	UserID  uint
	Images  string
	Content string
}

type Repository interface { // Data /Repository (berhubungan dg DB)
	Insert(newUser Core) (Core, error)
	Get(ID uint) (Core, error)
	Delete(ID uint) (Core, error)
	Edit(input Core) (Core, error)
	Login(input Core) (Core, error)
	GetByUsername(name Core) (Core, int)
	
}

type Service interface { // Bisnis logic
	AddUser(newUser Core) (Core, error)
	Get(ID uint) (Core, error)
	DeleteUser(ID uint) (Core, error)
	UpdateUser(input Core) (Core, error)
	Login(input Core) (Core, string, error)
}
