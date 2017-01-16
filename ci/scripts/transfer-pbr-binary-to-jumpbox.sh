#!/bin/bash

set -eu

./pcf-backup-and-restore-meta/unlock-ci.sh
chmod 400 pcf-backup-and-restore-meta/genesis-bosh/bosh.pem

bosh -n -t ${BOSH_TARGET} -u ${BOSH_CLIENT} -p ${BOSH_CLIENT_SECRET} \
  -d pcf-backup-and-restore-meta/deployments/acceptance-jump-box.yml \
  ssh --gateway_identity_file pcf-backup-and-restore-meta/genesis-bosh/bosh.pem \
  --gateway_user vcap --gateway_host genesis-bosh.backup-and-restore.cf-app.com \
  jump-box 0 \
  "sudo mkdir -p /var/vcap/store/pbr && \
   sudo chmod 775 /var/vcap/store/pbr && \
   sudo chown vcap:vcap /var/vcap/store/pbr
  "

ls release/pbr* | xargs -INAME bosh -n -t ${BOSH_TARGET} -u ${BOSH_CLIENT} -p ${BOSH_CLIENT_SECRET} \
  -d pcf-backup-and-restore-meta/deployments/acceptance-jump-box.yml \
  scp jump-box 0 NAME /var/vcap/store/pbr/ \
  --upload --gateway_identity_file pcf-backup-and-restore-meta/genesis-bosh/bosh.pem \
  --gateway_user vcap --gateway_host genesis-bosh.backup-and-restore.cf-app.com

ls release/pbr* | xargs -INAME basename NAME | rev | cut -d "." -f2- | rev | \
  xargs -INAME bosh -n -t ${BOSH_TARGET} -u ${BOSH_CLIENT} -p ${BOSH_CLIENT_SECRET} \
  -d pcf-backup-and-restore-meta/deployments/acceptance-jump-box.yml \
  ssh --gateway_identity_file pcf-backup-and-restore-meta/genesis-bosh/bosh.pem \
  --gateway_user vcap --gateway_host genesis-bosh.backup-and-restore.cf-app.com \
  jump-box 0 \
  "sudo chpst -u vcap:vcap mkdir -p /var/vcap/store/pbr/NAME && \
   sudo chpst -u vcap:vcap tar xvf /var/vcap/store/pbr/NAME.tar -C /var/vcap/store/pbr/NAME/ --strip-components 1 && \
   sudo rm -f /var/vcap/store/pbr/NAME.tar
  "