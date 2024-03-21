package model

import (
	"github.com/anujsharma13/model/constants"
)

type Workers struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	PhoneNumber string               `json:"phone,omitempty"`
	City        string               `json:"city,omitempty"`
	State       string               `json:"state,omitempty"`
	PinCode     string               `json:"pincode,omitempty"`
	WorkerType  constants.WorkerType `json:"workertype"`
}
