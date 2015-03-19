package main

var bidswitchBids = []string{`{
    "id": "{{ .GetRandomId }}",
    "imp": [
        {
            "id": "1",
            "banner": {
                "w": {{ .Dim.Width }},
                "h": {{ .Dim.Height }},
                "pos": 1,
                "topframe": 0,
                "mimes": [
                    "text/html",
                    "application/x-shockwave-flash"
                ],
                "ext": {
                    "extra_sizes": [
                        {
                            "h": 50,
                            "w": 300
                        },
                        {
                            "h": 75,
                            "w": 450
                        }
                    ]
                }
            },
            "ext": {
                "rubicon": {
                    "site_size_session_count": 5
                },
                "google": {
                    "excluded_attribute": [1,2,3],
                    "allowed_vendor_type": [10,11,12]
                },
                "yieldone": {
                    "allowed_creative_types": ["a", "b", "c"],
                    "allowed_creative_category_id": [1,2,3,4]
                },
                "viewability": 40,
                "inventory_class": 12345
            }
        }
    ],
    "site": {
        "id": "234563",
        "domain": "siteabcd.com",
        "page": "http://siteabcd.com/page.htm",
        "publisher": {
            "id": "rubicon_25"
        },
        "ext": {
            "mobile_site": 1
        }
    },
    "device": {
        "ip": "123.123.123.123",
        "geo": {
            "lat": 11.111111,
            "lon": -11.111111,
            "country": "US",
            "region": "CA",
            "city": "San Francisco",
            "zip": "11111"
        },
        "ua": "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
        "language": "en"
    },
    "user": {
        "id": "{{ .UserId }}",
        "buyeruid": "{{ .UserId }}"
    },
    "cur": [
        "USD"
    ],
    "at": 2,
    "ext": {
        "ssp": "rubicon",
        "is_secure": 1,
        "adtruth": {
            "tdl_millis": 19534993,
            "WEB_APP_BRIDGE_3_2": "FD5314A665D",
            "WEB_APP_BRIDGE_4_0": "4FD87B97753"
        },
        "ghostery": {
            "status": "verified",
            "domain": "siteabcd.com"
        },
        "google": {
            "detected_vertical": [{"id": 100, "weight": 2.1234}, {"id": 200, "weight":9.321}]
        },
        "creative_params": [
            {
                "type": "blacklist",
                "name": "ad_type",
                "value": [
                    "14011",
                    "14006"
                ]
            },
            {
                "type": "whitelist",
                "name": "ad_type",
                "value": [
                    "12345",
                    "67890"
                ]
            }
        ]
    }
}`, `{
    "id": "1234534625253",
    "imp": [
        {
            "id": "1",
            "video": {
                "mimes": [
                    "video/x-flv",
                    "video/mp4",
                    "application/x-shockwave-flash",
                    "application/javascript"
                ],
                "linearity": 1,
                "minduration": 5,
                "maxduration": 30,
                "protocols": [
                    2,
                    3,
                    5,
                    6
                ],
                "w": 640,
                "h": 480,
                "startdelay": 0,
                "battr": [
                    13,
                    14
                ],
                "minbitrate": 300,
                "maxbitrate": 1500,
                "api": [
                    1
                ],
                "companionad": [
                    {
                        "w": 300,
                        "h": 250,
                        "id": 1
                    }
                ],
                "companiontype": [
                    1,
                    2,
                    3
                ],
                "ext": {
                    "skippable": 1
                }
            }
        }
    ],
    "site": {
        "id": "234563",
        "domain": "siteabcd.com",
        "page": "http://siteabcd.com/page.htm",
        "publisher": {
            "id": "google_25"
        }
    },
    "device": {
        "ip": "123.123.123.123",
        "geo": {
            "country": "US",
            "region": "CA",
            "city": "San Francisco",
            "zip": "11111"
        },
        "ua": "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
        "language": "en"
    },
    "user": {
        "id": "{{ .UserId }}",
        "buyeruid": "{{ .UserId }}"
    },
    "cur": [
        "USD"
    ],
    "ext": {
        "ssp": "google"
    }
}`}
