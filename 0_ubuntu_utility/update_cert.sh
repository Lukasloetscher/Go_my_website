#this file needs to be run as sudo
#i recommend to run this script each night at around 2am, to minimize impact to users.

#stop docker so w can reload the certificates
docker stop lulo_portfolio
docker rm lulo_portfolio
#since we already stopped it, we can as well update the repo.
docker pull lukasloetscher/lulo_tryout:latest

#get new certificates
sudo certbot renew --force-renewal

#copy the certificates to the correct place
cp -r -L /etc/letsencrypt/live/* /srv/certificates/
#make the certificates accessible by non root users
chmod 644 /etc/letsencrypt/live/lukas.loetscher.swiss/privkey.pem

#restart docker
docker run -d -p 443:8080 -e PROD=true --restart always --name lulo_portfolio -v /srv/certificates:/mount/certificates:ro lukasloetscher/lulo_tryout