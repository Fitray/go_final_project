package core_domain

type NextDateParams struct {
	Type   string
	Day    int
	Days   []int
	Months []int
}

type NewTaskResponse struct {
	Id int `json:"id"`
}

type Task struct {
	Date    string  `json:"date"`
	Title   string  `json:"title"`
	Comment *string `json:"comment"`
	Repeat  string  `json:"repeat"`
	Id      string  `json:"id"`
}

type GetTasksResponse struct {
	Tasks []Task `json:"tasks"`
}

type SignInRequest struct {
	Password string `json:"password"`
}
