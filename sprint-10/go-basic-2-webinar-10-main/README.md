# go-basic-2-webinar-10

Linux

IP сервера 194.190.152.152

apt install nginx
apt install golang-1.21-go
apt install certbot python3-certbot-nginx

git clone git@github.com:vmarunin/go-basic-2-webinar-10.git

### Screen
startup_message off
defscrollback 102400
hardstatus alwayslastline "%H %w"

screen -t root 0 sudo su -
screen -t u 1

### .ssh/config
Host *
        ForwardAgent yes
Host srv
        User vmarunin
        HostName 194.190.152.152
        IdentityFile ~/.ssh/id_rsa

### Service

sudo cp config/nginx/vmarunin2.viewdns.net.conf /etc/nginx/conf.d/

sudo cp config/systemd/simpleback.service /etc/systemd/system/

sudo cp config/systemd/simpleback.service /etc/systemd/system/

systemctl status simpleback

journalctl -feu simpleback.service

### Domain

https://www.noip.com/members/dns/

Работает http://vmarunin2.viewdns.net/hello/

https://letsencrypt.org/

sudo certbot --nginx -d vmarunin2.viewdns.net

## Processes

* 'D' = UNINTERRUPTABLE_SLEEP
* 'R' = RUNNING & RUNNABLE
* 'S' = INTERRRUPTABLE_SLEEP
* 'T' = STOPPED
* 'Z' = ZOMBIE

sleep 20, kill, echo $? = 137 = 128+9

## Files

yes | head -n 2000000000 > big.txt
lsof -p

## iptables

iptables -L
iptables -A -I INPUT ! -i lo -p tcp --dport 8080 -j DROP

iptables-persistent

tcpdump -i lo -A -n -tt port 8080

