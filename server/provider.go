package main

import (
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewProvider() (client.ConfigProvider, error) {
	return session.NewSession()
}
