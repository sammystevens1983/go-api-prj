#!/bin/bash
curl -X POST http://localhost:5000/upload-file -F "file=@received_large_file.txt" -F "destDir=uploads" -F "destFilename=new_file.txt"