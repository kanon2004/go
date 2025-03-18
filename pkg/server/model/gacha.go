package model

import (
	"database/sql"
	"go-college/pkg/db"
)

type Gacha struct {
	Id    string
	Ratio int
}

// gacha_probabilityテーブルデータを取得する
func SelectGachaProbability() (*Gacha, error) {
	row := db.Conn.QueryRow("SELECT collection_item_id,ratio FROM gacha_probability")
	return convertToGacha(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToGacha(row *sql.Row) (*Gacha, error) {
	var gacha Gacha
	if err := row.Scan(&gacha.Id, &gacha.Ratio); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &gacha, nil
}
