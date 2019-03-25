package ReadQuickFix

import (
	"encoding/xml"
)

type FixDecl struct {
	XMLName     xml.Name          `xml:"fix"`
	DeclType    string            `xml:"type,attr"`
	Major       int               `xml:"major,attr"`
	Minor       int               `xml:"minor,attr"`
	ServicePack int               `xml:"servicepack,attr"`
	Trailer     Trailer           `xml:"trailer"`
	Header      Header            `xml:"header"`
	Components  Components        `xml:"components"`
	Fields      FieldDeclarations `xml:"fields"`
}
