# 动态路由
## 动态路由的定义

动态路由需要在文件名或者文件夹的前面（pages文件夹下的）加上下划线（_），名称可以随意起，符合规则就行

## 动态路由的使用

```html
index.vue
<template>
    <div>
        <ui>
            <li v-for="item in stuList" :key="item.id">
            <!--  采用以下两种方式在页面中使用动态路由 -->
                <nuxt-link :to="`/students/${item.id}`">this is the {{item.id}}</nuxt-link>
                <nuxt-link :to="{name: "students-id", params: {id: item.id}, query: {id: item.id}}">this is the {{item.id}} pages </nuxt-link>
            </li>
        </ui>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
            stuList: [
                {name: 'jack1', id: 1},
                {name: 'jack2', id: 2},
                {name: 'jack3', id: 3},
                {name: 'jack4', id: 4},
                {name: 'jack5', id: 5},
            ]
        }
    }
}
</script>
```

## 动态路由参数的使用

```html
<template>
    <div>
    <!-- 在template中使用¥route.params.id取出从上个页面携带过来的参数，可以不用加this，但如果不是在template中使用就要加this -->
        this is the {{ this.$route.params.id }} student's information
    </div>
</template>

<script>
export default {
  // validate对动态路由携带过来对参数进行校验
    validate(context) {
        return /^\d+$/.test(context.params.id)
    }
}
</script>
