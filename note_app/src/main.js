import Vue from 'vue'
import Router from 'vue-router'
import App from './App.vue'
import Index from './Index.vue'
import Note from './Note.vue'
Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/note/:id',
      name: 'note',
      component: Note,
      props: true
    }
  ]
})

new Vue({
  el: '#app',
  render: h => h(App),
  router
})
