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
        gitCommit = self.rev or "dev";

        currentDate = "2021-09-01T00:00:00Z";
        #currentDate = builtins.currentTime;
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

        packages.default = pkgs.buildGo122Module {
          pname = "${pname}";
          version = "${version}";
          src = ./.;

          # Override phases to define the command line entry point
          subPackages = [ "cmd" ];

          # When updating go.mod or go.sum, a new sha will need to be calculated,
          # update this if you have a mismatch after doing a change to thos files.
          vendorHash = "sha256-BNfSwK87GU2YO3x1AbxLwC+ByXkN4n/7OYX5mh04lP4=";

          doCheck = false;

          ldflags = let
            versionPkg = "github.com/brpaz/raindrop-images-dl/internal/version";
          in [
            "-X ${versionPkg}.Version=${version}"
            "-X ${versionPkg}.GitCommit=${gitCommit}"
            "-X ${versionPkg}.BuildDate=${currentDate}"
          ];

          postInstall = ''
            mv $out/bin/cmd $out/bin/raindrop-images-dl
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
