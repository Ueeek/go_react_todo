package router

import (
	"server/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

    //LOGIN APIS
    router.HandleFunc("/signup", middleware.Signup).Methods("POST")
    router.HandleFunc("/login", middleware.Login).Methods("POST")
    
    //TODO-APIS
	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET","OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/deleteDoneTask",middleware.DeleteDoneTask).Methods("DELETE","OPTIONS")
	return router
}
