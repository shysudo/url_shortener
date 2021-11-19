# url_shortener

  This is a bare-bones example of "url_shortener" application providing a REST API's to shortening the url.

## Build Application
  
  go build main.go -o url_short 
  
  or 
  
  go build
  
## Run test
  
  go test ./...

## Run Application
  
  go run ./main.go
  
  or
  
  ./url_short
 
 ## REST API
 
  The REST API to the example app is described below.

## Get short url of given url

### Request

    curl -request POST 'http://localhost:8080/url' \
         --header 'Content-Type: application/json' \
        --data-raw '{
            "url" : "https://developer.webex.com/docs1"
          }'
          
### Response
     
     short://7UenC-901603f2-2
