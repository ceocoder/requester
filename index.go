package main

var indexBids = []string{`{
    "device": {
        "ip": "123.231.123.231",
        "language": "PT",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.4.2; pt-pt; GT-P5210 Build/KOT49H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30"
    },
    "ext": {
        "edepth": 1,
        "sdepth": 1,
        "ssl": 0
    },
    "id": "{{ .GetRandomId }}",
    "imp": [
        {
            "banner": {
                "h": {{ .Dim.Height }},
                "pos": 3,
                "topframe": 0,
                "w": {{ .Dim.Width }}
            },
            "id": "1",
            "secure": 0
        }
    ],
    "site": {
        "cat": [
            "IAB14"
        ],
        "content": {
            "keywords": "1,6268"
        },
        "id": "144173",
        "page": "ebay.es"
    },
    "user": {
        "id": "{{ .UserId }}",
        "buyeruid": "{{ .UserId }}"
    }
}
`, `{
    "bcat": [
        "IAB7-41",
        "IAB9-30",
        "IAB13",
        "IAB20-18",
        "IAB14",
        "IAB22-2",
        "IAB11-2",
        "IAB14-1",
        "IAB13-7",
        "IAB11-4",
        "IAB8-5",
        "IAB23",
        "IAB20",
        "IAB20-3",
        "IAB24",
        "IAB13-2"
    ],
    "device": {
        "ip": "123.231.123.231",
        "ua": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.111 Safari/537.36"
    },
    "ext": {
        "edepth": 2,
        "sdepth": 2,
        "ssl": 0,
        "tz": 180
    },
    "id": "{{ .GetRandomId }}",
    "imp": [
        {
            "banner": {
                "h": {{ .Dim.Height }},
                "pos": 3,
                "topframe": 0,
                "w": {{ .Dim.Width }}
            },
            "ext": {
                "deal": [
                    {
                        "floor": 500,
                        "id": "12420143",
                        "type": "second"
                    }
                ]
            },
            "id": "1",
            "pmp": {
                "deals": [
                    {
                        "at": 2,
                        "bidfloor": 500,
                        "id": "{{ .GetRandomId }}"
                    }
                ]
            },
            "secure": 0
        }
    ],
    "site": {
        "cat": [
            "IAB20"
        ],
        "content": {
            "keywords": "1,252"
        },
        "domain": "foobar.com",
        "id": "{{ .GetRandomId }}",
        "publisher": {
            "name": "FooBar"
        }
    },
    "user": {
        "id": "{{ .UserId }}",
        "buyeruid": "{{ .UserId }}"
    }
}`}
