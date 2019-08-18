package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type Levels struct {
	xontrols.APICtrl
}

func (req *Levels) Get() {
	req.Serve(http.StatusNotImplemented, nil, nil)
}
