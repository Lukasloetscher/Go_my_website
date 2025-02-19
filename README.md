This is a simple go project to host wa websuite.
The Project is in first place about learning stuff. 
I restarted from the beginning, in order to incorporate Clean Code advice from the Book Clean Code from Robert C. Martins into it.

Secondly i try to make the project in a way, that i start with a workjing website fast, and then expand on it over time. (Instead of letting the project die in perfection.)
This also means that i will (or at least paln) to create the code fast,a nd then do multiple refractorings.

Currebntly the website is build with docker and then deployed to a Ubuntu VM

used commands:
 docker build --tag lukasloetscher/lulo_tryout .
 -> then use docker desktop to push it
 -> open cmd of ubuntu vm
 -> stop container 
   -> docker ps
   -> docker stop {name}
-> update the image
  -> docker pull lukasloetscher/lulo_tryout
  -> check with : docker images
-> start the docker: (maybe you need to change
  -> docker run -d -p 80:8080 lukasloetscher/lulo_tryout
  -> check with docker ps
  -> check by visitign the webpage

My notes:
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04
