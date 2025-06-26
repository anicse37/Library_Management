#!/bin/bash

echo "Hey Lazzzzzy"

git status 

read  -p "Enter yout commit: " msg

git add .
git status



git commit -m "$msg"



git push
