<template>
<div id="app" class="mdc-typography">
  <header class="mdc-toolbar mdc-toolbar--fixed mdc-toolbar--waterfall">
    <div class="mdc-toolbar__row">
      <section class="mdc-toolbar__section mdc-toolbar__section--align-start">
        <span class="mdc-toolbar__title">MONOLOG</span>
      </section>
      <section class="mdc-toolbar__section mdc-toolbar__section--align-end">
        <a @click="logout" class="material-icons mdc-toolbar__icon" aria-label="Logout" alt="Logout" v-if="!isLoginView()">exit_to_app</a>
      </section>
    </div>
  </header>
  <main class="mdc-toolbar-fixed-adjust">
    <div class="container">
      <router-view/>
    </div>
  </main>
</div>
</template>

<script>
export default {
  name: 'app',
  methods: {
    logout: function () {
      this.$router.push('/login', () => localStorage.clear())
    },
    isLoginView: function () {
      console.log(this.$route.path)
      return this.$route.path === '/login'
    }
  }
}
</script>

<style>
:root {
  --mdc-theme-background: #F5F5F5;
  --mdc-theme-primary: #CDDC39;
}

html,body {
  background-color: var(--mdc-theme-background,#F5F5F5);
}

.mdc-typography {
  font-family: Roboto Mono, Roboto, sans-serif;
}

.container {
  max-width: 1280px;
  margin: 0 auto;
}

.left-align {
  text-align: left;
}

.right-align {
  text-align: right;
}

.canter-align {
  text-align: center;
}

.ellipsis {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
