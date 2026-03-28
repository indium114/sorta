{
  description = "sorta devshell and package";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          name = "sorta-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
            just
          ];
        };

        packages.sorta = pkgs.buildGoModule {
          pname = "sorta";
          version = "2026.03.01-a";

          src = self;

          vendorHash = "sha256-aJllcMJduoi8VBWMJWsxm8swXtNonYZzX8etmNZePzc=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "An extensible search tool, inspired by Raycast and Vicinae";
            license = licenses.mit;
            platforms = platforms.all;
          };
        };

        apps.sorta = {
          type = "app";
          program = "${self.packages.${system}.sorta}/bin/sorta";
        };
      });
}
