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

// INSTALAR DOCKER
sudo apt update && sudo apt upgrade -y

sudo install -m 0755 -d /etc/apt/keyrings

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

sudo chmod a+r /etc/apt/keyrings/docker.gpg

echo \
"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
$(. /etc/os-release && echo $VERSION_CODENAME) stable" | \
sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update

sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

sudo usermod -aG docker $USER

newgrp docker

sudo service docker start

sudo service docker start

// INSTALAR NGINX
sudo apt install nginx -y

// VER LA IP
172.29.75.224

// PUERTOS UFW
sudo apt install ufw -y
sudo ufw status
sudo ufw enable

sudo ufw allow 80 443 8080 22
sudo ufw allow "Nginx Full"
sudo ufw status numbered

// FILEZILA
sudo apt update
sudo apt install openssh-server -y

docker-compose up --build
docker-compose down -v


docker run -d -p 8000:8000 -p 9443:9443 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce:lts

https://localhost:9443
