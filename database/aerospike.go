package database

import (
	"encoding/json"
	"github.com/aerospike/aerospike-client-go"
)

// AerospikeError provides a common struct for Aerospike Errors
type AerospikeError struct {
	Name  string
	Error error
}

// Aerospike is our primary struct
type Aerospike struct {
	Client          *aerospike.Client
	Hostname        string
	Port            int
	TLSName         string
	Policy          *aerospike.ClientPolicy
	GlobalNamespace string
	GlobalSetName   string
}

// Default creates a default config based on given parameters
func (a *Aerospike) Default(hostname string, port int) error {
	a.Hostname = hostname
	a.Port = port
	return a.Connect()
}

// String returns our Aerospike struct as a JSON string
func (a *Aerospike) String() string {
	marshaledStruct, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// Connect to the database
func (a *Aerospike) Connect() error {
	var err error
	if a.Policy != nil {
		if len(a.TLSName) > 0 {
			a.Client, err = aerospike.NewClientWithPolicyAndHost(a.Policy, &aerospike.Host{
				Name:    a.Hostname,
				TLSName: a.TLSName,
				Port:    a.Port,
			})
			return err
		}
		a.Client, err = aerospike.NewClientWithPolicy(a.Policy, a.Hostname, a.Port)
		return err
	}
	a.Client, err = aerospike.NewClient(a.Hostname, a.Port)
	return err
}

// Get a key value pair from our Redis DB
func (a *Aerospike) Get(key *aerospike.Key) (*aerospike.Record, error) {
	return a.Client.Get(nil, key)
}

// Set a key value pair in our Redis DB
func (a *Aerospike) Set(namespace, setName, key, bin, value string) error {
	if len(namespace) == 0 {
		a.GlobalNamespace = namespace
	}
	if len(setName) == 0 {
		a.GlobalSetName = setName
	}

	innerKey, err := aerospike.NewKey(namespace, setName, key)
	if err != nil {
		return err
	}

	return a.Client.PutBins(nil, innerKey, aerospike.NewBin(bin, value))
}
