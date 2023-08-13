#!/bin/bash

remote="me"
branch="master"

read -p "커밋 메시지를 입력하세요: " commit_message

echo add all changed File
git add .

echo commit File
git commit -m "$commit_message"

echo push All File to $branch branch

if git ls-remote --exit-code --heads $remote $branch; then
    git push $remote $branch
    echo "Push successful!"
else
    echo "Error: Remote branch $branch does not exist."
fi
