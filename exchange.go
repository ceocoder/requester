package main

import (
	"bytes"
	"fmt"
	"github.com/bsm/openrtb"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"text/template"
)

type Dimension struct {
	Height int `json:"w"`
	Width  int `json:"h"`
}

type Segment struct {
	Id    string
	Name  string
	Value string
}

type PublisherData struct {
	Url                           string
	SellerId, PublisherSettingsId int
	Publisher                     string
}

type BidRequest struct {
	Id        string
	Domain    string
	BuyerId   string
	UserId    string
	UserAgent string
	Dim       Dimension
	PubData   PublisherData
	Segments  []Segment
}

func (bidRequest BidRequest) GetRandomId() string {
	return randSeq(20)
}

var SEGMENTS = []Segment{
	{"121213", "provider1-age", "30-40"},
	{"121213", "provider2", "30-40"},
	{"121213", "provider3-age", "30-40"},
}

var DIMENSIONS = []Dimension{
	{468, 60},
	{120, 600},
	{728, 90},
	{300, 250},
	{250, 250},
	{336, 280},
	{120, 240},
	{125, 125},
	{160, 600},
	{180, 150},
	{110, 32},
	{120, 60},
	{180, 60},
	{420, 600},
	{420, 200},
	{234, 60},
	{200, 200},
}

var PUBLISHER_DATA = []PublisherData{
	{"http://www.youtube.com", 502, 32423234, "Youtube"},
	{"http://www.youtube.com/shows", 502, 32423234, "Youtube"},
	{"http://news.google.com", 10001, 56751341, "Google News"},
	{"http://news.google.com/news?pz=1&ned=us&topic=b&ict=ln", 10001, 12672383,
		"Google News"},
	{"http://www.google.com/finance?hl=en&ned=us&tab=ne", 1528, 84485234,
		"Google Finance"},
	{"http://www.nytimes.com/pages/technology/index.html", 936, 9034124,
		"New York Times"},
	{"http://some.gcn.site.com", 10002, 12002392, "GCN"},
}

var USER_AGENTS = []string{
	"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.2) ",
	"Gecko/2008092313 Ubuntu/8.04 (hardy) Firefox/3.1",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) ",
	"Gecko/20070118 Firefox/2.0.0.2pre",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1.7pre) Gecko/20070815 ",
	"Firefox/2.0.0.6 Navigator/9.0b3",
	"Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_4_11; en) AppleWebKit/528.5+",
	" (KHTML, like Gecko) Version/4.0 Safari/528.1",
	"Mozilla/5.0 (Macintosh; U; PPC Mac OS X; sv-se) AppleWebKit/419 ",
	"(KHTML, like Gecko) Safari/419.3",
	"Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 6.0; en-US)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0;)",
	"Mozilla/4.08 (compatible; MSIE 6.0; Windows NT 5.1)",
}

func exchange(requests int, bidSummary *BidSummary, wg *sync.WaitGroup) {

	//ensure that we close the wait group
	defer wg.Done()

	bidTpl := template.New("OpenRTB")
	switch *exchangeName {
	case "index":
		bidTpl.Parse(indexBids[rand.Intn(len(indexBids))])
		break
	case "rubicon":
		bidTpl.Parse(rubiconBids[rand.Intn(len(rubiconBids))])
		break
	case "bidswitch":
		bidTpl.Parse(bidswitchBids[rand.Intn(len(bidswitchBids))])
	}

	for i := 0; i < requests; i++ {

		pubData := PUBLISHER_DATA[rand.Intn(len(PUBLISHER_DATA))]
		u, _ := url.Parse(pubData.Url)

		bid := BidRequest{
			Id:        randSeq(10),
			Domain:    u.Host,
			BuyerId:   randSeq(10),
			UserId:    USER_IDS[rand.Intn(len(USER_IDS))],
			UserAgent: USER_AGENTS[rand.Intn(len(USER_AGENTS))],
			Dim:       DIMENSIONS[rand.Intn(len(DIMENSIONS))],
			PubData:   pubData,
			Segments:  SEGMENTS,
		}

		var b bytes.Buffer
		err := bidTpl.ExecuteTemplate(&b, "OpenRTB", bid)
		if err != nil {
			fmt.Printf("fError parsing template %v", err)
		}

		resp, err := http.Post(*bidderUrl,
			"application/json",
			bytes.NewBuffer(b.Bytes()))

		if err != nil {
			bidSummary.ERROR += 1
		} else if resp.StatusCode == 200 {
			buff := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
			buff.ReadFrom(resp.Body)
			orsp, err := openrtb.ParseResponseBytes(buff.Bytes())
			fmt.Println(string(buff.Bytes()))
			if err == nil {
				bidSummary.Html += generate(orsp)
				bidSummary.BIDS += 1
			}
		} else if resp.StatusCode == 204 {
			bidSummary.NOBIDS += 1
		} else {
			bidSummary.UNKNOWN += 1
		}
	}
}
