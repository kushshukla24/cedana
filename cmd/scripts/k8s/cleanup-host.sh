#!/bin/bash

chroot /host /bin/bash <<"EOT"
./cedana/reset.sh

rm -rf /cedana

rm -rf /criu

rm -f /usr/local/bin/cedana
rm -f /build-start-daemon.sh
EOT

echo "Clean up completed."
