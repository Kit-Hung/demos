/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/23 17:17
 * @Description： 认证 webhook 示例
 */
package main

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	authentication "k8s.io/api/authentication/v1beta1"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("authenticate", Authenticate)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tr authentication.TokenReview
	resp := map[string]interface{}{
		"apiVersion": "authentication.k8s.io/v1beta1",
		"kind":       "TokenReview",
	}

	err := decoder.Decode(&tr)
	if err != nil {
		handlerError(&w, err, http.StatusBadRequest, resp)
		return
	}

	log.Println("receiving request...")
	sts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tr.Spec.Token},
	)
	tc := oauth2.NewClient(context.Background(), sts)
	client := github.NewClient(tc)
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		handlerError(&w, err, http.StatusUnauthorized, resp)
		return
	}

	log.Println("[Success] login as: ", *user.Login)
	w.WriteHeader(http.StatusOK)
	resp["status"] = authentication.TokenReviewStatus{
		Authenticated: true,
		User: authentication.UserInfo{
			Username: *user.Login,
			UID:      *user.Login,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("[Error] encode response error: ", err)
		return
	}

	log.Println(http.ListenAndServe(":3000", nil))
}

func handlerError(w *http.ResponseWriter, err error, code int, resp map[string]interface{}) {
	log.Println("[Error] ", err)
	(*w).WriteHeader(code)
	resp["status"] = authentication.TokenReviewStatus{
		Authenticated: false,
	}
	err = json.NewEncoder(*w).Encode(resp)
	if err != nil {
		log.Println("[Error] encode response error: ", err)
	}
}
