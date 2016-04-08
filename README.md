# sgbus
## Rationale
Some components like the ESPert have very little computational memory, and cannot store a large amounts of bytes in memory.

The `sgbus` server provides a stripped-down API of lower degrees of nesting that will be a joy to use on IoT devices.

Additionally, heavy computation can be done on the server side, leaving the IoT device solely responsible for presentation.

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
You can get your key [here](http://www.mytransport.sg/content/mytransport/home/dataMall.html)

## API
1. `http://localhost:3000/bus-stop/123213`: Retrieves data for busstop with id `123123`

## License
MIT
