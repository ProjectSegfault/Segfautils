{ pkgs ? import <nixpkgs> {}
}:
pkgs.mkShell {
    name="dev";
    buildInputs = [
        pkgs.go
    ];
    shellHook = ''
        echo "Go installed, have your fun"
    '';
}