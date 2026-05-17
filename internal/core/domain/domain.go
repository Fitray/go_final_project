package core_domain

type NextDateParams struct {
	Type   string
	Day    int
	Days   []int
	Months []int
}

type TaskRequest struct {
	Date    string  `json:"date"`
	Title   string  `json:"title"`
	Comment *string `json:"comment"`
	Repeat  string  `json:"repeat"`
}

type TaskResponce struct {
	Id int `json:"id"`
}
