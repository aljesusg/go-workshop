curl -X GET localhost:8000/students -H "type: json"
curl -X POST localhost:8000/student/3 -H "type: json" -d '{"name":"Alberto"}'
curl -X DELETE localhost:8000/student/3 -H "type: json"
