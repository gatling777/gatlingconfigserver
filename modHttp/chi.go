package modHttp

import (
	"configServer/modUtility"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
)

type CChi struct {
	router *chi.Mux
}

var g_singleChi *CChi = &CChi{}

func getSingleChi() *CChi {
	return g_singleChi
}

func (pInst *CChi) Initialize() error {
	r := chi.NewRouter()
	r.Get("/", handlerHomepage)
	//r.Put("/info/{appid}", handlerInfoPut)
	//r.Put("/error/{appid}", handlerErrorPut)
	pInst.router = r

	return nil
}

func (pInst *CChi) start() error {
	listenStr := fmt.Sprintf(":%d", modUtility.G_HttpPort)
	fmt.Println("start listen: ", listenStr)
	err := http.ListenAndServe(listenStr, pInst.router)
	return err
}
