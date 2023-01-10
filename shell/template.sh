#!/bin/bash

readonly OK=0
readonly NONOK=1

CONFIG_FILE='/var/lib/kubelet/config.yaml'

function update_config() {
    if [ -f "${CONFIG_FILE}" ]; then
    	sed -i "s/cpu: 500m/cpu: 50m/g" ${CONFIG_FILE}
    	if [[ $? -ne 0 ]]; then
        	echo 'update config failed'
        	exit $NONOK
    	fi
    	cnt=$(cat ${CONFIG_FILE} | grep 'cpu: 50m' | wc -l)
    	if [[ $cnt -ne 2 ]]; then
        	echo 'match key word failed'
        	exit $NONOK
    	fi
    	return
    fi
    grep -E 'kube-reserved=cpu=[[:digit:]]*m' ${SERVICE_CONFIG}
    if [[ $? -eq 0 ]]; then
        sed -i 's/kube-reserved=cpu=[0-9]*m/kube-reserved=cpu=50m/g' ${SERVICE_CONFIG}
        sed -i 's/system-reserved=cpu=[0-9]*m/system-reserved=cpu=50m/g' ${SERVICE_CONFIG}
        return
    fi
    exit $NONOK   
}


function main() {
    update_config
}

main "$@"
