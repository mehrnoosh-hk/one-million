#Problem Requirements
We want to create a http-endpoint so that users can POST their work payload 
and receive the result back.
The nature of the works are time-consuming. We want to concurrently handle 
requests so that the process of payloads do not block other requests.
We do not want to spin up more than MAX_CONCURRENT number of goroutines at 
the same time. If MAX_CONCURRENT number is reached, Incoming request's 
payloads should be queued.


For Solving this problem we use Hexagonal Architecture.
###1-Create an endpoint
####1-1-Create and run a server on port 4242
####1-2-Create endpoint to receive POST requests.
End Point sends the received payload (with JSON) format and convert it to 
struct and send it to service to be processed. To do this we should first 
define required ports in the service package.


