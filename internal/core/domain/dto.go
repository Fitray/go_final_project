package core_domain

type NewTaskResponse struct {
	Id int `json:"id"`
}

type GetTasksResponse struct {
	Tasks []Task `json:"tasks"`
}

type SignInRequest struct {
	Password string `json:"password"`
}
