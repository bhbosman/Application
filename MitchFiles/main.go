package MitchFiles

//go:generate mitchgenerator   -out ./GeneratedFiles/MitchIdl.go -outType go -typesToUse Mitch   -packageName GeneratedFiles Mitch.idl
//go:generate gofmt -w  ./GeneratedFiles/MitchIdl.go

