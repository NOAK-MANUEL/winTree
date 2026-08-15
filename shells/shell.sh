# wintree shell integration
wintree-cd() {
    local dir

    dir="$(command wintree cd "$@")"

    if [[ -d "$dir" ]]; then
        cd -- "$dir"
    fi
}