package server

import (
	"context"

	"github.com/hashicorp/go-hclog")


type service struct{
	log hclog.Logger
}

func (s*service) getBMI(ctx context.Context, *protos.wahRequest) (*protos.bmiResponse, error){

		return &protos.bmiResponse{bmi: 1 }, nil
}