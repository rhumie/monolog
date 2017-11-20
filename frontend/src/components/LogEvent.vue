<template>
<div>
  <div role="progressbar" class="mdc-linear-progress mdc-linear-progress--indeterminate" v-show="isLoading">
    <div class="mdc-linear-progress__buffering-dots"></div>
    <div class="mdc-linear-progress__buffer"></div>
    <div class="mdc-linear-progress__bar mdc-linear-progress__primary-bar">
      <span class="mdc-linear-progress__bar-inner"></span>
    </div>
    <div class="mdc-linear-progress__bar mdc-linear-progress__secondary-bar">
      <span class="mdc-linear-progress__bar-inner"></span>
    </div>
  </div>
  <div class="scroll-area" @scroll="scroll">
    <div v-for="event in events" class="row">
      <span class="column" v-show="displayTimestamp">{{ dateToString(event.timestamp) }}</span><span class="column">{{event.message}}</span>
    </div>
  </div>
</div>
</template>

<script>
import axios from 'axios'
import Base from '@/components/Base'

export default {
  name: 'LogEvents',
  mixins: [
    Base
  ],
  props: {
    awsRegion: {
      type: String,
      required: true
    },
    logGroupName: {
      type: String,
      required: true
    },
    logStreamName: {
      type: String,
      required: true
    },
    startFromHead: {
      type: Boolean,
      default: this.filterPattern
    },
    filterPattern: {
      type: String,
      default: ''
    },
    startTime: {
      type: Date
    },
    endTime: {
      type: Date
    },
    displayTimestamp: {
      type: Boolean,
      default: true
    }
  },
  data () {
    return {
      events: [],
      nextBackwardToken: '',
      nextForwardToken: '',
      currentScrollTop: 0,
      isLoading: false
    }
  },
  methods: {
    async getLogEvents (customParams = {}) {
      if (this.isLoading) {
        return
      }
      this.isLoading = true
      const { res, error } = await axios.get('/api/logevents', {
        params: Object.assign({
          awsRegion: this.awsRegion,
          logStreamName: this.logStreamName,
          logGroupName: this.logGroupName,
          startFromHead: this.startFromHead,
          filterPattern: this.filterPattern,
          startTime: this.startTime ? this.startTime.getTime() : 0,
          endTime: this.endTime ? this.endTime.getTime() : 0
        }, customParams)
      })
      if (error) {
        if (error.status === 401) {
          this.$router.push('/login')
        }
      }

      switch (customParams.nextToken) {
        // Scroll down.
        case this.nextForwardToken:
          this.nextForwardToken = res.data.data.nextForwardToken
          this.events = [...this.events, ...res.data.data.logEvents]
          break
        // Scroll up.
        case this.nextBackwardToken:
          this.nextBackwardToken = res.data.data.nextBackwardToken
          this.events = [...res.data.data.logEvents, ...this.events]
          break
         // Initial
        default:
          this.nextBackwardToken = res.data.data.nextBackwardToken
          this.nextForwardToken = res.data.data.nextForwardToken
          this.events = [...res.data.data.logEvents]
          break
      }
      this.isLoading = false
    },
    async scroll (event) {
      // Do nothing when horizontal scroll.
      if (this.currentScrollTop === event.target.scrollTop) {
        return
      }
      // Scroll to top.
      if (event.target.scrollTop === 0 && this.nextBackwardToken !== '') {
        const scrollHeight = event.target.scrollHeight
        await this.getLogEvents({
          nextToken: this.nextBackwardToken
        })
        // Set scroll positon to the current row position.
        event.target.scrollTop = event.target.scrollHeight - scrollHeight
      }
      // Scroll to bottom.
      if (event.target.scrollTop + event.target.offsetHeight >= event.target.scrollHeight && this.nextForwardToken !== '') {
        await this.getLogEvents({
          nextToken: this.nextForwardToken
        })
      }
      this.currentScrollTop = event.target.scrollTop
    },
    clear () {
      this.events = []
      this.nextForwardToken = ''
      this.nextBackwardToken = ''
    },
    scrollTop () {
      const scroll = this.$el.querySelector('.scroll-area')
      scroll.scrollTop = 0
    },
    scrollBottom () {
      const scroll = this.$el.querySelector('.scroll-area')
      scroll.scrollTop = scroll.scrollHeight
    },
    async initialize () {
      this.clear()
      await this.getLogEvents()
      if (this.startFromHead) {
        this.scrollTop()
      } else {
        this.scrollBottom()
      }
    }
  },
  watch: {
    filterPattern (val) {
      this.initialize()
    },
    startTime (val) {
      this.initialize()
    },
    endTime (val) {
      this.initialize()
    },
    startFromHead (val) {
      this.initialize()
    }
  },
  mounted () {
    this.initialize()
  }
}
</script>

<style scoped>
.scroll-area {
  font-size: 0.75rem;
  height: 450px;
  overflow-x: auto;
  overflow-y: scroll;
}

.column {
  display: inline-block;
  white-space: pre;
  vertical-align: top;
}

.row {
  white-space: nowrap;
}
</style>
