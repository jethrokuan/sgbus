# sgbus
## Rationale
Components like the ESPert have small computational memory, and cannot store large amounts of bytes for further processing.

The `sgbus` server provides a stripped-down API of lower degrees of nesting that will be a joy to use on IoT devices.

The data is marshalled into structs, which are easily operated upon. Such operations include:

1. Calculating the time to the next bus

Heavy computation can be done on the server side, leaving the IoT device solely responsible for presentation.

## Instructions
``` bash
git clone https://github.com/jethrokuan/sgbus && cd sgbus
go get -u github.com/joho/godotenv
go get -u github.com/labstack/echo/...
go build
mv .env.sample .env
# After editing .env file...
./sgbus
```
## Getting Your Key
You can get your key [here](http://www.mytransport.sg/content/mytransport/home/dataMall.html).

## API
1. `http://localhost:3000/bus-stop/123213`: Retrieves data for busstop with id `123123`

## License
MIT
