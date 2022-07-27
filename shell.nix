{ pkgs ? import <nixpkgs> {}
}:
pkgs.mkShell {
    name="dev";
    buildInputs = [
        pkgs.go
    ];
    shellHook = ''
        export SEGFAUTILITIES_PORT=6893
        echo "Go installed, have your fun"
    '';
}