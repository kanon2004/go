package handler

import (
	"errors"
	"fmt"
	"go-college/pkg/server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RankGetReseponse struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Rank     int    `json:"rank"`
	Score    int    `json:"score"`
}

// 構造体配列、ＪＳＯＮにする前のデータを格納
type ReseponseRanking struct {
	Ranks []RankGetReseponse `json:"ranks"`
}

func HandleRankingGet() echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("0")
		var nc RankGetReseponse
		var ncr []RankGetReseponse

		fmt.Println("1")

		//関数呼び出し Userテーブルのデータ全部  戻り値スライス構造体
		start, _ := strconv.Atoi(c.QueryParam("start"))
		col, err := model.SelectUser(start)
		if err != nil {

			fmt.Printf("%v", err)
			
			return err
		}
		if col == nil {
			return errors.New("user not found")
		}

		var rank []*model.User
		rank = append(rank, col)

		fmt.Println("3")
		for index, value := range rank {

			//情報詰める
			nc.UserId = value.ID
			nc.UserName = value.Name
			nc.Rank = index
			nc.Score = int(value.HighScore)

			ncr = append(ncr, nc)

		}

		//response
		return c.JSON(http.StatusOK, ReseponseRanking{
			Ranks: ncr,
		})

	}
}
