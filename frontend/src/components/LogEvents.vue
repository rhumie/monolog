<template>
<div>
  <div class="mdc-layout-grid">
    <div class="mdc-layout-grid__inner">
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-4 mdc-layout-grid__cell--span-8-tablet">
        <div class="mdc-textfield mdc-textfield--fullwidth">
          <input type="text" id="search-string-textfield" class="mdc-textfield__input" v-model.lazy="searchString">
          <label class="mdc-textfield__label" for="search-string-textfield">Search String</label>
          <div class="mdc-textfield__bottom-line"></div>
        </div>
      </div>
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-4 mdc-layout-grid__cell--span-4-tablet">
        <date-time-picker v-model="searchDateFrom"
          id="search-date-from"
          label="From"
          :editable="false"
          :clearable="true">
        </date-time-picker>
      </div>
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-4 mdc-layout-grid__cell--span-4-tablet">
        <date-time-picker v-model="searchDateTo"
          id="search-date-to"
          label="To"
          :editable="false"
          :clearable="true">
        </date-time-picker>
      </div>
      <!-- <div class="button-area mdc-layout-grid__cell mdc-layout-grid__cell--span-2 mdc-layout-grid__cell--span-4-tablet right-align">
        <button class="mdc-button mdc-button--unelevated">
          <i class="material-icons mdc-button__icon">search</i>Search
        </button>
      </div> -->
    </div>
  </div>
  <div class="mdc-layout-grid mdc-theme--dark">
    <div class="mdc-layout-grid__inner event-tab">
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-12 mdc-layout-grid__cell--span-4-tablet">
        <div id="my-tab-bar-scroller" class="mdc-tab-bar-scroller">
          <div class="mdc-tab-bar-scroller__indicator mdc-tab-bar-scroller__indicator--back">
            <a class="mdc-tab-bar-scroller__indicator__inner material-icons" href="#" aria-label="scroll back button">navigate_before</a>
          </div>
          <div class="mdc-tab-bar-scroller__scroll-frame">
            <nav id="my-scrollable-tab-bar" class="mdc-tab-bar mdc-tab-bar-scroller__scroll-frame__tabs">
              <a v-for="log in selectedLogs" class="mdc-tab mdc-tab--active" @click="doScrollTabClickEvent(log)">{{ log.logStreamName }}</a>
              <span class="mdc-tab-bar__indicator"></span>
            </nav>
          </div>
          <div class="mdc-tab-bar-scroller__indicator mdc-tab-bar-scroller__indicator--forward">
            <a class="mdc-tab-bar-scroller__indicator__inner material-icons" href="#" aria-label="scroll forward button">navigate_next</a>
          </div>
        </div>
      </div>
    </div>
    <div class="mdc-layout-grid__inner event-title">
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-10-desktop mdc-layout-grid__cell--span-6-tablet mdc-layout-grid__cell--span-3-phone ellipsis">
        <span>{{currentLogName}}</span>
      </div>
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-2-desktop mdc-layout-grid__cell--span-2-tablet mdc-layout-grid__cell--span-1-phone right-align">
        <i class="mdc-icon-toggle material-icons" role="button" :aria-pressed="displayTimestamp"
          aria-label="Add to alerm" tabindex="0"
          data-toggle-on='{"label": "Remove from alerm", "content": "alarm_on"}'
          data-toggle-off='{"label": "Add to alerm", "content": "alarm_off"}'
          @MDCIconToggle:change="doDisplayTimestampChangeEvent">
          alarm_on
        </i>
        
        <i class="mdc-icon-toggle material-icons" role="button" :aria-pressed="searchFromHead"
          :class="{'mdc-icon-toggle--disabled':searchString}"
          aria-label="Add to trending" tabindex="0"
          data-toggle-on='{"label": "Remove from trending", "content": "vertical_align_bottom"}'
          data-toggle-off='{"label": "Add to trending", "content": "vertical_align_top"}'
          @MDCIconToggle:change="doSearchFromHeadChangeEvent">
          vertical_align_bottom
        </i>
      </div>
    </div>
    <div class="mdc-layout-grid__inner event-content">
      <div class="mdc-layout-grid__cell mdc-layout-grid__cell--span-12">
        <log-event v-for="log in selectedLogs"
          v-show="currentLog === log.logStreamId"
          :key="log.logStreamId"
          :awsRegion="log.awsRegion"
          :logGroupName="log.logGroupName"
          :logStreamName="log.logStreamName"
          :filterPattern="searchString"
          :startTime="searchDateFrom"
          :endTime="searchDateTo"
          :startFromHead="searchFromHead"
          :displayTimestamp="displayTimestamp">
        </log-event>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import {MDCTextfield} from '@material/textfield'
import {MDCTabBarScroller} from '@material/tabs'
import {MDCIconToggle} from '@material/icon-toggle'
import LogEvent from './LogEvent'
import DateTimePicker from './common/DateTimePicker'

export default {
  name: 'LogEvents',
  components: {
    LogEvent,
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
      searchString: '',
      searchDateFrom: null,
      searchDateTo: null,
      searchFromHead: false,
      displayTimestamp: true,
      currentLog: this.selectedLogs[0].logStreamId,
      currentLogGroupName: this.selectedLogs[0].logGroupName,
      currentLogStreamName: this.selectedLogs[0].logStreamName,
      eventStyle: {}
    }
  },
  created () {
    // window.addEventListener('resize', () => {
    //   const windowHeight = window.innerHeight || document.documentElement.clientHeight || 0
    //   const htmlHeight = document.querySelector('main').offsetHeight
    //   this.eventStyle = {
    //     height: `${windowHeight - htmlHeight}px`
    //   }
    // })
  },
  mounted () {
    MDCTabBarScroller.attachTo(document.querySelector('#my-tab-bar-scroller'))
    // Initialize text fields.
    this.$el.querySelectorAll('.mdc-textfield').forEach((element) => {
      MDCTextfield.attachTo(element)
    })
    // Initialize toggle icon.
    this.$el.querySelectorAll('.mdc-icon-toggle').forEach((element) => {
      MDCIconToggle.attachTo(element)
    })
  },
  computed: {
    currentLogName: function () {
      return `${this.currentLogGroupName} > ${this.currentLogStreamName}`
    }
  },
  methods: {
    doScrollTabClickEvent: function (log) {
      this.currentLog = log.logStreamId
      this.currentLogGroupName = log.logGroupName
      this.currentLogStreamName = log.logStreamName
    },
    doDisplayTimestampChangeEvent: function (event) {
      this.displayTimestamp = event.detail.isOn
    },
    doSearchFromHeadChangeEvent: function (event) {
      this.searchFromHead = event.detail.isOn
    }
  },
  watch: {
  }
}
</script>

<style scoped>
.mdc-textfield {
  height: 48px;
}

.mdc-tab {
  min-width: 160px;
  max-width: 160px;
  text-overflow: ellipsis;
}

.mdc-layout-grid {
  padding-top: 18px;
  padding-bottom: 0;
}

.mdc-tab-bar {
  margin: 0;
  height: 30px;
}

.button-area {
  position: relative;
}

.button-area .mdc-button {
  position: absolute;
  right: 0;
  bottom: 0;
}

.mdc-layout-grid__cell {
  width: 100%
}

.event-tab {
  background-color: #424242;
}

.event-title {
  background-color: #424242;
  color: #EEEEEE;
  font-size: 14px;
  overflow-x: auto;
  padding: 5px 24px 5px 24px;
}

.event-content {
  background-color: #EEEEEE;
}

.mdc-icon-toggle {
  padding: 5px 24px 5px 24px;
  text-align: right;
  display: inline;
  font-size: 18px;
  padding: 0;
  --mdc-ripple-fg-size: 18px !important;
  --mdc-ripple-left: 0 !important;
  --mdc-ripple-top: 0 !important;
  --mdc-ripple-fg-scale: 2 !important;
}

@media (min-width: 840px) {
  .mdc-tab-bar-scroller__indicator {
    width: var(--mdc-layout-grid-margin-desktop, 24px);
  }
}

@media (max-width: 839px) and (min-width: 480px){
  .mdc-tab-bar-scroller__indicator {
    width: var(--mdc-layout-grid-margin-tablet, 16px)
  }
}

@media (max-width: 479px) {
  .mdc-tab-bar-scroller__indicator {
    width: var(--mdc-layout-grid-margin-phone, 16px)
  }  
}

</style>
