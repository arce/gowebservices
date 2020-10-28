function test() {

service = fetchival
var posts = service('https://webservicegoapi.herokuapp.com/books')

//posts
posts.get()
posts.post({ title: 'Fetchival' })

//posts?category=javascript
posts.get({ category: 'javascript' })

//posts/1
posts(1).get()
posts(1).put({ title: 'Fetchival is simple' })
posts(1).patch({ title: 'Fetchival is simple' })
posts(1).delete()

var comments = posts('1/comments')

//posts/1/comments
comments.get()

//posts/1/comments/1
comments(1).get()
}