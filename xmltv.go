// Package xmltv provides structures for parsing XMLTV data.
package xmltv

import (
	"encoding/xml"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	t1, err := time.Parse("20060102150405 -0700", attr.Value)
	if err != nil {
		return err
	}

	*t = Time{t1}
	return nil
}

// check interface
var _ xml.UnmarshalerAttr = &Time{}

type Tv struct {
	Channels   []Channel   `xml:"channel"`
	Programmes []Programme `xml:"programme"`
}

type Channel struct {
	Id           string   `xml:"id,attr"`
	DisplayNames []string `xml:"display-name"`
	Icon         Icon     `xml:"icon"`
}

type Icon struct {
	Src string `xml:"src,attr"`
}

type Programme struct {
	ChannelId  string   `xml:"channel,attr"`
	Start      Time     `xml:"start,attr"`
	Stop       Time     `xml:"stop,attr"`
	Title      string   `xml:"title"`
	Desc       string   `xml:"desc"`
	Categories []string `xml:"category"`
}
