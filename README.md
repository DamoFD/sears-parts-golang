# Unofficial Sears Parts Direct API

Base URL: `https://catalog.searspartsdirect.com/graphql`

## Get Parts By Appliance ID

### URL

`?operationName=modelGetParts&variables=%7B%22id%22%3A%22{YOUR_APPLIANCE_ID}%22%2C%22from%22%3A{PAGE}%2C%22size%22%3A{PAGE_SIZE}%2C%22filters%22%3A%5B%7B%22name%22%3A%22SELLABLE%22%2C%22type%22%3A%22MATCH%22%2C%22values%22%3A%22true%22%7D%2C%7B%22name%22%3A%22RESTRICTION%22%2C%22type%22%3A%22NOT%22%2C%22values%22%3A%5B%2231%22%2C%2249%22%2C%2260%22%2C%224%22%2C%225%22%2C%229%22%2C%2210%22%2C%2211%22%2C%2213%22%2C%2221%22%2C%2222%22%2C%2225%22%2C%2233%22%2C%2234%22%2C%2212%22%2C%2217%22%2C%226%22%2C%2216%22%2C%2226%22%2C%2252%22%2C%2259%22%5D%7D%5D%2C%22orders%22%3A%5B%7B%22name%22%3A%22SUBSCRIBABLEPARTSONLY%22%2C%22order%22%3A%22DESC%22%7D%2C%7B%22name%22%3A%22RANK%22%2C%22order%22%3A%22DESC%22%7D%5D%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%222aeb68e92f826c68b8de9c8eb92bd17363b8767757f6af74c026232db4565ecc%22%7D%7D`

### URL Broken Down into JSON

```
"url": {
    "operationName": "modelGetParts", (string)
    "variables": {
        "id": "{YOUR_APPLIANCE_ID}", (string)
        "orders": [
            {
                "name": "SELLABLE", (string)
                "order": "DESC" (string)
            },
            {
                "name": "RANK", (string)
                "order": "DESC" (string)
            },
            {
                "name": "AVAILABILITY", (string)
                "order": "DESC" (string)
            },
        ],
        "filters": [
            {
                "name": "RESTRICTION", (string)
                "type": "NOT", (string)
                "values": ["31", "49", "60", "12", "17"]
            }
        ],
        "from": {PAGE}, (number)
        "size": {PAGE_SIZE}, (number)
        parentFilter: [
            {
                "name": "ISMAINCATEGORY", (string)
                "values": ["true"] ([]string)
            }
        ]
    },
    "extensions": {
        "persistedQuery": {
            "version": 1, (number)
            "sha256Hash": "2aeb68e92f826c68b8de9c8eb92bd17363b8767757f6af74c026232db4565ecc" (string)
        }
    }
}
```

### Headers

"Accept": "*/*"

"Accept-Language": "en-US,en;q=0.9"

"Cache-Control": "no-cache"

"Pragma": "no-cache"

"Origin": "https://www.searspartsdirect.com"

"Referer": "https://www.searspartsdirect.com/"

"content-type": "application/json"

"X-Api-Key": "tV6bZZfhUh3MQmZggG6iq6LjfrZgQgcR26Yv86En"

### JSON Response

```
{
    "data": {
        "model": {
            "id": "APPLIANCE_ID", (string)
            "number": "APPLIANCE_MODEL_NUMBER", (string)
            "title": "APPLIANCE_TYPE", (string)
            "brand": {
                "id": "APPLIANCE_BRAND_ID", (string)
                "name": "APPLIANCE_BRAND_NAME", (string)
                "__typename": "OBJECT_TYPE" (string)
            },
            "taxonomies": {
                "taxonomies": [
                    {
                        "id": "TAXONOMY_ID", (string)
                        "name": "TAXONOMY_NAME", (string)
                        "__typename": "OBJECT_TYPE" (string)
                    },
                ],
                "__typename": "OBJECT_TYPE" (string)
            },
            "parts": {
                "totalCount": "TOTAL_PART_COUNT", (number)
                "parts": [
                    {
                        "id": "PART_ID", (string)
                        "legacyId": "PART_LEGACY_ID", (string)
                        "title": "PART_TITLE", (string)
                        "number": "PART_NUMBER", (string)
                        "topSoldPosition": "TOP_SOLD_POSITION", (number)
                        "division": {
                            "id": "DIVISION_ID", (string)
                            "description": "DIVISION_DESCRIPTION", (string)
                            "__typename": "OBJECT_TYPE" (string)
                        },
                        "sourceId": "PART_SOURCE_ID", (string)
                        "media": {
                            "image": {
                                "urls": [
                                    "IMAGE_URL" (string)
                                ],
                                "__typename": "OBJECT_TYPE" (string)
                            },
                            "__typename": "OBJECT_TYPE" (string)
                        },
                        "clearance": "ON_CLEARANCE?", (boolean)
                        "returnable": "PART_RETURNABLE?", (boolean)
                        "substitutedByList": {
                            "parts": [PART_ARRAY], (<part>array)
                            "__typename": "OBJECT_TYPE" (string)
                        },
                        "pricing": {
                            "strikeThroughPriceType": "STRIKE_THROUGH?", (string)
                            "sell": PRICE, (float)
                            "list": PRICE, (float)
                            "availabilityInfo": {
                                "status": "AVAILABILITY", (string)
                                "inventories": [INVENTORY_ARRAY], (<inventory>array)
                                "__typename": "OBJECT_TYPE" (string)
                            },
                            "__typename": "OBJECT_TYPE" (string)
                        },
                        "restrictions": [RESTRICTION_ARRAY], (<restriction>array)
                        "subscribable": "SUBSCRIBABLE?", (boolean)
                        "subscriptions": [
                            {
                                "description": "SUBSCRIPTION_DESCRIPTION", (string)
                                "id": SUBSCRIPTION_ID, (number)
                                "interval": "SUBSCRIPTION_INTERVAL", (string)
                                "__typename": "OBJECT_TYPE" (string)
                            }
                        ],
                        "__typename": "OBJECT_TYPE" (string)
                    }
                ],
                "__typename": "OBJECT_TYPE" (string)
            },
            "__typename": "OBJECT_TYPE" (string)
        }
    }
}
```

## Search Models

### URL

`?operationName=modelSearch&variables=%7B%22q%22%3A%22{YOUR_SEARCH}%22%2C%22page%22%3A%7B%22from%22%3A{PAGE}%2C%22size%22%3A{PAGE_SIZE}%7D%2C%22priceFilter%22%3A%7B%22name%22%3A%22PRICE%22%2C%22type%22%3A%22RANGE%22%2C%22values%22%3A%5B%22%3E1%22%5D%7D%2C%22filters%22%3A%5B%5D%7D&extensions=%7B%22persistedQuery%22%3A%7B%22version%22%3A1%2C%22sha256Hash%22%3A%22eadec1e2e8cbfc0b7c3a4b87de9af960a4aa14df1cbea852fdff503e9740ad67%22%7D%7D`

### URL Broken Down into JSON

```
"url": {
    "operationName": "modelSearch", (string)
    "variables": {
        "q": "{YOUR_SEARCH_QUERY}", (string)
        "page": {
            "from": {PAGE}, (number)
            "size": {PAGE_SIZE} (number)
        },
        "priceFilter": {
            "name": "PRICE", (string)
            "type": "RANGE", (string)
            "values": [">1"] ([]string)
        },
        filters: []
    },
    "extensions": {
        "persistedQuery": {
            "version": 1, (number)
            "sha256Hash": "eadec1e2e8cbfc0b7c3a4b87de9af960a4aa14df1cbea852fdff503e9740ad67" (string)
        }
    }
}
```

### Headers

"Accept": "*/*"

"Accept-Language": "en-US,en;q=0.9"

"Cache-Control": "no-cache"

"Pragma": "no-cache"

"Origin": "https://www.searspartsdirect.com"

"Referer": "https://www.searspartsdirect.com/"

"content-type": "application/json"

"X-Api-Key": "tV6bZZfhUh3MQmZggG6iq6LjfrZgQgcR26Yv86En"

### JSON Response

```
{
    "data": {
        "modelSearch": {
            "totalCount": TOTAL_RESULT_COUNT, (number)
            "models": [
                {
                    "id": "MODEL_ID", (string)
                    "number": "MODEL_NUMBER", (string)
                    "title": "MODEL_TITLE", (string)
                    "brand": {
                        "id": "BRAND_ID", (string)
                        "name": "BRAND_NAME", (string)
                        "__typename": "OBJECT_TYPE" (string)
                    },
                    "taxonomies": {
                        "taxonomies": [
                            {
                                "id": "TAXONOMY_ID", (string)
                                "name": "TAXONOMY_NAME", (string)
                                "__typename": "OBJECT_TYPE" (string)
                            },
                        ],
                        "__typename": "OBJECT_TYPE" (string)
                    },
                    "__typename": "OBJECT_TYPE" (string)
                }
            ]
        }
    }
}
```
