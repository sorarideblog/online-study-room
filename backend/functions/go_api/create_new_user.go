package go_api
// todo 必要なければこの関数を消す。デバッグの為だけにgoogleログインサービスと紐付けないアカウントを作成するのは現状手間。
//import (
//	"cloud.google.com/go/firestore"
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//	"os"
//)
//
//type CreateNewUserResponseStruct struct {
//	Result string `json:"result"`
//	Message string `json:"message"`
//}
//
//func createNewUser(userId string, displayName string, client *firestore.Client, ctx context.Context) (*firestore.WriteResult, error) {
//	return client.Collection(USERS).Doc(userId).Set(ctx, map[string]interface{}{
//		"registration-date": firestore.ServerTimestamp,
//		"last-access":       firestore.ServerTimestamp,
//		"online":            false,
//		"status":            "",
//	})
//}
//
//// テストデバッグ用API。俺だけアクセスできるようにpasswordがいる。
//func CreateNewUser(w http.ResponseWriter, r *http.Request)  {
//	ctx, client := InitializeHttpFunc(&w)
//	defer client.Close()
//
//	userId, userName := r.FormValue(user_id), r.FormValue("user_name")
//	password := r.FormValue("password")
//	var apiResp CreateNewUserResponseStruct
//
//	if password == os.Getenv("password") {
//		if userId == "" || userName == "" {
//			apiResp.Result = ERROR
//			apiResp.Message = InvalidParams
//		} else {
//			if IsInUsers(userId, client, ctx) {
//				apiResp.Result = ERROR
//				apiResp.Message = "The user id is already used."
//			} else {
//				_, err := createNewUser(userId, userName, client, ctx)
//				if err != nil {
//					_, _ = fmt.Fprintf(w, "Failed creating the new user.")
//					log.Fatalln(err)
//				} else {
//					_, _ = fmt.Fprintf(w, "Successfully created the new user.")
//				}
//			}
//		}
//	} else {
//		apiResp.Result = ERROR
//		apiResp.Message = "Incorrect password."
//	}
//	bytes, _ := json.Marshal(apiResp)
//	_, _ = w.Write(bytes)
//}
