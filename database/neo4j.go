package database

import (
	"gopkg.in/jmcvetta/neoism.v1"
	"net/url"
)

type Neo4j struct {
	conn *neoism.Database
	DSN  url.URL
}

type CypherQuery struct {
	Results    interface{}
	Statement  string
	Parameters neoism.Props
}

// Connect attempts to connect to the DB
func (n *Neo4j) Connect() error {
	var err error
	n.conn, err = neoism.Connect(n.DSN.String())
	if err != nil {
		return err
	}
	return nil
}

// NewNode creates a new node
func (n *Neo4j) NewNode(node neoism.Props) (*neoism.Node, error) {
	return n.conn.CreateNode(node)
}

// CypherQuery perform a query - results will be populated with the query results - it must be a slice of structs.
func (n *Neo4j) CypherQuery(query CypherQuery) (interface{}, error) {
	err := n.conn.Cypher(&neoism.CypherQuery{
		Statement:  query.Statement,
		Parameters: query.Parameters,
		Result:     &query.Results,
	})

	if err != nil {
		return nil, err
	}

	return query.Results, nil
}

// TransactCypherQuery creates a CyperQuery Transaction
func (n *Neo4j) TransactCypherQuery(queries []CypherQuery) (interface{}, error) {
	var cypherQuery []*neoism.CypherQuery

	transaction, err := n.conn.Begin(cypherQuery)
	if err != nil {
		return nil, err
	}

	for _, query := range queries {
		query := neoism.CypherQuery{
			Statement:  query.Statement,
			Parameters: query.Parameters,
		}
		err = n.conn.Cypher(&query)
		if err != nil {
			return nil, err
		}
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}

	return &queries, nil
}
