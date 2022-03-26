package sthooks

import "context"

type Hook func(context.Context) context.Context

type ClientHooks struct {
	OnBegin               Hook
	OnStartPerformRequest Hook
	OnDonePerformRequest  Hook
	OnStartHandleResponse Hook
	OnDoneHandleResponse  Hook
	OnEnd                 Hook
}
