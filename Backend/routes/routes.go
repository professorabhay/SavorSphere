package routes

import (
   "net/http"
   "github.com/professorabhay/controllers"
)

func SetupRoutes() {
   http.HandleFunc("/auth/signup", controllers.SignUp)
}
