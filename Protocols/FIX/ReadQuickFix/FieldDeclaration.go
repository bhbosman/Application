package ReadQuickFix

import "encoding/xml"

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
	return nil
}
