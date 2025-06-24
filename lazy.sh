#!/bin/bash

echo "Hey Lazzzzzy"

git status 
sleep 2
read  -p "Enter yout commit: " msg

git add .
git status
sleep 2


git commit -m "$msg"

sleep 2

git push
