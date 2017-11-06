<template>
<div>
  <div class="mdc-layout-grid">
    <div class="mdc-layout-grid__inner">
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-3 mdc-layout-grid__cell--span-4-tablet">
        <div class="mdc-textfield mdc-textfield--fullwidth">
          <input type="text" id="log-group-name-textfield" class="mdc-textfield__input" v-model="searchLogGroupName">
          <label class="mdc-textfield__label" for="log-group-name-textfield">Log Group Name</label>
          <div class="mdc-textfield__bottom-line"></div>
        </div>
      </div>
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-3 mdc-layout-grid__cell--span-4-tablet">
        <div class="mdc-textfield mdc-textfield--fullwidth">
          <input type="text" id="log-stream-name-textfield" class="mdc-textfield__input" v-model="searchLogStreamName">
          <label class="mdc-textfield__label" for="log-stream-name-textfield">Log Stream Name</label>
          <div class="mdc-textfield__bottom-line"></div>
        </div>
      </div>
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-3 mdc-layout-grid__cell--span-4-tablet">
        <select-menu-aws-region id="select-menu-aws-region" :defaultRegion="this.searchAWSRegion" :fullwidth="true" @change="doChangeAWSRegionEvent">
        </select-menu-aws-region>
      </div>
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-3 mdc-layout-grid__cell--span-4-tablet">
        <date-time-picker v-model="searchDateFrom" id="search-from-date" label="Last Timestamp later than" :editable="false" :clearable="true">
        </date-time-picker>
      </div>
    </div>
  </div>
  <div class="mdc-list-group">
    <div v-for="logGroup in filteredLogGroups">
      <h3 class="mdc-list-group__subheader">{{ logGroup.name }}</h3>
      <ul class="mdc-list mdc-list--dense">
        <li class="mdc-list-item" v-for="logStream in logGroup.filteredLogStreams">
          <div class="mdc-checkbox mdc-list-item__start-detail">
            <input type="checkbox" :id="logStream.id" class="mdc-checkbox__native-control" :value="createValue(logGroup, logStream)" v-model="selectedLogStreams" />
            <div class="mdc-checkbox__background">
              <svg class="mdc-checkbox__checkmark" viewBox="0 0 24 24">
                  <path class="mdc-checkbox__checkmark__path" fill="none" stroke="white" d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
                </svg>
              <div class="mdc-checkbox__mixedmark"></div>
            </div>
          </div>
          <label v-bind:for="logStream.id">
              <span class="mdc-list-item__text">{{ logStream.name }}<span class="mdc-list-item__text__secondary">{{ dateToString(logStream.lastEventTimestamp) }}</span></span>
          </label>
          <!-- <button class="mdc-button mdc-button--dense mdc-list-item__end-detail "><i class="material-icons ">get_app</i></button> -->
        </li>
      </ul>
      <hr class="mdc-list-divider">
    </div>
  </div>
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
</div>
</template>

<script>
import axios from 'axios'
import { MDCTextfield } from '@material/textfield'
import SelectMenuAwsRegion from './common/SelectMenuAwsRegion'
import DateTimePicker from './common/DateTimePicker'

export default {
  name: 'LogIndex',
  components: {
    SelectMenuAwsRegion,
    DateTimePicker
  },
  props: {
    selectedLogs: {
      type: Array,
      required: true
    }
  },
  data () {
    return {
      searchLogGroupName: '',
      searchLogStreamName: '',
      searchAWSRegion: 'ap-northeast-1',
      searchDateFrom: null,
      resultLogGroups: [],
      selectedLogStreams: this.selectedLogs.map(elem => JSON.stringify(elem)),
      isLoading: true
    }
  },
  methods: {
    async searchLogs () {
      this.isLoading = true
      const { res, error } = await axios.get('/api/logs', {
        headers: {
          'X-ServiceToken': `Logs ${localStorage.getItem('token') || ''}`
        },
        params: {
          awsRegion: this.searchAWSRegion
        }
      })
      if (error) {
        if (error.status === 401) {
          this.$router.push('/login')
          return
        }
      }
      this.resultLogGroups = res.data.data.logGroups
      this.isLoading = false
    },
    async doChangeAWSRegionEvent (awsRegion) {
      this.selectedLogStreams = []
      this.resultLogGroups = []
      this.searchAWSRegion = awsRegion
      await this.searchLogs()
    },
    createValue (logGroup, logStream) {
      return JSON.stringify({
        awsRegion: this.searchAWSRegion,
        logGroupName: logGroup.name,
        logStreamName: logStream.name,
        logStreamId: logStream.id
      })
    },
    dateToString (val) {
      const date = new Date(val)
      const toDoubleDigits = function (num) {
        num += ''
        if (num.length === 1) {
          num = '0' + num
        }
        return num
      }
      const yyyy = date.getFullYear()
      const mm = toDoubleDigits(date.getMonth() + 1)
      const dd = toDoubleDigits(date.getDate())
      const hh = toDoubleDigits(date.getHours())
      const mi = toDoubleDigits(date.getMinutes())
      const ss = toDoubleDigits(date.getSeconds())
      return `${yyyy}/${mm}/${dd} ${hh}:${mi}:${ss}`
    }
  },
  computed: {
    filteredLogGroups: function () {
      const searchLogGroupName = this.searchLogGroupName.toLowerCase()
      const searchLogStreamName = this.searchLogStreamName.toLowerCase()
      // Filter by log group name.
      const filteredLogGroups = this.resultLogGroups.filter(elem => elem.name.toLowerCase().indexOf(searchLogGroupName) !== -1)
      // Filter by log stream name.
      filteredLogGroups.forEach(elem => {
        elem.filteredLogStreams = elem.logStreams.filter(elem => {
          return elem.name.toLowerCase().indexOf(searchLogStreamName) !== -1 && elem.lastEventTimestamp - this.searchDateFrom > 0
        })
      })
      return filteredLogGroups
    }
  },
  watch: {
    selectedLogStreams: function () {
      if (this.selectedLogStreams.length > 10) {}
      this.$emit('selectLogStream', this.selectedLogStreams.map(elem => JSON.parse(elem)))
    },
    searchAWSRegion: function () {
      this.$emit('changeAWSRegion', this.searchAWSRegion)
    }
  },
  created () {
    this.searchLogs()
  },
  mounted () {
    this.$el.querySelectorAll('.mdc-textfield').forEach((element) => {
      MDCTextfield.attachTo(element)
    })
    // snackbar = new MDCSnackbar(document.querySelector('.mdc-snackbar'))
  }
}
</script>

<style scoped>
.mdc-textfield {
  height: 48px;
  font-size: 0.936rem;
}

.mdc-list {
  font-family: inherit;
}

.mdc-layout-grid {
  padding-top: 18px;
  padding-bottom: 0;
}

.mdc-layout-grid__cell {
  height: 50px;
}

.mdc-linear-progress {
  top: 200px
}
</style>
