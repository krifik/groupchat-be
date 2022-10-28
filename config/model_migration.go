package config

import "mangojek-backend/entity"

type Entity struct {
	Entity interface{}
}

func RegisterEntities() []Entity {
	return []Entity{
		{Entity: entity.User{}},
		{Entity: entity.Product{}},
		{Entity: entity.Message{}},
	}
}
