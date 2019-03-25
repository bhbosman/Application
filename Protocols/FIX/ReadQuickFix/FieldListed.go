package ReadQuickFix

import "encoding/xml"

type FieldListed struct {
	Name     string
	Required bool
}

func (self *FieldListed) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "name":
			self.Name = attr.Value
			break

		case "required":
			if attr.Value == "Y" {
				self.Required = true
			}
		}
	}
	return d.Skip()
}
