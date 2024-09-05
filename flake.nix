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
          name = "api";
          src = ./.;
          buildInputs = [ pkgs.sqlcipher ];
          modules = ./gomod2nix.toml;
          preBuild = ''
            ${templ.packages.${system}.templ}/bin/templ generate
            ${pkgs.tailwindcss}/bin/tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify
          '';
        };
      in
      rec {
        formatter = pkgs.nixpkgs-fmt;
        packages.default = homepage;
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            jq
            go
            go-tools
            air
            gomod2nix.packages.${system}.default
            templ.packages.${system}.templ
            tailwindcss
            postgresql
            docker-compose
          ];
        };
        packages.container = pkgs.dockerTools.buildLayeredImage {
          name = "ldufour/goserver";
          tag = "latest";
          contents = [
            packages.default
            pkgs.postgresql
            pkgs.cacert
          ];
          config = {
            Cmd = [ "${packages.default}/bin/api" ];
            WorkingDir = "/";
            Volumes = {
              "/data" = { };
            };
            Env = [
              "IN_CONTAINER=true"
              "DB_HOST=localhost" # Added for PostgreSQL connection
              "DB_PORT=5433" # Added for PostgreSQL connection
              "DB_NAME=mydatabase" # Added for PostgreSQL connection
              "DB_USER=myuser" # Added for PostgreSQL connection
              "DB_PASSWORD=mypassword" # Added for PostgreSQL connection
            ];
          };
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
      }
    );
}
