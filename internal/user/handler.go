package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shokhzodbek/internal/handlers"
)


const (
	usersUrl="/users"
	userUrl= "/users/:uuid"
)

type handler struct {

}

func New() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router){
	router.GET(usersUrl,h.GetList)
	router.GET(userUrl,h.GetUser)
	router.POST(usersUrl,h.CreateUser)
	router.DELETE(userUrl,h.DeleteUser)
	router.PUT(userUrl,h.UpdateUser)
	router.PATCH(userUrl,h.UpdetePartUser)
}

func (h *handler) GetList(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	rw.Write([]byte("Get users"))

}

func (h *handler) GetUser(rw http.ResponseWriter,r *http.Request,p httprouter.Params) {
	
	rw.Write([]byte("Get one user"))
}

func (h *handler) CreateUser(rw http.ResponseWriter,r *http.Request,p httprouter.Params) {
	
	rw.Write([]byte("Create one user"))
}

func (h *handler) UpdetePartUser(rw http.ResponseWriter,r *http.Request,p httprouter.Params) {
	
	rw.Write([]byte("Update one user"))
}

func (h *handler) UpdateUser(rw http.ResponseWriter,r *http.Request,p httprouter.Params) {
	
	rw.Write([]byte("Update one user"))
}

func (h *handler) DeleteUser(rw http.ResponseWriter,r *http.Request,p httprouter.Params) {
	
	rw.Write([]byte("Delete one user"))
}