grpcurl -plaintext -d '{"ticket":{"source":"a","destination":"b","price":123},"user":{"email":"abc","firstName":"harshil","lastName":"gupta"}}' localhost:8080 cloudbees.Ticketing/Purchase
grpcurl -plaintext -d '{"ticket_id":0}' localhost:8080 cloudbees.Seating/Allocate
