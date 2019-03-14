package MitchFiles

////go:generate goidlgenerator -out outData03.json -typesToUse Mitch Mitch.idl

//go:generate goidlgenerator   -out ./GeneratedFiles/MitchIdl.go -outType go -typesToUse Mitch   -packageName GeneratedFiles Mitch.idl
//go:generate gofmt -w  ./GeneratedFiles/MitchIdl.go

////go:generate goidlgenerator   -out ./GeneratedFiles/MitchIdlDefault.go -outType go -typesToUse Mitch   -packageName GeneratedFiles -defaultTypes true   Mitch.idl
////go:generate gofmt -w  ./GeneratedFiles/MitchIdlDefault.go
