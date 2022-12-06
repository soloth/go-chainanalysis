package api

type API struct {
}

type SanctionsResponse struct {
	Identifications []struct {
		Category    string `json:"category"`
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"identifications"`
}
