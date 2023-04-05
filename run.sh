#remove unused containers
docker prune container  

#create network
docker create network grpc

#run server on 8080 
docker run -p 8080:8080 
--name server-service --network grpc server

#run client 
docker run --network grpc --name client-service client