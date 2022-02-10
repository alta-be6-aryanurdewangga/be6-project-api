package response

import "time"

type ProResponse struct {
	Id         uint      `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `jsonL:"updated_at"`
	Name       string    `json:"name"`
}
