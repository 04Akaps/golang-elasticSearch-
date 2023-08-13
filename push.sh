#!/bin/bash

branch="main"

read -p "커밋 메시지를 입력하세요: " commit_message

echo add all changed File
git add .

echo commit File
git commit -m "$commit_message"


echo push All File
git push me $branch
