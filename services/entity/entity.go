package entity

type ID = int32

type TableID struct {
	ID ID `json:"id" validate:"required,numeric"`
}
