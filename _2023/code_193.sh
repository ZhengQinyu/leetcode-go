#!/bin/bash
# 193. 有效电话号码
grep -E "^((\([0-9]{3}\)\ )|([0-9]{3}-))[0-9]{3}-[0-9]{4}$" file.txt