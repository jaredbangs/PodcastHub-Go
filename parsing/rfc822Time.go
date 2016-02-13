package parsing

import (
	"encoding/xml"
	"time"
)

type Rfc822Time struct {
	time.Time
}

func (r *Rfc822Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(time.RFC822, v)
	if err != nil {
		return err
	}
	*r = Rfc822Time{parse}
	return nil
}
