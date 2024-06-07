#!/bin/sh
 
IPT="/sbin/iptables"
IFACE_EXT="ens3"
IFACE_LOC="lo"
 
# Flushing iptable rules.
$IPT -F
$IPT -t nat -F
$IPT -t mangle -F
$IPT -X
$IPT -t nat -X
$IPT -t mangle -X
 
# Default politics
$IPT -P INPUT   DROP
$IPT -P FORWARD DROP
$IPT -P OUTPUT  ACCEPT

# Loopback interface 
$IPT -A INPUT  -i $IFACE_LOC -j ACCEPT
# SSH
$IPT -A INPUT -p tcp -m tcp -i $IFACE_EXT --dport 22 -j ACCEPT
# ICMP: разрешить все
$IPT -A INPUT -p icmp -i $IFACE_EXT -j ACCEPT
 
# Это правило обязательно если INPUT DROP.
# Разрешаем прохождение statefull-пакетов. Эта цепочка обязательная в любых настройках iptables,
# она разрешает прохождение пакетов в уже установленных 
# соединениях(ESTABLISHED), и на установление новых соединений от уже установленных (RELATED).
$IPT -A INPUT -i $IFACE_EXT -m state --state ESTABLISHED,RELATED -j ACCEPT

# allow nginx
iptables -A INPUT -p tcp -m multiport --destination-port 80,443 -j ACCEPT
# протколирование последнего правила в этом примере -P INPUT   DROP
$IPT -A INPUT -j LOG --log-level INFO --log-prefix "-P INPUT DROP: "
 