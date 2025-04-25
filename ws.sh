ws () {
    WS_CONFIG_DIR=~/.config/gitws
    ws-cmd "$@"
    if [ -f "${WS_CONFIG_DIR}/ws-post" ]; then
        . "${WS_CONFIG_DIR}/ws-post"
    fi
}
