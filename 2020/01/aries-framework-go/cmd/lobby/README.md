# Aries Framework Go Presentation Lobby server

This go project provides REST APIs to hold the presentation related data.

## APIs
### Create Demo
HTTP POST /demo

Request:
```
{ 
   "routerUrl":"http://example.com/router"
}
```

Response:

HTTP 200
```
{ 
   "uid":"15626e62-357f-4d7c-96a5-fb021a9381c4"
}
```

### Get Demo
HTTP GET /demo/{uid}

Response:

HTTP 200
```
{ 
   "routerUrl":"http://example.com/router"
}
```

### Post Invitation
HTTP POST /demo/{uid}/invitation

Request:
```
{ 
   "serviceEndpoint":"http://example.com/router",
   "recipientKeys":[ 
      "J8TfRBt7b2UPXaeiK1bbUTH5Y6d7wX44bMLkypXd4Cen"
   ],
   "@id":"56002b57-884e-4a57-981e-61b9e25ecc52",
   "label":"user-agent",
   "@type":"https://didcomm.org/didexchange/1.0/invitation"
}
```

Response:

HTTP 200
```
{ 
}
```

### Retrieve Invitations 
HTTP GET /demo/{uid}/invitations

Response:

HTTP 200
```
{ 
   "invitations":[ 
      { 
         "serviceEndpoint":"routing:endpoint",
         "recipientKeys":[ 
            "J8TfRBt7b2UPXaeiK1bbUTH5Y6d7wX44bMLkypXd4Cen"
         ],
         "@id":"56002b57-884e-4a57-981e-61b9e25ecc52",
         "label":"carl-agent",
         "@type":"https://didcomm.org/didexchange/1.0/invitation"
      }
   ]
}
```
