#!/bin/bash

BRANCH=$(git branch --show-current)

echo "Current date : $(date)"
echo "User : $(whoami)"
echo "Current branch : $BRANCH"

if [ "$BRANCH" = "feature/versioning" ]; then
  git checkout main
else
  git checkout feature/versioning
fi

BRANCH_CHANGED=$(git branch --show-current)
echo "Branch changed from $BRANCH to $BRANCH_CHANGED"
