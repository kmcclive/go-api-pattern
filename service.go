package goapipattern

type Service[M any] interface {
	Create(model M) error
	Delete(id uint) error
	FetchByID(id uint) (*M, error)
	List() (*[]M, error)
}
