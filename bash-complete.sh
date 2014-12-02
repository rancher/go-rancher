_rancher-bash-complete ()  
{  
    local cur

    COMPREPLY=()   # Array variable storing the possible completions.
    cur=${COMP_WORDS[COMP_CWORD]}

    case "$cur" in
    -*)
        COMPREPLY=( $( compgen -W '--url --accessKey --format' -- $cur ) );;
    esac

    return 0
}

complete -F _rancher-bash-complete -o filenames go-rancher
