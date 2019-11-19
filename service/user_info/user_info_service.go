package user_info

import (
	"github.com/shampoo6/beemongo/models/domains"
	"github.com/shampoo6/beemongo/models/dto"
)

func Insert(dto *dto.UserInfoDto) {
	domains.Insert(dto)
}
