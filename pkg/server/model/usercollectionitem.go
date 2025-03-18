package model

import (
	"database/sql"
	"fmt"
	"go-college/pkg/db"
)

// user_collection_item　テーブルテーブル
type UserCollectionItem struct {
	CollectionItemId string
}

// user_collection_item テーブルから所得//未完成、後で引数書いてな
func SelectUserCollectionItem(userId string) (*[]UserCollectionItem, error) {
	fmt.Println("！！！")
	rows, err := db.Conn.Query("SELECT collection_item_id FROM user_collection_item WHERE user_id=?", userId)
	fmt.Println("???")
	 if err != nil {
	 	fmt.Printf("%v", err)
		return nil, err
	 }
	fmt.Printf("%v",err)

	//スライスを初期化
	var useritems []UserCollectionItem

	//pにconvertToUserCollectionItem(rows)の実行結果を入れてスライスに格納
	for rows.Next() {
		p, err := convertToUserCollectionItem(rows)
		if err != nil {
			return nil, err
		}
		useritems = append(useritems, *p)
	}
	return &useritems, nil
}

func convertToUserCollectionItem(rows *sql.Rows) (*UserCollectionItem, error) {

	var useritem UserCollectionItem

	err := rows.Scan(&useritem.CollectionItemId)
	if err != nil {
		return nil, err
	}
	return &useritem, nil
}
