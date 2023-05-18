// Code generated by woco, DO NOT EDIT.

package oas

import (
	"github.com/google/uuid"
)

type DeleteSilenceRequest struct {
	UriParams DeleteSilenceRequestUriParams
}

type DeleteSilenceRequestUriParams struct {
	SilenceID uuid.UUID `binding:"required" uri:"silenceID"`
}

type GetSilenceRequest struct {
	UriParams GetSilenceRequestUriParams
}

type GetSilenceRequestUriParams struct {
	SilenceID uuid.UUID `binding:"required" uri:"silenceID"`
}

type GetSilencesRequest struct {
	Body []string `form:"filter"`
}

type PostSilencesRequest struct {
	Body PostableSilence
}

type PostSilencesResponse struct {
	SilenceID string `json:"silenceID,omitempty"`
}
