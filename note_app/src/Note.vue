<template>
<div class="note" v-if="note">
  <h1 class="note__title">{{  note.title }}</h1>
  <p class="note__body">{{ note.body }}</p>
  <p class="note__id">{{ note.id }}</p>
</div>  
</template>

<script>
import {HTTP} from './http-constants'

export default {
  props: ['id'],
  data() {
    return {
      note: null,
      endpoint:'/api/v1/notes/'
    };
  },
  methods: {
    getPost(id) {
      HTTP.get(this.endpoint+id).then(response => {
        this.note = response.data
      }).catch(error => {
          console.log('----- Error -----')
          console.log(error)
      })
    }
  },
  created() {
    this.getPost(this.id);
  },
  watch: {
    '$route'() {
      this.getPost(this.id);
    }
  }
};
</script>

<style></style>
