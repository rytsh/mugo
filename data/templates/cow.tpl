{{- $value := (default "Moooo!" .) }}
{{- $maxLen := (int (add (len $value) (int64 2))) -}}

/{{repeat $maxLen "-"}}\
| {{$value | printf "%s"}} |
\{{repeat $maxLen "-"}}/
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
