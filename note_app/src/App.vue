<template>
<div id="app">
  <header>
    <h1>Note SPA</h1>
  </header>
  <main>
    <aside class="sidebar">
      <router-link
          v-for="note in notes"
          active-class="is-active"
          class="link"
          :to="{ name: 'note', params: { id: note.id }}"
          :key="note.id">
        {{ note.id }}. {{ note.title }}
      </router-link>
    </aside>
    <div class="content">
      <router-view></router-view>
    </div>
  </main>
</div>
</template>

<script>
import {HTTP} from './http-constants'

export default {
  data() {
    return {
      notes: null,
      endpoint:'/api/v1/notes'
    };
  },
  created() {
    this.getAllNotes();
  },
  methods: {
    getAllNotes() {
      HTTP
        .get(this.endpoint)
        .then(response => {
          this.notes = response.data
        })
        .catch(error => {
          console.log('----- Error -----')
          console.log(error)
        });
    }
  }
};
</script>

<style>

</style>
