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
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ gomod2nix.overlays.default ];
        };
        homepage = pkgs.buildGoApplication {
          name = "homepage";
          src = ./.;
          modules = ./gomod2nix.toml;
          preBuild = ''
            ${templ.packages.${system}.templ}/bin/templ generate
            ${pkgs.tailwindcss}/bin/tailwindcss -i ./cmd/web/assets/css/input.css -o $TMPDIR/output.css --minify
            mkdir -p $out/cmd/web/assets/css
            cp $TMPDIR/output.css $out/cmd/web/assets/css/
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
        packages.container = pkgs.dockerTools.buildLayeredImage {
          name = "ldufour/goserver";
          tag = "latest";
          contents = [
            packages.default
            (pkgs.runCommand "assets" { } ''
              mkdir -p $out/cmd/web/assets
              cp -R ${./cmd/web/assets}/* $out/cmd/web/assets/
              mkdir -p $out/cmd/web/assets/css
              cp ${packages.default}/cmd/web/assets/css/output.css $out/cmd/web/assets/css/ || true
              chmod -R 755 $out/cmd/web/assets
              cp ${./postcss.config.js} $out/postcss.config.js
              cp ${./tailwind.config.js} $out/tailwind.config.js
            '')
          ];
          config = {
            Cmd = [ "${packages.default}/bin/api" ];
            WorkingDir = "/";
          };
        };
        packages.containerWithCachix =
          pkgs.runCommand "container-with-cachix" { buildInputs = [ pkgs.cachix ]; }
            ''
              # Build the container
              container=$(nix-build -E '(import ./. {}).packages.${system}.container' --no-out-link)

              # Push the container to Cachix
              cachix push mycache $container

              # Create a symlink to the built container
              ln -s $container $out
            '';
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
