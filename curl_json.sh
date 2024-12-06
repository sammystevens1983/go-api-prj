#!/bin/bash
curl -X POST http://localhost:5000/save-json \
-H "Content-Type: application/json" \
-d '{"message": "Hello, World!", "number": 42}'