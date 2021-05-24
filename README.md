# DAST API

- Basic Usage
   - Guest Hierarchies
   - Guest Judgements
- Authenticated Features
   - Users
   - Token Auth
   - Restrictions
- HTTP Requests
- HTTP Response Codes
- Endpoints
   - Hierarchy
      - Create
      - GET
      - Search by Username
   - Criteria
      - PUT
   - Pairwise Matrix
      - Generate
   - Judgements
      - GET
      - Search
      - Set
      - Resolve
   - Users
      - Create
      - Log In
      - Validate Token

### Introduction

> **PAY ATTENTION**: for a better understanding of this documentation it is important that the reader understands the analytical hierarchy process. 

DAST API is a RESTFull API written in GoLang that implements the analytic hierarchy process.

This API was developed in the context of a final computer engineering project by Ayrton Marini.

Provides the following features:

1. **CRUD hierarchies:** With DAST you can create your own AHP hierarchies.
2. **CRUD DAST Users:** If you need a history of troubleshooting, you can create your own user for the DAST API.
3. **Create Pairwise Comparison Matrix:** With a hierarchy you can create the well-known "pairwise comparison matrix". 
   This structure allows to store the judges and solve the multi-criteria problem.
4. **Solve the decision problem:** once you have a valid hierarchy, with your generated matrix, and
   the judges stored, then you can solve the problem and see the results.

### Support
For API support, please email mariniayrtond@gmail.com.

## Basic Usage

The DAST API has some basic features and it also provides special features.

### Guest Hierarchies

> If you are looking how to create a hierarchy, please jump to the [endpoints](#endpoints) section.

For an AHP process, a Hierarchy represents a fundamental concept. In DAST, a hierarchy is represented by a JSON
as follows:

```json
{
    "id": "UUID",
    "name": "new hierarchy",
    "description": "description",
    "owner": "guest",
    "alternatives": [
        "alternative1",
        "alternative2"
    ],
    "criteria": [
        {
            "level": 0,
            "id": "criterion1",
            "description": "criterion1",
            "parent": "",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 0,
            "id": "criterion2",
            "description": "criterion2",
            "parent": "",
            "score": {
                "local": 0,
                "global": 0
            }
        }
    ]
}
```

> All the public Hierarchies has **"guest"** owner

Each hierarchy has its own alternatives and its own criteria tree.

The alternatives are represented by a string JSON Array.

On the other hand, the criteria tree it's a JSON Array that contains JSON objects as follows:

```json
{
   "level": 0,
   "id": "criterion2",
   "description": "criterion2",
   "parent": "",
   "score": {
       "local": 0,
       "global": 0
   }
}
```

Each criterion has a level, a description and an id. If the criterion belongs to a deep level
in the criteria tree, you can specify it using the "parent" attribute.

### Guest Judgements

> If you have a Hierarchy, and you are looking how to create a judgement, please jump to the [endpoints](#endpoints) section.

Once you generate a valid Hierarchy, then you can generate Judgements. A judgment in DAST is an JSON
as follows:

```json
{
    "id": "UUID",
    "hierarchy_id": "Hierarchy_UUID",
    "status": "Complete",
    "date_created": "2021-05-23T18:13:35-03:00",
    "date_last_updated": "2021-05-23T18:13:46-03:00",
    "criteria_comparison": [
        {
            "level": 0,
            "matrix_context": {
                "compared_to": "",
                "elements": [
                    "criterion1",
                    "criterion2"
                ],
                "judgements": [
                    [
                        1,
                        2
                    ],
                    [
                        0.5,
                        1
                    ]
                ]
            }
        }
    ],
    "alternative_comparison": [
        {
            "compared_to": "criterio1",
            "elements": [
                "alternative1",
                "alternative2"
            ],
            "judgements": [
                [
                    1,
                    0.2
                ],
                [
                    5,
                    1
                ]
            ]
        },
        {
            "compared_to": "criterio2",
            "elements": [
                "alternative1",
                "alternative2"
            ],
            "judgements": [
                [
                    1,
                    0.3333333333333333
                ],
                [
                    3,
                    1
                ]
            ]
        }
    ],
    "results": {
        "alternative1": 0.19444444444444445,
        "alternative2": 0.8055555555555556
    }
}
```

This object is made up of three elements:

* **Criteria Comparison**: this JSON array contains the generated matrix for pairwise criterion comparison only.
* **Alternatives Comparison**: this JSON array contains the generated matrix for the pairwise comparison between alternatives vs
  each criterion.
* **Results**: this JSON object contains the result associated to each alternative, the sum total of the result values
must be 1.

## Authenticated Features

Registered user has access to special features such as:

* **Private Problems**: the hierarchies created by registered users can only be modified by the owner.
* **Search Hierarchies**: you have access to search the history of hierarchies resolved by DAST for you.
* **Search Judgements**: you have access to search the history of judgements made by you for specific hierarchies.

### Users

> If you need to register a new user, please jump to the [endpoints](#endpoints) section.

For the logged in user, DAST stores certain non-critical information in JSON format as follows:

```json
{
    "id": "user_id",
    "name": "user_name",
    "email": "user@mail.com"
}
```
 
### Token Auth

> If you need to create an auth token, please jump to the [endpoints](#endpoints) section.

Once you register a user, you can create an authentication token for that user.

The DAST Token is a `SHA256` hash.

### Restrictions

Each Hierarchy created by a no-guest user only can be modified by the owner. To authenticate your requests,
you have to send a header as follows:

```json
{
   "X-Auth-Token": "your SHA256 Hash"
}
```

## HTTP requests
All API requests are made by sending a secure HTTPS request using one of the following methods, depending on the action being taken:

* `POST` Create a resource
* `PUT` Update a resource
* `GET` Get a resource or list of resources
* `DELETE` Delete a resource

For PUT and POST requests the body of your request may include a JSON payload, and the URI being requested may include a query string specifying additional filters or commands, all of which are outlined in the following sections.

## HTTP Response Codes
Each response will be returned with one of the following HTTP status codes:

* `200` `OK` The request was successful.
* `201` `Created` The request was successful and also created a new resource.
* `400` `Bad Request` There was a problem with the request (security, malformed, data validation, etc.).
* `401` `Unauthorized` The supplied X-Auth-token is invalid.
* `404` `Not found` An attempt was made to access a resource that does not exist in the API.
* `405` `Method not allowed` The resource being accessed doesn't support the method specified (GET, POST, etc.).
* `500` `Server Error` An error on the server occurred.
* `503` `Server Unavailable` An error on the server network ocurred.

## Endpoints

## Create Hierarchy

    POST dast/hierarchy

Creates a [Hierarchy](#hierarchy-object) object.

### Parameters
#### URI Parameters
None

#### Headers
Field | Required | Description
--- | --- | ---
X-auth-token | Y | If hierarchy owner user is not equal to `guest`, then the X-auth-token header is required.
#### Body Parameters
Field | Required | Description
--- | --- | ---
name | Y | Name hierarchy is required
description | Y | Description hierarchy is required
owner | Y | You can modify this value to attach the hierarchy to specific registered user. If you want to do this, then you need to send the `X-auth-token` header. Otherwise always use `guest`.
objective | Y | Each hierarchy must have an objective. Objective hierarchy is required.
alternatives | Y | At least one is required.

### Example
#### URL

    POST https://dast-api.herokuapp.com/dast/hierarchy

#### Request Body
```json
{
    "name": "testing complex case",
    "description": "test hierarchy",
    "owner": "guest",
    "objective": "choose card",
    "alternatives": [
        "Ford",
        "Fiat",
        "Chevrolet"
    ]
}
```

#### Response

`Status 201: Created`

``` json
{
    "id": "60ac0e4cd920ddc480891cc9",
    "name": "testing complex case",
    "description": "esta es una jerarquía de prueba",
    "owner": "guest",
    "alternatives": [
        "Ford",
        "Fiat",
        "Chevrolet"
    ],
    "criteria": []
}
```

## GET Hierarchy

    GET dast/hierarchy/:id

Returns a [Hierarchy](#hierarchy-object) object.

### Parameters
None

### Example
#### URL

    GET https://dast-api.herokuapp.com/dast/hierarchy/60ac0e4cd920ddc480891cc9

#### Response

`Status 200: OK`

``` json
{
    "id": "60ac0e4cd920ddc480891cc9",
    "name": "testing complex case",
    "description": "esta es una jerarquía de prueba",
    "owner": "guest",
    "alternatives": [
        "Ford",
        "Fiat",
        "Chevrolet"
    ],
    "criteria": []
}
```

## Search User Hierarchies

    GET dast/hierarchies/:username/search

Returns a [][Hierarchy](#hierarchy-object) array object associated to a specific user.

### Parameters
None

### Example
#### URL

    GET https://dast-api.herokuapp.com/dast/hierarchies/prueba/search

#### Response

`Status 200: OK`

``` json
[
    {
        "id": "60ac0e4cd920ddc480891cc9",
        "name": "nueva jerarquía",
        "description": "una nueva jerarquía",
        "owner": "prueba",
        "alternatives": [
            "ganar",
            "mundial"
        ],
        "criteria": [
            {
                "level": 0,
                "id": "criterio1",
                "description": "criterio1",
                "parent": "",
                "score": {
                    "local": 0,
                    "global": 0
                }
            },
            {
                "level": 0,
                "id": "criterio2",
                "description": "criterio2",
                "parent": "",
                "score": {
                    "local": 0,
                    "global": 0
                }
            }
        ]
    }
]
```

## Criteria PUT

    PUT dast/hierarchy/:id/criteria

Attach a [][Criteria](#criterion-object) object array to a [Hierarchy](#hierarchy-object) object.

### Parameters
#### Body Parameters
Field | Required | Description
--- | --- | ---
level | Y | Level of the criterion.
description | Y | Description of the criterion.
id | Y | Id of the criterion. This id must be in snake-case.
parent | Y | Required if the criterion level is > 0.

#### Headers
Field | Required | Description
--- | --- | ---
X-auth-token | Y | If hierarchy owner user is not equal to `guest`, then the X-auth-token header is required.

### Example
#### URL

    PUT https://dast-api.herokuapp.com/dast/hierarchy/60ac0e4cd920ddc480891cc9/criteria

#### Request Body
```json
[
   {
      "level": 0,
      "description": "Velocidad",
      "id": "velocidad"
   },
   {
      "level": 1,
      "description": "Velocidad 1",
      "id": "velocidad_1",
      "parent": "velocidad"
   },
   {
      "level": 1,
      "description": "Velocidad 2",
      "id": "velocidad_2",
      "parent": "velocidad"
   },
   {
      "level": 0,
      "description": "Cilindrada",
      "id": "cilindrada"
   },
   {
      "level": 0,
      "description": "Aceleración",
      "id": "aceleracion"
   },
   {
      "level": 1,
      "description": "Aceleracion 1",
      "id": "aceleracion_1",
      "parent": "aceleracion"
   },
   {
      "level": 1,
      "description": "Aceleracion 2",
      "id": "aceleracion_2",
      "parent": "aceleracion"
   },
   {
      "level": 2,
      "description": "Aceleracion 21",
      "id": "aceleracion_21",
      "parent": "aceleracion_2"
   },
   {
      "level": 2,
      "description": "Aceleracion 22",
      "id": "aceleracion_22",
      "parent": "aceleracion_2"
   }
]
```

#### Response

`Status 200: OK`

``` json
{
    "id": "60ac0e4cd920ddc480891cc9",
    "name": "testing complex case",
    "description": "test hierarchy",
    "owner": "guest",
    "alternatives": [
        "Ford",
        "Fiat",
        "Chevrolet"
    ],
    "criteria": [
        {
            "level": 0,
            "id": "velocidad",
            "description": "Velocidad",
            "parent": "",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 1,
            "id": "velocidad_1",
            "description": "Velocidad 1",
            "parent": "velocidad",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 1,
            "id": "velocidad_2",
            "description": "Velocidad 2",
            "parent": "velocidad",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 0,
            "id": "cilindrada",
            "description": "Cilindrada",
            "parent": "",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 0,
            "id": "aceleracion",
            "description": "Aceleración",
            "parent": "",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 1,
            "id": "aceleracion_1",
            "description": "Aceleracion 1",
            "parent": "aceleracion",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 1,
            "id": "aceleracion_2",
            "description": "Aceleracion 2",
            "parent": "aceleracion",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 2,
            "id": "aceleracion_21",
            "description": "Aceleracion 21",
            "parent": "aceleracion_2",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 2,
            "id": "aceleracion_22",
            "description": "Aceleracion 22",
            "parent": "aceleracion_2",
            "score": {
                "local": 0,
                "global": 0
            }
        }
    ]
}
```

## Generate Pairwise Matrix

    POST dast/pairwise/:hierarchy_id/generate

Creates a [Judgement](#judgement-object) object related to specific [Hierarchy](#hierarchy-object).

### Parameters
None

#### Headers
Field | Required | Description
--- | --- | ---
X-auth-token | Y | If hierarchy owner user is not equal to `guest`, then the X-auth-token header is required.

### Example
#### URL

    POST https://dast-api.herokuapp.com/dast/pairwise/60ac0e4cd920ddc480891cc9/generate

#### Response

`Status 201: Created`

``` json
{
    "id": "60ac1334d920ddc480891ccb",
    "hierarchy_id": "60ac0e4cd920ddc480891cc9",
    "status": "Incomplete",
    "date_created": "2021-05-24T20:57:24Z",
    "date_last_updated": "2021-05-24T20:57:24Z",
    "criteria_comparison": [
        {
            "level": 0,
            "matrix_context": {
                "compared_to": "",
                "elements": [
                    "velocidad",
                    "cilindrada",
                    "aceleracion"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "velocidad",
                "elements": [
                    "velocidad_1",
                    "velocidad_2"
                ],
                "judgements": [
                    [
                        1,
                        0
                    ],
                    [
                        0,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "aceleracion",
                "elements": [
                    "aceleracion_1",
                    "aceleracion_2"
                ],
                "judgements": [
                    [
                        1,
                        0
                    ],
                    [
                        0,
                        1
                    ]
                ]
            }
        },
        {
            "level": 2,
            "matrix_context": {
                "compared_to": "aceleracion_2",
                "elements": [
                    "aceleracion_21",
                    "aceleracion_22"
                ],
                "judgements": [
                    [
                        1,
                        0
                    ],
                    [
                        0,
                        1
                    ]
                ]
            }
        }
    ],
    "alternative_comparison": [
        {
            "compared_to": "velocidad_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "velocidad_2",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "cilindrada",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_21",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_22",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        }
    ]
}
```

## GET Pairwise Matrix

    GET dast/pairwise/:hierarchy_id/judgements/:judgement_id

Gets a [Judgement](#judgement-object) object related to specific [Hierarchy](#hierarchy-object).

### Parameters
None
#### Headers
Field | Required | Description
--- | --- | ---
X-auth-token | Y | If hierarchy owner user is not equal to `guest`, then the X-auth-token header is required.

### Example
#### URL

    GET https://dast-api.herokuapp.com/dast/pairwise/60ac0e4cd920ddc480891cc9/judgements/60ac1334d920ddc480891ccb

#### Response

`Status 200: OK`

``` json
{
    "id": "60ac1334d920ddc480891ccb",
    "hierarchy_id": "60ac0e4cd920ddc480891cc9",
    "status": "Incomplete",
    "date_created": "2021-05-24T20:57:24Z",
    "date_last_updated": "2021-05-24T20:57:24Z",
    "criteria_comparison": [
        {
            "level": 0,
            "matrix_context": {
                "compared_to": "",
                "elements": [
                    "velocidad",
                    "cilindrada",
                    "aceleracion"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "velocidad",
                "elements": [
                    "velocidad_1",
                    "velocidad_2"
                ],
                "judgements": [
                    [
                        1,
                        0
                    ],
                    [
                        0,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "aceleracion",
                "elements": [
                    "aceleracion_1",
                    "aceleracion_2"
                ],
                "judgements": [
                    [
                        1,
                        0
                    ],
                    [
                        0,
                        1
                    ]
                ]
            }
        },
        {
            "level": 2,
            "matrix_context": {
                "compared_to": "aceleracion_2",
                "elements": [
                    "aceleracion_21",
                    "aceleracion_22"
                ],
                "judgements": [
                    [
                        1,
                        0
                    ],
                    [
                        0,
                        1
                    ]
                ]
            }
        }
    ],
    "alternative_comparison": [
        {
            "compared_to": "velocidad_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "velocidad_2",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "cilindrada",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_21",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_22",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0,
                    0
                ],
                [
                    0,
                    1,
                    0
                ],
                [
                    0,
                    0,
                    1
                ]
            ]
        }
    ]
}
```

## Search Pairwise Matrix By Hierarchy

    GET dast/pairwise/:hierarchy_id/search/judgements

Gets a [Judgement](#judgement-object) object array related to specific [Hierarchy](#hierarchy-object).

### Parameters
None

#### Headers
Field | Required | Description
--- | --- | ---
X-auth-token | Y | If hierarchy owner user is not equal to `guest`, then the X-auth-token header is required.

### Example
#### URL

     GET https://dast-api.herokuapp.com/dast/pairwise/60ac0e4cd920ddc480891cc9/search/judgements

#### Response

`Status 200: OK`

``` json
[
    {
        "id": "60ac1334d920ddc480891ccb",
        "hierarchy_id": "60ac0e4cd920ddc480891cc9",
        "status": "Incomplete",
        "date_created": "2021-05-24T20:57:24Z",
        "date_last_updated": "2021-05-24T20:57:24Z",
        "criteria_comparison": [
            {
                "level": 0,
                "matrix_context": {
                    "compared_to": "",
                    "elements": [
                        "velocidad",
                        "cilindrada",
                        "aceleracion"
                    ],
                    "judgements": [
                        [
                            1,
                            0,
                            0
                        ],
                        [
                            0,
                            1,
                            0
                        ],
                        [
                            0,
                            0,
                            1
                        ]
                    ]
                }
            },
            {
                "level": 1,
                "matrix_context": {
                    "compared_to": "velocidad",
                    "elements": [
                        "velocidad_1",
                        "velocidad_2"
                    ],
                    "judgements": [
                        [
                            1,
                            0
                        ],
                        [
                            0,
                            1
                        ]
                    ]
                }
            },
            {
                "level": 1,
                "matrix_context": {
                    "compared_to": "aceleracion",
                    "elements": [
                        "aceleracion_1",
                        "aceleracion_2"
                    ],
                    "judgements": [
                        [
                            1,
                            0
                        ],
                        [
                            0,
                            1
                        ]
                    ]
                }
            },
            {
                "level": 2,
                "matrix_context": {
                    "compared_to": "aceleracion_2",
                    "elements": [
                        "aceleracion_21",
                        "aceleracion_22"
                    ],
                    "judgements": [
                        [
                            1,
                            0
                        ],
                        [
                            0,
                            1
                        ]
                    ]
                }
            }
        ],
        "alternative_comparison": [
            {
                "compared_to": "velocidad_1",
                "elements": [
                    "Ford",
                    "Fiat",
                    "Chevrolet"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            },
            {
                "compared_to": "velocidad_2",
                "elements": [
                    "Ford",
                    "Fiat",
                    "Chevrolet"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            },
            {
                "compared_to": "cilindrada",
                "elements": [
                    "Ford",
                    "Fiat",
                    "Chevrolet"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            },
            {
                "compared_to": "aceleracion_1",
                "elements": [
                    "Ford",
                    "Fiat",
                    "Chevrolet"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            },
            {
                "compared_to": "aceleracion_21",
                "elements": [
                    "Ford",
                    "Fiat",
                    "Chevrolet"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            },
            {
                "compared_to": "aceleracion_22",
                "elements": [
                    "Ford",
                    "Fiat",
                    "Chevrolet"
                ],
                "judgements": [
                    [
                        1,
                        0,
                        0
                    ],
                    [
                        0,
                        1,
                        0
                    ],
                    [
                        0,
                        0,
                        1
                    ]
                ]
            }
        ]
    }
]
```

## Set Judgements

    PUT dast/pairwise/:hierarchy_id/judgements/:judgements_id

Set the judgements (`criteria_comparison` and `alternative_comparison`) to a [Judgement](#hierarchy-object) object.

### Parameters
#### Body Parameters
You have to take the nodes: `criteria_comparison` and `alternative_comparison` from the response int the api call shown in
[Generate Pairwise Matrix](#generate-pairwise-matrix).

Once you have the response, you have to fill the `judgements` of each matrix and send it to attach it to the judgement object.

#### Headers
Field | Required | Description
--- | --- | ---
X-auth-token | Y | If hierarchy owner user is not equal to `guest`, then the X-auth-token header is required.

### Example
#### URL

    PUT https://dast-api.herokuapp.com/dast/pairwise/60ac0e4cd920ddc480891cc9/judgements/60ac1334d920ddc480891ccb

#### Request Body
```json
{
   "criteria_comparison": [
      {
         "level": 0,
         "matrix_context": {
            "compared_to": "",
            "elements": [
               "velocidad",
               "cilindrada",
               "aceleracion"
            ],
            "judgements": [
               [
                  1,
                  2,
                  0.5
               ],
               [
                  0.5,
                  1,
                  3
               ],
               [
                  2,
                  0.3333333333333,
                  1
               ]
            ]
         }
      },
      {
         "level": 1,
         "matrix_context": {
            "compared_to": "velocidad",
            "elements": [
               "velocidad_1",
               "velocidad_2"
            ],
            "judgements": [
               [
                  1,
                  2
               ],
               [
                  0.5,
                  1
               ]
            ]
         }
      },
      {
         "level": 1,
         "matrix_context": {
            "compared_to": "aceleracion",
            "elements": [
               "aceleracion_1",
               "aceleracion_2"
            ],
            "judgements": [
               [
                  1,
                  5
               ],
               [
                  0.2,
                  1
               ]
            ]
         }
      },
      {
         "level": 2,
         "matrix_context": {
            "compared_to": "aceleracion_2",
            "elements": [
               "aceleracion_21",
               "aceleracion_22"
            ],
            "judgements": [
               [
                  1,
                  0.4
               ],
               [
                  2.5,
                  1
               ]
            ]
         }
      }
   ],
   "alternative_comparison": [
      {
         "compared_to": "velocidad_1",
         "elements": [
            "Ford",
            "Fiat",
            "Chevrolet"
         ],
         "judgements": [
            [
               1,
               0.5,
               3
            ],
            [
               2,
               1,
               0.5
            ],
            [
               0.333333333333,
               2,
               1
            ]
         ]
      },
      {
         "compared_to": "velocidad_2",
         "elements": [
            "Ford",
            "Fiat",
            "Chevrolet"
         ],
         "judgements": [
            [
               1,
               2,
               3
            ],
            [
               0.5,
               1,
               4
            ],
            [
               0.333333333333,
               0.25,
               1
            ]
         ]
      },
      {
         "compared_to": "cilindrada",
         "elements": [
            "Ford",
            "Fiat",
            "Chevrolet"
         ],
         "judgements": [
            [
               1,
               2,
               0.5
            ],
            [
               0.5,
               1,
               0.25
            ],
            [
               2,
               4,
               1
            ]
         ]
      },
      {
         "compared_to": "aceleracion_1",
         "elements": [
            "Ford",
            "Fiat",
            "Chevrolet"
         ],
         "judgements": [
            [
               1,
               2,
               0.5
            ],
            [
               0.5,
               1,
               0.33333333333333
            ],
            [
               2,
               3,
               1
            ]
         ]
      },
      {
         "compared_to": "aceleracion_21",
         "elements": [
            "Ford",
            "Fiat",
            "Chevrolet"
         ],
         "judgements": [
            [
               1,
               4,
               5
            ],
            [
               0.25,
               1,
               3
            ],
            [
               0.2,
               0.3333333333333,
               1
            ]
         ]
      },
      {
         "compared_to": "aceleracion_22",
         "elements": [
            "Ford",
            "Fiat",
            "Chevrolet"
         ],
         "judgements": [
            [
               1,
               2,
               1
            ],
            [
               0.5,
               1,
               0.3333333333333
            ],
            [
               1,
               3,
               1
            ]
         ]
      }
   ]
}
```

#### Response

`Status 200: OK`

``` json
{
    "id": "60ac1334d920ddc480891ccb",
    "hierarchy_id": "60ac0e4cd920ddc480891cc9",
    "status": "Incomplete",
    "date_created": "2021-05-24T20:57:24Z",
    "date_last_updated": "2021-05-24T22:03:10Z",
    "criteria_comparison": [
        {
            "level": 0,
            "matrix_context": {
                "compared_to": "",
                "elements": [
                    "velocidad",
                    "cilindrada",
                    "aceleracion"
                ],
                "judgements": [
                    [
                        1,
                        2,
                        0.5
                    ],
                    [
                        0.5,
                        1,
                        3
                    ],
                    [
                        2,
                        0.3333333333333,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "velocidad",
                "elements": [
                    "velocidad_1",
                    "velocidad_2"
                ],
                "judgements": [
                    [
                        1,
                        2
                    ],
                    [
                        0.5,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "aceleracion",
                "elements": [
                    "aceleracion_1",
                    "aceleracion_2"
                ],
                "judgements": [
                    [
                        1,
                        5
                    ],
                    [
                        0.2,
                        1
                    ]
                ]
            }
        },
        {
            "level": 2,
            "matrix_context": {
                "compared_to": "aceleracion_2",
                "elements": [
                    "aceleracion_21",
                    "aceleracion_22"
                ],
                "judgements": [
                    [
                        1,
                        0.4
                    ],
                    [
                        2.5,
                        1
                    ]
                ]
            }
        }
    ],
    "alternative_comparison": [
        {
            "compared_to": "velocidad_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0.5,
                    3
                ],
                [
                    2,
                    1,
                    0.5
                ],
                [
                    0.333333333333,
                    2,
                    1
                ]
            ]
        },
        {
            "compared_to": "velocidad_2",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    3
                ],
                [
                    0.5,
                    1,
                    4
                ],
                [
                    0.333333333333,
                    0.25,
                    1
                ]
            ]
        },
        {
            "compared_to": "cilindrada",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    0.5
                ],
                [
                    0.5,
                    1,
                    0.25
                ],
                [
                    2,
                    4,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    0.5
                ],
                [
                    0.5,
                    1,
                    0.33333333333333
                ],
                [
                    2,
                    3,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_21",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    4,
                    5
                ],
                [
                    0.25,
                    1,
                    3
                ],
                [
                    0.2,
                    0.3333333333333,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_22",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    1
                ],
                [
                    0.5,
                    1,
                    0.3333333333333
                ],
                [
                    1,
                    3,
                    1
                ]
            ]
        }
    ]
}
```

## Resolve Pairwise Matrix

    POST dast/pairwise/:hierarchy_id/judgements/:judgements_id/resolve

Resolve a [Judgement](#judgement-object) object related to specific [Hierarchy](#hierarchy-object). The results will be
placed at `results` node.

### Parameters
None

#### Headers
Field | Required | Description
--- | --- | ---
X-auth-token | Y | If hierarchy owner user is not equal to `guest`, then the X-auth-token header is required.

### Example
#### URL

    POST https://dast-api.herokuapp.com/dast/pairwise/60ac0e4cd920ddc480891cc9/judgements/60ac1334d920ddc480891ccb/resolve

#### Response

`Status 201: Created`

``` json
{
    "id": "60ac1334d920ddc480891ccb",
    "hierarchy_id": "60ac0e4cd920ddc480891cc9",
    "status": "Complete",
    "date_created": "2021-05-24T20:57:24Z",
    "date_last_updated": "2021-05-24T22:15:12Z",
    "criteria_comparison": [
        {
            "level": 0,
            "matrix_context": {
                "compared_to": "",
                "elements": [
                    "velocidad",
                    "cilindrada",
                    "aceleracion"
                ],
                "judgements": [
                    [
                        1,
                        2,
                        0.5
                    ],
                    [
                        0.5,
                        1,
                        3
                    ],
                    [
                        2,
                        0.3333333333333,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "velocidad",
                "elements": [
                    "velocidad_1",
                    "velocidad_2"
                ],
                "judgements": [
                    [
                        1,
                        2
                    ],
                    [
                        0.5,
                        1
                    ]
                ]
            }
        },
        {
            "level": 1,
            "matrix_context": {
                "compared_to": "aceleracion",
                "elements": [
                    "aceleracion_1",
                    "aceleracion_2"
                ],
                "judgements": [
                    [
                        1,
                        5
                    ],
                    [
                        0.2,
                        1
                    ]
                ]
            }
        },
        {
            "level": 2,
            "matrix_context": {
                "compared_to": "aceleracion_2",
                "elements": [
                    "aceleracion_21",
                    "aceleracion_22"
                ],
                "judgements": [
                    [
                        1,
                        0.4
                    ],
                    [
                        2.5,
                        1
                    ]
                ]
            }
        }
    ],
    "alternative_comparison": [
        {
            "compared_to": "velocidad_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    0.5,
                    3
                ],
                [
                    2,
                    1,
                    0.5
                ],
                [
                    0.333333333333,
                    2,
                    1
                ]
            ]
        },
        {
            "compared_to": "velocidad_2",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    3
                ],
                [
                    0.5,
                    1,
                    4
                ],
                [
                    0.333333333333,
                    0.25,
                    1
                ]
            ]
        },
        {
            "compared_to": "cilindrada",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    0.5
                ],
                [
                    0.5,
                    1,
                    0.25
                ],
                [
                    2,
                    4,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_1",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    0.5
                ],
                [
                    0.5,
                    1,
                    0.33333333333333
                ],
                [
                    2,
                    3,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_21",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    4,
                    5
                ],
                [
                    0.25,
                    1,
                    3
                ],
                [
                    0.2,
                    0.3333333333333,
                    1
                ]
            ]
        },
        {
            "compared_to": "aceleracion_22",
            "elements": [
                "Ford",
                "Fiat",
                "Chevrolet"
            ],
            "judgements": [
                [
                    1,
                    2,
                    1
                ],
                [
                    0.5,
                    1,
                    0.3333333333333
                ],
                [
                    1,
                    3,
                    1
                ]
            ]
        }
    ],
    "results": {
        "Chevrolet": 0.4424591229862888,
        "Fiat": 0.2162846496027526,
        "Ford": 0.34125622741095857
    }
}
```

## Create User

    POST dast/user/create

Creates a [User](#user-object) object.

### Parameters
#### URI Parameters
None

#### Body Parameters
Field | Required | Description
--- | --- | ---
name | Y | User name.
email | Y | User valid email.
password | Y | User password.

### Example
#### URL

    POST https://dast-api.herokuapp.com/dast/user/create

#### Request Body
```json
{
   "name": "rand",
   "email": "rand@gmail.com",
   "password": "1234"
}
```

#### Response

`Status 201: Created`

```json
{
    "id": "22j1isodjel2kssidjw23xjs",
    "name": "rand",
    "email": "rand@gmail.com"
}
```

## Log In User

    POST dast/user/login

Log in a [User](#user-object) into DAST.

### Parameters
#### URI Parameters
None

#### Body Parameters
Field | Required | Description
--- | --- | ---
name | Y | User name.
password | Y | User password.

### Example
#### URL

    POST https://dast-api.herokuapp.com/dast/user/login

#### Request Body
```json
{
   "name": "rand",
   "password": "1234"
}
```

#### Response

`Status 200: OK`

```json
{
   "message": "rand successful logged in",
   "username": "rand",
   "token": "SHA256 TOKEN"
}
```

## Validate Token

    POST dast/user/login

Validate if a SHA256 token is valid.

### Parameters
#### URI Parameters
None

#### Body Parameters
Field | Required | Description
--- | --- | ---
id | Y | Token username owner.
token | Y | SHA256.

### Example
#### URL

    POST https://dast-api.herokuapp.com/dast/user/validate

#### Request Body
```json
{
   "id": "rand",
   "token": "SHA256 TOKEN"
}
```

#### Response

`Status 204: No Content`

## Objects

### Hierarchy Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
id | string | Y | UUID for the hierarchy.
name | string | Y | Name of the hierarchy.
description | string | Y | Description of the hierarchy.
owner | string | Y | Owner for the hierarchy. `guest` is the default user.
alternatives | string array | Y | Alternatives for the hierarchy.
criteria | object array | Y | This array represents the Criteria Tree for the hierarchy. See [Criterion](#criterion-object) object.

### Criterion Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
level | int | Y | Tree level for the criterion.
id | string | Y | Identifier for the criterion.
description | string | Y | Description of the hierarchy.
parent | string | Y | If the criterion has a parent, then it will be appear here.
score | object | Y | The [Score](#score-object) object represents the weight of the tree node after the calculations have been performed.

### Score Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
local | float | Y | Local score for the node.
global | float | Y | Global score for the node.

### Judgement Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
id | string | Y | UUID for the judgement.
hierarchy_id | string | Y | UUID for the hierarchy
status | string | Y | Judgement status.
date_created | Date ISO | Y | Date created for the judgement.
date_last_updated | Date ISO | Y | Date last updated for the judgement.
criteria_comparison | object array | Y | This array contains each matrix generated for make the pairwise comparison between criteria. See [Criterion Pairwise Matrix](#criteria-pairwise-matrix-object) object.
alternative_comparison | object array | Y | This array contains each matrix generated for make the pairwise comparison between alternatives. See [Pairwise Matrix](#pairwise-matrix-object) object.

### Criteria Pairwise Matrix Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
level | int | Y | Matrix level in comparison context.
matrix_context | object | Y | Matrix values. See [Pairwise Matrix](#pairwise-matrix-object) object.

### Pairwise Matrix Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
compared_to | string | Y | Reference element with which the comparison is being made.
elements | string array | Y | Items being compared.
judgements | []float array | Y | Matrix values.

### User Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
name | string | Y | User name.
password | string | Y | User encrypted password.
email | string | Y | User email.

