package controller

import (
	"encoding/json"
	"github/Shitomo/producer/adapter"
	"github/Shitomo/producer/usecase/port"
	"net/http"
)

type UserRegisterController struct {
	input  port.UserRegisterInputPort
	output port.UserRegisterOutputPort
}

func NewUserRegisterController(
	input port.UserRegisterInputPort,
	output port.UserRegisterOutputPort,
) *UserRegisterController {
	return &UserRegisterController{
		input:  input,
		output: output,
	}
}

type UserRegisterPostBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDay  string `json:"birthday"`
}

func (c *UserRegisterController) Post(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(adapter.SetResWriter(r.Context(), w))

	if r.Method != http.MethodPost {
		c.output.ErrorRender(r.Context(), adapter.NewMethodNotAllowedError())

		return
	}

	var body UserRegisterPostBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	input, err := port.NewUserRegisterInput(body.FirstName, body.LastName, body.BirthDay)
	if err != nil {
		c.output.ErrorRender(r.Context(), adapter.NewBadRequestError(err.Error()))

		return
	}

	c.input.Execute(r.Context(), input)
}
