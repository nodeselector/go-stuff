package strest

import "github.com/nodeselector/go-stuff/sthttp"

type Options struct {
	Attempter sthttp.Attempter
}

type Option func(*Options)
