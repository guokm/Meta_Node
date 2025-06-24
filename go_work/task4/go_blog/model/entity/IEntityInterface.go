package entity

type Create[T any, R any] interface {
	Create(input T) R
}

type Delete[T any, R any] interface {
	update(id uint) R
}

type Update[T any, R any] interface {
	update(id uint) R
}

type Get[T any, R any] interface {
	update(id uint) R
}
