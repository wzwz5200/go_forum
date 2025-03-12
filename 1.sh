curl -X GET http://localhost:3000/api/posts/createpost \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your_jwt_token_here" \
-d '{
    "title": "测试帖子标题",
    "content": "这里是详细的测试内容...",
    "section_id": 1
}'