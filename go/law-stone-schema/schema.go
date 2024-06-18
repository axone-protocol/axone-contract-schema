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
//    askResponse, err := UnmarshalAskResponse(bytes)
//    bytes, err = askResponse.Marshal()
//
//    programResponse, err := UnmarshalProgramResponse(bytes)
//    bytes, err = programResponse.Marshal()
//
//    binary, err := UnmarshalBinary(bytes)
//    bytes, err = binary.Marshal()

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

func UnmarshalAskResponse(data []byte) (AskResponse, error) {
	var r AskResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AskResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalProgramResponse(data []byte) (ProgramResponse, error) {
	var r ProgramResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProgramResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Binary string

func UnmarshalBinary(data []byte) (Binary, error) {
	var r Binary
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Binary) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

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

func (x *QueryMsg) UnmarshalJSON(data []byte) error {
	x.Ask = nil
	x.Enum = nil
	var c Ask
	object, err := unmarshalUnion(data, nil, nil, nil, nil, false, nil, true, &c, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
		x.Ask = &c
	}
	return nil
}

func (x *QueryMsg) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, false, nil, x.Ask != nil, x.Ask, false, nil, x.Enum != nil, x.Enum, false)
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
