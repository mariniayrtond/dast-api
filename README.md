# DAST API

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

## POST Hierarchy

    POST dast/hierarchy

Creates a [Hierarchy](#hierarchy-object) object.

### Parameters
#### URI Parameters
None
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
    "id": "09da9c83dc7d4faaaca931a0bb644947",
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

## Objects

### Hierarchy Object

Field | Data Type | Read Only | Description
--- | --- | --- | ---
id | string | Y | UUID for the hierarchy.
name | string | N | Name of the hierarchy.
description | string | N | Description of the hierarchy.
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
