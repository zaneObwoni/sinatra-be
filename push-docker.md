
#login to dockerhub
docker login --username=<yourusername> (you will be prompted for a password)

#push to dockerhub
docker push zanetech/sinatra-be:latest

#build and run image
docker build . -t sinatra
docker run -i -t -p 8080:8080 sinatra