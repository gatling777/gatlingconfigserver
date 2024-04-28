package modHttp

import (
	"configServer/modUtility"
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

// func handlerInfoPut(w http.ResponseWriter, r *http.Request) {
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
// 	err = modLogMongodb.GetSingleLog().AddLogInfo(string(logStr), idStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadGateway)
// 		w.Write([]byte("add log error: " + err.Error()))
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("done"))
// }

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
