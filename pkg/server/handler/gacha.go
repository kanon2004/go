package handler

import (
	"fmt"
	"go-college/pkg/server/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetGachaReseponse struct {
}
type GachaReseponse struct {
	Results []GetGachaReseponse `json:"results"`
}

func HandleGachaGet() echo.HandlerFunc {
	return func(c echo.Context) error {

		var nc GetGachaReseponse
		var Ncr []GachaReseponse

		// gacha_probabilityテーブルデータを取得する
		if gacha, err := model.SelectGachaProbability(); err != nil {
			fmt.Printf("&v", err)
			return err
		}

		//1 ratioスライス
		gacha.Ratio
		//2 累積比率スライス
		
		//3　id+2のスライス構造体
		gacha.Id
		//4 ランダムな値と3のidを比較 = 引いたカードのidを決定(1 or 10)

		//5 4のidからname,rarityを所得(collection_itemテーブル)

		//6 isNew判定

		// レスポンスに必要な情報を詰めて返す
	return c.JSON(http.StatusOK,GachaReseponse {
			Results: Ncr,
		})
	}
}
