# 安装Consul
apt update; apt install consul

# 安装Fabio
mkdir /opt/fabio
export PATH=$PATH:/opt/fabio  #TODO
cd /opt/fabio
touch fabio.properties
wget https://githubfast.com/fabiolb/fabio/releases/download/v1.6.3/fabio-1.6.3-linux_amd64
chmod a+x fabio-1.6.3-linux_amd64
ln -s fabio-1.6.3-linux_amd64 fabio

# 启动Consul
consul agent -dev -ui -node=consul-dev -client=0.0.0.0

# 启动Fabio
/opt/fabio/fabio
