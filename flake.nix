{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
  };
  outputs = { self, flake-utils, nixpkgs, gomod2nix, }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ gomod2nix.overlays.default ];
        };
        homepage = pkgs.buildGoApplication {
          name = "api";
          src = ./.;
          modules = ./gomod2nix.toml;
        };
      in rec {
        formatter = pkgs.nixpkgs-fmt;
        packages.default = homepage;
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            jq
            go
            gopls
            vscode-langservers-extracted
            go-tools
            air
            gomod2nix.packages.${system}.default
            sqlite
          ];
        };
        apps = {
          default = flake-utils.lib.mkApp { drv = homepage; };
          run = flake-utils.lib.mkApp {
            drv = pkgs.writeShellScriptBin "run" ''
              if [ -f .env ]; then
                export $(cat .env | xargs)
              else
                echo ".env file not found"
                exit 1
              fi
              ${pkgs.go}/bin/go run cmd/api/main.go
            '';
          };
          test = flake-utils.lib.mkApp {
            drv = pkgs.writeShellScriptBin "test" ''
              echo "Testing..."
              ${pkgs.go}/bin/go test ./tests -v
            '';
          };
          watch = flake-utils.lib.mkApp {
            drv = pkgs.writeShellScriptBin "watch" ''
              make watch
            '';
          };
        };
      });
}
