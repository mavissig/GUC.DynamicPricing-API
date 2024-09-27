package domain

type UseCase struct {
	produce Producer
	cache   Cache
}

func New(
	produce Producer,
	cache Cache,
) *UseCase {
	return &UseCase{
		produce: produce,
		cache:   cache,
	}
}
