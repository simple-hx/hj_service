package datasource

import "hj_service/models"

var user = map[int]models.User{
	1: {
		ID:       1,
		Name:     "张三",
		Password: "123",
		Phone:    "12313131313",
	},
	2: {
		ID:       2,
		Name:     "李四",
		Password: "123",
		Phone:    "12313131313",
	},
}