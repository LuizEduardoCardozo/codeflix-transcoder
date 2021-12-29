#/usr/bin/env sh

git clone https://github.com/axiomatic-systems/Bento4.git /tmp/bento4 && cd /tmp/bento4 && \
    mkdir build && cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release .. && \
    make

ln -s /tmp/bento4/build/mp4fragment /usr/bin/mp4fragment
ln -s /tmp/bento4/build/mp4info /usr/bin/mp4info
ln -s /tmp/bento4/build/mp4dump /usr/bin/mp4dump

echo -e "#/bin/bash\npython3 /tmp/bento4/Source/Python/utils/mp4-dash.py \$@" > /usr/bin/mp4dash && chmod +x /usr/bin/mp4dash

mp4fragment

if [[ $? != 1 ]]; then
    exit 1
else
    exit 0
fi;
