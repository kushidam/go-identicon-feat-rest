package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)
	
	// IdentIconを作成するAPIエンドポイント（CreateIdentIconHandler）を追加
	e.POST("/identicon", s.CreateIdentIconHandler)

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
// IdentIconを作成するエンドポイント
// これはユーザが会員登録した際に、ユーザIDに応じて一意の画像を生成する。
// request：ユーザID
// response：成功/失敗ステータス
// 【処理内容】
// 1. ユーザIDを受け取る
// 2. ユーザIDを元にIdentIconを生成する
// 3. 生成したIdentIconをDBに格納する
// 4. 成功/失敗ステータスを返す
// 5. エラーが発生した場合はエラーメッセージを返す
// 6. ユーザIDが重複している場合はエラーメッセージを返す
// 7. ユーザIDが重複していない場合は成功ステータスを返す
func (s *Server) CreateIdentIconHandler(c echo.Context) error {
	// ユーザIDを受け取る
	
	// ユーザIDを元にIdentIconを生成する
	
	
	// 生成したIdentIconをDBに格納する
	
	
	// 成功/失敗ステータスを返す
	return c.JSON(http.StatusOK, "success")
}
