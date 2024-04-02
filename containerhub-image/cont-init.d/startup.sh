#!/command/with-contenv sh

echo $CONTAINERHUB_USER >/etc/ssh/ssh_principals
echo $CONTAINERHUB_CA_PUB_KEY | base64 -d >/etc/ssh/ssh_ca.pub
