package rpc

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      interface{}     `json:"id"`
}

type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *ErrorObj   `json:"error,omitempty"`
	ID      interface{} `json:"id"`
}

type ErrorObj struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HandlerFunc func(params json.RawMessage) (interface{}, *ErrorObj)

type Server struct {
	methods map[string]HandlerFunc
}

func NewServer() *Server {
	return &Server{
		methods: make(map[string]HandlerFunc),
	}
}

func (s *Server) Register(method string, fn HandlerFunc) {
	s.methods[method] = fn
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, nil, -32700, "parse error")
		return
	}

	if req.JSONRPC != "2.0" {
		writeError(w, req.ID, -32600, "invalid jsonrpc version")
		return
	}

	handler, ok := s.methods[req.Method]
	if !ok {
		writeError(w, req.ID, -32601, "method not found")
		return
	}

	result, rpcErr := handler(req.Params)
	if rpcErr != nil {
		writeError(w, req.ID, rpcErr.Code, rpcErr.Message)
		return
	}

	res := Response{
		JSONRPC: "2.0",
		Result:  result,
		ID:      req.ID,
	}

	json.NewEncoder(w).Encode(res)
}

func writeError(w http.ResponseWriter, id interface{}, code int, msg string) {
	res := Response{
		JSONRPC: "2.0",
		Error: &ErrorObj{
			Code:    code,
			Message: msg,
		},
		ID: id,
	}

	json.NewEncoder(w).Encode(res)
}