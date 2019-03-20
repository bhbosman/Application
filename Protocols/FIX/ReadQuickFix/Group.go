package ReadQuickFix

import "encoding/xml"

type Group struct {
	Fields []interface{}
}

func (self *Group) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, e := d.Token()
		if e != nil {
			break
		}
		switch val := t.(type) {
		case xml.StartElement:
			switch val.Name.Local {
			case "field":
				field := &FieldListed{}
				self.Fields = append(self.Fields, field)
				err := d.DecodeElement(field, &val)
				if err != nil {
					return err
				}

				break
			case "component":
				component := &Component{}
				err := d.DecodeElement(component, &val)
				self.Fields = append(self.Fields, component)
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

