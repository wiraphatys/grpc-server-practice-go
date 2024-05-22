package repositories

type AuthRepository interface {
	FindUserByEmailGRPC(email string)
}
