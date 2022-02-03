# slate-alexa-proxy

## Deployment

1. cd into this directory

2. Build

> sudo docker build -f ./cloud.Dockerfile --tag slate-cloud ../../.

4. Initialize 'lets encrypt'. YOU SHOULDN'T REALLY NEED TO DO THIS, SEE LETSENCRYPT! SECTION

> chmod +x ./init-letsencrypt.sh
>
> sudo ./init-letsencrypt.sh

5. Run compose, on rerun this will build the newly changed containers and redeploy

> Run sudo docker-compose up -d

## LETSENCRYPT!

Let'sEncrypt! is a free service that provides https tls certs that we need to run https. We need to run https, because
modern browsers will treat you like the plague if you don't...

You will need to edit the line `domains_list` if new subdomains, or actual domains are added to the NGNIX reverse proxy duties.
The variable is a space-delimited list of fully qualified domain names.

You should only need to run the `init-letsencrypt.sh` script once when updating this list. Make sure you get things right, or
else letsencrypt will get mad at you.

When rerunning `init-letsencrypt.sh`, if it already has a cert for a domain it'll let you know and ask if it should create a new
one to replace the old one. DO NOT CREATE A NEW ONE, see the last paragraph on why not :D

We run encryptbot in the compose script, which takes care of automatically updating the certs; you shouldn't need to do
anything on that front.
