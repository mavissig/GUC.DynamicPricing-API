package domain

import "github.com/google/uuid"

type Producer interface {
	AddData(client *Data) error
}

type Cache interface {
	GetDataByKey(key string) (*Data, error)
}

func (uc *UseCase) DataAdd(data *Data) (uuid.UUID, error) {
	idOperation := uuid.New()

	data.ID = idOperation
	err := uc.produce.AddData(data)
	if err != nil {
		return uuid.Nil, err
	}

	return idOperation, nil
}

func (uc *UseCase) DataGetByKey(key uuid.UUID) (*Data, error) {
	data, err := uc.cache.GetDataByKey(key.String())
	return data, err
}
