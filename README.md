# api-iso8583-to-JSON
Intermediate component to translate a JSON request into a ISO8583 transaction.
The microservice receives a request, converts it to an ISO 8583 message and forward it to a backend service. It then receives the response and converts it to a JSON to respond to the initial request.

## Request
The fields have the same restrictions as ISO8583.

When the microservice converts the message, it adds a left padding with blank spaces to 'string' fields and a left pading with '0' to 'number' fields to fill the lenght of the field as specified by ISO8583.
## Example request

```
{
    "mti": "0200",
    "fields": {
        "2": "4321123443211234",
        "3": "000000",
        "4": "000000012300",
        "7": "0304054133",
        "11": "001205",
        "14": "0205",
        "18": "5399",
        "22": "022",
        "25": "00",
        "35": "2312312332",
        "37": "206305000014",
        "41": "29110001",
        "42": "1001001",
        "49": "840"
    }
}
```

## Configuration
The configuration is in file [app.json](app.json).

| Key           | Description                                   | Default value             |
| -----------   | -----------                                   | ---------------------     |
| port          | The listen port                               | 8080                      |
| path          | The endpoint path                             | /iso2json                 |
| loglevel      | Log level                                     | debug                     |
| timeout       | Timeout to reach the backend svc              | 5                         |
| backend       | Endpoint to a backend svc that uses ISO8583   | mock                      |

*The backend configuration "mock" uses a mock client that simulates a call to a backend service using ISO8583 and it answer is just a copy of the request into a response and changin the MTI to a response code.*
## Deploy
docker run --name=iso2json -p 8080:8080 iso8583tojson

## Tests
go test ./... -cover