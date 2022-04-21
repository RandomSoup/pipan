{ pkgs }: {
    deps = [
        pkgs.go
        pkgs.xorg.libX11
        pkgs.gtk3-x11
    ];
}