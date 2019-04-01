go build

mv srv puffin-ipdb-srv 

# build srv img
docker build -t puffin-ipdb .
