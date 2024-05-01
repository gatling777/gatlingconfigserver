package modHttp

import (
	"configServer/modDatabase"
	"configServer/modUtility"
	"fmt"
	"io"

	//"io"
	"net/http"
	//"github.com/go-chi/chi/v5"
)

func handlerHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))

}

/*
	func routerInfo(r chi.Router) {
		r.Use(handlerInfo)
		r.Put("/")
	}

	func handlerInfo(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			appid := chi.URLParam(r, "appid")

			ctx := context.WithValue(r.Context(), "appid", appid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	func handlerInfoPut(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		appid, ok := ctx.Value("appid").(string)
		if !ok {
			http.Error(w, http.StatusText(422), 422)
			return
		}

}
*/
func checkXApiKey(w http.ResponseWriter, r *http.Request) bool {
	xkey := r.Header.Get("X_API_KEY")
	if xkey != modUtility.G_XApiKey {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid api key provided"))
		return false
	}

	return true
}

func handlerGetConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	xkey := r.Header.Get("X_API_KEY")
	if xkey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("ERROR"))
		return
	}
	config := modDatabase.DB_GetConfig(xkey)
	w.Write([]byte(config))

}
func handlerSetConfig(w http.ResponseWriter, r *http.Request) {
	xkey := r.Header.Get("X_API_KEY")
	if xkey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("ERROR"))
		return
	}
	if xkey != modUtility.G_XApiKey {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("WRONG"))
		return
	}
	key2 := r.Header.Get("X_API_KEY2")
	data2, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("WRONG2"))
		return
	}

	err = modDatabase.DB_InputConfig(key2, string(data2))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("WRONG23"))
		return
	}
	w.Write([]byte("OK"))
}

// func handlerErrorPut(w http.ResponseWriter, r *http.Request) {
// 	if !checkXApiKey(w, r) {
// 		return
// 	}
// 	idStr := chi.URLParam(r, "appid")
// 	logStr, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNoContent)
// 		w.Write([]byte("read content error"))
// 		return
// 	}
// 	err = modLogMongodb.GetSingleLog().AddLogError(string(logStr), idStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadGateway)
// 		w.Write([]byte("add log error: " + err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("done"))
// }
