#!/bin/bash

cd $1; 
find *.ts | sort -n | sed 's:\\ :\\\\\\ :g'| sed 's/^/file /' > fl.txt; 
ffmpeg -f concat -i fl.txt -c copy name.mp4; 
rm fl.txt *.ts
