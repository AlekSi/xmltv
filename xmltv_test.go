package xmltv

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func dummyReader(charset string, input io.Reader) (io.Reader, error) {
	return input, nil
}

func TestDecode(t *testing.T) {
	// Example downloaded from http://wiki.xmltv.org/index.php/XMLTVFormat
	// One may check it with `xmllint --noout --dtdvalid xmltv.dtd example.xml`
	f, err := os.Open("example.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var tv Tv
	dec := xml.NewDecoder(f)
	dec.CharsetReader = dummyReader
	err = dec.Decode(&tv)
	if err != nil {
		t.Fatal(err)
	}

	ch := Channel{
		Id:           "I10436.labs.zap2it.com",
		DisplayNames: []string{"13 KERA", "13 KERA TX42822:-", "13", "13 KERA fcc", "KERA", "KERA", "PBS Affiliate"},
		Icon:         Icon{Src: `file://C:\Perl\site/share/xmltv/icons/KERA.gif`},
	}
	if !reflect.DeepEqual(ch, tv.Channels[0]) {
		t.Errorf("\texpected: %#v\n\t\tactual:   %#v\n", ch, tv.Channels[0])
	}

	loc := time.FixedZone("", -6*60*60)
	pr := Programme{
		Id:         "someId",
		ChannelId:  "I10436.labs.zap2it.com",
		Start:      Time{time.Date(2008, 07, 15, 0, 30, 0, 0, loc)},
		Stop:       Time{time.Date(2008, 07, 15, 1, 0, 0, 0, loc)},
		Title:      "NOW on PBS",
		Desc:       "Jordan's Queen Rania has made job creation a priority to help curb the staggering unemployment rates among youths in the Middle East.",
		Categories: []string{"Newsmagazine", "Interview", "Public affairs", "Series"},
	}
	if !reflect.DeepEqual(pr, tv.Programmes[0]) {
		expected := fmt.Sprintf("\texpected: %#v\n\t\t\texpected start: %s\n\t\t\texpected stop : %s", pr, pr.Start, pr.Stop)
		actual__ := fmt.Sprintf("\tactual:   %#v\n\t\t\tactual start:   %s\n\t\t\tactual stop:    %s", tv.Programmes[0], tv.Programmes[0].Start, tv.Programmes[0].Stop)
		t.Errorf("%s\n%s\n", expected, actual__)
	}
}
