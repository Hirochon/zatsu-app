package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// テストサーバーのリクエストを作成
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// レスポンスを記録するレコーダーを作成
	rr := httptest.NewRecorder()

	// ハンドラを実行
	handler := http.HandlerFunc(handler)
	handler.ServeHTTP(rr, req)

	// ステータスコードを確認
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// レスポンスボディを確認
	expected := "こんにちは世界!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
} 