package main

import (
	"fmt"
	"go-rest-api/helpers"
	"go-rest-api/login"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func main() {
	//create chi router
	r := chi.NewRouter()
	//token verification and authorization
	utils.InitTokenAuth()
	r.Use(jwtauth.Verifier(utils.TokenAuth()))
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		uName := r.Form.Get("username")
		pwd := r.Form.Get("password")
		uNameCheck := helpers.IsEmpty(uName)
		pwdCheck := helpers.IsEmpty(pwd)
		if uNameCheck || pwdCheck {
			fmt.Fprintf(w, "There is empty data.")
			return
		}
		//login check
		if login.Login(uName, pwd) {
			role := "user"
			if uName == "admin" {
				role = "admin"
			}

			token, err := utils.GenerateToken(models.JwtModel{
				Name:     uName,
				Password: pwd,
				Role:     role,
			})
			if err != nil {
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Login Successful.\n Token:\n %s", token)
		} else {
			fmt.Fprintf(w, "Login failed")
		}
	})
	//users endpoint
	r.Route("/{username}", func(r chi.Router) {
		r.Get("/todos")
		r.Post("/todos")
		r.Put("/todos/{id}")
		r.Delete("/todos/{id}")
	})
	//admin endpoint
	r.Route("/admin", func(r chi.Router) {
		r.Get("/todos")
		r.Post("/todos")
		r.Put("/todos/{id}")
		r.Delete("/todos/{id}")

		// authorized endpoint
		r.Get("/users/{username}/todos")
		r.Post("/users/{username}/todos")
		r.Put("/users/{username}/todos/{id}")
		r.Delete("/users/{username}/todos/{id}")
	})
}
