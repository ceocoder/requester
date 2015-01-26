run:
	@go install
	@requester -requests 1 -numThreads 1 -url=http://localhost:8080/bid_request/index
