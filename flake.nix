{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    # cachix.url = "github:cachix/cachix";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
    templ = {
      url = "github:a-h/templ";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        gomod2nix.follows = "gomod2nix";
      };
    };
  };
  outputs =
    {
      self,
      flake-utils,
      nixpkgs,
      templ,
      gomod2nix,
    # cachix,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            gomod2nix.overlays.default
            # cachix.overlay.default
          ];
        };
        homepage = pkgs.buildGoApplication {
          name = "homepage";
          src = ./.;
          modules = ./gomod2nix.toml;
          preBuild = ''
            ${templ.packages.${system}.templ}/bin/templ generate
            ${pkgs.tailwindcss}/bin/tailwindcss -i ./cmd/web/assets/css/input.css -o ./cmd/web/assets/css/output.css --minify
          '';
        };
      in
      rec {
        formatter = pkgs.nixpkgs-fmt;
        packages.default = homepage;
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            go-tools
            air
            gomod2nix.packages.${system}.default
            templ.packages.${system}.templ
            tailwindcss
            docker-compose
          ];
        };
        packages.container = pkgs.dockerTools.buildImage {
          name = "ldufour/goserver";
          tag = "latest";
          created = "now";
          copyToRoot = pkgs.buildEnv {
            name = "image-root";
            paths = [ packages.default ];
            pathsToLink = [ "/bin" ];
          };
          config = {
            Cmd = [ "${packages.default}/bin/api" ]; # Adjust the binary name if necessary
          };
        };
        apps = {
          default = flake-utils.lib.mkApp { drv = homepage; };
          run = flake-utils.lib.mkApp {
            drv = pkgs.writeShellScriptBin "run" ''
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
              ${pkgs.air}/bin/air
            '';
          };
        };
      }
    );
}
