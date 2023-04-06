package configs

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Create a client to talk to the google sheet
func GetGoogleClient() (*sheets.Service, error) {
	ctx := context.Background()
	apiKey := os.Getenv("GOOGLE_API_KEY")

	srv, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve Sheets client: %v", err)

	}
	return srv, nil
}

// Read data from the google sheet via cell range
func ReadRange(srv *sheets.Service, spreadsheetID string, sheetName string, rangeName string) ([][]interface{}, error) {
	readRange := sheetName + "!" + rangeName
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("no data found")
	}

	// return resp.Values[0], nil
	var rows [][]interface{}
    for _, row := range resp.Values {
        rows = append(rows, row)
    }

    return rows, nil
}
