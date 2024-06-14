// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    executeMsg, err := UnmarshalExecuteMsg(bytes)
//    bytes, err = executeMsg.Marshal()
//
//    instantiateMsg, err := UnmarshalInstantiateMsg(bytes)
//    bytes, err = instantiateMsg.Marshal()
//
//    queryMsg, err := UnmarshalQueryMsg(bytes)
//    bytes, err = queryMsg.Marshal()
//
//    constructResponse, err := UnmarshalConstructResponse(bytes)
//    bytes, err = constructResponse.Marshal()
//
//    describeResponse, err := UnmarshalDescribeResponse(bytes)
//    bytes, err = describeResponse.Marshal()
//
//    selectResponse, err := UnmarshalSelectResponse(bytes)
//    bytes, err = selectResponse.Marshal()
//
//    storeResponse, err := UnmarshalStoreResponse(bytes)
//    bytes, err = storeResponse.Marshal()

package schema

import "bytes"
import "errors"

import "encoding/json"

func UnmarshalExecuteMsg(data []byte) (ExecuteMsg, error) {
	var r ExecuteMsg
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ExecuteMsg) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalInstantiateMsg(data []byte) (InstantiateMsg, error) {
	var r InstantiateMsg
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InstantiateMsg) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalQueryMsg(data []byte) (QueryMsg, error) {
	var r QueryMsg
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *QueryMsg) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalConstructResponse(data []byte) (ConstructResponse, error) {
	var r ConstructResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ConstructResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalDescribeResponse(data []byte) (DescribeResponse, error) {
	var r DescribeResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DescribeResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSelectResponse(data []byte) (SelectResponse, error) {
	var r SelectResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SelectResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStoreResponse(data []byte) (StoreResponse, error) {
	var r StoreResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StoreResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Execute messages
//
// Insert the data as RDF triples in the store. For already existing triples it acts as
// no-op.
//
// Only the smart contract owner (i.e. the address who instantiated it) is authorized to
// perform this action.
//
// Delete the data (RDF triples) from the store matching the patterns defined by the
// provided query. For non-existing triples it acts as no-op.
//
// Example: ```json { "prefixes": [ { "prefix": "foaf", "namespace":
// "http://xmlns.com/foaf/0.1/" } ], "delete": [ { "subject": { "variable": "s" },
// "predicate": { "variable": "p" }, "object": { "variable": "o" } } ], "where": [ {
// "simple": { "triplePattern": { "subject": { "variable": "s" }, "predicate": { "node": {
// "namedNode": {"prefixed": "foaf:givenName"} } }, "object": { "literal": { "simple":
// "Myrddin" } } } } }, { "simple": { "triplePattern": { "subject": { "variable": "s" },
// "predicate": { "variable": "p" }, "object": { "variable": "o" } } } } ] ```
//
// Only the smart contract owner (i.e. the address who instantiated it) is authorized to
// perform this action.
type ExecuteMsg struct {
	InsertData *InsertData `json:"insert_data,omitempty"`
	DeleteData *DeleteData `json:"delete_data,omitempty"`
}

type DeleteData struct {
	// Specifies the specific triple patterns to delete. If nothing is provided, the patterns                   
	// from the `where` clause are used for deletion.                                                           
	Delete                                                                                   []DeleteElement    `json:"delete"`
	// The prefixes used in the operation.                                                                      
	Prefixes                                                                                 []DeleteDataPrefix `json:"prefixes"`
	// Defines the patterns that data (RDF triples) should match in order for it to be                          
	// considered for deletion.                                                                                 
	Where                                                                                    []DeleteDataWhere  `json:"where"`
}

// Represents a triple pattern in a [SimpleWhereCondition].
type DeleteElement struct {
	// The object of the triple pattern.                            
	Object                                 DeleteVarOrNodeOrLiteral `json:"object"`
	// The predicate of the triple pattern.                         
	Predicate                              DeleteVarOrNode          `json:"predicate"`
	// The subject of the triple pattern.                           
	Subject                                DeleteVarOrNode          `json:"subject"`
}

// The object of the triple pattern.
//
// Represents either a variable, a node or a literal.
//
// A variable.
//
// A node, i.e. an IRI or a blank node.
//
// An RDF [literal](https://www.w3.org/TR/rdf11-concepts/#dfn-literal), i.e. a simple
// literal, a language-tagged string or a typed value.
type DeleteVarOrNodeOrLiteral struct {
	Variable *string        `json:"variable,omitempty"`
	Node     *PurpleNode    `json:"node,omitempty"`
	Literal  *PurpleLiteral `json:"literal,omitempty"`
}

// An RDF [literal](https://www.w3.org/TR/rdf11-concepts/#dfn-literal).
//
// A [simple literal](https://www.w3.org/TR/rdf11-concepts/#dfn-simple-literal) without
// datatype or language form.
//
// A [language-tagged
// string](https://www.w3.org/TR/rdf11-concepts/#dfn-language-tagged-string)
//
// A value with a datatype.
type PurpleLiteral struct {
	Simple               *string                     `json:"simple,omitempty"`
	LanguageTaggedString *PurpleLanguageTaggedString `json:"language_tagged_string,omitempty"`
	TypedValue           *PurpleTypedValue           `json:"typed_value,omitempty"`
}

type PurpleLanguageTaggedString struct {
	// The [language tag](https://www.w3.org/TR/rdf11-concepts/#dfn-language-tag).       
	Language                                                                      string `json:"language"`
	// The [lexical form](https://www.w3.org/TR/rdf11-concepts/#dfn-lexical-form).       
	Value                                                                         string `json:"value"`
}

type PurpleTypedValue struct {
	// The [datatype IRI](https://www.w3.org/TR/rdf11-concepts/#dfn-datatype-iri).          
	Datatype                                                                      PurpleIRI `json:"datatype"`
	// The [lexical form](https://www.w3.org/TR/rdf11-concepts/#dfn-lexical-form).          
	Value                                                                         string    `json:"value"`
}

// Represents an IRI.
//
// The [datatype IRI](https://www.w3.org/TR/rdf11-concepts/#dfn-datatype-iri).
//
// An IRI prefixed with a prefix. The prefixed IRI is expanded to a full IRI using the
// prefix definition specified in the query. For example, the prefixed IRI `rdf:type` is
// expanded to `http://www.w3.org/1999/02/22-rdf-syntax-ns#type`.
//
// A full IRI.
type PurpleIRI struct {
	Prefixed *string `json:"prefixed,omitempty"`
	Full     *string `json:"full,omitempty"`
}

// Represents either an IRI (named node) or a blank node.
//
// An RDF [IRI](https://www.w3.org/TR/rdf11-concepts/#dfn-iri).
//
// An RDF [blank node](https://www.w3.org/TR/rdf11-concepts/#dfn-blank-node).
type PurpleNode struct {
	NamedNode *PurpleIRI `json:"named_node,omitempty"`
	BlankNode *string    `json:"blank_node,omitempty"`
}

// The predicate of the triple pattern.
//
// Represents either a variable or a node.
//
// The subject of the triple pattern.
//
// A variable.
//
// A node, i.e. an IRI or a blank node.
type DeleteVarOrNode struct {
	Variable *string     `json:"variable,omitempty"`
	Node     *PurpleNode `json:"node,omitempty"`
}

// Represents a prefix, i.e. a shortcut for a namespace used in a query.
type DeleteDataPrefix struct {
	// The namespace associated with the prefix.       
	Namespace                                   string `json:"namespace"`
	// The prefix.                                     
	Prefix                                      string `json:"prefix"`
}

// Represents a condition in a [WhereClause].
//
// Represents a simple condition.
type DeleteDataWhere struct {
	Simple PurpleSimpleWhereCondition `json:"simple"`
}

// Represents a simple condition in a [WhereCondition].
//
// Represents a triple pattern, i.e. a condition on a triple based on its subject, predicate
// and object.
type PurpleSimpleWhereCondition struct {
	TriplePattern DeleteElement `json:"triple_pattern"`
}

type InsertData struct {
	// The data to insert. The data must be serialized in the format specified by the `format`         
	// field. And the data are subject to the limitations defined by the `limits` specified at         
	// contract instantiation.                                                                         
	Data                                                                                       string  `json:"data"`
	// The data format in which the triples are serialized. If not provided, the default format        
	// is [Turtle](https://www.w3.org/TR/turtle/) format.                                              
	Format                                                                                     *RDFXML `json:"format"`
}

// Instantiate message
type InstantiateMsg struct {
	// Limitations regarding store usage.                  
	Limits                               *StoreLimitsInput `json:"limits,omitempty"`
}

// Limitations regarding store usage.
//
// Contains requested limitations regarding store usages.
type StoreLimitsInput struct {
	// The maximum number of bytes the store can contain. The size of a triple is counted as the        
	// sum of the size of its subject, predicate and object, including the size of data types           
	// and language tags if any. Default to [Uint128::MAX] if not set, which can be considered          
	// as no limit.                                                                                     
	MaxByteSize                                                                                 *string `json:"max_byte_size,omitempty"`
	// The maximum number of bytes an insert data query can contain. Default to [Uint128::MAX]          
	// if not set, which can be considered as no limit.                                                 
	MaxInsertDataByteSize                                                                       *string `json:"max_insert_data_byte_size,omitempty"`
	// The maximum number of triples an insert data query can contain (after parsing). Default          
	// to [Uint128::MAX] if not set, which can be considered as no limit.                               
	MaxInsertDataTripleCount                                                                    *string `json:"max_insert_data_triple_count,omitempty"`
	// The maximum limit of a query, i.e. the maximum number of triples returned by a select            
	// query. Default to 30 if not set.                                                                 
	MaxQueryLimit                                                                               *int64  `json:"max_query_limit,omitempty"`
	// The maximum number of variables a query can select. Default to 30 if not set.                    
	MaxQueryVariableCount                                                                       *int64  `json:"max_query_variable_count,omitempty"`
	// The maximum number of bytes the store can contain for a single triple. The size of a             
	// triple is counted as the sum of the size of its subject, predicate and object, including         
	// the size of data types and language tags if any. The limit is used to prevent storing            
	// very large triples, especially literals. Default to [Uint128::MAX] if not set, which can         
	// be considered as no limit.                                                                       
	MaxTripleByteSize                                                                           *string `json:"max_triple_byte_size,omitempty"`
	// The maximum number of triples the store can contain. Default to [Uint128::MAX] if not            
	// set, which can be considered as no limit.                                                        
	MaxTripleCount                                                                              *string `json:"max_triple_count,omitempty"`
}

// Returns the resources matching the criteria defined by the provided query.
//
// Returns a description of the resource identified by the provided IRI as a set of RDF
// triples serialized in the provided format.
//
// Returns the resources matching the criteria defined by the provided query as a set of RDF
// triples serialized in the provided format.
type Select struct {
	Select    *SelectClass `json:"select,omitempty"`
	Describe  *Describe    `json:"describe,omitempty"`
	Construct *Construct   `json:"construct,omitempty"`
}

type Construct struct {
	// The format in which the triples are serialized. If not provided, the default format is               
	// [Turtle](https://www.w3.org/TR/turtle/) format.                                                      
	Format                                                                                   *RDFXML        `json:"format"`
	// The query to execute.                                                                                
	Query                                                                                    ConstructQuery `json:"query"`
}

// The query to execute.
//
// Represents a CONSTRUCT query over the triple store, allowing to retrieve a set of triples
// serialized in a specific format.
type ConstructQuery struct {
	// The triples to construct. If nothing is provided, the patterns from the `where` clause                     
	// are used for construction.                                                                                 
	Construct                                                                                  []ConstructElement `json:"construct"`
	// The prefixes used in the query.                                                                            
	Prefixes                                                                                   []QueryPrefix      `json:"prefixes"`
	// The WHERE clause. This clause is used to specify the triples to construct using variable                   
	// bindings.                                                                                                  
	Where                                                                                      []QueryWhere       `json:"where"`
}

// Represents a triple pattern in a [SimpleWhereCondition].
type ConstructElement struct {
	// The object of the triple pattern.                               
	Object                                 ConstructVarOrNodeOrLiteral `json:"object"`
	// The predicate of the triple pattern.                            
	Predicate                              ConstructVarOrNode          `json:"predicate"`
	// The subject of the triple pattern.                              
	Subject                                ConstructVarOrNode          `json:"subject"`
}

// The object of the triple pattern.
//
// Represents either a variable, a node or a literal.
//
// A variable.
//
// A node, i.e. an IRI or a blank node.
//
// An RDF [literal](https://www.w3.org/TR/rdf11-concepts/#dfn-literal), i.e. a simple
// literal, a language-tagged string or a typed value.
type ConstructVarOrNodeOrLiteral struct {
	Variable *string        `json:"variable,omitempty"`
	Node     *FluffyNode    `json:"node,omitempty"`
	Literal  *FluffyLiteral `json:"literal,omitempty"`
}

// An RDF [literal](https://www.w3.org/TR/rdf11-concepts/#dfn-literal).
//
// A [simple literal](https://www.w3.org/TR/rdf11-concepts/#dfn-simple-literal) without
// datatype or language form.
//
// A [language-tagged
// string](https://www.w3.org/TR/rdf11-concepts/#dfn-language-tagged-string)
//
// A value with a datatype.
type FluffyLiteral struct {
	Simple               *string                     `json:"simple,omitempty"`
	LanguageTaggedString *FluffyLanguageTaggedString `json:"language_tagged_string,omitempty"`
	TypedValue           *FluffyTypedValue           `json:"typed_value,omitempty"`
}

type FluffyLanguageTaggedString struct {
	// The [language tag](https://www.w3.org/TR/rdf11-concepts/#dfn-language-tag).       
	Language                                                                      string `json:"language"`
	// The [lexical form](https://www.w3.org/TR/rdf11-concepts/#dfn-lexical-form).       
	Value                                                                         string `json:"value"`
}

type FluffyTypedValue struct {
	// The [datatype IRI](https://www.w3.org/TR/rdf11-concepts/#dfn-datatype-iri).                  
	Datatype                                                                      VarOrNamedNodeIRI `json:"datatype"`
	// The [lexical form](https://www.w3.org/TR/rdf11-concepts/#dfn-lexical-form).                  
	Value                                                                         string            `json:"value"`
}

// Represents an IRI.
//
// The [datatype IRI](https://www.w3.org/TR/rdf11-concepts/#dfn-datatype-iri).
//
// An IRI prefixed with a prefix. The prefixed IRI is expanded to a full IRI using the
// prefix definition specified in the query. For example, the prefixed IRI `rdf:type` is
// expanded to `http://www.w3.org/1999/02/22-rdf-syntax-ns#type`.
//
// A full IRI.
type VarOrNamedNodeIRI struct {
	Prefixed *string `json:"prefixed,omitempty"`
	Full     *string `json:"full,omitempty"`
}

// Represents either an IRI (named node) or a blank node.
//
// An RDF [IRI](https://www.w3.org/TR/rdf11-concepts/#dfn-iri).
//
// An RDF [blank node](https://www.w3.org/TR/rdf11-concepts/#dfn-blank-node).
type FluffyNode struct {
	NamedNode *VarOrNamedNodeIRI `json:"named_node,omitempty"`
	BlankNode *string            `json:"blank_node,omitempty"`
}

// The predicate of the triple pattern.
//
// Represents either a variable or a node.
//
// The subject of the triple pattern.
//
// A variable.
//
// A node, i.e. an IRI or a blank node.
type ConstructVarOrNode struct {
	Variable *string     `json:"variable,omitempty"`
	Node     *FluffyNode `json:"node,omitempty"`
}

// Represents a prefix, i.e. a shortcut for a namespace used in a query.
type QueryPrefix struct {
	// The namespace associated with the prefix.       
	Namespace                                   string `json:"namespace"`
	// The prefix.                                     
	Prefix                                      string `json:"prefix"`
}

// Represents a condition in a [WhereClause].
//
// Represents a simple condition.
type QueryWhere struct {
	Simple FluffySimpleWhereCondition `json:"simple"`
}

// Represents a simple condition in a [WhereCondition].
//
// Represents a triple pattern, i.e. a condition on a triple based on its subject, predicate
// and object.
type FluffySimpleWhereCondition struct {
	TriplePattern ConstructElement `json:"triple_pattern"`
}

type Describe struct {
	// The format in which the triples are serialized. If not provided, the default format is              
	// [Turtle](https://www.w3.org/TR/turtle/) format.                                                     
	Format                                                                                   *RDFXML       `json:"format"`
	// The query to execute.                                                                               
	Query                                                                                    DescribeQuery `json:"query"`
}

// The query to execute.
//
// Represents a DESCRIBE query over the triple store, allowing to retrieve a description of
// a resource as a set of triples serialized in a specific format.
type DescribeQuery struct {
	// The prefixes used in the query.                                                                    
	Prefixes                                                                               []QueryPrefix  `json:"prefixes"`
	// The resource to describe given as a variable or a node.                                            
	Resource                                                                               VarOrNamedNode `json:"resource"`
	// The WHERE clause. This clause is used to specify the resource identifier to describe               
	// using variable bindings.                                                                           
	Where                                                                                  []QueryWhere   `json:"where"`
}

// The resource to describe given as a variable or a node.
//
// Represents either a variable or a named node (IRI).
//
// A variable.
//
// An RDF [IRI](https://www.w3.org/TR/rdf11-concepts/#dfn-iri).
type VarOrNamedNode struct {
	Variable  *string            `json:"variable,omitempty"`
	NamedNode *VarOrNamedNodeIRI `json:"named_node,omitempty"`
}

type SelectClass struct {
	// The query to execute.            
	Query                   SelectQuery `json:"query"`
}

// The query to execute.
//
// Represents a SELECT query over the triple store, allowing to select variables to return
// and to filter the results.
type SelectQuery struct {
	// The maximum number of results to return. If `None`, there is no limit. Note: the value of              
	// the limit cannot exceed the maximum query limit defined in the store limitations.                      
	Limit                                                                                       *int64        `json:"limit"`
	// The prefixes used in the query.                                                                        
	Prefixes                                                                                    []QueryPrefix `json:"prefixes"`
	// The items to select. Note: the number of items to select cannot exceed the maximum query               
	// variable count defined in the store limitations.                                                       
	Select                                                                                      []SelectItem  `json:"select"`
	// The WHERE clause. If `None`, there is no WHERE clause, i.e. all triples are returned                   
	// without filtering.                                                                                     
	Where                                                                                       []QueryWhere  `json:"where"`
}

// Represents an item to select in a [SelectQuery].
//
// Represents a variable.
type SelectItem struct {
	Variable string `json:"variable"`
}

// Represents the response of a [QueryMsg::Construct] query.
type ConstructResponse struct {
	// The data serialized in the specified format.       
	Data                                           string `json:"data"`
	// The format of the data.                            
	Format                                         RDFXML `json:"format"`
}

// Represents the response of a [QueryMsg::Describe] query.
type DescribeResponse struct {
	// The data serialized in the specified format.       
	Data                                           string `json:"data"`
	// The format of the data.                            
	Format                                         RDFXML `json:"format"`
}

// Represents the response of a [QueryMsg::Select] query.
type SelectResponse struct {
	// The head of the response, i.e. the set of variables mentioned in the results.        
	Head                                                                            Head    `json:"head"`
	// The results of the select query.                                                     
	Results                                                                         Results `json:"results"`
}

// The head of the response, i.e. the set of variables mentioned in the results.
//
// Represents the head of a [SelectResponse].
type Head struct {
	// The variables selected in the query.         
	Vars                                   []string `json:"vars"`
}

// The results of the select query.
//
// Represents the results of a [SelectResponse].
type Results struct {
	// The bindings of the results.                   
	Bindings                       []map[string]Value `json:"bindings"`
}

// Represents an IRI.
//
// Represents a literal S with optional language tag L or datatype IRI D.
//
// Represents a blank node.
type Value struct {
	Type                                Type        `json:"type"`
	// The value of the IRI.                        
	//                                              
	// The value of the literal.                    
	//                                              
	// The identifier of the blank node.            
	Value                               *ValueUnion `json:"value"`
	// The datatype of the literal.                 
	Datatype                            *Prefixed   `json:"datatype"`
	// The language tag of the literal.             
	XMLLang                             *string     `json:"xml:lang"`
}

// An IRI prefixed with a prefix. The prefixed IRI is expanded to a full IRI using the
// prefix definition specified in the query. For example, the prefixed IRI `rdf:type` is
// expanded to `http://www.w3.org/1999/02/22-rdf-syntax-ns#type`.
//
// A full IRI.
type Prefixed struct {
	Prefixed *string `json:"prefixed,omitempty"`
	Full     *string `json:"full,omitempty"`
}

// The value of the IRI.
//
// Represents an IRI.
//
// An IRI prefixed with a prefix. The prefixed IRI is expanded to a full IRI using the
// prefix definition specified in the query. For example, the prefixed IRI `rdf:type` is
// expanded to `http://www.w3.org/1999/02/22-rdf-syntax-ns#type`.
//
// A full IRI.
type ValueIRI struct {
	Prefixed *string `json:"prefixed,omitempty"`
	Full     *string `json:"full,omitempty"`
}

// Contains information related to triple store.
type StoreResponse struct {
	// The store limits.                   
	Limits                     StoreLimits `json:"limits"`
	// The store owner.                    
	Owner                      string      `json:"owner"`
	// The store current usage.            
	Stat                       StoreStat   `json:"stat"`
}

// The store limits.
//
// Contains limitations regarding store usages.
type StoreLimits struct {
	// The maximum number of bytes the store can contain. The size of a triple is counted as the       
	// sum of the size of its subject, predicate and object, including the size of data types          
	// and language tags if any.                                                                       
	MaxByteSize                                                                                 string `json:"max_byte_size"`
	// The maximum number of bytes an insert data query can contain.                                   
	MaxInsertDataByteSize                                                                       string `json:"max_insert_data_byte_size"`
	// The maximum number of triples an insert data query can contain (after parsing).                 
	MaxInsertDataTripleCount                                                                    string `json:"max_insert_data_triple_count"`
	// The maximum limit of a query, i.e. the maximum number of triples returned by a select           
	// query.                                                                                          
	MaxQueryLimit                                                                               int64  `json:"max_query_limit"`
	// The maximum number of variables a query can select.                                             
	MaxQueryVariableCount                                                                       int64  `json:"max_query_variable_count"`
	// The maximum number of bytes the store can contain for a single triple. The size of a            
	// triple is counted as the sum of the size of its subject, predicate and object, including        
	// the size of data types and language tags if any. The limit is used to prevent storing           
	// very large triples, especially literals.                                                        
	MaxTripleByteSize                                                                           string `json:"max_triple_byte_size"`
	// The maximum number of triples the store can contain.                                            
	MaxTripleCount                                                                              string `json:"max_triple_count"`
}

// The store current usage.
//
// Contains usage information about the triple store.
type StoreStat struct {
	// The total triple size in the store, in bytes.                 
	ByteSize                                                  string `json:"byte_size"`
	// The total number of IRI namespace present in the store.       
	NamespaceCount                                            string `json:"namespace_count"`
	// The total number of triple present in the store.              
	TripleCount                                               string `json:"triple_count"`
}

// Output in [RDF/XML](https://www.w3.org/TR/rdf-syntax-grammar/) format.
//
// Output in [Turtle](https://www.w3.org/TR/turtle/) format.
//
// Output in [N-Triples](https://www.w3.org/TR/n-triples/) format.
//
// Output in [N-Quads](https://www.w3.org/TR/n-quads/) format.
//
// The format of the data.
//
// Represents the format in which the data are serialized, for example when returned by a
// query or when inserted in the store.
type RDFXML string

const (
	NQuads       RDFXML = "n_quads"
	NTriples     RDFXML = "n_triples"
	RDFXMLRDFXML RDFXML = "rdf_xml"
	Turtle       RDFXML = "turtle"
)

// Returns information about the triple store.
type Store string

const (
	StoreStore Store = "store"
)

type Type string

const (
	BlankNode Type = "blank_node"
	Literal   Type = "literal"
	URI       Type = "uri"
)

// Query messages
type QueryMsg struct {
	Enum   *Store
	Select *Select
}

func (x *QueryMsg) UnmarshalJSON(data []byte) error {
	x.Select = nil
	x.Enum = nil
	var c Select
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.Select = &c
	}
	return nil
}

func (x *QueryMsg) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.Select != nil, x.Select, false, nil, x.Enum != nil, x.Enum, false)
}

type ValueUnion struct {
	String   *string
	ValueIRI *ValueIRI
}

func (x *ValueUnion) UnmarshalJSON(data []byte) error {
	x.ValueIRI = nil
	var c ValueIRI
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ValueIRI = &c
	}
	return nil
}

func (x *ValueUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.ValueIRI != nil, x.ValueIRI, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
			*pi = nil
	}
	if pf != nil {
			*pf = nil
	}
	if pb != nil {
			*pb = nil
	}
	if ps != nil {
			*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
			return false, err
	}

	switch v := tok.(type) {
	case json.Number:
			if pi != nil {
					i, err := v.Int64()
					if err == nil {
							*pi = &i
							return false, nil
					}
			}
			if pf != nil {
					f, err := v.Float64()
					if err == nil {
							*pf = &f
							return false, nil
					}
					return false, errors.New("Unparsable number")
			}
			return false, errors.New("Union does not contain number")
	case float64:
			return false, errors.New("Decoder should not return float64")
	case bool:
			if pb != nil {
					*pb = &v
					return false, nil
			}
			return false, errors.New("Union does not contain bool")
	case string:
			if haveEnum {
					return false, json.Unmarshal(data, pe)
			}
			if ps != nil {
					*ps = &v
					return false, nil
			}
			return false, errors.New("Union does not contain string")
	case nil:
			if nullable {
					return false, nil
			}
			return false, errors.New("Union does not contain null")
	case json.Delim:
			if v == '{' {
					if haveObject {
							return true, json.Unmarshal(data, pc)
					}
					if haveMap {
							return false, json.Unmarshal(data, pm)
					}
					return false, errors.New("Union does not contain object")
			}
			if v == '[' {
					if haveArray {
							return false, json.Unmarshal(data, pa)
					}
					return false, errors.New("Union does not contain array")
			}
			return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
			return json.Marshal(*pi)
	}
	if pf != nil {
			return json.Marshal(*pf)
	}
	if pb != nil {
			return json.Marshal(*pb)
	}
	if ps != nil {
			return json.Marshal(*ps)
	}
	if haveArray {
			return json.Marshal(pa)
	}
	if haveObject {
			return json.Marshal(pc)
	}
	if haveMap {
			return json.Marshal(pm)
	}
	if haveEnum {
			return json.Marshal(pe)
	}
	if nullable {
			return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
