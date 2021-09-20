package domain_user

import "encoding/json"

type PublicUser struct {
	Id        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName int64  `json:"first_name"`
	LastName  int64  `json:"last_name"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		userJson, _ := json.Marshal(user)
		var publicuser PrivateUser
		json.Unmarshal(userJson, &publicuser)
		return publicuser

	} else {
		userJson, _ := json.Marshal(user)
		var privateuser PrivateUser
		json.Unmarshal(userJson, &privateuser)
		return privateuser
	}
}
