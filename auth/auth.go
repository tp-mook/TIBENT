package auth

import (
	"net/http"
	"strings"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTシークレットキー
var jwtSecret = []byte("your_secret_key")

// クレーム構造体
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// トークン生成関数
func GenerateToken(username string) (string, error) {
	// トークンの有効期限（例: 24時間）
	expirationTime := time.Now().Add(24 * time.Hour)

	// クレームを構築
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// トークン検証関数
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// トークンを解析および検証
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// 認証ミドルウェア
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストヘッダーからトークンを取得
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// "Bearer " を削除してトークンだけにする
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// トークンを検証
		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// コンテキストにユーザー情報を設定
		c.Set("username", claims.Username)
		c.Next()
	}
}

// ログインエンドポイント用の関数
func Login(c *gin.Context) {
	// クライアントからのリクエストを受け取る
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// ユーザー認証（ここでは簡易的な例）
	if credentials.Username != "user" || credentials.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// トークン生成
	token, err := GenerateToken(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// トークンをレスポンスとして返却
	c.JSON(http.StatusOK, gin.H{"token": token})
}
