package handlers

import (
	"log"
	"net/http"

	"github.com/hoangminhtran94/room-bookings/internal/config"
	"github.com/hoangminhtran94/room-bookings/internal/form"
	"github.com/hoangminhtran94/room-bookings/internal/models"
	"github.com/hoangminhtran94/room-bookings/internal/render"
)

// TemplateData hold page data

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	remoteId := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteId)
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(res,req, "home.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About page handler
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	// stringMap := make(map[string]string)
	// stringMap["text"] = "st"
	render.RenderTemplate(res,req, "about.html", &models.TemplateData{})
}

func (m *Repository) Resevation(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res,req, "make-reservation.html", &models.TemplateData{
		Form: form.New(nil),
	})
}
//Handle post request to make reservation
func (m *Repository) PostResevation(res http.ResponseWriter, req *http.Request) {
	err :=req.ParseForm();
	if(err != nil) {
		log.Println(err)
		return
	}
	reservation:= models.Reservation {
		FirstName: req.Form.Get("first_name"),
		LastName: req.Form.Get("last_name"),
		Email: req.Form.Get("email"),
		Phone: req.Form.Get("phone"),
	}

	formdata:=form.New(req.PostForm);
	//Check if there is firstname in formdata
	formdata.Has("first_name")

	if(!formdata.Valid()) {
		data :=make(map[string]interface{})
		data["reservation"]  = reservation
		render.RenderTemplate(res,req, "make-reservation.html", &models.TemplateData{
			Form:formdata,
			Data:data,
	
		})
		return
	}
	

}
func (m *Repository) Availability(res http.ResponseWriter, req *http.Request) {
	// stringMap := make(map[string]string)
	// stringMap["text"] = "st"
	render.RenderTemplate(res,req, "availability.html", &models.TemplateData{})
}
func (m *Repository) PostAvailability(res http.ResponseWriter, req *http.Request) {
	// formData := req.Form
	// start := formData.Get("start")
	// end := formData.Get("end")
	render.RenderTemplate(res,req, "availability.html", &models.TemplateData{})
}

func (m *Repository) Generals(res http.ResponseWriter, req *http.Request) {
	// stringMap := make(map[string]string)
	// stringMap["text"] = "st"
	render.RenderTemplate(res,req, "generals.html", &models.TemplateData{})
}
func (m *Repository) MajorSuite(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res,req,"majors.html",&models.TemplateData{})
}

func (m *Repository) Contact(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res,req,"contact.html",&models.TemplateData{})
}