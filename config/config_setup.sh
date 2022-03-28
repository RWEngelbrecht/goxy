touch .env
if [ -z ${proxy_port+x} ]
then
	echo -n "PORT to LISTEN on? "
	read -r proxy_port
fi
if [ -z ${app_url+x} ]
then
	echo -n "URL:PORT to PROXY to? (e.g. localhost || docker container name || docker-compose service name)"
	read -r app_url
fi
if [ -z ${user_id+x} ]
then
	echo -n "What is your USER ID? "
	read -r user_id
fi
if [ -z ${insight_url+x} ]
then
	echo -n "URL:PORT that insight-proxy is on? (e.g. localhost || docker container name || docker-compose service name)"
	read -r insight_url
fi
echo "Adding variables to .env..."
echo "PROXY_PORT=$proxy_port\nAPP_URL=http://$app_url\nSIGNING_CERT=./certs/jwter-signing-key.pem\nPUBLIC_CERT=./certs/jwter-pub-cert.pem\nUSER_ID=$user_id\nINSIGHT_URL=$insight_url" >> .env
echo "Done"
