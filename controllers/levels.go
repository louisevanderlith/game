package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Levels struct {
}

func (req *Levels) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
}
