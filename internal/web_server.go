package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type JSONRPCRequestBody struct {
	JSONRPCVersion string `json:"jsonrpc,omitempty"`
	Method         string `json:"method,omitempty"`
	Params         []any  `json:"params,omitempty"`
	ID             int64  `json:"id,omitempty"`
}

func StartServer() error {
	http.HandleFunc("/", handleJSONRPCRequest)
	return http.ListenAndServe(":8080", nil)
}

func handleJSONRPCRequest(responseWriter http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		respond(responseWriter, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	headerContentType := req.Header.Get("Content-Type")
	// Content-Type SHOULD be 'application/json-rpc' but MAY be
	// 'application/json' or 'application/jsonrequest'.
	// See https://www.jsonrpc.org/historical/json-rpc-over-http.html.
	if !slices.Contains([]string{"application/json", "application/json-rpc", "application/jsonrequest"}, headerContentType) {
		respond(responseWriter, "Content-Type not supported.", http.StatusUnsupportedMediaType)
		return
	}

	body, err := decodeBody(req)
	if err != nil {
		respond(responseWriter, fmt.Sprintf("Request body could not be parsed, err: %s", err.Error()), http.StatusBadRequest)
		return
	}

	zap.L().Info("Request received.", zap.String("method", req.Method), zap.String("path", req.URL.Path), zap.String("query", req.URL.RawQuery), zap.Any("body", body))
}

func decodeBody(req *http.Request) (JSONRPCRequestBody, error) {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	var body JSONRPCRequestBody

	err := decoder.Decode(&body)

	return body, err
}

func respond(responseWriter http.ResponseWriter, message string, httpStatusCode int) {
	responseWriter.WriteHeader(httpStatusCode)

	if message == "" {
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)

	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	_, err := responseWriter.Write(jsonResp)

	if err != nil {
		zap.L().Error("Failed to write response.", zap.Error(err))
	}
}