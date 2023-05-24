package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/config"
	"example.com/model"
)

// GetDB returns a database object
func GetDB() (*sql.DB) {
    db,err := config.ConnectDB()
    if err!=nil{
        log.Fatal(err)
    }
    return db
}

// AllProject = Select Project API
func AllProject(w http.ResponseWriter, r *http.Request) {
    var Project model.Project
    var response model.Response
    var arrProject []model.Project
 
    db := GetDB()
    defer db.Close()
 
    rows, err := db.Query("SELECT id, name,title, date,sapnumber,notes,branch_id,status_id,service_id,service_name FROM project")
 
    if err != nil {
        log.Print(err)
    }
 
    for rows.Next() {
        err = rows.Scan(&Project.Id, &Project.Name,&Project.Title ,&Project.Date, &Project.SAPNumber,&Project.Notes,&Project.BranchID,&Project.StatusID,&Project.ServiceID,&Project.ServiceName)
        
        if err != nil {
            log.Fatal(err.Error())
        } else {
            arrProject = append(arrProject, Project)
        }
    }
 
    response.Status = 200
    response.Message = "Success"
    response.Data = arrProject
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}

// InsertProject = Insert Project API
func InsertProject(w http.ResponseWriter, r *http.Request) {
    var response model.Response
 
    db := GetDB()
    defer db.Close()
 
    err := r.ParseMultipartForm(4096)
    if err != nil {
        panic(err)
    }
    name := r.FormValue("name")
    date := r.FormValue("date")
    title :=r.FormValue("title")
    sapnumber :=r.FormValue("sapnumber")
    notes :=r.FormValue("notes")
    branch_id :=r.FormValue("branch_id")
    status_id :=r.FormValue("status_id")
    service_id :=r.FormValue("service_id")
    service_name :=r.FormValue("service_name")

    _, err = db.Exec("INSERT INTO Project(id,name,title, date,sapnumber,notes,branch_id,status_id,service_id,service_name) VALUES(?, ?,?,?,?,?,?,?,?,?)", name,title, date,sapnumber,notes,branch_id,status_id,service_id,service_name )
 
    if err != nil {
        log.Print(err)
        return
    }
    response.Status = 200
    response.Message = "Insert data successfully"
    fmt.Print("Insert data to database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}

// GetProjectByID = Get Project by ID API
func GetProjectByID(w http.ResponseWriter, r *http.Request) {
    var Project model.Project
    var response model.Response
 
    db := GetDB()
    defer db.Close()
 
    id := r.URL.Query().Get("id")
 
    err := db.QueryRow("SELECT id, name,title, date,sapnumber,notes,branch_id,status_id,service_id,service_name FROM project WHERE id = ?", id).Scan(&Project.Id, &Project.Name,&Project.Title ,&Project.Date, &Project.SAPNumber,&Project.Notes,&Project.BranchID,&Project.StatusID,&Project.ServiceID,&Project.ServiceName )
 
    switch {
    case err == sql.ErrNoRows:
        log.Printf("No Project with that ID.")
        response.Status = 404
        response.Message = "Project not found"
    case err != nil:
        log.Fatal(err.Error())
    default:
        response.Status = 200
        response.Message = "Success"
        response.Data = Project
    }
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}

// UpdateProject = Update Project API
func UpdateProject(w http.ResponseWriter, r *http.Request) {
    var response model.Response
 
    db := GetDB()
    defer db.Close()
 
    err := r.ParseMultipartForm(4096)
    if err != nil {
        panic(err)
    }
    id := r.FormValue("id")
    name := r.FormValue("name")
    title := r.FormValue("title")
 
    _, err = db.Exec("UPDATE Project SET name = ?, title =? WHERE id = ?", name, title, id)
 
    if err != nil {
        log.Print(err)
        return
    }
    response.Status = 200
    response.Message = "Update data successfully"
    fmt.Print("Update data to database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}

// DeleteProject = Delete Project API
func DeleteProject(w http.ResponseWriter, r *http.Request) {
    var response model.Response
 
    db := GetDB()
    defer db.Close()
 
    id := r.URL.Query().Get("id")
 
    _, err := db.Exec("DELETE FROM Project WHERE id = ?", id)
 
    if err != nil {
        log.Print(err)
        return
    }
    response.Status = 200
    response.Message = "Delete data successfully"
    fmt.Print("Delete data from database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}
