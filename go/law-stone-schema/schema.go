package schema

type Binary string

// Instantiate message
type InstantiateMsg struct {
	// The Prolog program carrying law rules and facts.                                 
	Program                                                                      string `json:"program"`
	// The `okp4-objectarium` contract address on which to store the law program.       
	StorageAddress                                                               string `json:"storage_address"`
}

// If not broken, ask the logic module the provided query with the law program loaded.
type Ask struct {
	Ask AskClass `json:"ask"`
}

type AskClass struct {
	Query string `json:"query"`
}

type AskResponse struct {
	Answer  *Answer `json:"answer"`
	GasUsed int64   `json:"gas_used"`
	Height  int64   `json:"height"`
}

type Answer struct {
	HasMore   bool     `json:"has_more"`
	Results   []Result `json:"results"`
	Success   bool     `json:"success"`
	Variables []string `json:"variables"`
}

type Result struct {
	Substitutions []Substitution `json:"substitutions"`
}

type Substitution struct {
	Term     Term   `json:"term"`
	Variable string `json:"variable"`
}

type Term struct {
	Arguments []Term `json:"arguments"`
	Name      string `json:"name"`
}

// ProgramResponse carry elements to locate the program in a `okp4-objectarium` contract.
type ProgramResponse struct {
	// The program object id in the `okp4-objectarium` contract.                         
	ObjectID                                                                      string `json:"object_id"`
	// The `okp4-objectarium` contract address on which the law program is stored.       
	StorageAddress                                                                string `json:"storage_address"`
}

// Execute messages
//
// Break the stone making this contract unusable, by clearing all the related resources: -
// Unpin all the pinned objects on `okp4-objectarium` contracts, if any. - Forget the main
// program (i.e. or at least unpin it). Only the contract admin is authorized to break it,
// if any. If already broken, this is a no-op.
type ExecuteMsg string

const (
	BreakStone ExecuteMsg = "break_stone"
)

// If not broken, returns the law program location information.
//
// ProgramCode returns the law program code.
type Program string

const (
	ProgramCode    Program = "program_code"
	ProgramProgram Program = "program"
)

// Query messages
type QueryMsg struct {
	Ask  *Ask
	Enum *Program
}
