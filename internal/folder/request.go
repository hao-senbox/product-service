package folder

type CreateFolderRequest struct {
	Name     string  `json:"name"`
	ParentID *string `json:"parent_id"`
}

type UpdateFolderRequest struct {
	Name     string  `json:"name"`
	ParentID *string `json:"parent_id"`
}	