package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type EditBannerInput struct {
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	Status         string `json:"status"`
	Tags           string `json:"tags"`
	MediaType      string `json:"mediaType"`
	MediaReference string `json:"mediaReference"`
	ExternalLink   string `json:"externalLink"`
	Priority       string `json:"priority"`
}

func (input EditBannerInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Title, validation.Required, validation.Length(5, 100)),
		validation.Field(&input.Summary, validation.Length(10, 150)),
		validation.Field(&input.Status, validation.Required),
		validation.Field(&input.Tags, validation.Required),
		validation.Field(&input.MediaType, validation.Required),
		validation.Field(&input.MediaReference, validation.Required, is.URL),
		validation.Field(&input.ExternalLink, is.URL),
		validation.Field(&input.Priority, validation.Required, is.Digit),
	)
}
