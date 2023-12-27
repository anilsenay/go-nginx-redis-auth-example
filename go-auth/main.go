package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []UserInfo{{"anil", "123"}}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		body := UserInfo{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		if !slices.Contains(users, body) {
			w.WriteHeader(401)
			return
		}

		token := tokenGenerator()
		val, _ := json.Marshal(body)

		_, err = rdb.Set(context.Background(), token, val, time.Hour*24).Result()
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Header().Add("Authorization", "Bearer "+token)
		w.WriteHeader(201)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(401)
			return
		}
		parts := strings.Split(token, " ")
		if parts[0] != "Bearer" {
			w.WriteHeader(401)
			return
		}

		auth_token := parts[1]
		_, err := rdb.Get(context.Background(), auth_token).Result()
		if err != nil {
			w.WriteHeader(401)
			return
		}

		w.WriteHeader(302)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen error: ", err)
	}
}

func tokenGenerator() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
