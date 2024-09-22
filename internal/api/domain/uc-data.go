package domain

type ClientRepo interface {
	//ClientAdd(client *Client) error
	//ClientDeleteById(id uint64) error
	//ClientGet(name, surname string, page, pageSize int) ([]*Client, error)
	//ClientChangeAddress(id uint64, address *Address) error
}

func (uc *UseCase) DataAdd(client *Data) error {
	//err := uc.clientRepo.ClientAdd(client)
	return nil
}

//
//func (uc *UseCase) ClientGet(name, surname string, page, pageSize int) ([]*Client, error) {
//	return uc.clientRepo.ClientGet(name, surname, page, pageSize)
//}
//
//func (uc *UseCase) ClientDelete(id uint64) error {
//	err := uc.clientRepo.ClientDeleteById(id)
//	return err
//}
//
//func (uc *UseCase) ClientChangeAddress(id uint64, address *Address) error {
//	return uc.clientRepo.ClientChangeAddress(id, address)
//}
