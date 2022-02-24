#!/usr/bin/env bash
echo 'Runing migrations...'
/app/migrate up
echo 'Start application...'
/app/main