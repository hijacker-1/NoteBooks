# asyncData函数

asyncData函数主要用来在页面组件加载之前向后端请求数据，然后在页面加载之前进行渲染

> 注意： asyncData函数只能在页面和子页面组件当中使用，而不能在components中的组件使用

## asyncData发送get请求

```js
/*asyncData函数中的参数使用context，而此处是把axios从context中解构出来，当然也可以添加别的参数，如route
*/
async asyncData({$axios}) {
  //async搭配await使用，此处get方法的参数只是一个url，也就是请求地址
   let res = await $axios.get('http://localhost:3000/data.json')
   console.log(res.data.data.bookList)
   return {bookList: res.data.data.bookList}
 }

async asyncData({$axios, route}) {
  let id = route.query.id
  // 这里是asyncData携带参数对后端进行访问，所有携带的参数添加到params中
  await $axios.get('http://localhost:8080', {
    params: {
      id
    }
  }).then((res) => {
    return {bookList: res.data.data.bookList}
  })
}
```

## asyncData发送post请求
 ```js
 async asyncData({$axios}) {
   await $axios.post('http:localhost:8080', {
     data: {
       id = this.id
       username = this.username
       password = this.password
     }
   }).then((response) => {
     // 对返回结果进行处理
   })
 }
