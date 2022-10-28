package config

import "github.com/krifik/groupchat-be/entity"

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
