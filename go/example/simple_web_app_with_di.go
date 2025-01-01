// Goの暗黙のインタフェースによる依存性注入を理解するための簡単なウェブアプリケーション
package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 小さなユーティリティ関数として、ログを記録する関数を定義
func LogOutput(message string) {
	fmt.Print(message)
}

// 簡単なデータ保存場所
type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// SimpleDataStore のインスタンスを生成するファクトリ関数
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// --------------------
// ビジネスロジックの作成
// --------------------

// ユーザーを探して「こんにちは」や「さようなら」を言うビジネスロジックの作成
// ビジネスロジックにはデータが必要なためデータ保存ロジックに依存する
// また、ビジネスロジック起動時にログ記録関数に依存する
// ログやデータ保存に LogOutput や SimpleDataStore 以外の仕組みを使いたくなる可能性があるため、これらへの依存を強制しない

// データ保存ロジックが何に依存するかを説明するインタフェース
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

// ログ記録が何に依存するかを説明するインタフェース
type Logger interface {
	Log(message string)
}

// 関数 LogOutput が Logger インタフェースに適合するように、
// Log メソッドを持った関数型を定義
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// 定義された依存性をフィールドとして持つ構造体を定義
// 具象型については何も触れていないため、依存はない
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

// ユーザーを探して「こんにちは」を言うビジネスロジック
func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("SayHello(" + userID + ")\n")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("不明なユーザー")
	}
	return name + "さんこんにちは。\n", nil
}

// ユーザーを探して「さようなら」を言うビジネスロジック
func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("SayGoodbye(" + userID + ")\n")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("不明なユーザー")
	}
	return name + "さんさようなら。\n", nil
}

// SimpleLogic のインスタンスを作成するファクトリ関数
// インタフェースを渡すと構造体を返す
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// --------------------
// API の作成
// --------------------

// コントローラー用のビジネスロジックインタフェースを定義
type Logic interface {
	SayHello(userID string) (string, error)
	SayGoodbye(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayHello内: ")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func (c Controller) SayGoodbye(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayGoodbye内: ")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayGoodbye(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// Controller のインスタンスを作成するファクトリ関数
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

// すべてのコンポーネントを結びつけ、サーバーを起動
// 異なる実装に入れ替えたい場合は、ここだけ書き換えればよい
func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.HandleFunc("/goodbye", c.SayGoodbye)
	http.ListenAndServe(":8080", nil)
}
