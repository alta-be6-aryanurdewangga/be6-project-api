package response

import "time"

type BookResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name  string `json:"name"`
	Publisher string `json:"publisher"`
	Author string `json:"author"`
}
