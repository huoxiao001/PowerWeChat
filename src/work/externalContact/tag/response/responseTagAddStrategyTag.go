package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseTagAddStrategyTag struct {
	*response.ResponseWork

	ExternalContactList *power.HashMap `json:"tag_group"`
}