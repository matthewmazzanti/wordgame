{ pkgs ? import <nixpkgs> {}, ... }:
with pkgs;
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    compile-daemon
    nodejs_latest
    docker-compose
    parallel
    mysql
    (python3.withPackages (pkgs: with pkgs; [
      mysql-connector
    ]))
  ];
}
