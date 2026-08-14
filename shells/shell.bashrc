# wintree shell integration
wintree-cd() {
    local dir

    dir="$(command winTree cd "$@")"

    if [[ -d "$dir" ]]; then
        cd -- "$dir"
    fi
}