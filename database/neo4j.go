package database

import (
	mantisError "github.com/sphireinc/mantis/error"
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
		mantisError.HandleError("error in Neo4j Connect()", err)
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
		mantisError.HandleError("error in Neo4j CypherQuery()", err)
		return nil, err
	}

	return query.Results, nil
}

// TransactCypherQuery creates a CyperQuery Transaction
func (n *Neo4j) TransactCypherQuery(queries []CypherQuery) (interface{}, error) {
	var cypherQuery []*neoism.CypherQuery

	transaction, err := n.conn.Begin(cypherQuery)
	if err != nil {
		mantisError.HandleError("error in Neo4j TransactCypherQuery() transaction begin", err)
		return nil, err
	}

	for _, query := range queries {
		query := neoism.CypherQuery{
			Statement:  query.Statement,
			Parameters: query.Parameters,
		}
		err = n.conn.Cypher(&query)
		if err != nil {
			mantisError.HandleError("error in Neo4j TransactCypherQuery() query build", err)
			return nil, err
		}
	}

	err = transaction.Commit()
	if err != nil {
		mantisError.HandleError("error in Neo4j TransactCypherQuery() transaction commit", err)
		return nil, err
	}

	return &queries, nil
}
