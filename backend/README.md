URL_Frontend=
ServerDB=
PortDB=
UserDB=
PasswordBD=
NameDB=
PortServer=
Secret_Key=
Source_Path=
IMAGEN=
PDF=
Url=

docker-compose down -v
docker-compose up --build

docker run -d -p 8000:8000 -p 9443:9443 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce:lts

https://localhost:9443
