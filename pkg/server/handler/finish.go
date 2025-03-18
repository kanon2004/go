package handler

import (
	"errors"
	"fmt"
	"go-college/pkg/context/auth"
	"go-college/pkg/server/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FinishRequest struct {
	Score int `json:"score"`
}

// HandleUserGet ユーザ情報取得処理
func HandleScoreGet() echo.HandlerFunc {
	return func(c echo.Context) error {

		//reqで今回のスコアを所得
		req := &FinishRequest{}
		if err := c.Bind(req); err != nil {
			fmt.Printf("%v", err)
			return err
		}

		//user_idで認証してるIDを所得
		ctx := c.Request().Context()
		userID := auth.GetUserIDFromContext(ctx)
		if userID == "" {
			return errors.New("userID is empty")
		}

		//finishでAさんのいままでのhigh_score,coinを所得
		finish, err := model.SelectScoreScore(userID)
		if err != nil {
			fmt.Printf("%v", err)
			return err
		}
		if finish == nil {
			return errors.New("score,coin not found")
		}

		// high_score更新
		if req.Score > int(finish.HighScore) {
			finish.HighScore = int32(req.Score)
			model.UpdateScore(finish)
		}

		// coin更新
		coin := req.Score*10 + 100

		finish.Coin += int32(coin)
		model.UpdateCoin(finish)
		
		// レスポンスに必要な情報を詰めて返却
		return c.JSON(http.StatusOK, &CoinResponse{
			Coin: coin,
		})
	}
}

type CoinResponse struct {
	Coin int `json:"coin"`
}
