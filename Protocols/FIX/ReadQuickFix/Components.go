package ReadQuickFix

import "encoding/xml"

type Components struct {
	Fields []interface{}
}

func (self *Components) addField(field interface{}) {
	if self.Fields == nil {
		self.Fields = make([]interface{}, 0, 8)
	}
	self.Fields = append(self.Fields, field)
}

func (self *Components) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				self.addField(field)
				err := d.DecodeElement(field, &val)
				if err != nil {
					return err
				}
				break
			case "group":
				group := &Group{}
				self.addField(group)
				err := d.DecodeElement(group, &val)
				if err != nil {
					return err
				}
				break
			case "component":
				component := &Component{}
				self.addField(component)
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
