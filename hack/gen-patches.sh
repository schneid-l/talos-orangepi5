#!/usr/bin/env sh

set -e

REPO_DIR=$1
PATCH_DIR=$2

if [ -z "$REPO_DIR" ] || [ -z "$PATCH_DIR" ]; then
    echo "Usage: $0 <repo-dir> <patch-dir>"
    exit 1
fi

if [ ! -d "$REPO_DIR" ]; then
    echo "Repository directory does not exist: $REPO_DIR"
    exit 1
fi

if [ -d "$PATCH_DIR" ]; then
    echo "Patch directory already exists: $PATCH_DIR"
    read -p "Do you want to erase it? [y/N] " -r
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -rf "$PATCH_DIR"
    else
        echo "Aborting..."
        exit 1
    fi
fi
mkdir -p "$PATCH_DIR"

git -C "$REPO_DIR" format-patch --stdout HEAD^1 > "$PATCH_DIR/talos-orangepi5.patch"
