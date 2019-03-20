package ReadQuickFix

import (
	"encoding/xml"
)

type FixDecl struct {
	XMLName     xml.Name `xml:"fix"`
	DeclType    string   `xml:"type,attr"`
	Major       int      `xml:"major,attr"`
	Minor       int      `xml:"minor,attr"`
	ServicePack int      `xml:"servicepack,attr"`
	Trailer    Trailer           `xml:"trailer"`
	Header     Header            `xml:"header"`
	Components Components        `xml:"components"`
	Fields     FieldDeclarations `xml:"fields"`
}

type FieldDeclaration struct {
}

func (self *FieldDeclaration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, e := d.Token()
		if e != nil {
			break
		}
		switch t.(type) {
		case xml.StartElement:
			break
		}
	}
	return nil}

type FieldDeclarations struct {
	Fields []interface{}
}

func (self *FieldDeclarations) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, e := d.Token()
		if e != nil {
			break
		}
		switch val := t.(type) {
		case xml.StartElement:
			switch val.Name.Local {
			case "field":
				field := &FieldDeclaration{}
				self.Fields = append(self.Fields, field)
				err := d.DecodeElement(field, &val)
				if err != nil {
					return err
				}
				break
			default:
				panic("element not handled")
			}
			break
		}
	}
	return nil
}
