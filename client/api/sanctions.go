package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (a *API) IsSanctioned(walletAddress string, apiKey string) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://public.chainalysis.com/api/v1/address/" + walletAddress, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	var data SanctionsResponse
	_data, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(_data, &data)
	if err != nil {
		return 0, err
	}
	return len(data.Identifications), nil
}
