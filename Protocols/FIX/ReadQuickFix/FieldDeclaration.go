package ReadQuickFix

import (
	"encoding/xml"
)

type FieldValue struct {
	EnumValue       string
	EnumDescription string
}

func (self *FieldValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "enum":
			self.EnumValue = attr.Value
			break
		case "description":
			self.EnumDescription = attr.Value
		}
	}
	return d.Skip()
}

type FieldDeclaration struct {
	Number string
	Name   string
	Type   string
	Values []*FieldValue
}

func (self *FieldDeclaration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "number":
			self.Number = attr.Value
			break
		case "name":
			self.Name = attr.Value
			break
		case "type":
			self.Type = attr.Value
			break
		}
	}

	for {
		t, e := d.Token()
		if e != nil {
			break
		}
		switch val := t.(type) {
		case xml.StartElement:
			switch val.Name.Local {
			case "value":
				field := &FieldValue{}
				self.Values = append(self.Values, field)
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
