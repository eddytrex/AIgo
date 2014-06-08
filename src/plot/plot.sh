#!/bin/bash

# note: you need GNU plot

for FILE in *.csv; do 
  ncolumns=`awk -F ',' '{print NF}' f0_layer.csv |head -n 1`
  #echo "$ncolumns"
  c=""
  for i in $(seq 2 $ncolumns); do
     a=`echo " \"${FILE}\" using 1:${i} with lines,"`
     c="${c}${a}"
  done    
    
  gnuplot <<__EOF
    set datafile sep ","
    set term png
    set output "${FILE}.png"
    plot ${c}
__EOF
  c=""
done
