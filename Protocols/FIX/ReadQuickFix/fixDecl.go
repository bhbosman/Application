package ReadQuickFix

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
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

func (self FixDecl) EnumsToIds(writer io.Writer) {
	sb := bufio.NewWriter(writer)
	for _, field := range self.Fields.Fields {
		if len(field.Values) > 0 {
			_, _ = sb.WriteString(fmt.Sprintf("enum %v // %v\n", field.Name, field.Type))
			_, _ = sb.WriteString(fmt.Sprintf("{\n"))
			_, _ = sb.WriteString(fmt.Sprintf("\n};\n"))
		}
	}
	sb.Flush()
}
