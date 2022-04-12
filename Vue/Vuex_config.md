# Vuex状态树共享属性配置

## store文件夹下的配置
```js
// 需要全局共享的数据全都放在state里
// state最终返回一个方法,  此处使用箭头函数简化代码书写
export const state = () => ({
    count: 0, // count就是需要全局共享的属性，可以在这里赋一个初始值
})

// 用来修改state中全局数据的方法
export const mutations = {
  // add方法第一个参数一定是上一部分中的state
    add(state) {
        state.count++ // 函数体中对state中的属性进行操作
    },
    del(state) {
        state.count--
    },

    // 这里面写的函数第一个参数永远是state， 后面的就是从其他地方传过来的参数
    addN(state, step) {
        state.count += step // step是由用户传过来的参数
}
}

// actions中进行的是异步操作
export const actions = {
    addAsync(context) {
        setTimeout(()=> {
            // 在actions中不能直接对state进行更改，必须先用commit调用mutatinos中的方法
            context.commit('add') // 第一个参数写mutations中想要调用的函数名，第二个参数写用户传过来的参数
        }, 1000)
    }
}
```

## 在组件中的使用

```js
<template>
    <div>
        <h2>the value of count is {{this.$store.state.count}}</h2> // 采用这种方法取出vuex状态树中的全局属性
        <button @click="handle1">+1</button>
    </div>
</template>

<script>
import {mapState, mapMutations} from 'vuex' // 从vuex中导入这两个东西
export default {
    data: function() {
        return {

        }
    },

    computed: {
        ...mapState(['count']) // 使用展开运算符将状态树中的属性展开到计算属性中，就可以像使用data属性一样进行使用
    },
    methods: {
        handle1() {
          // 在组件方法中用this.$store.commit调用mutations中的方法
            this.$store.commit('add');
        }
        handle2() {
            this.del()
        },
        handle3() {
            // commit的作用就是触发一个mutation
            // this.$store.commit('addN', 4)
            // 这里的dispatch函数就是用来触发actions中的方法
            this.$store.dispatch('addAsync')
        },
        ...mapMutations(['addN','add', 'del']) // 使用展开运算符将mutations中的方法展开到methods中，就可以直接调用
    }
}
</script>
```
