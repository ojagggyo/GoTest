package main

import (
	"encoding/json"
	"net/http"

	"rpcserver/rpc"
)

func main() {
	server := rpc.NewServer()

	// config取得API
	server.Register("database_api.get_config", func(params json.RawMessage) (interface{}, *rpc.ErrorObj) {

		config := map[string]interface{}{
			"IS_TEST_NET":      false,
			"ENABLE_FEATURE_X": true,
			"CHAIN_ID":         "test-chain",
		}

		return config, nil
	})

	// echo API
	server.Register("echo", func(params json.RawMessage) (interface{}, *rpc.ErrorObj) {
		var v interface{}
		json.Unmarshal(params, &v)
		return v, nil
	})

	http.Handle("/rpc", server)

	http.ListenAndServe(":8111", nil)
}