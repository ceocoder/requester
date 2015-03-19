package main

import (
	"bytes"
	b64 "encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

//variables for cli flags
var (
	bidderUrl            = flag.String("url", "http://localhost:8080/bid/openrtb", "Bidder Endpoint")
	sampleEncryptedPrice = flag.String("sampleEncryptedPrice", "foo", "Sample encrypted price")
	buyerUserIdsFile     = flag.String("buyerUserIdsFile", "none", "Buyer's userIds")
	encodeUserIds        = flag.Bool("encodeUserIds", false, "Encode useIds using base64 encoding")
	exchangeName         = flag.String("exchangeName", "index", "Exchange type to use - possible values rubicon, index")
	maxQps               = flag.Int("maxQps", 100, "Maximum queries per second")
	seconds              = flag.Int("seconds", 0, "Number of seconds to send reqeusts")
	requests             = flag.Int("requests", 0, "Number of requests to send to bidder")
	numThreads           = flag.Int("numThreads", 20, "Number of concurrent requests")
	threadInterval       = flag.Float64("threadInterval", 0.2, "Interval between concurrent rampup")
	standAlone           = flag.Bool("standAlone", false, "Run a standalone dummy bidder that responds with NOBID")
)

//computed variables
var (
	USER_IDS = []string{}
)

func parseAndValidate() {
	flag.Parse()
	readUsersList()

	//ensure that only seconds OR reqests are specified
	if (*seconds == 0 && *requests == 0) || (*seconds > 0 && *requests > 0) {
		fmt.Errorf("Please specify either number of seconds OR number of requests")
		flag.PrintDefaults()
		os.Exit(1)
	}

}

func readUsersList() {
	if *buyerUserIdsFile != "none" {
		users, err := ioutil.ReadFile(*buyerUserIdsFile)
		if err != nil {
			log.Fatalf("Unable to open userlist file %v", *buyerUserIdsFile)
		}
		for _, user := range strings.Split(string(users), "\n") {
			if *encodeUserIds {
				USER_IDS = append(USER_IDS, b64.URLEncoding.EncodeToString([]byte(user)))
			} else {
				USER_IDS = append(USER_IDS, user)
			}
		}
	}
}

func main() {
	parseAndValidate()

	bidSummaries := make([]BidSummary, 10)

	//seed random
	rand.Seed(time.Now().UTC().UnixNano())

	//temporary bidder standin
	if *standAlone {
		go bidderServer()
	}

	var wg sync.WaitGroup
	for i := 0; i < *numThreads; i++ {
		wg.Add(1)
		go exchange(*requests / *numThreads, &bidSummaries[i], &wg)
		//ramp up requests
		time.Sleep(time.Duration(*threadInterval) * time.Second)
	}

	// wait for all bids to finish
	wg.Wait()

	bidSummary := BidSummary{}
	var html bytes.Buffer
	for _, bs := range bidSummaries {
		bidSummary.BIDS += bs.BIDS
		bidSummary.NOBIDS += bs.NOBIDS
		bidSummary.ERROR += bs.ERROR
		html.WriteString(bs.Html)
	}
	//merge all bid summaries into one.

	fmt.Println("============ BID SUMMARY ============")
	fmt.Printf("NOBIDS: %v\n", bidSummary.NOBIDS)
	fmt.Printf("BIDS: %v\n", bidSummary.BIDS)
	fmt.Printf("ERRORS: %v\n", bidSummary.ERROR)
	ioutil.WriteFile("summary.html", []byte(renderFinalHtml(html.String())), 0644)
}
