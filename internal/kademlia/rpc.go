package kademlia

import (
	"encoding/json"
	"errors"

	"github.com/viktorfrom/d7024e-kademlia/pkg/randarr"
)

// RPCType type definition
type RPCType string

// RPC type declaration
const (
	Ping      = RPCType("PING")
	Store     = RPCType("STORE")
	FindValue = RPCType("FIND_VALUE")
	FindNode  = RPCType("FIND_NODE")
	OK        = RPCType("OK")
)

const (
	errWrongType = "unexpected rpc type given"
)

var rpcTypes = []RPCType{Ping, Store, FindValue, FindNode, OK}

// RPC contains the type of the RPC, the payload (data) as well as a quasi random ID for
// that RPC
type RPC struct {
	Type     *RPCType `json:"type"`
	Payload  *Payload `json:"payload"`
	ID       *string  `json:"id"`
	SenderIP *string  `json:"senderIP"`
}

// Payload contains the data sent in RPCs. Can contain a value and/or a list of contacts.
type Payload struct {
	Value    *string   `json:"value"`
	Contacts []Contact `json:"contacts"`
}

// NewRPC creates a new RPC with a random ID added to it
func NewRPC(rpc RPCType, senderID string, payload Payload) (*RPC, error) {
	err := validateRPCType(rpc)
	if err != nil {
		return nil, err
	}

	randomStr := randarr.RandomHexString(20)
	randomID := string(randomStr)
	newRPC := RPC{&rpc, &payload, &randomID, &senderID}

	return &newRPC, nil
}

func validateRPCType(rpc RPCType) error {
	for _, rpcType := range rpcTypes {
		if rpcType == rpc {
			return nil
		}
	}
	return errors.New(errWrongType)
}

// MarshalRPC serializes the RPC struct and returns the result as a byte array
func MarshalRPC(rpc RPC) ([]byte, error) {
	var data []byte
	data, err := json.Marshal(rpc)

	return data, err
}

// UnmarshalRPC deserializes the given byte array and returns an RPC
func UnmarshalRPC(data []byte) (*RPC, error) {
	rpc := RPC{}
	err := json.Unmarshal(data, &rpc)
	if err != nil {
		return nil, err
	}

	return &rpc, nil
}