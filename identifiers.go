package easyparser

// Report ...
type Report struct {
	Init      bool // Set to true if `init()` was detected
	Constants map[string]Constant
	Variables map[string]Variable
	Functions map[string]Function
	Structs   map[string]Struct
	Interface map[string]Interface
}

// Constant ...
type Constant struct {
	Name  string
	Type  string
	Value string
}

// Variable ...
type Variable struct {
	Name string
	Type string
}

// Function ...
type Function struct {
	Name string
	In   []Variable
	Out  []Variable
}

// Parameter ...
type Parameter struct {
	Name string
	Type string
}

// Struct ...
type Struct struct {
	Name   string
	Fields []Field
}

// Field ...
type Field struct {
	Name string
	Type string
	Tags map[string]string
}

// Interface ...
type Interface struct {
	Name    string
	Methods []Function
}
