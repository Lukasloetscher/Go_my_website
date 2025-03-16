#!/bin/bash

echo $(date -u)

apt update

apt upgrade -y

reboot now