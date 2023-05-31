package models

import (
	pb_models "github.com/pocketbase/pocketbase/models"
)

type Collection struct {
	pb_models.BaseModel

	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Schema     []Schema `json:"schema"`
	ListRule   string   `json:"listRule"`
	ViewRule   string   `json:"viewRule"`
	CreateRule string   `json:"createRule"`
	UpdateRule string   `json:"updateRule"`
	DeleteRule string   `json:"deleteRule"`
	Options    any      `json:"options"`
	Indexes    []string `json:"indexes"`
}

type Schema struct {
	System   bool   `json:"system"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Unique   bool   `json:"unique"`
	Options  any    `json:"options"`
}

type CollectionList struct {
	Page        int          `json:"page"`
	PerPage     int          `json:"perPage"`
	TotalItems  int          `json:"totalItems"`
	TotalPages  int          `json:"totalPages"`
	Collections []Collection `json:"items"`
}

type CollectionCreateRequest struct {
	// optional id
	Id *string `json:"id,omitempty"`
	// required name
	Name string `json:"name"`
	// required type
	Type string `json:"type"`
	// required schema
	Schema []Schema `json:"schema"`
	// optional system
	System *bool `json:"system,omitempty"`
	// optional listRule
	ListRule *string `json:"listRule,omitempty"`
	// optional viewRule
	ViewRule *string `json:"viewRule,omitempty"`
	// optional createRule
	CreateRule *string `json:"createRule,omitempty"`
	// optional updateRule
	UpdateRule *string `json:"updateRule,omitempty"`
	// optional deleteRule
	DeleteRule *string `json:"deleteRule,omitempty"`
	// optional options
	Options *any `json:"options,omitempty"`
	// optional indexes
	Indexes *[]string `json:"indexes,omitempty"`
}

type CollectionUpdateRequest struct {
	// required name
	Name *string `json:"name,omitempty"`
	// required type
	Type *string `json:"type,omitempty"`
	// required schema
	Schema *[]Schema `json:"schema,omitempty"`
	// optional system
	System *bool `json:"system,omitempty"`
	// optional listRule
	ListRule *string `json:"listRule,omitempty"`
	// optional viewRule
	ViewRule *string `json:"viewRule,omitempty"`
	// optional createRule
	CreateRule *string `json:"createRule,omitempty"`
	// optional updateRule
	UpdateRule *string `json:"updateRule,omitempty"`
	// optional deleteRule
	DeleteRule *string `json:"deleteRule,omitempty"`
	// optional options
	Options *any `json:"options,omitempty"`
	// optional indexes
	Indexes *[]string `json:"indexes,omitempty"`
}

type CollectionImportRequest struct {
	// required collections
	Collections []Collection `json:"collections"`
	// optional deleteMissing bool
	DeleteMissing *bool `json:"deleteMissing,omitempty"`
}
