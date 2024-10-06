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
        pname = "raindrop-images-dl";
        version = "0.1.0";
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

        packages.default = pkgs.stdenv.mkDerivation {
          pname = "${pname}";
          version = "${version}";
          src = ./.;
          buildInputs = [ pkgs.go ];

          buildPhase = ''
            # this line removes a bug where value of $HOME is set to a non-writable /homeless-shelter dir
            export HOME=$(pwd)
            export CGO_ENABLED=0
            go build -o $out main.go
          '';

          meta = with pkgs.lib; {
            homepage = "https://github.com/brpaz/raindrop-images-dl";
            description = "Raindrop Images Downloader";
            platforms = platforms.all;
            license = licenses.mit;
            maintainers = with maintainers; [ brpaz ];
          };
        };
    });
}
