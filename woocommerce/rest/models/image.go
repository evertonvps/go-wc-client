package models

type Image struct {
	ID              int    `json:"id"`
	DateCreated     string `json:"date_created"`
	DateCreatedGmt  string `json:"date_created_gmt"`
	DateModified    string `json:"date_modified"`
	DateModifiedGmt string `json:"date_modified_gmt"`
	Src             string `json:"src"`
	Name            string `json:"name"`
	Alt             string `json:"alt"`
}
