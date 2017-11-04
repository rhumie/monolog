<template>
<div>
  <select-menu
    :label="label"
    :options="selectOptions"
    :fullwidth="fullwidth"
    :id="id"
    @change="change"></select-menu>
</div>
</template>

<script>
import SelectMenu from './SelectMenu'

export default {
  name: 'select-menu-aws-region',
  props: {
    id: {
      type: String,
      required: true
    },
    label: {
      type: String,
      default: 'AWS Region'
    },
    fullwidth: {
      type: Boolean
    },
    defaultRegion: {
      type: String
    }
  },
  data () {
    return {
      selectOptions: [{
        id: 'us-east-1',
        label: 'us-east-1',
        selected: this.isDefaultRegion('us-east-1')
      }, {
        id: 'us-west-1',
        label: 'us-west-1',
        selected: this.isDefaultRegion('us-west-1')
      }, {
        id: 'us-west-2',
        label: 'us-west-2',
        selected: this.isDefaultRegion('us-west-2')
      }, {
        id: 'eu-west-1',
        label: 'eu-west-1',
        selected: this.isDefaultRegion('eu-west-1')
      }, {
        id: 'eu-central-1',
        label: 'eu-central-1',
        selected: this.isDefaultRegion('eu-west-1')
      }, {
        id: 'ap-southeast-1',
        label: 'ap-southeast-1',
        selected: this.isDefaultRegion('ap-southeast-1')
      }, {
        id: 'ap-southeast-2',
        label: 'ap-southeast-2',
        selected: this.isDefaultRegion('ap-southeast-2')
      }, {
        id: 'ap-northeast-1',
        label: 'ap-northeast-1',
        selected: this.isDefaultRegion('ap-northeast-1')
      }, {
        id: 'sa-east-1',
        label: 'sa-east-1',
        selected: this.isDefaultRegion('sa-east-1')
      }]
    }
  },
  components: {
    SelectMenu
  },
  computed: {
  },
  methods: {
    isDefaultRegion (val) {
      return val === this.defaultRegion
    },
    change (val) {
      this.$emit('change', val)
    }
  }
}
</script>
