{ pkgs ? import ../../../nix { } }:
let qied = (pkgs.callPackage ../../../. { });
in
qied.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-qied.patch
  ];
})
