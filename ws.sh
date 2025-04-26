ws () {
    GITWS_CONFIG_DIR=~/.config/gitws
    if ! [ -d ${GITWS_CONFIG_DIR} ]; then 
        mkdir -p ${GITWS_CONFIG_DIR}
    fi
    
    export GITWS_TMP_CMD="/tmp/.gitws-$(hexdump -e '/1 "%02x"' -n4 < /dev/random).tmp"
    ${GITWS_CONFIG_DIR}/ws-cmd "$@"
    if [ -f "${GITWS_TMP_CMD}" ]; then
        . "${GITWS_TMP_CMD}"
    fi
}
