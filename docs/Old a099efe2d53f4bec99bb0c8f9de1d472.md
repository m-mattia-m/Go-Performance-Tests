# Old

## ID / Int

```go
var users []User
id := c.Param("id")

i := sort.Search(len(users), func(i int) bool { return id <= users[i].Id })
if i < len(users) && users[i].Id == id {
	c.JSON(200, users[i])
} else {
	c.JSON(400, gin.H{"error": "No user found with the id: " + id})
}
```