<template>
<div>
  <div class="mdc-toolbar-fixed-adjust">
    <component :is="currentView" :selectedLogs="selectedLogs" @selectLogStream="doSelectLogStreamEvent">
    </component>
  </div>
  <div class="mdc-toolbar mdc-toolbar--fixed mdc-toolbar--waterfall">
    <div class="mdc-toolbar__row">
      <section class="mdc-toolbar__section mdc-toolbar__section--align-start" v-show="isLogEventsView()">
        <button class="mdc-button mdc-button--stroked" @click="transitionLogIndexView">
          <i class="material-icons mdc-button__icon">arrow_back</i>Back to Index
        </button>
      </section>
      <section class="mdc-toolbar__section mdc-toolbar__section--align-end" v-show="isLogIndexView()">
        <button class="mdc-button mdc-button--stroked" @click="transitionLogEventsView" :disabled="selectedLogs.length === 0">
          <i class="material-icons mdc-button__icon">visibility</i>View Log Events
        </button>
      </section>
    </div>
  </div>
</div>
</template>

<script>
import LogIndex from './LogIndex'
import LogEvents from './LogEvents'

export default {
  name: 'LogDashboard',
  props: {},
  components: {
    LogIndex,
    LogEvents
  },
  data () {
    return {
      currentView: LogIndex,
      selectedLogs: []
    }
  },
  methods: {
    doSelectLogStreamEvent (val) {
      this.selectedLogs = val
    },
    isLogIndexView () {
      return this.currentView === LogIndex
    },
    isLogEventsView () {
      return this.currentView === LogEvents
    },
    transitionLogIndexView () {
      this.currentView = LogIndex
    },
    transitionLogEventsView () {
      this.currentView = LogEvents
    }
  },
  watch: {}
}
</script>

<style scoped>
.mdc-textfield {
  height: 48px;
}

.mdc-toolbar {
  background-color: var(--mdc-theme-background, white)
}

.mdc-toolbar--fixed {
  top: initial;
  bottom: 0;
}

.mdc-toolbar__section {
  padding: 10px 50px 10px 50px;
}

.mdc-toolbar-fixed-adjust {
  margin-top: 0px;
  margin-bottom: 60px;
}
</style>
