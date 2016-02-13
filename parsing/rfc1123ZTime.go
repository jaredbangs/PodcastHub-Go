package parsing

import (
	"encoding/xml"
	"time"
)

type Rfc1123ZTime struct { // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	time.Time
}

func (r *Rfc1123ZTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(time.RFC1123Z, v)
	if err != nil {
		return err
	}
	*r = Rfc1123ZTime{parse}
	return nil
}
