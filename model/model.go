package model

import (

	
)

type Project struct {
	Id          string `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Title       string `form:"title" json:"title"`
	Date        string `form:"date" json:"date"`
	SAPNumber   string `form:"sapnumber" json:"sapnumber"`
	Notes       string `form:"notes" json:"notes"`
	BranchID    string `form:"branch_id" json:"branch_id"`
	StatusID    string `form:"status_id" json:"status_id"`
	ServiceID   string `form:"service_id" json:"service_id"`
	ServiceName string `form:"service_name" json:"service_name"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

