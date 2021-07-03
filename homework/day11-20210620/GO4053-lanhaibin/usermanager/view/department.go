package view

import (
	"net/http"
	"strconv"
	"usermanager/auth"
	"usermanager/model"
	"usermanager/service"
	"usermanager/template"
)

func HandleCreateDepartment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.PostForm.Get("name")
		err := service.Ds.Create(name)
		if err != nil {
			template.RenderTemplate(w, "department_create.html", struct {
				Current model.User
				Err     string
			}{auth.GetCurrentUser(r), err.Error()})
			return
		}
		http.Redirect(w, r, "/department", http.StatusFound)
		return
	}
	template.RenderTemplate(w, "department_create.html", struct {
		Current model.User
		Err     string
	}{auth.GetCurrentUser(r), ""})

}

func HandleDeleteDepartment(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	ids := r.FormValue("id")
	id, _ := strconv.Atoi(ids)
	service.Ds.Delete(id)

	http.Redirect(w, r, "/department", http.StatusFound)

}

func HandleQueryDepartment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s := r.Form.Get("q")
	departmentList := service.Ds.Query(s)
	template.RenderTemplate(w, "department.html", struct {
		Current     model.User
		Departments []model.Department
	}{auth.GetCurrentUser(r), departmentList})
}

func HandleModifyDepartment(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		r.ParseForm()
		ids := r.PostForm.Get("id")
		id, _ := strconv.Atoi(ids)
		name := r.PostForm.Get("name")

		service.Ds.Modify(id, name)
		http.Redirect(w, r, "/department", http.StatusFound)
	} else {
		ids := r.FormValue("id")
		id, _ := strconv.Atoi(ids)
		department, err := service.Ds.Get(id)
		if err != nil {
			http.Redirect(w, r, "/department", http.StatusFound)
			return
		}
		template.RenderTemplate(w, "department_modify.html", struct {
			Current    model.User
			Department model.Department
		}{auth.GetCurrentUser(r), department})
	}
}
