# wintree shell integration
function wintree-cd
        set -l dir (command wintree cd $argv)

        if test -d "$dir"
            cd "$dir"
        end


end