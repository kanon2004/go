package handler

import (
	"errors"
	"fmt"

	"go-college/pkg/context/auth"
	"go-college/pkg/server/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CollectionGetReseponse struct {
	CollectionID string `json:"collectionID"`
	Name         string `json:"name"`
	Rarity       int    `json:"rarity"`
	HasItem      bool   `json:"hasItem"`
}

// 構造体配列、ＪＳＯＮにする前のデータを格納
type Reseponse struct {
	Collections []CollectionGetReseponse `json:"collections"`
}

// HandleCollectionGet ユーザ情報取得処理
func HandleCollectionGet() echo.HandlerFunc {
	return func(c echo.Context) error {

		// Contextから認証済みのユーザIDを取得
		ctx := c.Request().Context()
		userID := auth.GetUserIDFromContext(ctx)
		if userID == "" {
			return errors.New("userID is empty")
		}

		//関数呼び出し、上で認証した人が持ってるcollectionitemを所得(スライス)
		user, err := model.SelectUserCollectionItem(userID)
		if err != nil {
			fmt.Printf("%v", err)
			return err
		}
		if user == nil {
			return errors.New("collectionitem not found")
		}

		// //関数呼び出し collectionテーブルのデータ全部  戻り値スライス構造体
		col, err := model.SelectCollectionItem()
		if err != nil {
			return err
		}
		if col == nil {
			return errors.New("collection not found")
		}

		//sort

		var nc CollectionGetReseponse
		var Ncr []CollectionGetReseponse

		for _, value := range *col {
			var hantei bool
			for _, users := range *user {
				if value.Id == users.CollectionItemId {
					hantei = true
					goto jump
				} else {
					hantei = false
				}
			}
			jump:
				nc.CollectionID = value.Id
				nc.Name = value.Name
				nc.Rarity = value.Rarity
				nc.HasItem = hantei
				Ncr = append(Ncr, nc)
			
		}
		fmt.Printf("%v", Ncr)
		// レスポンスに必要な情報を詰めて返すす
		return c.JSON(http.StatusOK, Reseponse{
			Collections: Ncr,
		})
	}
}
