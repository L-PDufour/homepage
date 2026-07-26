{
  config,
  lib,
  pkgs,
  ...
}:
let
  cfg = config.services.homepage;
in
{
  options.services.homepage = {
    enable = lib.mkEnableOption "homepage web service";

    package = lib.mkOption {
      type = lib.types.package;
      description = "The homepage package to run.";
    };

    port = lib.mkOption {
      type = lib.types.port;
      default = 8003;
    };

    environment = lib.mkOption {
      type = lib.types.attrsOf lib.types.str;
      default = { };
      example = {
        CF_TEAM_DOMAIN = "lpdufour";
        GO_ENV = "production";
      };
      description = "Extra non-secret environment variables.";
    };

    secrets = lib.mkOption {
      type = lib.types.attrsOf lib.types.path;
      default = { };
      example = {
        CF_AUD_TAG = "/run/agenix/cf-aud-tag";
        CF_API_KEY = "/run/agenix/cf-api-key";
      };
      description = ''
        Env var name -> path to a file containing just that secret's raw
        value. Loaded via systemd LoadCredential, never joined into a
        plaintext file on disk.
      '';
    };

    environmentFile = lib.mkOption {
      type = lib.types.nullOr lib.types.path;
      default = null;
      description = ''
        Optional extra EnvironmentFile (KEY=VALUE lines), e.g. from
        sops-nix/agenix if you'd rather manage secrets that way instead
        of (or alongside) `secrets`.
      '';
    };
  };

  config = lib.mkIf cfg.enable {
    systemd.services.homepage = {
      description = "Homepage Go binary";
      wantedBy = [ "multi-user.target" ];
      after = [ "network.target" ];

      environment = cfg.environment // {
        PORT = toString cfg.port;
      };

      serviceConfig = {
        ExecStart = pkgs.writeShellScript "homepage-start" ''
          set -euo pipefail
          ${lib.concatStrings (
            lib.mapAttrsToList (name: _: ''
              export ${name}="$(cat "$CREDENTIALS_DIRECTORY/${name}")"
            '') cfg.secrets
          )}
          exec ${cfg.package}/bin/api
        '';
        LoadCredential = lib.mapAttrsToList (name: path: "${name}:${path}") cfg.secrets;

        DynamicUser = true;
        StateDirectory = "homepage";
        WorkingDirectory = "/var/lib/homepage";
        Restart = "on-failure";

        NoNewPrivileges = true;
        PrivateTmp = true;
        ProtectSystem = "strict";
        ProtectHome = true;
        ProtectKernelTunables = true;
        ProtectKernelModules = true;
        ProtectControlGroups = true;
        RestrictSUIDSGID = true;
        LockPersonality = true;
        RestrictRealtime = true;
      }
      // lib.optionalAttrs (cfg.environmentFile != null) {
        EnvironmentFile = cfg.environmentFile;
      };
    };
  };
}
