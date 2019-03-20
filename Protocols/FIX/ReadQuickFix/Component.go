package ReadQuickFix

import "encoding/xml"

type Component struct {
	Fields []interface{}
}


func (self *Component) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "group":
				group := &Group{}
				self.Fields = append(self.Fields, group)
				err := d.DecodeElement(group, &val)
				if err != nil {
					return err
				}
				break
			case "component":
				component := &Component{}
				self.Fields = append(self.Fields, component)
				err := d.DecodeElement(component, &val)
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

