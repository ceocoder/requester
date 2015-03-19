package main

import (
	"bytes"
	"fmt"
	"html"
	"strings"
	"text/template"

	b64 "encoding/base64"

	"github.com/bsm/openrtb"
)

const (
	//Auction Macros
	AUCTION_ID       = "${AUCTION_ID}"
	AUCTION_IMP_ID   = "${AUCTION_IMP_ID}"
	AUCTION_SEAT_ID  = "${AUCTION_SEAT_ID}"
	AUCTION_AD_ID    = "${AUCTION_AD_ID}"
	AUCTION_PRICE    = "${AUCTION_PRICE}"
	AUCTION_CURRENCY = "${IMAGE}"
	// Tempate for One AdUnit
	AD_TEMPLATE = `
<li>
<h3>Bid Response</h3>
<h2>{{ .Bid.Id }}</h2>
<h3>{{ .Bid.Price }}</h3>
<p>{{ .EscapedAdm }}</p>
<h3>Rendered Snippet</h3>
<iframe src="data:text/html;base64,
{{ .Base64Snip }}" scrolling=no marginwidth=0 marginheight=0></iframe>
</li>`
)

type Bid struct {
	Bid openrtb.Bid
}

func (bid Bid) Base64Snip() string {
	return b64.StdEncoding.EncodeToString([]byte(*bid.Bid.Adm))
}

func (bid Bid) EscapedAdm() string {
	return html.EscapeString(*bid.Bid.Adm)
}

var (
	responseTpl, _ = template.New("Ad").Parse(AD_TEMPLATE)
)

//BidSummary can be used as shared bid summary across routines OR create one per routine and
//merge stats after all routines have returned
type BidSummary struct {
	BIDS, NOBIDS, ERROR, UNKNOWN int
	Html                         string
}

// take a valid response and turn it into HTML to be rendered
func generate(resp *openrtb.Response) string {
	var buffer bytes.Buffer
	for _, seatbid := range resp.Seatbid {
		for _, bid := range seatbid.Bid {
			replaceMacros(bid.Adm)
			responseTpl.Execute(&buffer, Bid{bid}) //TODO: replaceMacros(bid.Adm))
		}
	}
	return buffer.String()
}

func replaceMacros(adm *string) {
	//TODO: write a real func
	*adm = strings.Replace(*adm, AUCTION_ID, randSeq(10), -1)
	*adm = strings.Replace(*adm, AUCTION_IMP_ID, randSeq(10), -1)
	*adm = strings.Replace(*adm, AUCTION_SEAT_ID, randSeq(10), -1)
	*adm = strings.Replace(*adm, AUCTION_AD_ID, randSeq(10), -1)
	*adm = strings.Replace(*adm, AUCTION_CURRENCY, randSeq(7), -1)
	*adm = strings.Replace(*adm, AUCTION_PRICE, randPrice(), -1)
}

func renderFinalHtml(body string) string {

	return fmt.Sprintf(`<html>
<head>
</head>
<body>
<ul>
%s
</ul>
</body>
</html>`, body)
}
