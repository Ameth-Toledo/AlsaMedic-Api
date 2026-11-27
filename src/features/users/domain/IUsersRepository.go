package domain

type IUsersRepository interface {
	Save()
	Update()
	Delete()
}
