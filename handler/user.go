package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"example.com/myapi/auth"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    // CORSヘッダー（ここにも必要）
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

    // 👇 OPTIONSの場合すぐ返す（ここが重要！）
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // 認証ヘッダー確認
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
        return
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    fmt.Println("🔑 tokenString:", tokenString) // ← ここでログが出るはず！

    token, claims, err := auth.VerifyToken(tokenString)
    if err != nil || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    email, ok := claims["email"].(string)
    if !ok {
        http.Error(w, "Email not found in token claims", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "message": fmt.Sprintf("ようこそ、%s さん！", email),
    })
}
