package asyncapiv2

// Info is a representation of the corresponding asyncapi object filled
// from an asyncapi specification that will be used to generate code.
// Source: https://www.asyncapi.com/docs/reference/specification/v2.6.0#infoObject
type Info struct {
	// --- AsyncAPI fields -----------------------------------------------------

	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`

	// --- Non AsyncAPI fields -------------------------------------------------
}
