{{ exec "echo -n 'Hello World'" | map.Get "stdout" | codec.ByteToString }}
