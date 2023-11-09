package model

import (
	"github.com/CoulsonChen/BitoAPI/internal/constant"
)

type Person struct {
	Id     int `json:"-"`
	Name   string
	Height float64
	Gender constant.Gender
	Dates  int
}
