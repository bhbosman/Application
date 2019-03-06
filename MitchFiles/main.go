package MitchFiles

//go:generate goidlgenerator -out outData03.json  Mitch.idl

//go:generate goidlgenerator -out ./GeneratedFiles/MitchIdl.go -outType go  -packageName GeneratedFiles Mitch.idl
