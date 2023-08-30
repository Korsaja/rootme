package printer

import (
	"korsaj.io/rootme/pkg/models"
)

type Printer interface {
	PrintText(text string)
	PrintError(text string)
	PrintProfile(profile *models.Profile)
}
