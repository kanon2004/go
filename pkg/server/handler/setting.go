package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-college/pkg/constant"
)

// HandleSettingGet ゲーム設定情報取得処理
func HandleSettingGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, &settingGetResponse{
			GachaCoinConsumption: constant.GachaCoinConsumption,
		})
	}
}

type settingGetResponse struct {
	GachaCoinConsumption int32 `json:"gachaCoinConsumption"`
}
