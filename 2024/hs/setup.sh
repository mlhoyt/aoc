#!/usr/bin/env sh

nix shell \
   nixpkgs#ghc \
   nixpkgs#haskellPackages.cabal-install \
   nixpkgs#haskellPackages.haskell-language-server \
   nixpkgs#haskellPackages.fourmolu \
   nixpkgs#haskellPackages.hlint \
   nixpkgs#haskellPackages.stan
