package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Levels struct {
}

func (req *Levels) Get(ctx context.Contexer) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
