#!/bin/bash
#
# swagger.sh
#
# Copyright (c) 2019 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#

# Generate Sia client with Swagger
swagger generate client -f https://raw.githubusercontent.com/jkawamoto/sia-swagger/master/swagger.yaml -t .

# Fix the generated code
sed -i -e "s/Message \\*string/Message string/" ./models/standard_error.go