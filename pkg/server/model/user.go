package model

import (
	"database/sql"
	"fmt"

	"go-college/pkg/db"
)

// User userテーブルデータ
type User struct {
	ID        string
	AuthToken string
	Name      string
	HighScore int32
	Coin      int32
}

// InsertUser データベースをレコードを登録する
func InsertUser(record *User) error {
	if _, err := db.Conn.Exec(
		"INSERT INTO user (id, auth_token, name, high_score, coin) values (?, ?, ?, ?, ?)",
		record.ID,
		record.AuthToken,
		record.Name,
		record.HighScore,
		record.Coin,
	); err != nil {
		return err
	}
	return nil
}

// ここだけ自分で追加
func SelectUser(start int) (*User, error) {
	fmt.Println("2")
	query := fmt.Sprintf("SELECT id,high_score FROM user ORDER BY high_score,id LIMIT 10 OFFSET %d", start)
	row := db.Conn.QueryRow(query)
	return convertToUsers(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToUsers(row *sql.Row) (*User, error) {
	var user User
	if err := row.Scan(&user.ID, &user.HighScore); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// SelectUserByAuthToken auth_tokenを条件にレコードを取得する
func SelectUserByAuthToken(authToken string) (*User, error) {
	row := db.Conn.QueryRow("SELECT * FROM user WHERE auth_token=?", authToken)
	return convertToUser(row)
}

// SelectUserByPrimaryKey 主キーを条件にレコードを取得する
func SelectUserByPrimaryKey(userID string) (*User, error) {
	row := db.Conn.QueryRow("SELECT * FROM user WHERE id=?", userID)
	return convertToUser(row)
}

// UpdateUserByPrimaryKey 主キーを条件にレコードを更新する
func UpdateUserByPrimaryKey(record *User) error {
	if _, err := db.Conn.Exec(
		"UPDATE user SET name=? WHERE id=?",
		record.Name,
		record.ID,
	); err != nil {
		return err
	}
	return nil
}

// 自作 IDを条件にレコードを取得する
func SelectScoreScore(userID string) (*User, error) {
	row := db.Conn.QueryRow("SELECT id,high_score,coin FROM user WHERE id=?", userID)
	return convertToScore(row)
}
// convertToUser rowデータをUserデータへ変換する
func convertToScore(row *sql.Row) (*User, error) {
	var user User
	if err := row.Scan(&user.ID,&user.Coin, &user.HighScore); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// 自作　high_score更新
func UpdateScore(record *User) error {
	if _, err := db.Conn.Exec(
		"UPDATE user SET high_score=? WHERE id=?",
		record.HighScore,
		record.ID,
	); err != nil {
		return err
	}
	return nil
}


// 自作 IDを条件にcoinを更新する
func UpdateCoin(record *User) error {
	if _, err := db.Conn.Exec(
		"UPDATE user SET coin=? WHERE id=?",
		record.Coin,
		record.ID,
	); err != nil {
		return err
	}
	return nil
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*User, error) {
	var user User
	if err := row.Scan(&user.ID, &user.AuthToken, &user.Name, &user.HighScore, &user.Coin); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
