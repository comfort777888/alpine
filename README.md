# ASCII ART WEB DOCKERIZE#

to initialize project:

**~./Start_Docker_Project.sh**

After that the project starts on server - localhost:8080.

Purpose of project:

###### Ascii-art-web is a program which consists in receiving a string as an argument in text area and outputting the string in a graphic representation using ASCII.The result of program will be in html form.

**allowed range of ASCII is :** 32-126 (all alpa-numeric symbols in latin alphabet + additional common signs).

## Docker:

Build image (You can specify a repository and tag at which to save the new image if the build succeeds):

**~docker build -t ascii-art-web .**

Run image:

**~docker run -d -p 8080:8080 --name ascii-web ascii-art-web**

**~docker images**

**~docker container ls**

Use this to delete everything:

_~docker system prune -a --volumes_

To start existing but not working right now container:

_~docker start <container ID or name>_

To stop existing and working right now container:

_~docker stop <container ID or name>_

To show all containers created in the system:

_~docker ps -a_

The project was written by **_Asya(m.a_k)_** and **_Aray(araya)_**

students of **_Alem school_**, 6-flow students :grinning:
