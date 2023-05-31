package models

import (
	pb_models "github.com/pocketbase/pocketbase/models"
)

type Admin struct {
	pb_models.BaseModel

	Avatar int    `json:"avatar"`
	Email  string `json:"email"`
}

type AdminList struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"perPage"`
	TotalItems int     `json:"totalItems"`
	TotalPages int     `json:"totalPages"`
	Admins     []Admin `json:"items"`
}

type AdminCreateRequest struct {
	// optional id
	Id *string `json:"id,omitempty"`
	// required email
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	// optional avatar
	Avatar *int `json:"avatar,omitempty"`
}

type AdminUpdateRequest struct {
	// optional email
	Email *string `json:"email,omitempty"`
	// optional password
	Password        *string `json:"password,omitempty"`
	PasswordConfirm *string `json:"passwordConfirm,omitempty"`
	// optional avatar
	Avatar *int `json:"avatar,omitempty"`
}
