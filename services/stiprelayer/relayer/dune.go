package relayer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"os"
)

// DuneAPIKey is the API key for Dune, fetched from the environment variables.
var DuneAPIKey = os.Getenv("DUNE_API_KEY")

// ExecuteDuneQuery executes a predefined query on the Dune API and returns the http response.
func (s *STIPRelayer) ExecuteDuneQuery(parentCtx context.Context, queryType string) (executionID string, err error) {
	ctx, span := s.handler.Tracer().Start(parentCtx, "ExecuteDuneQuery", trace.WithAttributes(attribute.String("queryType", queryType)))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	client := &http.Client{}
	var queryID string
	if queryType == "bridge" {
		queryID = "3345214"
	} else if queryType == "rfq" {
		queryID = "3348161"
	}
	s.handler.ConfigureHTTPClient(client)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://api.dune.com/api/v1/query/%s/execute", queryID), bytes.NewBufferString(`{"performance": "large"}`))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Dune-API-Key", DuneAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute Dune query: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v", err)
		}
	}()

	fmt.Println("EXECUTING DUNE QUERY")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	var ok bool
	executionID, ok = result["execution_id"]
	if !ok {
		return "", fmt.Errorf("no execution_id found in response")
	}

	return executionID, nil
}

// GetExecutionResults fetches the results of a Dune query execution using the provided execution ID.
func (s *STIPRelayer) GetExecutionResults(parentCtx context.Context, executionID string) (_ *QueryResult, err error) {
	ctx, span := s.handler.Tracer().Start(parentCtx, "ExecuteDuneQuery", trace.WithAttributes(attribute.String("executionID", executionID)))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	client := &http.Client{}
	s.handler.ConfigureHTTPClient(client)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://api.dune.com/api/v1/execution/%s/results", executionID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Dune-API-Key", DuneAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get execution results: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v", err)
		}
	}()
	fmt.Println("GETTING EXECUTION RESULTS")

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	getResultsBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read execution results body: %w", err)
	}

	var jsonResult QueryResult
	err = json.Unmarshal(getResultsBody, &jsonResult)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &jsonResult, nil
}
