package main

type customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var cs_1 = customer{ID: "1", Name: "Luu Thang 1", Role: "Admin", Email: "thanglq1@gmail.com", Phone: "0703333191", Contacted: false}
var cs_2 = customer{ID: "2", Name: "Luu Thang 2", Role: "Staff", Email: "thanglq2@gmail.com", Phone: "0703333192", Contacted: true}
var cs_3 = customer{ID: "3", Name: "Luu Thang 3", Role: "Guest", Email: "thanglq3@gmail.com", Phone: "0703333193", Contacted: false}
var cs_4 = customer{ID: "4", Name: "Luu Thang 4", Role: "Admin", Email: "thanglq4@gmail.com", Phone: "0703333194", Contacted: true}
var cs_5 = customer{ID: "5", Name: "Luu Thang 5", Role: "Staff", Email: "thanglq5@gmail.com", Phone: "0703333195", Contacted: false}
var cs_6 = customer{ID: "6", Name: "Luu Thang 6", Role: "Guest", Email: "thanglq6@gmail.com", Phone: "0703333196", Contacted: true}

// Mock Database to store customer data
var mock_database = map[string]customer{
	"1": cs_1,
	"2": cs_2,
	"3": cs_3,
	"4": cs_4,
	"5": cs_5,
	"6": cs_6}
