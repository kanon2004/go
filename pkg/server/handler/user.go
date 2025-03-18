package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"go-college/pkg/context/auth"
	"go-college/pkg/server/model"
)

// HandleUserCreate ユーザ情報作成処理
func HandleUserCreate() echo.HandlerFunc {
	return func(c echo.Context) error {

		req := &userCreateRequest{}
		if err := c.Bind(req); err != nil {
			fmt.Printf("%v", err)
			return err
		}

		// UUIDでユーザIDを生成する
		userID, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		// UUIDで認証トークンを生成する
		authToken, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		// データベースにユーザデータを登録する
		if err := model.InsertUser(&model.User{
			ID:        userID.String(),
			AuthToken: authToken.String(),
			Name:      req.Name,
			HighScore: 0,
			Coin:      0,
		}); err != nil {
			return err
		}

		// 生成した認証トークンを返却
		return c.JSON(http.StatusOK, &userCreateResponse{Token: authToken.String()})
	}
}

type userCreateRequest struct {
	Name string `json:"name"`
}

type userCreateResponse struct {
	Token string `json:"token"`
}

// HandleUserGet ユーザ情報取得処理
func HandleUserGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Contextから認証済みのユーザIDを取得
		ctx := c.Request().Context()
		userID := auth.GetUserIDFromContext(ctx)
		if userID == "" {
			return errors.New("userID is empty")
		}

		// ユーザデータの取得処理を実装
		user, err := model.SelectUserByPrimaryKey(userID)
		if err != nil {
			return err
		}
		if user == nil {
			return errors.New("user not found")
		}

		// レスポンスに必要な情報を詰めて返却
		return c.JSON(http.StatusOK, &userGetResponse{
			ID:        user.ID,
			Name:      user.Name,
			HighScore: user.HighScore,
			Coin:      user.Coin,
		})
	}
}

type userGetResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	HighScore int32  `json:"highScore"`
	Coin      int32  `json:"coin"`
}

// HandleUserUpdate ユーザ情報更新処理
func HandleUserUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストBodyから更新後情報を取得
		req := &userUpdateRequest{}
		if err := c.Bind(req); err != nil {
			return err
		}

		// Contextから認証済みのユーザIDを取得
		ctx := c.Request().Context()
		userID := auth.GetUserIDFromContext(ctx)
		if userID == "" {
			return errors.New("userID is empty")
		}

		// TODO: ユーザデータの取得処理と存在チェックを実装 (ヒント: model.SelectUserByPrimaryKeyを使用する)

		usr, err := model.SelectUserByPrimaryKey(userID)
		if err != nil {
			return err
		}
		// TODO: userテーブルの更新処理を実装 (ヒント: model.UpdateUserByPrimaryKeyを使用する)
		usr.Name = req.Name
		model.UpdateUserByPrimaryKey(usr)

		return c.NoContent(http.StatusOK)
	}
}

type userUpdateRequest struct {
	Name string `json:"name"`
}
