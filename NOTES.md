OSX
sudo easy_install pycap
#xgo - TODO trouble with pcap.h

In order to use cgo when cross-compiling, you need to set the CC environment variable to a Darwin -> GNU/Linux cross-compiler. You are using a Darwin native compiler, which can not work. Even if the build succeeded, the resulting code would not run on GNU/Linux.


#####################
LINUX
sudo apt-get update
sudo apt-get install build-essential
sudo apt-get install libpcap0.8-dev

sudo apt-get install tcpdump

# change permissions - not working 
sudo groupadd pcap
sudo usermod -a -G pcap user
sudo chgrp pcap /usr/sbin/tcpdump
sudo chmod 750 /usr/sbin/tcpdump

# can go install w/ CGO_ENABLED=1 go install