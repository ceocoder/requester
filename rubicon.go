package main

var rubiconBids = []string{`{
    "id": "{{ .GetRandomId }}",
    "at": 2,
    "tmax": 129,
    "imp": [{
        "id": "1",
        "tagid": "{{ .GetRandomId }}",
        "banner": {
            "w": 728,
            "h": 90,
			"battr": [9, 1, 2],
            "api": []
        },
        "iframebuster": []
    }],
    "site": {
        "id": "15756",
        "domain": "http://zoopla.co.uk",
        "name": "Zoopla",
        "cat": ["IAB21"],
        "page": "http://www.zoopla.co.uk/for-sale/details/31611794?search_identifier=e12d33a11efeb568768a845e59d51dfc",
        "publisher": {
            "id": "9208"
        }
    },
    "device": {
        "ua": "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; InfoPath.3; .NET4.0C; .NET4.0E; MS-RTC LM 8)",
        "ip": "192.168.1.1",
        "language": "en",
        "devicetype": 2,
        "js": 1,
        "geo": {
            "country": "GBR",
            "region": "HNS"
        }
    },
    "user": {
        "id": "64c017a3d756e2566596cc1499294d1b602c88a3",
        "buyeruid": "7be2ed3d-a245-4045-af05-f15c9771a73e",
        "ext": {
            "sessiondepth": 5
        }
    }
}`}
