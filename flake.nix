{
  description = "Devenv";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.gotools
            pkgs.gotestsum
            pkgs.golangci-lint
            pkgs.go-task
            pkgs.gomarkdoc
            pkgs.lefthook
            pkgs.direnv
            pkgs.hadolint
            pkgs.delve
            pkgs.commitlint
            pkgs.goreleaser
          ];

          shellHook = ''
            export CGO_ENABLED=0
            export GOROOT=${pkgs.go}/share/go
            lefthook install
            go mod download
          '';

          # Optional: setup Go environment variables
          GO111MODULE = "on";
          GOROOT = "${pkgs.go}/share/go";
        };
    });
}
