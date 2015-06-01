package goauth0Request

import (
	"fmt"

	"github.com/franela/goreq"
)

type Get struct {
	Domain     string
	Client_id  string
	Username   string
	Password   string
	Connection string
	Grant_type string
	Scope      string
	Uri        string
}

type Auth0Prepost struct {
	Client_id  string `json:"client_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Connection string `json:"connection"`
	Grant_type string `json:"grant_type"`
	Scope      string `json:"scope"`
	Device     string `json:"device"`
}

func (getR Get) Do() (*goreq.Response, error) {
	type Bearer struct {
		Id_token string `json:"id_token"`
	}

	item := Auth0Prepost{
		getR.Client_id,
		getR.Username,
		getR.Password,
		getR.Connection,
		getR.Grant_type,
		getR.Scope,
		"",
	}
	res, err := goreq.Request{
		Method:      "POST",
		Uri:         getR.Domain,
		Accept:      "application/json",
		ContentType: "application/json",
		Body:        item,
	}.Do()
	if err != nil {
		return nil, err
	}

	var token Bearer
	res.Body.FromJsonTo(&token)

	req := goreq.Request{
		Method: "GET",
		Uri:    getR.Uri,
	}
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", token.Id_token))
	return req.Do()
}
