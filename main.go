package main

import (
	"net/http"

	"github.com/emicklei/go-restful"
)

// User user
type User struct {
	Id   string `xml:"id" json:"id"`
	Name string `xml:"name" json:"name"`
}
type Users struct {
	Users []User `xml:"user" json:"user"`
}

func main() {
	ws := new(restful.WebService)
	ws.
		Path("/user").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(userFind))
	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

func userFind(req *restful.Request, res *restful.Response) {
	// TODO
	id := req.PathParameter("user-id")
	usr := &User{Id: id, Name: "荣锋亮"}

	res.WriteHeaderAndEntity(http.StatusOK, Users{Users: []User{*usr}})
}
