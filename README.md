# dockerizeapp
Application to create a docker container out of a binary

#To test
curl -X POST http://localhost:8080/upload -F "file=@/root/go/src/GAutoDocker/binaries/a_x86.out"   -H "Content-Type: multipart/form-data"
