## how to use

### add note
```
curl -i -X POST -H "Content-Type: application/json" -d '{"text": "first note"}' http://localhost:8080/note/
curl -i -X POST -H "Content-Type: application/json" -d '{"text": "second note"}' http://localhost:8080/note/
curl -i -X POST -H "Content-Type: application/json" -d '{"text": "third note"}' http://localhost:8080/note/
```

### get all notes
```
curl -i -X GET http://localhost:8080/note/ 
```

### get note by id 
```
curl -i -X GET http://localhost:8080/note/<id>
```

### get first or last note
```
curl -i -X GET http://localhost:8080/note/<last or first>
```

### delete all notes
```
curl -i -X DELETE http://localhost:8080/note/
```

### delete note by id
```
curl -i -X DELETE http://localhost:8080/note/<id>
```

