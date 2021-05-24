# DAST API

### Introduction

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

```
{
    "id": "60aac572a7fd79e56a31bf33",
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
            "id": "criteria1",
            "description": "criteria1",
            "parent": "",
            "score": {
                "local": 0,
                "global": 0
            }
        },
        {
            "level": 0,
            "id": "criteria2",
            "description": "criteria2",
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

```
{
   "level": 0,
   "id": "criteria2",
   "description": "criteria2",
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

Once 

## Authenticated Features

### User CRUD
Access to the API is granted by providing your username and password using HTTP basic authentication.  The username and password used, is the same username and password you use to access the Zingle web interface.

```no-highlight
GET https://api.zingle.me/v1/

{
    "status": {
        "text": "OK",
        "status_code": 200,
        "description": null
    },
    "auth": {
        "id": "4c11f5e3-50b6-4995-b471-b8ef0015488d",
        "email": "joe@example.com",
        "first_name": "Joe",
        "last_name": "Smith",
        "title": null,
        "authorization_class": "contact"
    }
}
```
 
### Token Auth

### Historic

### Restrictions

## Endpoints
