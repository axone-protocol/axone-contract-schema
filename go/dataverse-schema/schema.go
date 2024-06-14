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
//    dataverseResponse, err := UnmarshalDataverseResponse(bytes)
//    bytes, err = dataverseResponse.Marshal()

package schema

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

func UnmarshalDataverseResponse(data []byte) (DataverseResponse, error) {
	var r DataverseResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DataverseResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// `ExecuteMsg` defines the set of possible actions that can be performed on the dataverse.
//
// This enum provides variants for registering services, datasets, and other operations
// related to the dataverse.
//
// Submits new claims about a resource to the dataverse.
//
// The SubmitClaims message is a pivotal component in the dataverse, enabling entities to
// contribute new claims about various resources. A claim represents a statement made by an
// entity, referred to as the issuer, which could be a person, organization, or service.
// These claims pertain to a diverse range of resources, including digital resources,
// services, zones, or individuals, and are asserted as factual by the issuer.
//
// #### Format
//
// Claims are injected into the dataverse through Verifiable Presentations (VPs). These
// presentations effectively amalgamate and showcase multiple credentials, thus providing a
// cohesive and comprehensive view of the assertions being made.
//
// While the data in a VP typically revolves around a common subject, it accommodates an
// unlimited number of subjects and issuers. This flexibility allows for a broad spectrum of
// claims to be represented.
//
// Primarily, the claims leverage the OKP4 ontology, which facilitates articulating
// assertions about widely acknowledged resources in the dataverse, including digital
// services, digital resources, zones, governance, and more.
//
// Additionally, other schemas may also be employed to supplement and enhance the validated
// knowledge contributed to these resources.
//
// #### Preconditions
//
// To maintain integrity and coherence in the dataverse, several preconditions are set for
// the submission of claims:
//
// 1. **Format Requirement**: Claims must be encapsulated within Verifiable Presentations
// (VPs).
//
// 2. **Unique Identifier Mandate**: Each Verifiable Credential within the dataverse must
// possess a unique identifier.
//
// 3. **Issuer Signature**: Claims must bear the issuer's signature. This signature must be
// verifiable, ensuring authenticity and credibility.
//
// Revoke or withdraw a previously submitted claims.
//
// #### Preconditions:
//
// 1. **Identifier Existance**: The identifier of the claims must exist in the dataverse.
type ExecuteMsg struct {
	SubmitClaims *SubmitClaims `json:"submit_claims,omitempty"`
	RevokeClaims *RevokeClaims `json:"revoke_claims,omitempty"`
}

type RevokeClaims struct {
	// The unique identifier of the claims to be revoked.       
	Identifier                                           string `json:"identifier"`
}

type SubmitClaims struct {
	// RDF format in which the metadata is represented. If not provided, the default format is        
	// [Turtle](https://www.w3.org/TR/turtle/) format.                                                
	Format                                                                                    *RDFXML `json:"format"`
	// The serialized metadata intended for attachment. This metadata should adhere to the            
	// format specified in the `format` field.                                                        
	Metadata                                                                                  string  `json:"metadata"`
}

// `InstantiateMsg` is used to initialize a new instance of the dataverse.
type InstantiateMsg struct {
	// A unique name to identify the dataverse instance.                        
	Name                                                      string            `json:"name"`
	// The configuration used to instantiate the triple store.                  
	TriplestoreConfig                                         TripleStoreConfig `json:"triplestore_config"`
}

// The configuration used to instantiate the triple store.
//
// `TripleStoreConfig` represents the configuration related to the management of the triple
// store.
type TripleStoreConfig struct {
	// The code id that will be used to instantiate the triple store contract in which to store                       
	// dataverse semantic data. It must implement the cognitarium interface.                                          
	CodeID                                                                                     string                 `json:"code_id"`
	// Limitations regarding triple store usage.                                                                      
	Limits                                                                                     TripleStoreLimitsInput `json:"limits"`
}

// Limitations regarding triple store usage.
//
// Contains requested limitations regarding store usages.
type TripleStoreLimitsInput struct {
	// The maximum number of bytes the store can contain. The size of a triple is counted as the        
	// sum of the size of its subject, predicate and object, including the size of data types           
	// and language tags if any. Default to [Uint128::MAX] if not set, which can be considered          
	// as no limit.                                                                                     
	MaxByteSize                                                                                 *string `json:"max_byte_size"`
	// The maximum number of bytes an insert data query can contain. Default to [Uint128::MAX]          
	// if not set, which can be considered as no limit.                                                 
	MaxInsertDataByteSize                                                                       *string `json:"max_insert_data_byte_size"`
	// The maximum number of triples an insert data query can contain (after parsing). Default          
	// to [Uint128::MAX] if not set, which can be considered as no limit.                               
	MaxInsertDataTripleCount                                                                    *string `json:"max_insert_data_triple_count"`
	// The maximum limit of a query, i.e. the maximum number of triples returned by a select            
	// query. Default to 30 if not set.                                                                 
	MaxQueryLimit                                                                               *int64  `json:"max_query_limit"`
	// The maximum number of variables a query can select. Default to 30 if not set.                    
	MaxQueryVariableCount                                                                       *int64  `json:"max_query_variable_count"`
	// The maximum number of bytes the store can contain for a single triple. The size of a             
	// triple is counted as the sum of the size of its subject, predicate and object, including         
	// the size of data types and language tags if any. The limit is used to prevent storing            
	// very large triples, especially literals. Default to [Uint128::MAX] if not set, which can         
	// be considered as no limit.                                                                       
	MaxTripleByteSize                                                                           *string `json:"max_triple_byte_size"`
	// The maximum number of triples the store can contain. Default to [Uint128::MAX] if not            
	// set, which can be considered as no limit.                                                        
	MaxTripleCount                                                                              *string `json:"max_triple_count"`
}

// `QueryMsg` defines the set of possible queries that can be made to retrieve information
// about the dataverse.
//
// This enum provides variants for querying the dataverse's details and other related
// information.
//
// Retrieves information about the current dataverse instance.
type QueryMsg struct {
	Dataverse Dataverse `json:"dataverse"`
}

type Dataverse struct {
}

// DataverseResponse is the response of the Dataverse query.
type DataverseResponse struct {
	// The name of the dataverse.       
	Name                         string `json:"name"`
}

// RDF/XML Format
//
// RDF/XML is a syntax to express RDF information in XML. See the [official RDF/XML
// specification](https://www.w3.org/TR/rdf-syntax-grammar/).
//
// Turtle (Terse RDF Triple Language) Format
//
// Turtle is a textual format for representing RDF triples in a more compact and
// human-readable way compared to RDF/XML. See the [official Turtle
// specification](https://www.w3.org/TR/turtle/).
//
// N-Triples Format
//
// N-Triples is a line-based, plain text format for encoding an RDF graph. Each line
// corresponds to a single RDF triple. See the [official N-Triples
// specification](https://www.w3.org/TR/n-triples/).
//
// N-Quads Format
//
// N-Quads is an extension of N-Triples to support RDF datasets by adding an optional fourth
// element to represent the graph name. See the [official N-Quads
// specification](https://www.w3.org/TR/n-quads/).
type RDFXML string

const (
	NQuads       RDFXML = "n_quads"
	NTriples     RDFXML = "n_triples"
	RDFXMLRDFXML RDFXML = "rdf_xml"
	Turtle       RDFXML = "turtle"
)
