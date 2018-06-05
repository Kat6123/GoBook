package main

import (
	"html/template"
)

var templates = template.Must(template.ParseFiles("contributors.html"))

type (
	contributors []*contributor

	contributor struct {
		Login         string `json:"login"`
		Url           string `json:"html_url"`
		Admin         bool   `json:"site_admin"`
		Type          string `json:"type"`
		Contributions int    `json:"contributions"`
	}

	contribPage struct {
		Contributors  contributors
		TotalNumber   int
		CurrentNumber int
	}
)

//type issues []*issue
//type milestones []*milestone

//type issue struct {
//	Url   string `json:"html_url"`
//	Title string
//	State string
//	Body  string
//}
//
//type milestone struct {
//	Url          string `json:"html_url"`
//	Title        string
//	Description  string
//	OpenIssues   int `json:"open_issues"`
//	ClosedIssues int `json:"closed_issues"`
//	State        string
//}
