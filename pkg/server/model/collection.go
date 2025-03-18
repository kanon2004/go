package model

import (
	"database/sql"
	"go-college/pkg/db"
)

// user_collection_item　テーブルテーブル
type CollectionItem struct {
	Id     string
	Name   string
	Rarity int
}

// collection_item テーブルから所得
func SelectCollectionItem() (*[]CollectionItem, error) {

	rows, err := db.Conn.Query("SELECT id,name,rarity FROM collection_item")
	if err != nil {
		//
	}

	//このスライスにとってきたデータ入れる
	var items []CollectionItem

	//データがなくなるまで繰り返す
	for rows.Next() {
		p, err := convertToCollectionItem(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *p)
	}
	return &items, nil
}

func convertToCollectionItem(rows *sql.Rows) (*CollectionItem, error) {

	var item CollectionItem

	err := rows.Scan(&item.Id, &item.Name, &item.Rarity)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
